// Package app JWT权限处理，例如生成、校验token
package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/golang-minibear2333/gin-blog/global"
	"github.com/golang-minibear2333/gin-blog/pkg/util"
	"time"
)

// Claims JMT 数据结构，用于计算token
type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	// 这是jwt库里预定义的，也就是JWT的规范，可以点进去看源码
	jwt.StandardClaims
}

// GetJWTSecret 获取该项目的 JWT Secret，目前我们是直接使用配置所配置的 Secret
func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

// GenerateToken 生成 JWT Token ，JWT的核心函数
func GenerateToken(appKey, appSecret string) (string, error) {
	// 传输传入客户端传入的 AppKey 和 AppSecret
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			// 获取项目配置中所设置的签发者（Issuer）和过期时间（ExpiresAt）
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}
	// 根据 Claims 结构体创建 Token 实例，参数为加密算法方案的枚举和Claims结构体
	// 书中说：Claims，主要是用于传递用户所预定义的一些权利要求，便于后续的加密、校验等行为。
	// 我的理解是 Claims 中可以自定义的增加一些字段，表示用户的权限内容
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成签名字符串，根据所传入 Secret 不同，进行签名并返回标准的 Token
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

// ParseToken 解析和校验 Token，与GenerateToken相对
// 其函数流程主要是解析传入的 Token，然后根据 Claims 的相关属性要求进行校验
func ParseToken(token string) (*Claims, error) {
	// 方法内部主要是具体的解码和校验的过程，最终返回 *Token。
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 函数传递，用来获取配置中的Secret
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		// Valid 验证基于时间的声明 例如：过期时间（ExpiresAt）、签发者（Issuer）、生效时间（Not Before）
		// 需要注意的是，如果没有任何声明在令牌中，仍然会被认为是有效的
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
