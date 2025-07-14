package utils

import (
	"context"
	"errors"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	jwt "github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenValid            = errors.New("")
	TokenExpired          = errors.New("token")
	TokenNotValidYet      = errors.New("token")
	TokenMalformed        = errors.New("token")
	TokenSignatureInvalid = errors.New("")
	TokenInvalid          = errors.New("token")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.GVA_CONFIG.JWT.SigningKey),
	}
}

func (j *JWT) CreateClaims(baseClaims request.BaseClaims) request.CustomClaims {
	bf, _ := ParseDuration(global.GVA_CONFIG.JWT.BufferTime)
	claims := request.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bf / time.Second), // 1 token
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"GVA"},                   //
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)), //
			// ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),    // 注释掉过期时间，让token永不过期
			Issuer: global.GVA_CONFIG.JWT.Issuer, //
		},
	}
	return claims
}

// CreateToken token
func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken token token
func (j *JWT) CreateTokenByOldToken(oldToken string, claims request.CustomClaims) (string, error) {
	v, err, _ := global.GVA_Concurrency_Control.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// ParseToken  token
func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})

	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenExpired):
			return nil, TokenExpired
		case errors.Is(err, jwt.ErrTokenMalformed):
			return nil, TokenMalformed
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			return nil, TokenSignatureInvalid
		case errors.Is(err, jwt.ErrTokenNotValidYet):
			return nil, TokenNotValidYet
		default:
			return nil, TokenInvalid
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, TokenValid
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetRedisJWT
//@description: jwtredis
//@param: jwt string, userName string
//@return: err error

func SetRedisJWT(jwt string, userName string) (err error) {
	// jwt永不过期
	err = global.GVA_REDIS.Set(context.Background(), userName, jwt, 0).Err()
	return err
}
