// The business logic.
// versions:
// - protoc-gen-gin v0.0.1
// - protoc  v3.12.4
// source: example/v1/greeter.proto

package v1

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GreeterSayHelloBusinessHandler0(req *HelloRequest, c *gin.Context) (HelloReply, error) {
	// Here can put your business logic, can use ORM:github.com/go-woo/protoc-gen-ent
	// //INSERT_POINT: DO NOT DELETE THIS LINE!
	return HelloReply{}, nil
}

func GreeterCreateUserBusinessHandler0(req *CreateUserRequest, c *gin.Context) (CreateUserReply, error) {
	// Here can put your business logic, can use ORM:github.com/go-woo/protoc-gen-ent
	// //INSERT_POINT: DO NOT DELETE THIS LINE!
	return CreateUserReply{}, nil
}

func GreeterLoginBusinessHandler0(req *LoginRequest, c *gin.Context) (LoginReply, error) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "hello" && password == "world" {
		token, _ := genToken(MyCustomClaims{1, "hello", jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * time.Hour).Unix(),
			Issuer:    "hello",
		}})
		return LoginReply{Token: "Bearer " + token}, nil
	} else {
		return LoginReply{}, fmt.Errorf("username or password error")
	}
}

func GreeterUpdateUserBusinessHandler0(req *UpdateUserRequest, c *gin.Context) (UpdateUserReply, error) {
	if u, ok := c.Get("userid"); ok {
		fmt.Printf("get userid %v ok", u)
	} else {
		return UpdateUserReply{}, fmt.Errorf("username or password error")
	}

	// Here can put your business logic, can use ORM:github.com/go-woo/protoc-gen-ent
	// //INSERT_POINT: DO NOT DELETE THIS LINE!
	return UpdateUserReply{}, nil
}

func GreeterDeleteUserBusinessHandler0(req *UserRequest, c *gin.Context) (UserReply, error) {
	if u, ok := c.Get("userid"); ok {
		fmt.Printf("get userid %v ok", u)
	} else {
		return UserReply{}, fmt.Errorf("username or password error")
	}

	// Here can put your business logic, can use ORM:github.com/go-woo/protoc-gen-ent
	// //INSERT_POINT: DO NOT DELETE THIS LINE!
	return UserReply{}, nil
}

func GreeterListUsersBusinessHandler0(req *UserRequest, c *gin.Context) (UserReplys, error) {
	// Here can put your business logic, can use ORM:github.com/go-woo/protoc-gen-ent
	// //INSERT_POINT: DO NOT DELETE THIS LINE!
	return UserReplys{}, nil
}
