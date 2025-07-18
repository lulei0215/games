package system

import (
	"errors"
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Register
//@description:
//@param: u model.SysUser
//@return: userInter system.SysUser, err error

type UserService struct{}

var UserServiceApp = new(UserService)

func (userService *UserService) Register(u system.SysUser) (userInter system.SysUser, err error) {
	var user system.SysUser
	if !errors.Is(global.GVA_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { //
		return userInter, errors.New("")
	}
	//  uuid hash
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.New()
	err = global.GVA_DB.Create(&u).Error
	return u, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: Login
//@description:
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func (userService *UserService) Login(u *system.SysUser) (userInter *system.SysUser, err error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}

	var user system.SysUser
	err = global.GVA_DB.Where("username = ?", u.Username).Preload("Authorities").Preload("Authority").First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("password error")
		}
		MenuServiceApp.UserAuthorityDefaultRouter(&user)
	}
	return &user, err
}

func (userService *UserService) ApiLogin(u *system.SysUser) (userInter *system.SysUser, err error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}
	var user system.SysUser
	err = global.GVA_DB.Where("username = ?", u.Username).Preload("Authorities").Preload("Authority").First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("password error")
		}
	}
	return &user, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ChangePassword
//@description:
//@param: u *model.SysUser, newPassword string
//@return: userInter *model.SysUser,err error

func (userService *UserService) ChangePassword(u *system.SysUser, newPassword string) (userInter *system.SysUser, err error) {
	fmt.Println("ChangePassword", u.ID)
	var user system.SysUser
	if err = global.GVA_DB.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return nil, errors.New("old password error")
	}
	user.Password = utils.BcryptHash(newPassword)
	err = global.GVA_DB.Save(&user).Error
	return &user, err

}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUserInfoList
//@description:
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (userService *UserService) GetUserInfoList(info systemReq.GetUserList) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysUser{})
	var userList []system.SysUser

	if info.NickName != "" {
		db = db.Where("nick_name LIKE ?", "%"+info.NickName+"%")
	}
	if info.Phone != "" {
		db = db.Where("phone LIKE ?", "%"+info.Phone+"%")
	}
	if info.Username != "" {
		db = db.Where("username LIKE ?", "%"+info.Username+"%")
	}
	if info.Email != "" {
		db = db.Where("email LIKE ?", "%"+info.Email+"%")
	}
	db = db.Where("id > 1")
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
	return userList, total, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserAuthority
//@description:
//@param: uuid uuid.UUID, authorityId string
//@return: err error

func (userService *UserService) SetUserAuthority(id uint, authorityId uint) (err error) {

	assignErr := global.GVA_DB.Where("sys_user_id = ? AND sys_authority_authority_id = ?", id, authorityId).First(&system.SysUserAuthority{}).Error
	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
		return errors.New("")
	}

	var authority system.SysAuthority
	err = global.GVA_DB.Where("authority_id = ?", authorityId).First(&authority).Error
	if err != nil {
		return err
	}
	var authorityMenu []system.SysAuthorityMenu
	var authorityMenuIDs []string
	err = global.GVA_DB.Where("sys_authority_authority_id = ?", authorityId).Find(&authorityMenu).Error
	if err != nil {
		return err
	}

	for i := range authorityMenu {
		authorityMenuIDs = append(authorityMenuIDs, authorityMenu[i].MenuId)
	}

	var authorityMenus []system.SysBaseMenu
	err = global.GVA_DB.Preload("Parameters").Where("id in (?)", authorityMenuIDs).Find(&authorityMenus).Error
	if err != nil {
		return err
	}
	hasMenu := false
	for i := range authorityMenus {
		if authorityMenus[i].Name == authority.DefaultRouter {
			hasMenu = true
			break
		}
	}
	if !hasMenu {
		return errors.New(",")
	}

	err = global.GVA_DB.Model(&system.SysUser{}).Where("id = ?", id).Update("authority_id", authorityId).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserAuthorities
//@description:
//@param: id uint, authorityIds []string
//@return: err error

func (userService *UserService) SetUserAuthorities(adminAuthorityID, id uint, authorityIds []uint) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var user system.SysUser
		TxErr := tx.Where("id = ?", id).First(&user).Error
		if TxErr != nil {
			global.GVA_LOG.Debug(TxErr.Error())
			return errors.New("")
		}
		TxErr = tx.Delete(&[]system.SysUserAuthority{}, "sys_user_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}
		var useAuthority []system.SysUserAuthority
		for _, v := range authorityIds {
			e := AuthorityServiceApp.CheckAuthorityIDAuth(adminAuthorityID, v)
			if e != nil {
				return e
			}
			useAuthority = append(useAuthority, system.SysUserAuthority{
				SysUserId: id, SysAuthorityAuthorityId: v,
			})
		}
		TxErr = tx.Create(&useAuthority).Error
		if TxErr != nil {
			return TxErr
		}
		TxErr = tx.Model(&user).Update("authority_id", authorityIds[0]).Error
		if TxErr != nil {
			return TxErr
		}
		//  nil
		return nil
	})
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUser
//@description:
//@param: id float64
//@return: err error

func (userService *UserService) DeleteUser(id int) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&system.SysUser{}).Error; err != nil {
			return err
		}
		if err := tx.Delete(&[]system.SysUserAuthority{}, "sys_user_id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserInfo
//@description:
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser

func (userService *UserService) SetUserInfo(req system.SysUser) error {
	return global.GVA_DB.Model(&system.SysUser{}).
		Select("updated_at", "nick_name", "header_img", "phone", "email", "enable", "level").
		Where("id=?", req.ID).
		Updates(map[string]interface{}{
			"updated_at": time.Now(),
			"nick_name":  req.NickName,
			"header_img": req.HeaderImg,
			"phone":      req.Phone,
			"email":      req.Email,
			"level":      req.Level,
			"enable":     req.Enable,
		}).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetSelfInfo
//@description:
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser

func (userService *UserService) SetSelfInfo(req system.SysUser) error {
	return global.GVA_DB.Model(&system.SysUser{}).
		Where("id=?", req.ID).
		Updates(req).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetSelfSetting
//@description:
//@param: req datatypes.JSON, uid uint
//@return: err error

func (userService *UserService) SetSelfSetting(req common.JSONMap, uid uint) error {
	return global.GVA_DB.Model(&system.SysUser{}).Where("id = ?", uid).Update("origin_setting", req).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: GetUserInfo
//@description:
//@param: uuid uuid.UUID
//@return: err error, user system.SysUser

func (userService *UserService) GetUserInfo(uuid uuid.UUID) (user system.SysUser, err error) {
	var reqUser system.SysUser
	err = global.GVA_DB.Preload("Authorities").Preload("Authority").First(&reqUser, "uuid = ?", uuid).Error
	if err != nil {
		return reqUser, err
	}
	MenuServiceApp.UserAuthorityDefaultRouter(&reqUser)
	return reqUser, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: id
//@param: id int
//@return: err error, user *model.SysUser

func (userService *UserService) FindUserById(id int) (user *system.SysUser, err error) {
	var u system.SysUser
	err = global.GVA_DB.Where("id = ?", id).First(&u).Error
	return &u, err
}
func (userService *UserService) FindUserByUId(id uint) (user *system.SysUser, err error) {
	var u system.SysUser
	err = global.GVA_DB.Where("id = ?", id).First(&u).Error
	return &u, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserByUuid
//@description: uuid
//@param: uuid string
//@return: err error, user *model.SysUser

func (userService *UserService) FindUserByUuid(uuid string) (user *system.SysUser, err error) {
	var u system.SysUser
	if err = global.GVA_DB.Where("uuid = ?", uuid).First(&u).Error; err != nil {
		return &u, errors.New("")
	}
	return &u, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ResetPassword
//@description:
//@param: ID uint
//@return: err error

func (userService *UserService) ResetPassword(ID uint, password string) (err error) {
	err = global.GVA_DB.Model(&system.SysUser{}).Where("id = ?", ID).Update("password", utils.BcryptHash(password)).Error
	return err
}
func (userService *UserService) ResetWithdrawPassword(ID uint, password string) (err error) {
	fmt.Println("ResetWithdrawPassword", ID, password)
	err = global.GVA_DB.Model(&system.SysUser{}).Where("id = ?", ID).Update("withdraw_password", utils.BcryptHash(password)).Error
	return err
}
func (userService *UserService) ChangeWithdrawPassword(u *system.SysUser, newPassword string) (userInter *system.SysUser, err error) {
	var user system.SysUser
	if err = global.GVA_DB.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(u.Password, user.WithdrawPassword); !ok {
		return nil, errors.New("old password error")
	}
	
	// 直接使用Update方法更新withdraw_password字段，更高效
	err = global.GVA_DB.Model(&system.SysUser{}).
		Where("id = ?", u.ID).
		Update("withdraw_password", utils.BcryptHash(newPassword)).Error
	
	if err != nil {
		return nil, err
	}
	
	// 返回更新后的用户信息
	user.WithdrawPassword = utils.BcryptHash(newPassword)
	return &user, err
}
func (userService *UserService) VerifyWithdrawPassword(u *system.SysUser, newPassword string) (err error) {
	var user system.SysUser
	if err = global.GVA_DB.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return err
	}
	if ok := utils.BcryptCheck(u.Password, user.WithdrawPassword); !ok {
		return errors.New("WITHDRAW_PASSWORD_ERROR")
	}

	return nil

}
func (userService *UserService) SetWithdrawPassword(u *system.SysUser, newPassword string, loginPassword string) (userInter system.SysUser, err error) {
	var user system.SysUser
	if err = global.GVA_DB.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return user, err
	}
	if ok := utils.BcryptCheck(loginPassword, user.Password); !ok {
		return user, errors.New("LOGIN_PASSWORD_ERROR")
	}
	user.WithdrawPassword = utils.BcryptHash(newPassword)
	err = global.GVA_DB.Save(&user).Error
	return user, err

}
func (userService *UserService) ApiRegister(u system.SysUser) (userInter system.SysUser, err error) {
	var user system.SysUser
	if !errors.Is(global.GVA_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { //
		return userInter, errors.New("USERNAME_DUPLICATE")
	}
	//  uuid hash
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.New()
	err = global.GVA_DB.Create(&u).Error
	return u, err
}
func (userService *UserService) BindEmail(ID uint, email string) (err error) {
	err = global.GVA_DB.Model(&system.SysUser{}).Where("id = ?", ID).Update("email", email).Error
	return err
}
func (userService *UserService) CheckEmail(email string) (err error) {
	var user system.SysUser
	if !errors.Is(global.GVA_DB.Where("email = ?", email).First(&user).Error, gorm.ErrRecordNotFound) { //
		return errors.New("EMAIL_DUPLICATE")
	}
	return nil
}
func (userService *UserService) GetRobot(number int) (users_res []system.ApiSysUser, err error) {
	var users []system.SysUser
	err = global.GVA_DB.Where("robot = ?", 1).Limit(number).Find(&users).Error
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, fmt.Errorf("no available robots found")
	}
	var userIds []uint
	for _, user := range users {
		userIds = append(userIds, user.ID)
	}

	err = global.GVA_DB.Model(&system.SysUser{}).
		Where("id IN ?", userIds).
		Update("robot", 2).Error
	if err != nil {
		return nil, fmt.Errorf("update robot status failed: %v", err)
	}
	for _, user := range users {
		apiUser := system.ApiSysUser{
			UUID:             user.UUID,
			Username:         user.Username,
			Password:         user.Password,
			NickName:         user.NickName,
			HeaderImg:        user.HeaderImg,
			Phone:            user.Phone,
			Email:            user.Email,
			Enable:           user.Enable,
			WithdrawPassword: user.WithdrawPassword,
			Balance:          user.Balance,
			Birthday:         user.Birthday,
			Facebook:         user.Facebook,
			Whatsapp:         user.Whatsapp,
			Telegram:         user.Telegram,
			Twitter:          user.Twitter,
			VipLevel:         user.VipLevel,
			VipExpireTime:    user.VipExpireTime,
			UserType:         user.UserType,
			Level:            user.Level,
		}
		users_res = append(users_res, apiUser)
	}
	return users_res, nil
}

func (userService *UserService) UserRelation(ID uint, email string) (err error) {
	fmt.Println("UserRelation", ID, email)
	err = global.GVA_DB.Model(&system.SysUser{}).Where("id = ?", ID).Update("email", email).Error
	return err
}
