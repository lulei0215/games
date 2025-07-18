package utils

var (
	IdVerify                        = Rules{"ID": []string{NotEmpty()}}
	ApiVerify                       = Rules{"Path": {NotEmpty()}, "Description": {NotEmpty()}, "ApiGroup": {NotEmpty()}, "Method": {NotEmpty()}}
	MenuVerify                      = Rules{"Path": {NotEmpty()}, "Name": {NotEmpty()}, "Component": {NotEmpty()}, "Sort": {Ge("0")}}
	MenuMetaVerify                  = Rules{"Title": {NotEmpty()}}
	LoginVerify                     = Rules{"CaptchaId": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	RegisterVerify                  = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityId": {NotEmpty()}}
	ApiRegisterVerify               = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}}
	PageInfoVerify                  = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	CustomerVerify                  = Rules{"CustomerName": {NotEmpty()}, "CustomerPhoneData": {NotEmpty()}}
	AutoCodeVerify                  = Rules{"Abbreviation": {NotEmpty()}, "StructName": {NotEmpty()}, "PackageName": {NotEmpty()}}
	AutoPackageVerify               = Rules{"PackageName": {NotEmpty()}}
	AuthorityVerify                 = Rules{"AuthorityId": {NotEmpty()}, "AuthorityName": {NotEmpty()}}
	AuthorityIdVerify               = Rules{"AuthorityId": {NotEmpty()}}
	OldAuthorityVerify              = Rules{"OldAuthorityId": {NotEmpty()}}
	ChangePasswordVerify            = Rules{"Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	SetUserAuthorityVerify          = Rules{"AuthorityId": {NotEmpty()}}
	ApiSendEmailCodeVerify          = Rules{"email": {NotEmpty()}}
	ApiBettingVerify                = Rules{"coin": {NotEmpty(), Gt("0")}}
	SettleVerify                    = Rules{"SessionId": {NotEmpty()}, "Gid": {NotEmpty()}}
	CreateTradeVerify               = Rules{"TotalAmount": {Gt("0")}}
	CreateTrade2Verify              = Rules{"TotalAmount": {Gt("0")}}
	AddUserWithdrawalAccountsVerify = Rules{"AccountName": {NotEmpty()}, "AccountType": {NotEmpty()}, "AccountNumber": {NotEmpty()}, "CpfNumber": {NotEmpty()}}
	AddWithdrawAccountRequestVerify = Rules{"BankCode": {NotEmpty()}, "BankAcctName": {NotEmpty()}, "BankFirstName": {NotEmpty()}, "BankLastName": {NotEmpty()}, "BankAcctNo": {NotEmpty()}, "AccPhone": {NotEmpty()}, "IdentityNo": {NotEmpty()}, "IdentityType": {NotEmpty()}}
	TradeCallbackRequestVerify      = Rules{"AccountName": {NotEmpty()}, "AccountType": {NotEmpty()}, "AccountNumber": {NotEmpty()}, "CpfNumber": {NotEmpty()}}
)
