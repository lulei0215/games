package api

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/api"
	apiReq "github.com/flipped-aurora/gin-vue-admin/server/model/api/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"go.uber.org/zap"
)

type MonitorTransfersApi struct{}

// CreateMonitorTransfers monitorTransfers
// @Tags MonitorTransfers
// @Summary monitorTransfers
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body api.MonitorTransfers true "monitorTransfers"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /monitorTransfers/createMonitorTransfers [post]
func (monitorTransfersApi *MonitorTransfersApi) CreateMonitorTransfers(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	var monitorTransfers api.MonitorTransfers
	err := c.ShouldBindJSON(&monitorTransfers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = monitorTransfersService.CreateMonitorTransfers(ctx, &monitorTransfers)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// DeleteMonitorTransfers monitorTransfers
// @Tags MonitorTransfers
// @Summary monitorTransfers
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body api.MonitorTransfers true "monitorTransfers"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /monitorTransfers/deleteMonitorTransfers [delete]
func (monitorTransfersApi *MonitorTransfersApi) DeleteMonitorTransfers(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	id := c.Query("id")
	err := monitorTransfersService.DeleteMonitorTransfers(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// DeleteMonitorTransfersByIds monitorTransfers
// @Tags MonitorTransfers
// @Summary monitorTransfers
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} ""
// @Router /monitorTransfers/deleteMonitorTransfersByIds [delete]
func (monitorTransfersApi *MonitorTransfersApi) DeleteMonitorTransfersByIds(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := monitorTransfersService.DeleteMonitorTransfersByIds(ctx, ids)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// UpdateMonitorTransfers monitorTransfers
// @Tags MonitorTransfers
// @Summary monitorTransfers
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body api.MonitorTransfers true "monitorTransfers"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /monitorTransfers/updateMonitorTransfers [put]
func (monitorTransfersApi *MonitorTransfersApi) UpdateMonitorTransfers(c *gin.Context) {
	// ctxcontext
	ctx := c.Request.Context()

	var monitorTransfers api.MonitorTransfers
	err := c.ShouldBindJSON(&monitorTransfers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = monitorTransfersService.UpdateMonitorTransfers(ctx, monitorTransfers)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// FindMonitorTransfers idmonitorTransfers
// @Tags MonitorTransfers
// @Summary idmonitorTransfers
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "idmonitorTransfers"
// @Success 200 {object} response.Response{data=api.MonitorTransfers,msg=string} ""
// @Router /monitorTransfers/findMonitorTransfers [get]
func (monitorTransfersApi *MonitorTransfersApi) FindMonitorTransfers(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	id := c.Query("id")
	remonitorTransfers, err := monitorTransfersService.GetMonitorTransfers(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithData(remonitorTransfers, c)
}

// GetMonitorTransfersList monitorTransfers
// @Tags MonitorTransfers
// @Summary monitorTransfers
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query apiReq.MonitorTransfersSearch true "monitorTransfers"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} ""
// @Router /monitorTransfers/getMonitorTransfersList [get]
func (monitorTransfersApi *MonitorTransfersApi) GetMonitorTransfersList(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	var pageInfo apiReq.MonitorTransfersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := monitorTransfersService.GetMonitorTransfersInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "", c)
}

// GetMonitorTransfersPublic monitorTransfers
// @Tags MonitorTransfers
// @Summary monitorTransfers
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} ""
// @Router /monitorTransfers/getMonitorTransfersPublic [get]
func (monitorTransfersApi *MonitorTransfersApi) GetMonitorTransfersPublic(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	//
	// ，C，
	monitorTransfersService.GetMonitorTransfersPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "monitorTransfers",
	}, "", c)
}

// 充值sol
func (monitorTransfersApi *MonitorTransfersApi) Recharge(c *gin.Context) {

	var r apiReq.DecryptRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		response.FailWithMessage("Invalid request format: "+err.Error(), c)
		return
	}

	jsonData, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		response.FailWithMessage("Failed to process request data", c)
		return
	}

	decrypted, err := utils.CBCDecrypt(string(jsonData))
	if err != nil {
		global.GVA_LOG.Error("Decryption failed", zap.Error(err))
		response.FailWithMessage("Decryption failed: "+err.Error(), c)
		return
	}
	decryptedStr, ok := decrypted.(string)
	if !ok {
		response.FailWithMessage("Decryption result is not a string", c)
		return
	}
	var der api.MonitorTransfers
	if err := json.Unmarshal([]byte(decryptedStr), &der); err != nil {
		response.FailWithMessage("Failed to unmarshal decrypted data: "+err.Error(), c)
		return
	}

	monitorTransfersService.CreateAndCheckMonitorTransfers(der)

	response.OkWithMessage("ok", c)
}

const (
	WalletDir      = "./config/wallet"  //
	WalletFile     = "wallet.json"      //
	AutoWalletFile = "auto_wallet.json" //
)

// 提现sol
func (monitorTransfersApi *MonitorTransfersApi) Transfer(c *gin.Context) {

	var r apiReq.DecryptRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		response.FailWithMessage("Invalid request format: "+err.Error(), c)
		return
	}
	jsonData, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		response.FailWithMessage("Failed to process request data", c)
		return
	}
	decrypted, err := utils.CBCDecrypt(string(jsonData))
	if err != nil {
		global.GVA_LOG.Error("Decryption failed", zap.Error(err))
		response.FailWithMessage("Decryption failed: "+err.Error(), c)
		return
	}
	decryptedStr, ok := decrypted.(string)
	if !ok {
		response.FailWithMessage("Decryption result is not a string", c)
		return
	}
	var der apiReq.MonitorTransferApi
	if err := json.Unmarshal([]byte(decryptedStr), &der); err != nil {
		response.FailWithMessage("Failed to unmarshal decrypted data: "+err.Error(), c)
		return
	}

	_, err = os.Getwd()
	if err != nil {
		response.FailWithMessage("get wd fail", c)
		return
	}

	if err := os.MkdirAll(WalletDir, 0700); err != nil {
		response.FailWithMessage("create wallet dir fail: "+err.Error(), c)
		return
	}

	walletPath := filepath.Join(WalletDir, WalletFile)

	absWalletPath, err := filepath.Abs(walletPath)
	if err != nil {
		response.FailWithMessage("get wallet file absolute path fail: "+err.Error(), c)
		return
	}

	if _, err := os.Stat(walletPath); os.IsNotExist(err) {
		response.FailWithMessage("wallet file not exist: "+absWalletPath, c)
		return
	}

	if _, err := os.ReadFile(walletPath); err != nil {
		response.FailWithMessage("wallet file read fail: "+err.Error(), c)
		return
	}

	wm, err := loadWallet(walletPath)
	if err != nil {
		response.FailWithMessage("wallet loading fail: "+err.Error(), c)
		return
	}

	txhash, status, err := secureTransfer(*wm, der.To, uint64(der.Amount), der.Password, der.TotpCode)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 获取hash 保存 提现记录
	fmt.Println("txhash:", txhash, status)
	response.OkWithMessage("ok", c)
}

// 安全转账 - 双重验证
func secureTransfer(wm WalletManager, toAddress string, amount uint64, password string, totpCode string) (string, string, error) {

	err := verifyPassword(wm, password)
	if err != nil {
		return "", "", fmt.Errorf("password fail: %v", err)
	}

	privateKey, err := decryptPrivateKeyWithTOTP(wm, totpCode)
	if err != nil {
		return "", "", fmt.Errorf("TOTP fail: %v", err)
	}

	txHash, status, err := transferSOL(wm, privateKey, toAddress, amount)
	if err != nil {
		return "", "", fmt.Errorf("transfer fail: %v", err)
	}

	fmt.Printf("✓ 转账成功! 交易哈希: %s\n", txHash)
	return txHash, status, nil
}

func verifyTOTPStrict(secret, code string) bool {

	now := time.Now().Unix()
	currentWindow := now / 30
	currentCode, err := totp.GenerateCodeCustom(secret, time.Unix(currentWindow*30, 0), totp.ValidateOpts{
		Period:    30,
		Skew:      0,
		Digits:    6,
		Algorithm: otp.AlgorithmSHA1,
	})
	if err != nil {
		fmt.Printf("code fail: %v\n", err)
		return false
	}

	isValid := code == currentCode

	return isValid
}

type EncryptedWallet struct {
	EncryptedPrivateKey string `json:"encrypted_private_key"`
	PublicKey           string `json:"public_key"`
	Salt                string `json:"salt"`
	IV                  string `json:"iv"`
	EncryptedTOTPKey    string `json:"encrypted_totp_key"`
}

type WalletManager struct {
	client *rpc.Client
	wallet *EncryptedWallet
}

func loadWallet(walletPath string) (*WalletManager, error) {
	walletBytes, err := os.ReadFile(walletPath)
	if err != nil {
		return nil, fmt.Errorf("read wallet fail: %v", err)
	}

	var walletData map[string]interface{}
	err = json.Unmarshal(walletBytes, &walletData)
	if err != nil {
		return nil, fmt.Errorf("read wallet fail1: %v", err)
	}

	walletJSON, err := json.Marshal(walletData["wallet"])
	if err != nil {
		return nil, fmt.Errorf("read wallet fail2: %v", err)
	}

	var encryptedWallet EncryptedWallet
	err = json.Unmarshal(walletJSON, &encryptedWallet)
	if err != nil {
		return nil, fmt.Errorf("read wallet fail3: %v", err)
	}

	return &WalletManager{
		client: rpc.New(rpc.DevNet_RPC),
		wallet: &encryptedWallet,
	}, nil
}

func verifyPassword(wm WalletManager, password string) error {
	salt, err := base64.StdEncoding.DecodeString(wm.wallet.Salt)
	if err != nil {
		return fmt.Errorf("wallet file fail")
	}

	iv, err := base64.StdEncoding.DecodeString(wm.wallet.IV)
	if err != nil {
		return fmt.Errorf("wallet file fail1")
	}

	encryptedTOTPKey, err := base64.StdEncoding.DecodeString(wm.wallet.EncryptedTOTPKey)
	if err != nil {
		return fmt.Errorf("wallet file fail2")
	}

	passwordKey := generateAESKey(password, salt)
	block, err := aes.NewCipher(passwordKey)
	if err != nil {
		return fmt.Errorf("password error")
	}

	totpKey := make([]byte, len(encryptedTOTPKey))
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(totpKey, encryptedTOTPKey)

	ciphertext, err := base64.StdEncoding.DecodeString(wm.wallet.EncryptedPrivateKey)
	if err != nil {
		return fmt.Errorf("wallet file fail3")
	}

	block2, err := aes.NewCipher(totpKey)
	if err != nil {
		return fmt.Errorf("password error1")
	}

	plaintext := make([]byte, len(ciphertext))
	stream2 := cipher.NewCFBDecrypter(block2, iv)
	stream2.XORKeyStream(plaintext, ciphertext)

	if len(plaintext) != 64 {
		return fmt.Errorf("password error2")
	}

	var isValid bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				isValid = false
			}
		}()

		privateKey := solana.PrivateKey(plaintext)
		publicKey := privateKey.PublicKey()
		if publicKey.String() == wm.wallet.PublicKey {
			isValid = true
		} else {
			isValid = false
		}
	}()

	if !isValid {
		return fmt.Errorf("password error3")
	}

	return nil
}

// totp decode private key
func decryptPrivateKeyWithTOTP(wm WalletManager, totpCode string) (*solana.PrivateKey, error) {
	walletPath := filepath.Join(WalletDir, WalletFile)
	walletBytes, err := os.ReadFile(walletPath)
	if err != nil {
		return nil, fmt.Errorf("read wallet fail6: %v", err)
	}

	var walletData map[string]interface{}
	if err := json.Unmarshal(walletBytes, &walletData); err != nil {
		return nil, fmt.Errorf("wallet file fail4: %v", err)
	}

	totpSecret, ok := walletData["totp_secret"].(string)
	if !ok {
		return nil, fmt.Errorf("TOTP secret fail")
	}

	if !verifyTOTPStrict(totpSecret, totpCode) {
		return nil, fmt.Errorf("TOTP verify fail")
	}

	salt, err := base64.StdEncoding.DecodeString(wm.wallet.Salt)
	if err != nil {
		return nil, fmt.Errorf("wallet file fail7")
	}

	iv, err := base64.StdEncoding.DecodeString(wm.wallet.IV)
	if err != nil {
		return nil, fmt.Errorf("wallet file fail8")
	}

	ciphertext, err := base64.StdEncoding.DecodeString(wm.wallet.EncryptedPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("wallet file fail9")
	}

	totpKey := generateAESKey(totpSecret, salt)

	block, err := aes.NewCipher(totpKey)
	if err != nil {
		return nil, fmt.Errorf("decrypt fail")
	}

	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(plaintext, ciphertext)

	if len(plaintext) != 64 {
		return nil, fmt.Errorf("TOTP len fail")
	}

	var privateKey *solana.PrivateKey
	var isValid bool

	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("private key verify panic: %v\n", r)
				isValid = false
			}
		}()

		pk := solana.PrivateKey(plaintext)
		publicKey := pk.PublicKey()

		if publicKey.String() == wm.wallet.PublicKey {
			privateKey = &pk
			isValid = true
		} else {
			isValid = false
		}
	}()

	if !isValid {
		return nil, fmt.Errorf("TOTP verify fail")
	}

	return privateKey, nil
}

// transferSOL
func transferSOL(wm WalletManager, fromPrivateKey *solana.PrivateKey, toAddress string, amount uint64) (string, string, error) {
	toPubKey, err := solana.PublicKeyFromBase58(toAddress)
	if err != nil {
		return "", "", fmt.Errorf("invalid receive address: %v", err)
	}

	fromPubKey := fromPrivateKey.PublicKey()

	instruction := system.NewTransferInstruction(
		amount,
		fromPubKey,
		toPubKey,
	).Build()

	recent, err := wm.client.GetLatestBlockhash(context.Background(), rpc.CommitmentFinalized)
	if err != nil {
		return "", "", fmt.Errorf("get blockhash fail: %v", err)
	}

	tx, err := solana.NewTransaction(
		[]solana.Instruction{instruction},
		recent.Value.Blockhash,
		solana.TransactionPayer(fromPubKey),
	)
	if err != nil {
		return "", "", fmt.Errorf("create transaction fail: %v", err)
	}

	_, err = tx.Sign(
		func(key solana.PublicKey) *solana.PrivateKey {
			if key.String() == fromPubKey.String() {
				return fromPrivateKey
			}
			return nil
		},
	)
	if err != nil {
		return "", "", fmt.Errorf("sign transaction fail: %v", err)
	}

	sig, err := wm.client.SendTransaction(context.Background(), tx)
	if err != nil {
		return "", "", fmt.Errorf("send transaction fail: %v", err)
	}

	txHash := sig.String()

	time.Sleep(2 * time.Second)

	status, err := getTransactionStatus(txHash)
	if err != nil {
		fmt.Printf("警告：无法获取交易状态: %v\n", err)
		status = "sent"
	}

	return txHash, status, nil
}

func generateAESKey(password string, salt []byte) []byte {
	hash := sha256.Sum256([]byte(password + string(salt)))
	return hash[:]
}

// 查询交易状态
func getTransactionStatus(txHash string) (string, error) {
	client := rpc.New(rpc.DevNet_RPC)
	ctx := context.Background()

	// 解析交易哈希
	signature := solana.MustSignatureFromBase58(txHash)

	// 获取交易详情
	tx, err := client.GetTransaction(ctx, signature, &rpc.GetTransactionOpts{
		Encoding:   solana.EncodingBase64,
		Commitment: rpc.CommitmentConfirmed,
	})
	if err != nil {
		return "", fmt.Errorf("获取交易详情失败: %v", err)
	}

	if tx == nil {
		return "pending", nil
	}

	if tx.Meta != nil && tx.Meta.Err != nil {
		return "failed", nil
	}

	return "confirmed", nil
}
