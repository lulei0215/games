package payment

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// 创建交易响应结构体
type CreateTradeResponse struct {
	Success   bool   `json:"success"`
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
	Data      struct {
		MerchantId      interface{} `json:"merchantId"`
		MerchantOrderNo string      `json:"merchantOrderNo"`
		OrderNo         string      `json:"orderNo"`
		PayUrl          string      `json:"payUrl"`
		PayRaw          string      `json:"payRaw,omitempty"`
		Amount          int64       `json:"amount"`
		Status          string      `json:"status"`
		Currency        string      `json:"currency"`
		PayType         string      `json:"payType"`
	} `json:"data"`
}

func (c CreateTradeResponse) OkWithDetailed(h gin.H, s string, param3 *gin.Context) {
	panic("unimplemented")
}

// 查询交易订单响应结构体
type QueryTradeResponse struct {
	Success   bool   `json:"success"`
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
	Data      struct {
		MerchantId      interface{} `json:"merchantId"`
		MerchantOrderNo string      `json:"merchantOrderNo"`
		OrderNo         string      `json:"orderNo"`
		Amount          int64       `json:"amount"`
		Status          string      `json:"status"`
		Currency        string      `json:"currency"`
		PayType         string      `json:"payType"`
		RefCpf          string      `json:"ref_cpf,omitempty"`
		RefName         string      `json:"ref_name,omitempty"`
	} `json:"data"`
}

// 交易回调通知结构体
type TradeCallbackRequest struct {
	Success   bool   `json:"success"`
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
	Data      struct {
		MerchantId      interface{} `json:"merchantId"`
		MerchantOrderNo string      `json:"merchantOrderNo"`
		OrderNo         string      `json:"orderNo"`
		Amount          int64       `json:"amount"`
		Status          string      `json:"status"`
		Currency        string      `json:"currency"`
		PayType         string      `json:"payType"`
		RefCpf          string      `json:"ref_cpf,omitempty"`
		RefName         string      `json:"ref_name,omitempty"`
		Sign            string      `json:"sign"`
	} `json:"data"`
}

// 创建提现响应结构体
type CreatePaymentResponse struct {
	Success   bool   `json:"success"`
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
	Data      struct {
		MerchantId      interface{} `json:"merchantId"`
		MerchantOrderNo string      `json:"merchantOrderNo"`
		OrderNo         string      `json:"orderNo"`
		Amount          int64       `json:"amount"`
		Status          string      `json:"status"`
		Currency        string      `json:"currency"`
		ErrorMsg        string      `json:"errorMsg,omitempty"`
	} `json:"data"`
}

// 查询提现响应结构体
type QueryPaymentResponse struct {
	Success   bool   `json:"success"`
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
	Data      struct {
		MerchantId      interface{} `json:"merchantId"`
		MerchantOrderNo string      `json:"merchantOrderNo"`
		OrderNo         string      `json:"orderNo"`
		Amount          int64       `json:"amount"`
		Status          string      `json:"status"`
		Currency        string      `json:"currency"`
		ErrorMsg        string      `json:"errorMsg,omitempty"`
	} `json:"data"`
}

// 提现回调通知结构体
type PaymentCallbackRequest struct {
	Success   bool   `json:"success"`
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
	Data      struct {
		MerchantId      interface{} `json:"merchantId"`
		MerchantOrderNo string      `json:"merchantOrderNo"`
		OrderNo         string      `json:"orderNo"`
		Amount          int64       `json:"amount"`
		Status          string      `json:"status"`
		Currency        string      `json:"currency"`
		ErrorMsg        string      `json:"errorMsg,omitempty"`
		Sign            string      `json:"sign"`
	} `json:"data"`
}

// 余额查询响应结构体
type BalanceResponse struct {
	Success   bool   `json:"success"`
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
	Data      struct {
		Balance          int64  `json:"balance"`
		UnsettledBalance int64  `json:"unsettledBalance"`
		FrozenAmount     int64  `json:"frozenAmount"`
		Currency         string `json:"currency"`
	} `json:"data"`
}

// 提现反查响应结构体
type ReversePaymentResponse struct {
	Success   bool   `json:"success"`
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
	Data      struct {
		MerchantId      interface{} `json:"merchantId"`
		MerchantOrderNo string      `json:"merchantOrderNo"`
		OrderNo         string      `json:"orderNo"`
		Amount          int64       `json:"amount"`
		Status          string      `json:"status"`
		Currency        string      `json:"currency"`
		ErrorMsg        string      `json:"errorMsg,omitempty"`
	} `json:"data"`
}

// 支付客户端
type PaymentClient struct {
	BaseURL    string
	MerchantId string
	SecretKey  string
}

// 回调去重管理器
type CallbackDeduplicator struct {
	processedCallbacks map[string]bool
	mutex              sync.RWMutex
}

func NewCallbackDeduplicator() *CallbackDeduplicator {
	return &CallbackDeduplicator{
		processedCallbacks: make(map[string]bool),
	}
}

func (d *CallbackDeduplicator) IsProcessed(merchantOrderNo, orderNo string) bool {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	key := fmt.Sprintf("%s_%s", merchantOrderNo, orderNo)
	return d.processedCallbacks[key]
}

func (d *CallbackDeduplicator) MarkProcessed(merchantOrderNo, orderNo string) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	key := fmt.Sprintf("%s_%s", merchantOrderNo, orderNo)
	d.processedCallbacks[key] = true
}

// 表单签名生成方法，添加详细的调试信息
func (pc *PaymentClient) GenerateFormSign(params url.Values) string {
	fmt.Println("\n🔐 === 表单签名生成过程 ===")

	// 1. 获取所有非空参数
	var keys []string
	validParams := make(map[string]string)

	fmt.Println("1️⃣ 基本参数:")
	for key, values := range params {
		if len(values) > 0 && values[0] != "" {
			validParams[key] = values[0]
			keys = append(keys, key)
			fmt.Printf("   %s=\"%s\"\n", key, values[0])
		}
	}

	// 2. 按key的ASCII码升序排序
	sort.Strings(keys)

	// 3. 拼接参数字符串
	var paramPairs []string
	fmt.Println("\n2️⃣ 排序之后:")
	for _, key := range keys {
		paramPairs = append(paramPairs, key+"="+validParams[key])
	}
	paramString := strings.Join(paramPairs, "&")
	fmt.Printf("   %s\n", paramString)

	// 4. 拼接密钥
	signString := paramString + "&secret=" + pc.SecretKey
	fmt.Println("\n3️⃣ 拼接密钥:")
	fmt.Printf("   %s\n", signString)

	// 5. MD5签名
	hash := md5.Sum([]byte(signString))
	signature := hex.EncodeToString(hash[:])
	fmt.Println("\n4️⃣ MD5签名:")
	fmt.Printf("   %s\n", signature)

	return signature
}

// 回调签名验证（不包含URL参数）
func (pc *PaymentClient) VerifyCallbackSign(data map[string]interface{}, receivedSign string) bool {
	fmt.Println("\n🔍 === 回调签名验证过程 ===")

	// 排除sign字段
	validParams := make(map[string]string)
	for k, v := range data {
		if k != "sign" && v != nil && v != "" {
			validParams[k] = fmt.Sprintf("%v", v)
		}
	}

	// 按key的ASCII码升序排序
	var keys []string
	for k := range validParams {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 构建签名字符串
	var paramPairs []string
	for _, key := range keys {
		paramPairs = append(paramPairs, key+"="+validParams[key])
	}
	paramString := strings.Join(paramPairs, "&")

	// 拼接密钥
	signString := paramString + "&secret=" + pc.SecretKey

	// MD5签名
	hash := md5.Sum([]byte(signString))
	calculatedSign := hex.EncodeToString(hash[:])

	fmt.Printf("计算签名: %s\n", calculatedSign)
	fmt.Printf("接收签名: %s\n", receivedSign)

	return calculatedSign == receivedSign
}

// 处理交易成功回调
func (pc *PaymentClient) processTradeSuccess(data TradeCallbackRequest) {
	fmt.Println("✅ 处理交易成功回调")
	fmt.Printf("   商户订单号: %s\n", data.Data.MerchantOrderNo)
	fmt.Printf("   平台订单号: %s\n", data.Data.OrderNo)
	fmt.Printf("   交易金额: %d 分\n", data.Data.Amount)
	fmt.Printf("   交易状态: %s\n", data.Data.Status)
	// 这里可以添加业务逻辑，如更新数据库、发送通知等
}

// 处理交易失败回调
func (pc *PaymentClient) processTradeFailure(data TradeCallbackRequest) {
	fmt.Println("❌ 处理交易失败回调")
	fmt.Printf("   商户订单号: %s\n", data.Data.MerchantOrderNo)
	fmt.Printf("   平台订单号: %s\n", data.Data.OrderNo)
	fmt.Printf("   交易金额: %d 分\n", data.Data.Amount)
	fmt.Printf("   交易状态: %s\n", data.Data.Status)
	// 这里可以添加业务逻辑，如更新数据库、发送通知等
}

// 处理提现成功回调
func (pc *PaymentClient) processPaymentSuccess(data PaymentCallbackRequest) {
	fmt.Println("✅ 处理提现成功回调")
	fmt.Printf("   商户订单号: %s\n", data.Data.MerchantOrderNo)
	fmt.Printf("   平台订单号: %s\n", data.Data.OrderNo)
	fmt.Printf("   提现金额: %d 分\n", data.Data.Amount)
	fmt.Printf("   提现状态: %s\n", data.Data.Status)
	// 这里可以添加业务逻辑，如更新数据库、发送通知等
}

// 处理提现失败回调
func (pc *PaymentClient) processPaymentFailure(data PaymentCallbackRequest) {
	fmt.Println("❌ 处理提现失败回调")
	fmt.Printf("   商户订单号: %s\n", data.Data.MerchantOrderNo)
	fmt.Printf("   平台订单号: %s\n", data.Data.OrderNo)
	fmt.Printf("   提现金额: %d 分\n", data.Data.Amount)
	fmt.Printf("   提现状态: %s\n", data.Data.Status)
	if data.Data.ErrorMsg != "" {
		fmt.Printf("   失败原因: %s\n", data.Data.ErrorMsg)
	}
	// 这里可以添加业务逻辑，如更新数据库、发送通知等
}

// 1. 创建交易
func (pc *PaymentClient) CreateTrade() {
	fmt.Println("=== 创建交易 ===")

	// 使用固定参数进行测试
	formData := url.Values{}
	formData.Set("merchantId", pc.MerchantId)
	formData.Set("merchantOrderNo", fmt.Sprintf("ORDER_%d", time.Now().Unix()))
	formData.Set("amount", "100")
	formData.Set("payType", "PIX_QRCODE")
	formData.Set("currency", "BRL")
	formData.Set("content", "测试订单")
	formData.Set("clientIp", "192.168.1.100")
	formData.Set("callback", "https://your-domain.com/callback")
	formData.Set("redirect", "https://your-domain.com/success")
	formData.Set("reqTimesTamp", strconv.FormatInt(time.Now().UTC().Unix(), 10))

	fmt.Println("\n📝 请求参数:")
	for k, v := range formData {
		fmt.Printf("  %s: %s\n", k, v[0])
	}

	// 生成签名
	signature := pc.GenerateFormSign(formData)
	formData.Set("sign", signature)

	// 发送请求
	resp, err := http.PostForm(pc.BaseURL+"/api/open/merchant/trade/create", formData)
	if err != nil {
		fmt.Printf("❌ 请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ 读取响应失败: %v\n", err)
		return
	}

	fmt.Printf("\n📡 响应状态: %s\n", resp.Status)
	fmt.Printf("📄 响应内容: %s\n", string(body))

	// 解析响应
	var response CreateTradeResponse
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Printf("❌ 解析响应失败: %v\n", err)
		return
	}

	if response.Success {
		fmt.Println("✅ 创建交易成功!")
		fmt.Printf("   平台订单号: %s\n", response.Data.OrderNo)
		fmt.Printf("   支付链接: %s\n", response.Data.PayUrl)
		fmt.Printf("   订单状态: %s\n", response.Data.Status)
	} else {
		fmt.Printf("❌ 创建交易失败: %s\n", response.Message)
	}
}

// 2. 查询交易订单
func (pc *PaymentClient) QueryTrade() {
	fmt.Println("=== 查询交易订单 ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("请输入商户订单号: ")
	merchantOrderNo, _ := reader.ReadString('\n')
	merchantOrderNo = strings.TrimSpace(merchantOrderNo)
	if merchantOrderNo == "" {
		fmt.Println("❌ 商户订单号不能为空")
		return
	}

	formData := url.Values{}
	formData.Set("merchantId", pc.MerchantId)
	formData.Set("merchantOrderNo", merchantOrderNo)
	formData.Set("reqTimesTamp", strconv.FormatInt(time.Now().UTC().Unix(), 10))

	// 生成签名
	signature := pc.GenerateFormSign(formData)
	formData.Set("sign", signature)

	// 发送请求
	resp, err := http.PostForm(pc.BaseURL+"/api/open/merchant/trade/query", formData)
	if err != nil {
		fmt.Printf("❌ 请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ 读取响应失败: %v\n", err)
		return
	}

	fmt.Printf("\n📡 响应状态: %s\n", resp.Status)
	fmt.Printf("📄 响应内容: %s\n", string(body))

	// 解析响应
	var response QueryTradeResponse
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Printf("❌ 解析响应失败: %v\n", err)
		return
	}

	if response.Success {
		fmt.Println("✅ 查询成功!")
		fmt.Printf("   平台订单号: %s\n", response.Data.OrderNo)
		fmt.Printf("   订单金额: %d 分\n", response.Data.Amount)
		fmt.Printf("   订单状态: %s\n", response.Data.Status)
		fmt.Printf("   币种: %s\n", response.Data.Currency)
	} else {
		fmt.Printf("❌ 查询失败: %s\n", response.Message)
	}
}

// 3. 交易回调通知接收服务器
func (pc *PaymentClient) StartTradeCallbackServer() {
	fmt.Println("=== 启动交易回调通知服务器 ===")

	deduplicator := NewCallbackDeduplicator()

	http.HandleFunc("/trade/callback", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "只允许POST请求", http.StatusMethodNotAllowed)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("❌ 读取请求体失败: %v\n", err)
			http.Error(w, "读取请求体失败", http.StatusBadRequest)
			return
		}

		fmt.Printf("📨 收到交易回调通知: %s\n", string(body))

		var callbackData TradeCallbackRequest
		if err := json.Unmarshal(body, &callbackData); err != nil {
			fmt.Printf("❌ 解析回调数据失败: %v\n", err)
			http.Error(w, "解析数据失败", http.StatusBadRequest)
			return
		}

		// 检查是否已处理过
		if deduplicator.IsProcessed(callbackData.Data.MerchantOrderNo, callbackData.Data.OrderNo) {
			fmt.Printf("⚠️ 重复回调，跳过处理: %s\n", callbackData.Data.MerchantOrderNo)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success":   true,
				"errorCode": "success",
				"message":   "回调已处理",
			})
			return
		}

		// 验证签名（回调签名不包含URL参数）
		dataMap := map[string]interface{}{
			"merchantId":      callbackData.Data.MerchantId,
			"merchantOrderNo": callbackData.Data.MerchantOrderNo,
			"orderNo":         callbackData.Data.OrderNo,
			"amount":          callbackData.Data.Amount,
			"status":          callbackData.Data.Status,
			"currency":        callbackData.Data.Currency,
			"payType":         callbackData.Data.PayType,
		}

		if callbackData.Data.RefCpf != "" {
			dataMap["ref_cpf"] = callbackData.Data.RefCpf
		}
		if callbackData.Data.RefName != "" {
			dataMap["ref_name"] = callbackData.Data.RefName
		}

		if !pc.VerifyCallbackSign(dataMap, callbackData.Data.Sign) {
			fmt.Println("❌ 回调签名验证失败")
			http.Error(w, "签名验证失败", http.StatusUnauthorized)
			return
		}

		// 标记为已处理
		deduplicator.MarkProcessed(callbackData.Data.MerchantOrderNo, callbackData.Data.OrderNo)

		// 根据状态处理回调
		switch callbackData.Data.Status {
		case "PAID":
			pc.processTradeSuccess(callbackData)
		case "PAY_FAILED":
			pc.processTradeFailure(callbackData)
		default:
			fmt.Printf("📝 其他状态回调: %s\n", callbackData.Data.Status)
		}
		// 返回成功响应
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success":   true,
			"errorCode": "success",
			"message":   "回调处理成功",
		})
	})

	fmt.Println("🚀 交易回调服务器启动在 http://localhost:8081/trade/callback")
	fmt.Println("💡 使用 Ctrl+C 停止服务器")

	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Printf("❌ 服务器启动失败: %v\n", err)
	}
}

// 4. 提现申请
func (pc *PaymentClient) CreatePayment() {
	fmt.Println("=== 提现申请 ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("请输入商户订单号: ")
	merchantOrderNo, _ := reader.ReadString('\n')
	merchantOrderNo = strings.TrimSpace(merchantOrderNo)
	if merchantOrderNo == "" {
		merchantOrderNo = fmt.Sprintf("PAY_%d", time.Now().Unix())
		fmt.Printf("使用默认订单号: %s\n", merchantOrderNo)
	}

	fmt.Print("请输入提现金额/分: ")
	amountStr, _ := reader.ReadString('\n')
	amountStr = strings.TrimSpace(amountStr)
	if amountStr == "" {
		amountStr = "100"
		fmt.Printf("使用默认金额: %s 分\n", amountStr)
	}

	fmt.Print("请输入币种 (BRL/USD/CNY): ")
	currency, _ := reader.ReadString('\n')
	currency = strings.TrimSpace(currency)
	if currency == "" {
		currency = "BRL"
		fmt.Printf("使用默认币种: %s\n", currency)
	}

	fmt.Print("请输入账户类型 (PERSONAL_BANK/COMPANY_BANK): ")
	accountType, _ := reader.ReadString('\n')
	accountType = strings.TrimSpace(accountType)
	if accountType == "" {
		accountType = "PERSONAL_BANK"
		fmt.Printf("使用默认账户类型: %s\n", accountType)
	}

	fmt.Print("请输入账号: ")
	accountNo, _ := reader.ReadString('\n')
	accountNo = strings.TrimSpace(accountNo)
	if accountNo == "" {
		accountNo = "123456789"
		fmt.Printf("使用默认账号: %s\n", accountNo)
	}

	fmt.Print("请输入账户名: ")
	accountName, _ := reader.ReadString('\n')
	accountName = strings.TrimSpace(accountName)
	if accountName == "" {
		accountName = "Test User"
		fmt.Printf("使用默认账户名: %s\n", accountName)
	}

	formData := url.Values{}
	formData.Set("merchantId", pc.MerchantId)
	formData.Set("merchantOrderNo", merchantOrderNo)
	formData.Set("amount", amountStr)
	formData.Set("currency", currency)
	formData.Set("accountType", accountType)
	formData.Set("accountNo", accountNo)
	formData.Set("accountName", accountName)
	formData.Set("reqTimesTamp", strconv.FormatInt(time.Now().UTC().Unix(), 10))

	// 生成签名
	signature := pc.GenerateFormSign(formData)
	formData.Set("sign", signature)

	// 发送请求
	resp, err := http.PostForm(pc.BaseURL+"/api/open/merchant/payment/create", formData)
	if err != nil {
		fmt.Printf("❌ 请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ 读取响应失败: %v\n", err)
		return
	}

	fmt.Printf("\n📡 响应状态: %s\n", resp.Status)
	fmt.Printf("📄 响应内容: %s\n", string(body))

	// 解析响应
	var response CreatePaymentResponse
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Printf("❌ 解析响应失败: %v\n", err)
		return
	}

	if response.Success {
		fmt.Println("✅ 提现申请成功!")
		fmt.Printf("   平台订单号: %s\n", response.Data.OrderNo)
		fmt.Printf("   提现金额: %d 分\n", response.Data.Amount)
		fmt.Printf("   提现状态: %s\n", response.Data.Status)
	} else {
		fmt.Printf("❌ 提现申请失败: %s\n", response.Message)
	}
}

// 5. 提现查询
func (pc *PaymentClient) QueryPayment() {
	fmt.Println("=== 提现查询 ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("请输入商户订单号: ")
	merchantOrderNo, _ := reader.ReadString('\n')
	merchantOrderNo = strings.TrimSpace(merchantOrderNo)
	if merchantOrderNo == "" {
		fmt.Println("❌ 商户订单号不能为空")
		return
	}

	formData := url.Values{}
	formData.Set("merchantId", pc.MerchantId)
	formData.Set("merchantOrderNo", merchantOrderNo)
	formData.Set("reqTimesTamp", strconv.FormatInt(time.Now().UTC().Unix(), 10))

	// 生成签名
	signature := pc.GenerateFormSign(formData)
	formData.Set("sign", signature)

	// 发送请求
	resp, err := http.PostForm(pc.BaseURL+"/api/open/merchant/payment/query", formData)
	if err != nil {
		fmt.Printf("❌ 请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ 读取响应失败: %v\n", err)
		return
	}

	fmt.Printf("\n📡 响应状态: %s\n", resp.Status)
	fmt.Printf("📄 响应内容: %s\n", string(body))

	// 解析响应
	var response QueryPaymentResponse
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Printf("❌ 解析响应失败: %v\n", err)
		return
	}

	if response.Success {
		fmt.Println("✅ 查询成功!")
		fmt.Printf("   平台订单号: %s\n", response.Data.OrderNo)
		fmt.Printf("   提现金额: %d 分\n", response.Data.Amount)
		fmt.Printf("   提现状态: %s\n", response.Data.Status)
		if response.Data.ErrorMsg != "" {
			fmt.Printf("   错误信息: %s\n", response.Data.ErrorMsg)
		}
	} else {
		fmt.Printf("❌ 查询失败: %s\n", response.Message)
	}
}

// 6. 提现订单回调接收服务器
func (pc *PaymentClient) StartPaymentCallbackServer() {
	fmt.Println("=== 启动提现回调通知服务器 ===")

	deduplicator := NewCallbackDeduplicator()

	http.HandleFunc("/payment/callback", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "只允许POST请求", http.StatusMethodNotAllowed)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("❌ 读取请求体失败: %v\n", err)
			http.Error(w, "读取请求体失败", http.StatusBadRequest)
			return
		}

		fmt.Printf("📨 收到提现回调通知: %s\n", string(body))

		var callbackData PaymentCallbackRequest
		if err := json.Unmarshal(body, &callbackData); err != nil {
			fmt.Printf("❌ 解析回调数据失败: %v\n", err)
			http.Error(w, "解析数据失败", http.StatusBadRequest)
			return
		}

		// 检查是否已处理过
		if deduplicator.IsProcessed(callbackData.Data.MerchantOrderNo, callbackData.Data.OrderNo) {
			fmt.Printf("⚠️ 重复回调，跳过处理: %s\n", callbackData.Data.MerchantOrderNo)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success":   true,
				"errorCode": "success",
				"message":   "回调已处理",
			})
			return
		}

		// 验证签名（回调签名不包含URL参数）
		dataMap := map[string]interface{}{
			"merchantId":      callbackData.Data.MerchantId,
			"merchantOrderNo": callbackData.Data.MerchantOrderNo,
			"orderNo":         callbackData.Data.OrderNo,
			"amount":          callbackData.Data.Amount,
			"status":          callbackData.Data.Status,
			"currency":        callbackData.Data.Currency,
		}

		if callbackData.Data.ErrorMsg != "" {
			dataMap["errorMsg"] = callbackData.Data.ErrorMsg
		}

		if !pc.VerifyCallbackSign(dataMap, callbackData.Data.Sign) {
			fmt.Println("❌ 回调签名验证失败")
			http.Error(w, "签名验证失败", http.StatusUnauthorized)
			return
		}

		// 标记为已处理
		deduplicator.MarkProcessed(callbackData.Data.MerchantOrderNo, callbackData.Data.OrderNo)

		// 根据状态处理回调
		switch callbackData.Data.Status {
		case "SUCCESS", "COMPLETED":
			pc.processPaymentSuccess(callbackData)
		case "FAILED", "REJECTED":
			pc.processPaymentFailure(callbackData)
		default:
			fmt.Printf("📝 其他状态回调: %s\n", callbackData.Data.Status)
		}

		// 返回成功响应
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success":   true,
			"errorCode": "success",
			"message":   "回调处理成功",
		})
	})

	fmt.Println("🚀 提现回调服务器启动在 http://localhost:8081/payment/callback")
	fmt.Println("💡 使用 Ctrl+C 停止服务器")

	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Printf("❌ 服务器启动失败: %v\n", err)
	}
}

// 7. 余额查询
func (pc *PaymentClient) QueryBalance() {
	fmt.Println("=== 余额查询 ===")

	formData := url.Values{}
	formData.Set("merchantId", pc.MerchantId)

	// 生成签名
	signature := pc.GenerateFormSign(formData)
	formData.Set("sign", signature)

	// 发送请求
	resp, err := http.PostForm(pc.BaseURL+"/api/open/merchant/balance/query", formData)
	if err != nil {
		fmt.Printf("❌ 请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ 读取响应失败: %v\n", err)
		return
	}

	fmt.Printf("\n📡 响应状态: %s\n", resp.Status)
	fmt.Printf("📄 响应内容: %s\n", string(body))

	// 解析响应
	var response BalanceResponse
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Printf("❌ 解析响应失败: %v\n", err)
		return
	}

	if response.Success {
		fmt.Println("✅ 余额查询成功!")
		fmt.Printf("   可用余额: %d 分\n", response.Data.Balance)
		fmt.Printf("   待结算金额: %d 分\n", response.Data.UnsettledBalance)
		fmt.Printf("   冻结金额: %d 分\n", response.Data.FrozenAmount)
		fmt.Printf("   币种: %s\n", response.Data.Currency)
	} else {
		fmt.Printf("❌ 余额查询失败: %s\n", response.Message)
	}
}

// 8. 提现反查
func (pc *PaymentClient) ReversePayment() {
	fmt.Println("=== 提现反查 ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("请输入平台订单号: ")
	orderNo, _ := reader.ReadString('\n')
	orderNo = strings.TrimSpace(orderNo)
	if orderNo == "" {
		fmt.Println("❌ 平台订单号不能为空")
		return
	}

	formData := url.Values{}
	formData.Set("merchantId", pc.MerchantId)
	formData.Set("orderNo", orderNo)

	// 生成签名
	signature := pc.GenerateFormSign(formData)
	formData.Set("sign", signature)

	// 发送请求
	resp, err := http.PostForm(pc.BaseURL+"/api/open/merchant/payment/reverse", formData)
	if err != nil {
		fmt.Printf("❌ 请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ 读取响应失败: %v\n", err)
		return
	}

	fmt.Printf("\n📡 响应状态: %s\n", resp.Status)
	fmt.Printf("📄 响应内容: %s\n", string(body))

	// 解析响应
	var response ReversePaymentResponse
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Printf("❌ 解析响应失败: %v\n", err)
		return
	}

	if response.Success {
		fmt.Println("✅ 反查成功!")
		fmt.Printf("   商户订单号: %s\n", response.Data.MerchantOrderNo)
		fmt.Printf("   提现金额: %d 分\n", response.Data.Amount)
		fmt.Printf("   提现状态: %s\n", response.Data.Status)
	} else {
		fmt.Printf("❌ 反查失败: %s\n", response.Message)
	}
}

// 签名算法测试
func (pc *PaymentClient) TestSignature() {
	fmt.Println("=== 表单签名算法测试 ===")

	// 使用示例中的测试数据
	fmt.Println("📋 使用示例验证数据:")

	formData := url.Values{}
	formData.Set("apple", "red")
	formData.Set("banana", "yellow")
	formData.Set("orange", "orange color")
	formData.Set("weight", "123")

	// 临时设置密钥为示例密钥
	originalSecret := pc.SecretKey
	pc.SecretKey = "asecretkey"

	signature := pc.GenerateFormSign(formData)
	expectedSign := "5cbef8dddb54e753714857162eba1bed"

	fmt.Printf("期望签名: %s\n", expectedSign)
	fmt.Printf("实际签名: %s\n", signature)

	if signature == expectedSign {
		fmt.Println("✅ 签名验证成功!")
	} else {
		fmt.Println("❌ 签名验证失败!")
	}

	// 恢复原始密钥
	pc.SecretKey = originalSecret

	// 测试实际交易参数
	fmt.Println("\n📋 测试实际交易参数:")
	realFormData := url.Values{}
	realFormData.Set("merchantId", pc.MerchantId)
	realFormData.Set("merchantOrderNo", "TEST_ORDER_123")
	realFormData.Set("amount", "100")
	realFormData.Set("payType", "PIX_QRCODE")
	realFormData.Set("currency", "BRL")
	realFormData.Set("content", "测试订单")
	realFormData.Set("clientIp", "192.168.1.100")

	realSignature := pc.GenerateFormSign(realFormData)
	fmt.Printf("实际交易签名: %s\n", realSignature)
}

func InitPayment() *PaymentClient {
	return &PaymentClient{
		BaseURL:    "https://gateway.novavexis.com",    // 替换为实际的API地址
		MerchantId: "100100",                           // 替换为实际的商户ID
		SecretKey:  "RiFagbDcHXVFcXcLHkAqMdcqXPtZRdYK", // 替换为实际的密钥
	}
}
