package system

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
)

// Login
// @Tags     Base
// @Summary
// @Produce   application/json
// @Param    data  body      systemReq.Login                                                true  ", , "
// @Success  200   {object}  response.Response{data=systemRes.LoginResponse,msg=string}  ",token,"
// @Router   /base/login [post]
func (b *BaseApi) Login(c *gin.Context) {
	var l systemReq.Login
	err := c.ShouldBindJSON(&l)
	key := c.ClientIP()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(l, utils.LoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	//
	openCaptcha := global.GVA_CONFIG.Captcha.OpenCaptcha               //
	openCaptchaTimeOut := global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut //
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	var oc bool = false
	if openCaptcha > 0 {
		oc = true
	} else if ok {
		if count, ok := v.(int); ok && count > 5 {
			oc = true
		}
	}

	if !oc || (l.CaptchaId != "" && l.Captcha != "" && store.Verify(l.CaptchaId, l.Captcha, true)) {
		u := &system.SysUser{Username: l.Username, Password: l.Password}
		user, err := userService.Login(u)
		if err != nil {
			global.GVA_LOG.Error("! !", zap.Error(err))
			// +1
			global.BlackCache.Increment(key, 1)
			response.FailWithMessage("", c)
			return
		}
		if user.Enable != 1 {
			global.GVA_LOG.Error("! !")
			// +1
			global.BlackCache.Increment(key, 1)
			response.FailWithMessage("", c)
			return
		}
		b.TokenNext(c, *user)
		return
	}
	// +1
	global.BlackCache.Increment(key, 1)
	response.FailWithMessage("", c)
}

// TokenNext jwt
func (b *BaseApi) TokenNext(c *gin.Context, user system.SysUser) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} //
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.UUID,
		ID:          uint(user.ID),
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("token!", zap.Error(err))
		response.FailWithMessage("token", c)
		return
	}
	if global.GVA_CONFIG.System.UseMultipoint {
		if err := utils.SetRedisJWT(token, user.Username); err != nil {
			global.GVA_LOG.Error("!", zap.Error(err))
			response.FailWithMessage("", c)
			return
		}
	}
	utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
	response.OkWithDetailed(systemRes.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
	}, "ok", c)
}

// Register
// @Tags     SysUser
// @Summary
// @Produce   application/json
// @Param    data  body      systemReq.Register                                            true  ", , , ID"
// @Success  200   {object}  response.Response{data=systemRes.SysUserResponse,msg=string}  ","
// @Router   /user/admin_register [post]
func (b *BaseApi) Register(c *gin.Context) {
	var r systemReq.Register
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(r, utils.RegisterVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var authorities []system.SysAuthority
	for _, v := range r.AuthorityIds {
		authorities = append(authorities, system.SysAuthority{
			AuthorityId: v,
		})
	}
	user := &system.SysUser{Username: r.Username, NickName: r.NickName, Password: r.Password, HeaderImg: r.HeaderImg, AuthorityId: r.AuthorityId, Authorities: authorities, Enable: r.Enable, Phone: r.Phone, Email: r.Email}
	userReturn, err := userService.Register(*user)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithDetailed(systemRes.SysUserResponse{User: userReturn}, "", c)
		return
	}
	response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "", c)
}

// ChangePassword
// @Tags      SysUser
// @Summary
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      systemReq.ChangePasswordReq    true  ", , "
// @Success   200   {object}  response.Response{msg=string}  ""
// @Router    /user/changePassword [post]
func (b *BaseApi) ChangePassword(c *gin.Context) {
	var req systemReq.ChangePasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(req, utils.ChangePasswordVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := utils.GetUserID(c)
	u := &system.SysUser{GVA_MODEL: global.GVA_MODEL{ID: uid}, Password: req.Password}
	_, err = userService.ChangePassword(u, req.NewPassword)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage("，", c)
		return
	}
	response.OkWithMessage("", c)
}
func (b *BaseApi) ChangeWithdrawPassword(c *gin.Context) {
	fmt.Println("ChangeWithdrawPassword")

	var req systemReq.ChangePasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(req, utils.ChangePasswordVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.FailWithMessage("ID", c)
		return
	}

	u := &system.SysUser{GVA_MODEL: global.GVA_MODEL{ID: uid}, Password: req.Password}
	_, err = userService.ChangeWithdrawPassword(u, req.NewPassword)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage("，", c)
		return
	}
	response.OkWithMessage("", c)
}

// GetUserList
// @Tags      SysUser
// @Summary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.GetUserList                                        true  ", "
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  ",,,,"
// @Router    /user/getUserList [post]
func (b *BaseApi) GetUserList(c *gin.Context) {
	var pageInfo systemReq.GetUserList
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userService.GetUserInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage("", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "", c)
}

// SetUserAuthority
// @Tags      SysUser
// @Summary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.SetUserAuth          true  "UUID, ID"
// @Success   200   {object}  response.Response{msg=string}  ""
// @Router    /user/setUserAuthority [post]
func (b *BaseApi) SetUserAuthority(c *gin.Context) {
	var sua systemReq.SetUserAuth
	err := c.ShouldBindJSON(&sua)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if UserVerifyErr := utils.Verify(sua, utils.SetUserAuthorityVerify); UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	err = userService.SetUserAuthority(userID, sua.AuthorityId)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	claims := utils.GetUserInfo(c)
	claims.AuthorityId = sua.AuthorityId
	token, err := utils.NewJWT().CreateToken(*claims)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	c.Header("new-token", token)
	c.Header("new-expires-at", strconv.FormatInt(claims.ExpiresAt.Unix(), 10))
	utils.SetToken(c, token, int((claims.ExpiresAt.Unix()-time.Now().Unix())/60))
	response.OkWithMessage("", c)
}

// SetUserAuthorities
// @Tags      SysUser
// @Summary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.SetUserAuthorities   true  "UUID, ID"
// @Success   200   {object}  response.Response{msg=string}  ""
// @Router    /user/setUserAuthorities [post]
func (b *BaseApi) SetUserAuthorities(c *gin.Context) {
	var sua systemReq.SetUserAuthorities
	err := c.ShouldBindJSON(&sua)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	authorityID := utils.GetUserAuthorityId(c)
	err = userService.SetUserAuthorities(authorityID, sua.ID, sua.AuthorityIds)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage("", c)
		return
	}
	response.OkWithMessage("", c)
}

// DeleteUser
// @Tags      SysUser
// @Summary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "ID"
// @Success   200   {object}  response.Response{msg=string}  ""
// @Router    /user/deleteUser [delete]
func (b *BaseApi) DeleteUser(c *gin.Context) {
	var reqId request.GetById
	err := c.ShouldBindJSON(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(reqId, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	jwtId := utils.GetUserID(c)
	if jwtId == uint(reqId.ID) {
		response.FailWithMessage(", 。", c)
		return
	}
	err = userService.DeleteUser(reqId.ID)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage("", c)
		return
	}
	response.OkWithMessage("", c)
}

// SetUserInfo
// @Tags      SysUser
// @Summary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysUser                                             true  "ID, , , "
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  ""
// @Router    /user/setUserInfo [put]
func (b *BaseApi) SetUserInfo(c *gin.Context) {
	var user systemReq.ChangeUserInfo
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(user, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if len(user.AuthorityIds) != 0 {
		authorityID := utils.GetUserAuthorityId(c)
		err = userService.SetUserAuthorities(authorityID, user.ID, user.AuthorityIds)
		if err != nil {
			global.GVA_LOG.Error("!", zap.Error(err))
			response.FailWithMessage("", c)
			return
		}
	}
	err = userService.SetUserInfo(system.SysUser{
		GVA_MODEL: global.GVA_MODEL{
			ID: user.ID,
		},
		NickName:  user.NickName,
		HeaderImg: user.HeaderImg,
		Phone:     user.Phone,
		Email:     user.Email,
		Enable:    user.Enable,
	})
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage("", c)
		return
	}
	response.OkWithMessage("", c)
}

// SetSelfInfo
// @Tags      SysUser
// @Summary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysUser                                             true  "ID, , , "
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  ""
// @Router    /user/SetSelfInfo [put]
func (b *BaseApi) SetSelfInfo(c *gin.Context) {
	var user systemReq.ChangeUserInfo
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user.ID = utils.GetUserID(c)
	err = userService.SetSelfInfo(system.SysUser{
		GVA_MODEL: global.GVA_MODEL{
			ID: user.ID,
		},
		NickName:  user.NickName,
		HeaderImg: user.HeaderImg,
		Phone:     user.Phone,
		Email:     user.Email,
		Enable:    user.Enable,
	})
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage("", c)
		return
	}
	response.OkWithMessage("", c)
}

// SetSelfSetting
// @Tags      SysUser
// @Summary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      map[string]interface{}  true  ""
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  ""
// @Router    /user/SetSelfSetting [put]
func (b *BaseApi) SetSelfSetting(c *gin.Context) {
	var req common.JSONMap
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = userService.SetSelfSetting(req, utils.GetUserID(c))
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage("", c)
		return
	}
	response.OkWithMessage("", c)
}

// GetUserInfo
// @Tags      SysUser
// @Summary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  ""
// @Router    /user/getUserInfo [get]
func (b *BaseApi) GetUserInfo(c *gin.Context) {
	uuid := utils.GetUserUuid(c)
	ReqUser, err := userService.GetUserInfo(uuid)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage("", c)
		return
	}
	response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "", c)
}

// ResetPassword
// @Tags      SysUser
// @Summary
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      system.SysUser                 true  "ID"
// @Success   200   {object}  response.Response{msg=string}  ""
// @Router    /user/resetPassword [post]
func (b *BaseApi) ResetPassword(c *gin.Context) {
	var rps systemReq.ResetPassword
	err := c.ShouldBindJSON(&rps)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.ResetPassword(rps.ID, rps.Password)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(""+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}
func (b *BaseApi) ResetWithdrawPassword(c *gin.Context) {
	// ID
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.FailWithMessage("ID", c)
		return
	}

	var rps systemReq.WithdrawPassword
	err := c.ShouldBindJSON(&rps)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	rps.ID = uid
	fmt.Println("ResetWithdrawPassword", uid)
	err = userService.ResetWithdrawPassword(rps.ID, rps.Password)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(""+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}
func (b *BaseApi) ApiLogin(c *gin.Context) {
	var l systemReq.ApiLogin
	err := c.ShouldBindJSON(&l)
	key := c.ClientIP()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// err = utils.Verify(l, utils.LoginVerify)
	// if err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// 	return
	// }
	u := &system.SysUser{Username: l.Username, Password: l.Password}
	user, err := userService.ApiLogin(u)
	if err != nil {
		global.GVA_LOG.Error("! !", zap.Error(err))
		global.BlackCache.Increment(key, 1)
		response.FailWithMessage("", c)
		return
	}
	if user.Enable != 1 {
		global.GVA_LOG.Error("! !")
		// +1
		response.FailWithMessage("", c)
		return
	}
	b.ApiTokenNext(c, *user)
	return
}

// TokenNext jwt
func (b *BaseApi) ApiTokenNext(c *gin.Context, user system.SysUser) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)}
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:     user.UUID,
		ID:       uint(user.ID),
		NickName: user.NickName,
		Username: user.Username,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("get token error!", zap.Error(err))
		response.FailWithMessage("get token error", c)
		return
	}

	// token
	if err := utils.SetRedisJWT(token, user.Username); err != nil {
		global.GVA_LOG.Error("set fail!", zap.Error(err))
		response.FailWithMessage("set fail", c)
		return
	}

	utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
	response.OkWithDetailed(systemRes.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
	}, "ok", c)
}
func (b *BaseApi) ApiRegister(c *gin.Context) {
	var r systemReq.Register
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(r, utils.ApiRegisterVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var authorities []system.SysAuthority
	for _, v := range r.AuthorityIds {
		authorities = append(authorities, system.SysAuthority{
			AuthorityId: v,
		})
	}

	r.AuthorityId = 888
	r.AuthorityIds = []uint{888}
	r.Enable = 1

	user := &system.SysUser{Username: r.Username, NickName: r.NickName, Password: r.Password, HeaderImg: r.HeaderImg, AuthorityId: r.AuthorityId, Authorities: authorities, Enable: r.Enable, Phone: r.Phone, Email: r.Email}
	userReturn, err := userService.ApiRegister(*user)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithDetailed(systemRes.SysUserResponse{User: userReturn}, "", c)
		return
	}
	response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "", c)
}

func GenerateVerificationCode() string {
	const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const length = 4

	rand.Seed(time.Now().UnixNano())

	code := make([]byte, length)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}

	return string(code)
}
func (b *BaseApi) SendCode(c *gin.Context) {
	uid := utils.GetRedisUserID(c)
	if uid == 0 {
		response.Result(401, nil, "", c)
		return
	}
	var r systemReq.SendEmailCodeRequest
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(r, utils.ApiSendEmailCodeVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.CheckEmail(r.Email)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	code := GenerateVerificationCode()

	err = global.GVA_REDIS.Set(c, fmt.Sprintf("email_code_%d:%s", uid, r.Email), code, 5*time.Minute).Err()
	if err != nil {
		global.GVA_LOG.Error("save fail !", zap.Error(err))
		response.FailWithMessage("send fail", c)
		return
	}

	err = SendVerificationEmail(r.Email, code)
	if err != nil {
		global.GVA_LOG.Error("send fail!", zap.Error(err))
		response.FailWithMessage("send fail", c)
		return
	}

	response.OkWithMessage("send ok", c)
}

func SendVerificationEmail(email, code string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "lul0215@163.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "verification code")
	m.SetBody("text/plain", fmt.Sprintf("Your verification code is: %s, valid for 5 minutes.", code))

	d := gomail.NewDialer("smtp.163.com", 465, "lul0215@163.com", "PNgReYFe54tQNej6")
	d.SSL = true

	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("send fail: %v\n", err)
		return err
	}

	return nil
}
func (b *BaseApi) BindeMail(c *gin.Context) {
	uid := utils.GetRedisUserID(c)
	if uid == 0 {
		response.Result(401, nil, "user fail", c)
		return
	}
	var r systemReq.BindEmailRequest
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	storedCode, err := global.GVA_REDIS.Get(c, fmt.Sprintf("email_code_%d:%s", uid, r.Email)).Result()
	if err != nil {
		response.FailWithMessage("code error", c)
		return
	}

	if storedCode != r.Code {
		response.FailWithMessage("code error", c)
		return
	}

	global.GVA_REDIS.Del(c, fmt.Sprintf("email_code:%s", r.Email))

	err = userService.BindEmail(uid, r.Email)
	if err != nil {
		global.GVA_LOG.Error("bind email fail!", zap.Error(err))
		response.FailWithDetailed(nil, "", c)
		return
	}

	response.OkWithDetailed(nil, "bind email ok", c)
}
