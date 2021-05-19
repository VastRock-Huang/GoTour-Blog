package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/vastrock-huang/gotour-blogservice/global"
	"github.com/vastrock-huang/gotour-blogservice/pkg/util"
	"time"
)

//JWT声明
type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims	//JWT库中的JWT声明,即JWT的Payload部分
	//Audience  string `json:"aud,omitempty"`	//JWT 受众
	//ExpiresAt int64  `json:"exp,omitempty"`	//过期时间
	//Id        string `json:"jti,omitempty"`	//JWT ID
	//IssuedAt  int64  `json:"iat,omitempty"`	//签发时间
	//Issuer    string `json:"iss,omitempty"`	//签发者
	//NotBefore int64  `json:"nbf,omitempty"`	//生效时间
	//Subject   string `json:"sub,omitempty"`	//主题
}

//获取JWT Secret
func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

//生成JWT的令牌
func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)	//过期时间
	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}
	//生成token,传入的实参是加密方法和payload部分
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//生成签名字符串,使用密钥签名
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

//解析校验Token
func ParseToken(token string) (*Claims, error) {
	//使用密钥解析Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{},
	func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		//判断Token是否有效
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}