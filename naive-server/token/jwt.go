package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("123")

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

const expire_time = 30 * time.Minute

// 生成token
func GenerateToken(username, password string) string {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //设置过期时间
			Issuer:    "admin",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := tokenClaims.SignedString(jwtSecret)
	return token
}

// 验证token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, e := tokenClaims.Claims.(*Claims); e && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
