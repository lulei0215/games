package payment

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	apiReq "github.com/flipped-aurora/gin-vue-admin/server/model/api/request"
)

// 代收接口请求结构体
type PayinCreateRequest struct {
	MerNo        string `json:"merNo"`                // 平台唯一标识，即商户号
	CurrencyCode string `json:"currencyCode"`         // 币种编码
	PayType      string `json:"payType"`              // 支付类型编码
	RandomNo     string `json:"randomNo"`             // 随机数
	OutTradeNo   string `json:"outTradeNo"`           // 订单号
	TotalAmount  string `json:"totalAmount"`          // 订单金额
	NotifyUrl    string `json:"notifyUrl"`            // 交易异步回调地址
	Sign         string `json:"sign"`                 // 加密字符串
	PayCardNo    string `json:"payCardNo"`            // 付款账号
	PayBankCode  string `json:"payBankCode"`          // 银行编码/付款方式
	PayName      string `json:"payName"`              // 付款人姓名
	PayEmail     string `json:"payEmail"`             // 邮箱
	PayPhone     string `json:"payPhone"`             // 付款人手机号
	PayViewUrl   string `json:"payViewUrl,omitempty"` // 支付后同步跳转的地址
}

// 代收接口响应结构体
type PayinCreateResponse struct {
	ResultCode  string `json:"resultCode"`            // 状态码
	StateInfo   string `json:"stateInfo"`             // 状态描述
	MerNo       string `json:"merNo,omitempty"`       // 平台唯一标识，即商户号
	OutTradeNo  string `json:"outTradeNo,omitempty"`  // 订单号
	TotalAmount string `json:"totalAmount,omitempty"` // 订单金额
	PayURL      string `json:"payURL,omitempty"`      // 支付地址
	PayOrderNo  string `json:"payOrderNo,omitempty"`  // 平台订单号
	PayParams   string `json:"payParams,omitempty"`   // 支付原生串
	Sign        string `json:"sign"`                  // 加密字符串
}

// 支付客户端
type PaymentClient2 struct {
	BaseURL    string
	MerchantId string
	SecretKey  string
}

// 代收接口 - 创建支付订单
func (pc *PaymentClient2) CreatePayin(r apiReq.CreateTradeData2, OutTradeNo string) (int, string, PayinCreateResponse) {
	fmt.Println("=== Payment Interface - Create Payment Order ===")

	// Build request parameters - Merchant's own system order number, please ensure uniqueness
	requestData := PayinCreateRequest{
		MerNo:        pc.MerchantId,
		CurrencyCode: "BRL",
		PayType:      "PIX",
		RandomNo:     OutTradeNo, // 14位随机数
		OutTradeNo:   OutTradeNo,
		TotalAmount:  r.TotalAmount,
		NotifyUrl:    "https://api.bzgame777.com/callback/trade2",
		PayCardNo:    r.PayCardNo,
		PayBankCode:  r.PayBankCode,
		PayName:      r.PayName,
		PayEmail:     r.PayEmail,
		PayPhone:     r.PayPhone,
		PayViewUrl:   "https://www.bzgame777.com",
	}

	// If there is a redirect URL, add it to the request
	if r.PayViewUrl != "" {
		requestData.PayViewUrl = r.PayViewUrl
	}

	// Generate signature
	signature := pc.GeneratePayinSign(requestData)
	requestData.Sign = signature

	if requestData.PayViewUrl != "" {
		fmt.Printf("  Redirect URL: %s\n", requestData.PayViewUrl)
	}
	fmt.Printf("  Signature: %s\n", requestData.Sign)

	// Convert to JSON
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return -1, "JSON serialization failed: " + err.Error(), PayinCreateResponse{}
	}

	// Send POST request
	resp, err := http.Post(pc.BaseURL+"/payin/create", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("❌ Request failed: %v\n", err)
		return -1, "Network request failed: " + err.Error(), PayinCreateResponse{}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ Failed to read response: %v\n", err)
		return -1, "Failed to read response: " + err.Error(), PayinCreateResponse{}
	}

	fmt.Printf("\n📡 Response status: %s\n", resp.Status)
	fmt.Printf("📄 Response content: %s\n", string(body))

	// Parse response
	var response PayinCreateResponse
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Printf("❌ Failed to parse response: %v\n", err)
		return -1, "Failed to parse response: " + err.Error(), PayinCreateResponse{}
	}

	// response := PayinCreateResponse{
	// 	ResultCode:  "0000",
	// 	StateInfo:   "订单创建成功",
	// 	MerNo:       OutTradeNo,
	// 	OutTradeNo:  OutTradeNo,
	// 	PayOrderNo:  "1234567890",
	// 	PayURL:      "https://www.baidu.com",
	// 	TotalAmount: "100",
	// 	PayParams:   "1234567890",
	// 	Sign:        "1234567890",
	// }
	// Check response status code
	if response.ResultCode == "0000" {
		return 1, "success", response
	} else {
		fmt.Printf("❌ Payment order creation failed: %s - %s\n", response.ResultCode, response.StateInfo)
		return -1, fmt.Sprintf("create error: %s - %s", response.ResultCode, response.StateInfo), response
	}
}

// 代收接口签名生成方法（使用JSON格式）
func (pc *PaymentClient2) GeneratePayinSign(request PayinCreateRequest) string {
	fmt.Println("\n🔐 === 代收接口签名生成过程 ===")

	// 1. 构建参数映射（排除sign字段）
	params := make(map[string]string)
	params["merNo"] = request.MerNo
	params["currencyCode"] = request.CurrencyCode
	params["payType"] = request.PayType
	params["randomNo"] = request.RandomNo
	params["outTradeNo"] = request.OutTradeNo
	params["totalAmount"] = request.TotalAmount
	params["notifyUrl"] = request.NotifyUrl
	params["payCardNo"] = request.PayCardNo
	params["payBankCode"] = request.PayBankCode
	params["payName"] = request.PayName
	params["payEmail"] = request.PayEmail
	params["payPhone"] = request.PayPhone

	// 如果有跳转地址，添加到参数中
	if request.PayViewUrl != "" {
		params["payViewUrl"] = request.PayViewUrl
	}

	// 2. 按key的ASCII码升序排序
	var keys []string
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// 3. 构建JSON字符串
	var jsonPairs []string
	fmt.Println("1️⃣ 排序后的参数:")
	for _, key := range keys {
		jsonPairs = append(jsonPairs, fmt.Sprintf("\"%s\":\"%s\"", key, params[key]))
		fmt.Printf("   %s: %s\n", key, params[key])
	}
	jsonString := "{" + strings.Join(jsonPairs, ",") + "}"
	fmt.Printf("\n2️⃣ JSON字符串:\n   %s\n", jsonString)

	// 4. 拼接密钥
	signString := jsonString + pc.SecretKey
	fmt.Printf("\n3️⃣ 拼接密钥:\n   %s\n", signString)

	// 5. MD5签名并转换为大写
	hash := md5.Sum([]byte(signString))
	signature := strings.ToUpper(hex.EncodeToString(hash[:]))
	fmt.Printf("\n4️⃣ MD5签名(大写):\n   %s\n", signature)

	return signature
}

// 代收回调处理
func (pc *PaymentClient2) StartPayinCallbackServer(r apiReq.TradeCallback2Request) bool {

	receivedSign := r.Sign

	params := make(map[string]interface{})
	params["merNo"] = r.MerNo
	params["currencyCode"] = r.CurrencyCode
	params["payType"] = r.PayType
	params["outTradeNo"] = r.OutTradeNo
	params["totalAmount"] = r.TotalAmount
	params["payOrderNo"] = r.PayOrderNo
	params["payState"] = r.PayState
	params["payDate"] = r.PayDate

	if !pc.VerifyPayinCallbackSign(params, receivedSign) {
		fmt.Println("❌ 代收回调签名验证失败")
		return false
	}
	return true

}

// 代收回调签名验证
func (pc *PaymentClient2) VerifyPayinCallbackSign(data map[string]interface{}, receivedSign string) bool {
	fmt.Println("\n🔍 === 代收回调签名验证过程 ===")

	// 按key的ASCII码升序排序
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 构建JSON字符串
	var jsonPairs []string
	fmt.Println("1️⃣ 排序后的参数:")
	for _, key := range keys {
		jsonPairs = append(jsonPairs, fmt.Sprintf("\"%s\":\"%s\"", key, data[key]))
		fmt.Printf("   %s: %s\n", key, data[key])
	}
	jsonString := "{" + strings.Join(jsonPairs, ",") + "}"
	fmt.Printf("\n2️⃣ JSON字符串:\n   %s\n", jsonString)

	// 拼接密钥
	signString := jsonString + pc.SecretKey
	fmt.Printf("\n3️⃣ 拼接密钥:\n   %s\n", signString)

	// MD5签名并转换为大写
	hash := md5.Sum([]byte(signString))
	calculatedSign := strings.ToUpper(hex.EncodeToString(hash[:]))
	fmt.Printf("\n4️⃣ MD5签名(大写):\n   %s\n", calculatedSign)
	fmt.Printf("接收签名: %s\n", receivedSign)

	return calculatedSign == receivedSign
}

// 代收订单查询请求结构体
type PayinQueryRequest struct {
	MerNo      string `json:"merNo"`      // 平台唯一标识，即商户号
	OutTradeNo string `json:"outTradeNo"` // 订单号
	Sign       string `json:"sign"`       // 加密字符串
}

// 代收订单查询响应结构体
type PayinQueryResponse struct {
	ResultCode  string `json:"resultCode"`            // 状态码
	StateInfo   string `json:"stateInfo"`             // 状态描述
	MerNo       string `json:"merNo,omitempty"`       // 平台唯一标识，即商户号
	OutTradeNo  string `json:"outTradeNo,omitempty"`  // 订单号
	PayOrderNo  string `json:"payOrderNo,omitempty"`  // 平台订单号
	PayState    string `json:"payState,omitempty"`    // 订单状态
	TotalAmount string `json:"totalAmount,omitempty"` // 订单金额
	Sign        string `json:"sign"`                  // 加密字符串
}

// 代收订单查询
func (pc *PaymentClient2) QueryPayin(outTradeNo string) (int, string, PayinQueryResponse) {
	fmt.Println("=== 代收订单查询 ===")

	// 构建请求参数
	requestData := PayinQueryRequest{
		MerNo:      pc.MerchantId,
		OutTradeNo: outTradeNo,
	}

	// 生成签名
	signature := pc.GeneratePayinQuerySign(requestData)
	requestData.Sign = signature

	fmt.Println("\n📝 请求参数:")
	fmt.Printf("  商户号: %s\n", requestData.MerNo)
	fmt.Printf("  订单号: %s\n", requestData.OutTradeNo)
	fmt.Printf("  签名: %s\n", requestData.Sign)

	// 转换为JSON
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Printf("❌ JSON序列化失败: %v\n", err)
		return -1, err.Error(), PayinQueryResponse{}
	}

	// 发送POST请求
	resp, err := http.Post(pc.BaseURL+"/query/payOrder", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("❌ 请求失败: %v\n", err)
		return -1, err.Error(), PayinQueryResponse{}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ 读取响应失败: %v\n", err)
		return -1, err.Error(), PayinQueryResponse{}
	}

	fmt.Printf("\n📡 响应状态: %s\n", resp.Status)
	fmt.Printf("📄 响应内容: %s\n", string(body))

	// 解析响应
	var response PayinQueryResponse
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Printf("❌ 解析响应失败: %v\n", err)
		return -1, err.Error(), PayinQueryResponse{}
	}

	// 处理响应结果
	switch response.ResultCode {
	case "0000":
		fmt.Println("✅ 查询成功!")
		fmt.Printf("   商户号: %s\n", response.MerNo)
		fmt.Printf("   订单号: %s\n", response.OutTradeNo)
		fmt.Printf("   平台订单号: %s\n", response.PayOrderNo)
		fmt.Printf("   订单金额: %s 元\n", response.TotalAmount)
		fmt.Printf("   订单状态: %s (%s)\n", response.PayState, pc.getPayStateInfo(response.PayState))
		return 0, "查询成功", response
	default:
		fmt.Printf("❌ 查询失败: %s - %s\n", response.ResultCode, response.StateInfo)
		return -1, response.StateInfo, response
	}
}

// 生成代收订单查询签名
func (pc *PaymentClient2) GeneratePayinQuerySign(request PayinQueryRequest) string {
	fmt.Println("\n🔐 === 代收订单查询签名生成过程 ===")

	// 1. 构建参数映射（排除sign字段）
	params := make(map[string]string)
	params["merNo"] = request.MerNo
	params["outTradeNo"] = request.OutTradeNo

	// 2. 按key的ASCII码升序排序
	var keys []string
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// 3. 构建JSON字符串
	var jsonPairs []string
	fmt.Println("1️⃣ 排序后的参数:")
	for _, key := range keys {
		jsonPairs = append(jsonPairs, fmt.Sprintf("\"%s\":\"%s\"", key, params[key]))
		fmt.Printf("   %s: %s\n", key, params[key])
	}
	jsonString := "{" + strings.Join(jsonPairs, ",") + "}"
	fmt.Printf("\n2️⃣ JSON字符串:\n   %s\n", jsonString)

	// 4. 拼接密钥
	signString := jsonString + pc.SecretKey
	fmt.Printf("\n3️⃣ 拼接密钥:\n   %s\n", signString)

	// 5. MD5签名并转换为大写
	hash := md5.Sum([]byte(signString))
	signature := strings.ToUpper(hex.EncodeToString(hash[:]))
	fmt.Printf("\n4️⃣ MD5签名(大写):\n   %s\n", signature)

	return signature
}

// 获取订单状态信息
func (pc *PaymentClient2) getPayStateInfo(payState string) string {
	switch payState {
	case "99":
		return "待支付"
	case "00":
		return "支付成功"
	case "01":
		return "支付失败"
	case "04":
		return "未知错误"
	default:
		return "未知状态"
	}
}

// 测试代收订单查询签名
func (pc *PaymentClient2) TestPayinQuerySignature() {
	fmt.Println("=== 测试代收订单查询签名 ===")

	// 模拟查询请求
	testRequest := PayinQueryRequest{
		MerNo:      "100100",
		OutTradeNo: "TEST_ORDER_123",
	}

	// 生成签名
	signature := pc.GeneratePayinQuerySign(testRequest)
	fmt.Printf("生成的签名: %s\n", signature)

	// 验证签名
	dataMap := map[string]interface{}{
		"merNo":      testRequest.MerNo,
		"outTradeNo": testRequest.OutTradeNo,
	}

	isValid := pc.VerifyPayinQuerySign(dataMap, signature)
	fmt.Printf("签名验证结果: %t\n", isValid)
}

// 验证代收订单查询签名
func (pc *PaymentClient2) VerifyPayinQuerySign(data map[string]interface{}, receivedSign string) bool {
	fmt.Println("\n🔍 === 代收订单查询签名验证过程 ===")

	// 按key的ASCII码升序排序
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 构建JSON字符串
	var jsonPairs []string
	fmt.Println("1️⃣ 排序后的参数:")
	for _, key := range keys {
		jsonPairs = append(jsonPairs, fmt.Sprintf("\"%s\":\"%s\"", key, data[key]))
		fmt.Printf("   %s: %s\n", key, data[key])
	}
	jsonString := "{" + strings.Join(jsonPairs, ",") + "}"
	fmt.Printf("\n2️⃣ JSON字符串:\n   %s\n", jsonString)

	// 拼接密钥
	signString := jsonString + pc.SecretKey
	fmt.Printf("\n3️⃣ 拼接密钥:\n   %s\n", signString)

	// MD5签名并转换为大写
	hash := md5.Sum([]byte(signString))
	calculatedSign := strings.ToUpper(hex.EncodeToString(hash[:]))
	fmt.Printf("\n4️⃣ MD5签名(大写):\n   %s\n", calculatedSign)
	fmt.Printf("接收签名: %s\n", receivedSign)

	return calculatedSign == receivedSign
}

// 代付申请请求结构体
type CashoutCreateRequest struct {
	MerNo         string `json:"merNo"`                   // 平台唯一标识，即商户号
	RandomNo      string `json:"randomNo"`                // 随机数
	CurrencyCode  string `json:"currencyCode"`            // 币种编码
	TotalAmount   string `json:"totalAmount"`             // 订单金额
	OutTradeNo    string `json:"outTradeNo"`              // 订单号
	BankCode      string `json:"bankCode"`                // 银行编码
	BankAcctName  string `json:"bankAcctName,omitempty"`  // 收款人姓名（全名）
	BankFirstName string `json:"bankFirstName,omitempty"` // 收款人FirstName
	BankLastName  string `json:"bankLastName,omitempty"`  // 收款人LastName
	BankAcctNo    string `json:"bankAcctNo"`              // 收款人账号
	AccPhone      string `json:"accPhone"`                // 收款人手机号
	AccEmail      string `json:"accEmail,omitempty"`      // 邮箱
	NotifyUrl     string `json:"notifyUrl"`               // 交易异步回调地址
	IdentityNo    string `json:"identityNo"`              // 证件号码/税号
	IdentityType  string `json:"identityType"`            // 证件类型/收款类型编码
	ReqTimesTamp  string `json:"reqTimesTamp"`            // 证件类型/收款类型编码
	Sign          string `json:"sign"`                    // 加密字符串
}

// 代付申请响应结构体
type CashoutCreateResponse struct {
	ResultCode   string `json:"resultCode"`             // 状态码
	ResultMsg    string `json:"resultMsg"`              // 状态描述
	MerNo        string `json:"merNo,omitempty"`        // 平台唯一标识，即商户号
	OutTradeNo   string `json:"outTradeNo,omitempty"`   // 订单号
	RemitOrderNo string `json:"remitOrderNo,omitempty"` // 平台订单号
	TotalAmount  string `json:"totalAmount,omitempty"`  // 订单金额
	Sign         string `json:"sign,omitempty"`         // 加密字符串
}

// 代付申请
func (pc *PaymentClient2) CreateCashout(formData CashoutCreateRequest) (int, string, CashoutCreateResponse) {
	fmt.Println("=== 代付申请 ===")

	// 构建请求参数
	requestData := CashoutCreateRequest{
		MerNo:         pc.MerchantId,
		RandomNo:      formData.RandomNo, // 14位随机数
		CurrencyCode:  formData.CurrencyCode,
		TotalAmount:   formData.TotalAmount,
		OutTradeNo:    formData.OutTradeNo,
		BankCode:      formData.BankCode,
		BankAcctName:  formData.BankAcctName,
		BankAcctNo:    formData.BankAcctNo,
		AccPhone:      formData.AccPhone,
		NotifyUrl:     formData.NotifyUrl,
		IdentityNo:    formData.IdentityNo,
		IdentityType:  formData.IdentityType,
		BankFirstName: formData.BankFirstName,
		BankLastName:  formData.BankLastName,
		ReqTimesTamp:  strconv.FormatInt(time.Now().UTC().UnixMilli(), 10),
	}

	// 如果有邮箱，添加到请求中
	if formData.AccEmail != "" {
		requestData.AccEmail = formData.AccEmail
	}

	// 生成签名
	signature := pc.GenerateCashoutSign(requestData)
	requestData.Sign = signature

	fmt.Println("\n📝 请求参数:")
	fmt.Printf("  商户号: %s\n", requestData.MerNo)
	fmt.Printf("  随机数: %s\n", requestData.RandomNo)
	fmt.Printf("  币种: %s\n", requestData.CurrencyCode)
	fmt.Printf("  金额: %s\n", requestData.TotalAmount)
	fmt.Printf("  订单号: %s\n", requestData.OutTradeNo)
	fmt.Printf("  银行编码: %s\n", requestData.BankCode)
	fmt.Printf("  收款人姓名: %s\n", requestData.BankAcctName)
	fmt.Printf("  收款人账号: %s\n", requestData.BankAcctNo)
	fmt.Printf("  收款人手机号: %s\n", requestData.AccPhone)
	if requestData.AccEmail != "" {
		fmt.Printf("  邮箱: %s\n", requestData.AccEmail)
	}
	fmt.Printf("  回调地址: %s\n", requestData.NotifyUrl)
	fmt.Printf("  证件号码: %s\n", requestData.IdentityNo)
	fmt.Printf("  证件类型: %s\n", requestData.IdentityType)
	fmt.Printf("  签名: %s\n", requestData.Sign)

	// 转换为JSON
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Printf("❌ JSON序列化失败: %v\n", err)
		return -1, "JSON序列化失败: " + err.Error(), CashoutCreateResponse{}
	}
	fmt.Println("jsonData", string(jsonData))
	// 发送POST请求
	resp, err := http.Post(pc.BaseURL+"/cashOut/create", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("❌ 请求失败: %v\n", err)
		return -1, "网络请求失败: " + err.Error(), CashoutCreateResponse{}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ 读取响应失败: %v\n", err)
		return -1, "读取响应失败: " + err.Error(), CashoutCreateResponse{}
	}

	fmt.Printf("\n📡 响应状态: %s\n", resp.Status)
	fmt.Printf("📄 响应内容: %s\n", string(body))

	// 解析响应
	var response CashoutCreateResponse
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Printf("❌ 解析响应失败: %v\n", err)
		return -1, "解析响应失败: " + err.Error(), CashoutCreateResponse{}
	}

	// 处理响应结果
	switch response.ResultCode {
	case "0000":
		fmt.Println("✅ 代付申请提交成功!")
		fmt.Printf("   平台订单号: %s\n", response.RemitOrderNo)
		fmt.Printf("   订单金额: %s 元\n", response.TotalAmount)
		fmt.Printf("   状态描述: %s\n", response.ResultMsg)
		return 0, "Payment application submitted successfully", response

	case "E0001":
		fmt.Println("⚠️ 代付驳回")
		fmt.Printf("   状态描述: %s\n", response.ResultMsg)
		fmt.Println("   注意：已提交到银行处理，需等待处理结果")
		return 1, "Payment rejected", response

	case "E0002":
		fmt.Println("⚠️ 银行网络波动")
		fmt.Printf("   状态描述: %s\n", response.ResultMsg)
		fmt.Println("   注意：已提交到银行处理，需等待处理结果")
		return 1, "Bank network fluctuation", response

	case "E0003":
		fmt.Println("⚠️ 银行验证波动")
		fmt.Printf("   状态描述: %s\n", response.ResultMsg)
		fmt.Println("   注意：已提交到银行处理，需等待处理结果")
		return 1, "Bank verification fluctuation", response

	case "9999":
		fmt.Println("❌ 参数校验有误")
		fmt.Printf("   状态描述: %s\n", response.ResultMsg)
		fmt.Println("   可以直接退款")
		return -1, "Parameter validation error", response

	case "99":
		fmt.Println("❌ 业务校验失败")
		fmt.Printf("   状态描述: %s\n", response.ResultMsg)
		fmt.Println("   可以直接退款")
		return -1, "Business validation failed", response

	default:
		fmt.Printf("❌ 代付申请失败: %s - %s\n", response.ResultCode, response.ResultMsg)
		return -1, fmt.Sprintf("Payment application failed: %s - %s", response.ResultCode, response.ResultMsg), response
	}
}

// 验证代付创建响应签名
func (pc *PaymentClient2) VerifyCashoutCreateResponseSign(data map[string]interface{}, receivedSign string) bool {
	fmt.Println("\n🔍 === 代付创建响应签名验证过程 ===")

	// 按key的ASCII码升序排序
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 构建JSON字符串
	var jsonPairs []string
	fmt.Println("1️⃣ 排序后的参数:")
	for _, key := range keys {
		jsonPairs = append(jsonPairs, fmt.Sprintf("\"%s\":\"%s\"", key, data[key]))
		fmt.Printf("   %s: %s\n", key, data[key])
	}
	jsonString := "{" + strings.Join(jsonPairs, ",") + "}"
	fmt.Printf("\n2️⃣ JSON字符串:\n   %s\n", jsonString)

	// 拼接密钥
	signString := jsonString + pc.SecretKey
	fmt.Printf("\n3️⃣ 拼接密钥:\n   %s\n", signString)

	// MD5签名并转换为大写
	hash := md5.Sum([]byte(signString))
	calculatedSign := strings.ToUpper(hex.EncodeToString(hash[:]))
	fmt.Printf("\n4️⃣ MD5签名(大写):\n   %s\n", calculatedSign)
	fmt.Printf("接收签名: %s\n", receivedSign)

	return calculatedSign == receivedSign
}

// 生成代付申请签名
func (pc *PaymentClient2) GenerateCashoutSign(request CashoutCreateRequest) string {
	fmt.Println("\n🔐 === 代付申请签名生成过程 ===")

	// 1. 构建参数映射（排除sign字段）
	params := make(map[string]string)
	params["merNo"] = request.MerNo
	params["randomNo"] = request.RandomNo
	params["currencyCode"] = request.CurrencyCode
	params["totalAmount"] = request.TotalAmount
	params["outTradeNo"] = request.OutTradeNo
	params["bankCode"] = request.BankCode
	params["bankAcctName"] = request.BankAcctName
	params["bankAcctNo"] = request.BankAcctNo
	params["accPhone"] = request.AccPhone
	params["notifyUrl"] = request.NotifyUrl
	params["identityNo"] = request.IdentityNo
	params["identityType"] = request.IdentityType
	params["bankFirstName"] = request.BankFirstName
	params["bankLastName"] = request.BankLastName
	params["reqTimesTamp"] = request.ReqTimesTamp

	// 如果有邮箱，添加到参数中
	if request.AccEmail != "" {
		params["accEmail"] = request.AccEmail
	}

	// 1. 过滤空值参数
	filtered := make(map[string]string)
	for k, v := range params {
		if v != "" {
			filtered[k] = v
		}
	}

	// 2. 按ASCII码升序排序
	keys := make([]string, 0, len(filtered))
	for k := range filtered {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 3. 组装JSON字符串
	var jsonPairs []string
	for _, k := range keys {
		jsonPairs = append(jsonPairs, fmt.Sprintf("\"%s\":\"%s\"", k, filtered[k]))
	}
	jsonString := "{" + strings.Join(jsonPairs, ",") + "}"

	// 4. 拼接密钥
	signString := jsonString + pc.SecretKey

	// 5. MD5加密并转大写
	md5Sum := md5.Sum([]byte(signString))
	signature := strings.ToUpper(hex.EncodeToString(md5Sum[:]))

	return signature
}

// 测试代付申请签名
func (pc *PaymentClient2) TestCashoutSignature() {
	fmt.Println("=== 测试代付申请签名 ===")

	// 模拟代付请求
	testRequest := CashoutCreateRequest{
		MerNo:        "100100",
		RandomNo:     "12345678901234",
		CurrencyCode: "BRL",
		TotalAmount:  "10.00",
		OutTradeNo:   "TEST_CASHOUT_123",
		BankCode:     "PIX",
		BankAcctName: "Test User",
		BankAcctNo:   "12345678901",
		AccPhone:     "+551234567890",
		NotifyUrl:    "https://your-domain.com/cashout/callback",
		IdentityNo:   "12345678901",
		IdentityType: "CPF",
	}

	// 生成签名
	signature := pc.GenerateCashoutSign(testRequest)
	fmt.Printf("生成的签名: %s\n", signature)

	// 验证签名
	dataMap := map[string]interface{}{
		"merNo":        testRequest.MerNo,
		"randomNo":     testRequest.RandomNo,
		"currencyCode": testRequest.CurrencyCode,
		"totalAmount":  testRequest.TotalAmount,
		"outTradeNo":   testRequest.OutTradeNo,
		"bankCode":     testRequest.BankCode,
		"bankAcctName": testRequest.BankAcctName,
		"bankAcctNo":   testRequest.BankAcctNo,
		"accPhone":     testRequest.AccPhone,
		"notifyUrl":    testRequest.NotifyUrl,
		"identityNo":   testRequest.IdentityNo,
		"identityType": testRequest.IdentityType,
	}

	isValid := pc.VerifyCashoutSign(dataMap, signature)
	fmt.Printf("签名验证结果: %t\n", isValid)
}

// 验证代付申请签名
func (pc *PaymentClient2) VerifyCashoutSign(data map[string]interface{}, receivedSign string) bool {
	fmt.Println("\n🔍 === 代付申请签名验证过程 ===")

	// 按key的ASCII码升序排序
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 构建JSON字符串
	var jsonPairs []string
	fmt.Println("1️⃣ 排序后的参数:")
	for _, key := range keys {
		jsonPairs = append(jsonPairs, fmt.Sprintf("\"%s\":\"%s\"", key, data[key]))
		fmt.Printf("   %s: %s\n", key, data[key])
	}
	jsonString := "{" + strings.Join(jsonPairs, ",") + "}"
	fmt.Printf("\n2️⃣ JSON字符串:\n   %s\n", jsonString)

	// 拼接密钥
	signString := jsonString + pc.SecretKey
	fmt.Printf("\n3️⃣ 拼接密钥:\n   %s\n", signString)

	// MD5签名并转换为大写
	hash := md5.Sum([]byte(signString))
	calculatedSign := strings.ToUpper(hex.EncodeToString(hash[:]))
	fmt.Printf("\n4️⃣ MD5签名(大写):\n   %s\n", calculatedSign)
	fmt.Printf("接收签名: %s\n", receivedSign)

	return calculatedSign == receivedSign
}

// 代付订单查询请求结构体
type CashoutQueryRequest struct {
	MerNo      string `json:"merNo"`      // 平台唯一标识，即商户号
	OutTradeNo string `json:"outTradeNo"` // 订单号
	Sign       string `json:"sign"`       // 加密字符串
}

// 代付订单查询响应结构体
type CashoutQueryResponse struct {
	ResultCode   string `json:"resultCode"`             // 状态码
	ResultMsg    string `json:"resultMsg"`              // 状态描述
	MerNo        string `json:"merNo,omitempty"`        // 平台唯一标识，即商户号
	OutTradeNo   string `json:"outTradeNo,omitempty"`   // 订单号
	RemitOrderNo string `json:"remitOrderNo,omitempty"` // 平台订单号
	RemitState   string `json:"remitState,omitempty"`   // 订单状态
	TotalAmount  string `json:"totalAmount,omitempty"`  // 订单金额
	Sign         string `json:"sign"`                   // 加密字符串
}

// 代付订单查询
func (pc *PaymentClient2) QueryCashout(outTradeNo string) (int, string, CashoutQueryResponse) {
	fmt.Println("=== 代付订单查询 ===")

	// 构建请求参数
	requestData := CashoutQueryRequest{
		MerNo:      pc.MerchantId,
		OutTradeNo: outTradeNo,
	}

	// 生成签名
	signature := pc.GenerateCashoutQuerySign(requestData)
	requestData.Sign = signature

	fmt.Println("\n📝 请求参数:")
	fmt.Printf("  商户号: %s\n", requestData.MerNo)
	fmt.Printf("  订单号: %s\n", requestData.OutTradeNo)
	fmt.Printf("  签名: %s\n", requestData.Sign)

	// 转换为JSON
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Printf("❌ JSON序列化失败: %v\n", err)
		return -1, "JSON序列化失败: " + err.Error(), CashoutQueryResponse{}
	}

	// 发送POST请求
	resp, err := http.Post(pc.BaseURL+"/query/remitOrder", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("❌ 请求失败: %v\n", err)
		return -1, "网络请求失败: " + err.Error(), CashoutQueryResponse{}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ 读取响应失败: %v\n", err)
		return -1, "读取响应失败: " + err.Error(), CashoutQueryResponse{}
	}

	fmt.Printf("\n📡 响应状态: %s\n", resp.Status)
	fmt.Printf("📄 响应内容: %s\n", string(body))

	// 解析响应
	var response CashoutQueryResponse
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Printf("❌ 解析响应失败: %v\n", err)
		return -1, "解析响应失败: " + err.Error(), CashoutQueryResponse{}
	}

	// 处理响应结果
	switch response.ResultCode {
	case "0000":
		fmt.Println("✅ 查询成功!")
		fmt.Printf("   商户号: %s\n", response.MerNo)
		fmt.Printf("   订单号: %s\n", response.OutTradeNo)
		fmt.Printf("   平台订单号: %s\n", response.RemitOrderNo)
		fmt.Printf("   订单金额: %s 元\n", response.TotalAmount)

		// 解析订单状态
		stateInfo := pc.getCashoutStateInfo(response.RemitState)
		fmt.Printf("   订单状态: %s (%s)\n", response.RemitState, stateInfo)
		return 0, "查询成功", response

	case "1001":
		fmt.Println("⚠️ 没有查询到对应订单")
		fmt.Printf("   状态描述: %s\n", response.ResultMsg)
		return 1, "没有查询到对应订单", response

	default:
		fmt.Printf("❌ 查询失败: %s - %s\n", response.ResultCode, response.ResultMsg)
		return -1, response.ResultMsg, response
	}
}

// 生成代付订单查询签名
func (pc *PaymentClient2) GenerateCashoutQuerySign(request CashoutQueryRequest) string {
	fmt.Println("\n🔐 === 代付订单查询签名生成过程 ===")

	// 1. 构建参数映射（排除sign字段）
	params := make(map[string]string)
	params["merNo"] = request.MerNo
	params["outTradeNo"] = request.OutTradeNo

	// 2. 按key的ASCII码升序排序
	var keys []string
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// 3. 构建JSON字符串
	var jsonPairs []string
	fmt.Println("1️⃣ 排序后的参数:")
	for _, key := range keys {
		jsonPairs = append(jsonPairs, fmt.Sprintf("\"%s\":\"%s\"", key, params[key]))
		fmt.Printf("   %s: %s\n", key, params[key])
	}
	jsonString := "{" + strings.Join(jsonPairs, ",") + "}"
	fmt.Printf("\n2️⃣ JSON字符串:\n   %s\n", jsonString)

	// 4. 拼接密钥
	signString := jsonString + pc.SecretKey
	fmt.Printf("\n3️⃣ 拼接密钥:\n   %s\n", signString)

	// 5. MD5签名并转换为大写
	hash := md5.Sum([]byte(signString))
	signature := strings.ToUpper(hex.EncodeToString(hash[:]))
	fmt.Printf("\n4️⃣ MD5签名(大写):\n   %s\n", signature)

	return signature
}

// 获取代付订单状态信息
func (pc *PaymentClient2) getCashoutStateInfo(remitState string) string {
	switch remitState {
	case "00":
		return "代付成功（出款成功）"
	case "01":
		return "代付提交成功（出款中）"
	case "02":
		return "代付失败（出款失败）- 出款情况待确认"
	case "03":
		return "代付异常（出款情况待确认）"
	case "04":
		return "代付异常（出款情况待确认）"
	case "05":
		return "未知状态（出款情况待确认）"
	case "06":
		return "待审核（未扣款）"
	case "07":
		return "卡受限（未扣款）- 可直接退款"
	case "08":
		return "商户号受限（未扣款）- 可直接退款"
	case "11":
		return "出款不明确（出款情况待确认）"
	case "12":
		return "代付驳回（出款失败）"
	case "13":
		return "代付取消（未扣款）- 可直接退款"
	case "50":
		return "网络异常（出款情况待确认）"
	case "1000":
		return "代付失败，并退款"
	default:
		return "未知状态"
	}
}

// 测试代付订单查询签名
func (pc *PaymentClient2) TestCashoutQuerySignature() {
	fmt.Println("=== 测试代付订单查询签名 ===")

	// 模拟查询请求
	testRequest := CashoutQueryRequest{
		MerNo:      "100100",
		OutTradeNo: "TEST_CASHOUT_123",
	}

	// 生成签名
	signature := pc.GenerateCashoutQuerySign(testRequest)
	fmt.Printf("生成的签名: %s\n", signature)

	// 验证签名
	dataMap := map[string]interface{}{
		"merNo":      testRequest.MerNo,
		"outTradeNo": testRequest.OutTradeNo,
	}

	isValid := pc.VerifyCashoutQuerySign(dataMap, signature)
	fmt.Printf("签名验证结果: %t\n", isValid)
}

// 验证代付订单查询签名
func (pc *PaymentClient2) VerifyCashoutQuerySign(data map[string]interface{}, receivedSign string) bool {
	fmt.Println("\n🔍 === 代付订单查询签名验证过程 ===")

	// 按key的ASCII码升序排序
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 构建JSON字符串
	var jsonPairs []string
	fmt.Println("1️⃣ 排序后的参数:")
	for _, key := range keys {
		jsonPairs = append(jsonPairs, fmt.Sprintf("\"%s\":\"%s\"", key, data[key]))
		fmt.Printf("   %s: %s\n", key, data[key])
	}
	jsonString := "{" + strings.Join(jsonPairs, ",") + "}"
	fmt.Printf("\n2️⃣ JSON字符串:\n   %s\n", jsonString)

	// 拼接密钥
	signString := jsonString + pc.SecretKey
	fmt.Printf("\n3️⃣ 拼接密钥:\n   %s\n", signString)

	// MD5签名并转换为大写
	hash := md5.Sum([]byte(signString))
	calculatedSign := strings.ToUpper(hex.EncodeToString(hash[:]))
	fmt.Printf("\n4️⃣ MD5签名(大写):\n   %s\n", calculatedSign)
	fmt.Printf("接收签名: %s\n", receivedSign)

	return calculatedSign == receivedSign
}

// 代付状态分类处理
func (pc *PaymentClient2) ProcessCashoutState(remitState string) {
	fmt.Printf("处理代付状态: %s\n", remitState)

	switch remitState {
	case "00":
		fmt.Println("✅ 代付成功 - 可以确认用户到账")

	case "01":
		fmt.Println("⏳ 代付处理中 - 需要等待最终结果")

	case "02", "03", "04", "05", "11", "12", "50":
		fmt.Println("⚠️ 出款情况待确认 - 需要与客服确认最终状态")

	case "06":
		fmt.Println("⏳ 待审核 - 等待平台审核")

	case "07", "08", "13":
		fmt.Println(" 可直接退款 - 未扣款，可以安全退款")

	case "1000":
		fmt.Println("✅ 代付失败并已退款 - 用户已收到退款")

	default:
		fmt.Println("❓ 未知状态 - 需要进一步确认")
	}
}

// 代付结果异步回调请求结构体
type CashoutCallbackRequest struct {
	MerNo        string `json:"merNo"`        // 平台唯一标识，即商户号
	CurrencyCode string `json:"currencyCode"` // 币种编码
	OutTradeNo   string `json:"outTradeNo"`   // 订单号
	TotalAmount  string `json:"totalAmount"`  // 订单金额
	RemitOrderNo string `json:"remitOrderNo"` // 平台订单号
	RemitState   string `json:"remitState"`   // 代付状态
	RemitDate    string `json:"remitDate"`    // 代付完成时间
	OrderMessage string `json:"orderMessage"` // 代付详情
	Sign         string `json:"sign"`         // 加密字符串
}

// 代付结果异步回调处理
func (pc *PaymentClient2) StartCashoutCallbackServer(r apiReq.PaymentCallback2FormRequest) bool {
	fmt.Println("=== 代付结果异步回调处理 ===")
	fmt.Printf("商户号: %s\n", r.MerNo)
	fmt.Printf("币种: %s\n", r.CurrencyCode)
	fmt.Printf("订单号: %s\n", r.OutTradeNo)
	fmt.Printf("订单金额: %s\n", r.TotalAmount)
	fmt.Printf("平台订单号: %s\n", r.RemitOrderNo)
	fmt.Printf("代付状态: %s\n", r.RemitState)
	fmt.Printf("代付完成时间: %s\n", r.RemitDate)
	fmt.Printf("代付详情: %s\n", r.OrderMessage)
	fmt.Printf("签名: %s\n", r.Sign)

	receivedSign := r.Sign

	params := make(map[string]interface{})
	params["merNo"] = r.MerNo
	params["currencyCode"] = r.CurrencyCode
	params["outTradeNo"] = r.OutTradeNo
	params["totalAmount"] = r.TotalAmount
	params["remitOrderNo"] = r.RemitOrderNo
	params["remitState"] = r.RemitState
	params["remitDate"] = r.RemitDate
	params["orderMessage"] = r.OrderMessage

	if !pc.VerifyCashoutCallbackSign(params, receivedSign) {
		fmt.Println("❌ 代付回调签名验证失败")
		return false
	}

	fmt.Println("✅ 代付回调签名验证成功")
	return true
}

// 代付回调签名验证
func (pc *PaymentClient2) VerifyCashoutCallbackSign(data map[string]interface{}, receivedSign string) bool {
	fmt.Println("\n🔍 === 代付回调签名验证过程 ===")

	// 按key的ASCII码升序排序
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 构建JSON字符串
	var jsonPairs []string
	fmt.Println("1️⃣ 排序后的参数:")
	for _, key := range keys {
		jsonPairs = append(jsonPairs, fmt.Sprintf("\"%s\":\"%s\"", key, data[key]))
		fmt.Printf("   %s: %s\n", key, data[key])
	}
	jsonString := "{" + strings.Join(jsonPairs, ",") + "}"
	fmt.Printf("\n2️⃣ JSON字符串:\n   %s\n", jsonString)

	// 拼接密钥
	signString := jsonString + pc.SecretKey
	fmt.Printf("\n3️⃣ 拼接密钥:\n   %s\n", signString)

	// MD5签名并转换为大写
	hash := md5.Sum([]byte(signString))
	calculatedSign := strings.ToUpper(hex.EncodeToString(hash[:]))
	fmt.Printf("\n4️⃣ MD5签名(大写):\n   %s\n", calculatedSign)
	fmt.Printf("接收签名: %s\n", receivedSign)

	return calculatedSign == receivedSign
}

// 获取代付状态信息
func (pc *PaymentClient2) getCashoutCallbackStateInfo(remitState string) string {
	switch remitState {
	case "00":
		return "出款成功"
	case "02":
		return "出款失败"
	case "11":
		return "出款不明确（出款情况待确认）"
	case "12":
		return "代付驳回（出款失败）"
	case "13":
		return "代付取消（未扣款）- 可直接退款"
	case "1000":
		return "代付失败，并退款（出款失败，金额退回余额）"
	default:
		return "未知状态"
	}
}

// 处理代付回调状态
func (pc *PaymentClient2) ProcessCashoutCallbackState(remitState string) {
	fmt.Printf("处理代付回调状态: %s\n", remitState)

	switch remitState {
	case "00":
		fmt.Println("✅ 出款成功 - 可以确认用户到账")

	case "02", "12":
		fmt.Println("❌ 出款失败 - 需要处理失败情况")

	case "11":
		fmt.Println("⚠️ 出款不明确 - 需要与客服确认最终状态")

	case "13":
		fmt.Println("✅ 代付取消 - 未扣款，可以安全退款")

	case "1000":
		fmt.Println("✅ 代付失败并已退款 - 用户已收到退款")

	default:
		fmt.Println("❓ 未知状态 - 需要进一步确认")
	}
}
func InitPayment2() *PaymentClient2 {
	return &PaymentClient2{
		BaseURL:    "https://api.donepay.cc",           // 替换为实际的API地址
		MerchantId: "DO250715070648667",                // 替换为实际的商户ID
		SecretKey:  "4DA9D4EC0EA34D8A28FA947760A49E5E", // 替换为实际的密钥
	}
}
