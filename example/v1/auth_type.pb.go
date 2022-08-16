// Auth use data type.
// versions:
// - protoc-gen-gin v0.0.1
// - protoc  v3.12.4
// source: example/v1/greeter.proto

package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

type MyCustomClaims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var mySigningKey = []byte("dangerous")

func genToken(claims MyCustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signToken, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return signToken, nil

}

func parserToken(signToken string) (*MyCustomClaims, error) {
	t := strings.Split(signToken, " ")
	if len(t) != 2 || t[0] != "Bearer" {
		return nil, fmt.Errorf("token format invalid")
	}
	var claims MyCustomClaims
	token, err := jwt.ParseWithClaims(t[1], &claims, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if token.Valid {
		return &claims, nil
	} else {
		return nil, err
	}
}

func JWTAuthMiddleware(c *gin.Context) {
	signToken := c.Request.Header.Get("Authorization")
	if signToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"reason": "Authorization can't null",
			"msg":    "",
		})
		c.Abort()
		return
	}
	myclaims, err := parserToken(signToken)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"reason": "Token is invalid",
			"msg":    "",
		})
		c.Abort()
		return
	}
	c.Set("userid", myclaims.Id)
	c.Next()
}
