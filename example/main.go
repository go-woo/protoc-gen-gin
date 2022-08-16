package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
	"time"
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
func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	RegisterRouter(r)
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("Server starting error:%v\n", err)
	}
}

func RegisterRouter(r *gin.Engine) {
	r.POST("/login", Login)

	r.GET("/secure", JWTAuthMiddleware, Secure)

	r.GET("/secure/health", Health)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "hello" && password == "world" {
		token, _ := genToken(MyCustomClaims{1, "hello", jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * time.Hour).Unix(),
			Issuer:    "hello",
		}})
		c.JSON(http.StatusOK, gin.H{"token": "Bearer " + token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"reason": "Username or password error",
			"msg":    "",
		})
	}
}

func Secure(c *gin.Context) {
	if u, ok := c.Get("userid"); ok {
		c.String(http.StatusOK, "/secure Got user %v OK!", u)
	} else {
		c.Status(http.StatusInternalServerError)
	}
}

func Health(c *gin.Context) {
	c.String(http.StatusOK, "/secure/health OK!")
}
