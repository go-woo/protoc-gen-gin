// Auth use data type.
// versions:
// - protoc-gen-gin v0.0.1
// - protoc  v3.12.4
// source: example/v1/greeter.proto

package v1

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	entadapter "github.com/casbin/ent-adapter"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
	"strings"
)

type MyCustomClaims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var mySigningKey = []byte("dangerous")

var enforcer *casbin.Enforcer

func init() {
	cu := os.Getenv("CASBIN_URL")
	if cu == "" {
		cu = "root:123456@tcp(127.0.0.1:3306)/casbin"
	}
	a, err := entadapter.NewAdapter("mysql", cu)
	if err != nil {
		panic(err)
	}
	if enforcer, err = casbin.NewEnforcer("./rbac_with_domains_model.conf", a); err != nil {
		panic(err)
	}
}
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

func WooAuthMiddleware(c *gin.Context) {
	signToken := c.Request.Header.Get("Authorization")
	if signToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"reason": "Authorization can't null",
			"msg":    "",
		})
		return
	}
	myclaims, err := parserToken(signToken)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"reason": "Token is invalid",
			"msg":    "",
		})
		return
	}
	c.Set("userid", myclaims.Id)
	c.Set("username", myclaims.Username)
	if has, err := enforcer.Enforce(myclaims.Username, c.Request.RequestURI, c.Request.Method); err != nil || !has {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"reason": "casbin did not permission",
			"msg":    "forbidden",
		})
		return
	}

	c.Next()
}
