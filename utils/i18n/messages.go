package i18n

import (
	"fmt"
	"strings"
)

// Language constants
const (
	LangEnglish    = "en"
	LangPortuguese = "pt"
	DefaultLang    = LangEnglish
)

// MessageKey 消息键类型
type MessageKey string

// 定义消息键常量
const (
	// 通用消息
	MsgSuccess           MessageKey = "success"
	MsgFailed            MessageKey = "failed"
	MsgUserNotFound      MessageKey = "user_not_found"
	MsgInsufficientFunds MessageKey = "insufficient_funds"
	MsgInvalidRequest    MessageKey = "invalid_request"

	// 登录相关消息
	MsgLoginPasswordError    MessageKey = "login_password_error"
	MsgWithdrawPasswordError MessageKey = "withdraw_password_error"
	MsgOldPasswordError      MessageKey = "old_password_error"
	MsgPasswordError         MessageKey = "password_error"
	MsgEmailDuplicate        MessageKey = "email_duplicate"
	MsgUsernameDuplicate     MessageKey = "username_duplicate"

	// 支付相关消息
	MsgPaymentSuccess    MessageKey = "payment_success"
	MsgPaymentFailed     MessageKey = "payment_failed"
	MsgPaymentPending    MessageKey = "payment_pending"
	MsgPaymentProcessing MessageKey = "payment_processing"
	MsgPaymentCancelled  MessageKey = "payment_cancelled"

	// 提现相关消息
	MsgWithdrawalSuccess      MessageKey = "withdrawal_success"
	MsgWithdrawalFailed       MessageKey = "withdrawal_failed"
	MsgWithdrawalPending      MessageKey = "withdrawal_pending"
	MsgWithdrawalCancelled    MessageKey = "withdrawal_cancelled"
	MsgAccountNotFound        MessageKey = "account_not_found"
	MsgInvalidAmount          MessageKey = "invalid_amount"
	MsgRecordNotFound         MessageKey = "record_not_found"
	MsgUnauthorized           MessageKey = "unauthorized"
	MsgCannotCancelWithdrawal MessageKey = "cannot_cancel_withdrawal"

	// 错误消息
	MsgCreateRecordFailed MessageKey = "create_record_failed"
	MsgUpdateStatusFailed MessageKey = "update_status_failed"
	MsgRedisUpdateFailed  MessageKey = "redis_update_failed"
	MsgDatabaseError      MessageKey = "database_error"
	MsgSystemBusy         MessageKey = "system_busy"
	MsgSystemError        MessageKey = "system_error"
)

// Messages 多语言消息映射
var Messages = map[string]map[MessageKey]string{
	LangEnglish: {
		// 通用消息
		MsgSuccess:           "Success",
		MsgFailed:            "Failed",
		MsgUserNotFound:      "User not found",
		MsgInsufficientFunds: "Insufficient funds",
		MsgInvalidRequest:    "Invalid request",

		// 登录相关消息
		MsgLoginPasswordError:    "Login password error",
		MsgWithdrawPasswordError: "Withdraw password error",
		MsgOldPasswordError:      "Old password error",
		MsgPasswordError:         "Password error",
		MsgEmailDuplicate:        "Email already exists",
		MsgUsernameDuplicate:     "Username already exists",

		// 支付相关消息
		MsgPaymentSuccess:    "Payment successful",
		MsgPaymentFailed:     "Payment failed",
		MsgPaymentPending:    "Payment pending",
		MsgPaymentProcessing: "Payment processing",
		MsgPaymentCancelled:  "Payment cancelled",

		// 提现相关消息
		MsgWithdrawalSuccess:      "Withdrawal successful",
		MsgWithdrawalFailed:       "Withdrawal failed",
		MsgWithdrawalPending:      "Withdrawal pending",
		MsgWithdrawalCancelled:    "Withdrawal cancelled successfully",
		MsgAccountNotFound:        "Withdrawal account not found",
		MsgInvalidAmount:          "Invalid amount",
		MsgRecordNotFound:         "Record not found",
		MsgUnauthorized:           "Unauthorized access",
		MsgCannotCancelWithdrawal: "Cannot cancel withdrawal with current status",

		// 错误消息
		MsgCreateRecordFailed: "Failed to create transaction record",
		MsgUpdateStatusFailed: "Failed to update transaction status",
		MsgRedisUpdateFailed:  "Failed to update user balance",
		MsgDatabaseError:      "Database operation failed",
		MsgSystemBusy:         "System is busy, please try again later",
		MsgSystemError:        "System error, please contact support",
	},
	LangPortuguese: {
		// 通用消息
		MsgSuccess:           "Sucesso",
		MsgFailed:            "Falhou",
		MsgUserNotFound:      "Usuário não encontrado",
		MsgInsufficientFunds: "Saldo insuficiente",
		MsgInvalidRequest:    "Solicitação inválida",

		// 登录相关消息
		MsgLoginPasswordError:    "Erro na senha de login",
		MsgWithdrawPasswordError: "Erro na senha de saque",
		MsgOldPasswordError:      "Erro na senha antiga",
		MsgPasswordError:         "Erro na senha",
		MsgEmailDuplicate:        "Email já existe",
		MsgUsernameDuplicate:     "Nome de usuário já existe",

		// 支付相关消息
		MsgPaymentSuccess:    "Pagamento bem-sucedido",
		MsgPaymentFailed:     "Pagamento falhou",
		MsgPaymentPending:    "Pagamento pendente",
		MsgPaymentProcessing: "Processando pagamento",
		MsgPaymentCancelled:  "Pagamento cancelado",

		// 提现相关消息
		MsgWithdrawalSuccess:      "Saque bem-sucedido",
		MsgWithdrawalFailed:       "Saque falhou",
		MsgWithdrawalPending:      "Saque pendente",
		MsgWithdrawalCancelled:    "Saque cancelado com sucesso",
		MsgAccountNotFound:        "Conta de saque não encontrada",
		MsgInvalidAmount:          "Quantia inválida",
		MsgRecordNotFound:         "Registro não encontrado",
		MsgUnauthorized:           "Acesso não autorizado",
		MsgCannotCancelWithdrawal: "Não é possível cancelar o saque com o status atual",

		// 错误消息
		MsgCreateRecordFailed: "Falha ao criar registro de transação",
		MsgUpdateStatusFailed: "Falha ao atualizar status da transação",
		MsgRedisUpdateFailed:  "Falha ao atualizar saldo do usuário",
		MsgDatabaseError:      "Operação de banco de dados falhou",
		MsgSystemBusy:         "Sistema ocupado, tente novamente mais tarde",
		MsgSystemError:        "Erro do sistema, entre em contato com o suporte",
	},
}

// GetMessage 获取指定语言的消息
func GetMessage(lang string, key MessageKey) string {
	// 规范化语言代码
	lang = NormalizeLang(lang)

	// 检查语言是否存在
	if messages, exists := Messages[lang]; exists {
		if message, found := messages[key]; found {
			return message
		}
	}

	// 回退到默认语言
	if messages, exists := Messages[DefaultLang]; exists {
		if message, found := messages[key]; found {
			return message
		}
	}

	// 如果都找不到，返回键名
	return string(key)
}

// NormalizeLang 规范化语言代码
func NormalizeLang(lang string) string {
	if lang == "" {
		return DefaultLang
	}

	// 转换为小写
	lang = strings.ToLower(lang)

	// 处理常见的语言代码变体
	switch {
	case strings.HasPrefix(lang, "pt"):
		return LangPortuguese
	case strings.HasPrefix(lang, "en"):
		return LangEnglish
	default:
		return DefaultLang
	}
}

// GetLangFromHeader 从请求头获取语言代码
func GetLangFromHeader(acceptLang string) string {
	if acceptLang == "" {
		return DefaultLang
	}

	// 解析 Accept-Language 头
	// 例如: "pt-BR,pt;q=0.9,en;q=0.8"
	langs := strings.Split(acceptLang, ",")
	for _, lang := range langs {
		// 去除权重信息 (q=0.9)
		lang = strings.Split(strings.TrimSpace(lang), ";")[0]
		normalized := NormalizeLang(lang)

		// 如果是支持的语言，直接返回
		if _, exists := Messages[normalized]; exists {
			return normalized
		}
	}
	fmt.Println("DefaultLang", DefaultLang)
	return DefaultLang
}
