package common

//实现用户认真和维持用户登陆状态
import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"lzh.practice/ginessential/model"
)

//加密密钥
//用来hash
var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	//token的有效时间
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		//gorm里有id
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			//token的发放时间
			IssuedAt: time.Now().Unix(),
			//发放人
			Issuer: "lzh.practice",
			//主题
			Subject: "user token",
		},
	}
	//HS256对称加密 RS256非对称加密
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//从tokenstring中解析claim然后返回
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
