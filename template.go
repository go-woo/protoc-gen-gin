package main

import (
	"bytes"
	"strings"
	"text/template"
)

var routerTemplate = `
import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-woo/protoc-gen-gin/runtime"
)
{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}
{{$hasJwt := .HasJwt}}
func Register{{.ServiceType}}Router(r *gin.Engine) {
	{{- range .Methods}}
	{{- if .RequireToken}}
	r.{{.Method}}("{{.Path}}", JWTAuthMiddleware, _{{$svrType}}_{{.Name}}{{.Num}}_HTTP_Handler)
	{{- else}}
	r.{{.Method}}("{{.Path}}", _{{$svrType}}_{{.Name}}{{.Num}}_HTTP_Handler)
	{{end}}
	{{- end}}
}
{{range .Methods}}
func _{{$svrType}}_{{.Name}}{{.Num}}_HTTP_Handler(c *gin.Context) {
	if c.Request.ParseForm() != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "form data format error",
			"msg":    "",
		})
	}
	uv := c.Request.Form

	var req *{{.Request}} = new({{.Request}})
	{{- range .Fields}}
	uv.Add("{{.ProtoName}}", c.Param("{{.ProtoName}}"))
	{{- end}}
	if len(uv) > 0 {
		if err := runtime.BindValues(req, uv); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"reason": "request data format error",
				"msg":    "",
			})
		}
	}
	reply, err := {{$svrType}}{{.Name}}BusinessHandler{{.Num}}(req, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": "BusinessHandler error",
			"msg":    "",
		})
	}
	c.JSON(http.StatusOK, &reply)
}
{{end}}
`

var handlerTemplate = `
import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gin-gonic/gin"
)
{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}
{{$hasJwt := .HasJwt}}
{{range .Methods}}
func {{$svrType}}{{.Name}}BusinessHandler{{.HandlerNum}}(req *{{.Request}}, c *gin.Context) ({{.Reply}}, error) {
	{{- if .IsLogin}}
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "hello" && password == "world" {
		token, _ := genToken(MyCustomClaims{1, "hello", jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * time.Hour).Unix(),
			Issuer:    "hello",
		}})
		return {{.Reply}}{Token: "Bearer " + token}, nil
	} else {
		return {{.Reply}}{}, fmt.Errorf("username or password error")
	}
	{{- else}}
	// Here can put your business logic, can use ORM:github.com/go-woo/protoc-gen-ent
	// Below is example business logic code
	{{- if .RequireToken}}
	if u, ok := c.Get("userid"); ok {
		fmt.Printf("get userid %v ok", u)
	} else {
		return {{.Reply}}{}, fmt.Errorf("username or password error")
	}
	{{end}}
	return {{.Reply}}{}, nil{{end}}
}
{{end}}
`
var authTypeTemplate = `

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

type MyCustomClaims struct {
	Id       int    ` + "`json:\"id\"`" + `
	Username string ` + "`json:\"username\"`" + `
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
`

type serviceDesc struct {
	ServiceType string // Greeter
	ServiceName string // example.Greeter
	Metadata    string // example/v1/greeter.proto
	Methods     []*methodDesc
	MethodSets  map[string]*methodDesc
	LoginUrl    string
	HasJwt      bool
	//JwtRootPaths []*JwtRootPath
}

type JwtRootPath struct {
	RootPath string
}

type methodDesc struct {
	Name         string
	OriginalName string // The parsed original name
	Num          int
	HandlerNum   int
	Request      string
	Reply        string
	Path         string
	Method       string
	HasVars      bool
	HasBody      bool
	Body         string
	ResponseBody string
	Fields       []*RequestField
	LoginUrl     string
	RequireToken bool
	IsLogin      bool
}

type RequestField struct {
	ProtoName string
	GoName    string
	GoType    string
	ConvExpr  string
}

func (s *serviceDesc) execute(tpl string) string {
	s.MethodSets = make(map[string]*methodDesc)
	for _, m := range s.Methods {
		s.MethodSets[m.Name] = m
	}

	buf := new(bytes.Buffer)
	tmpl, err := template.New("http").Parse(strings.TrimSpace(tpl))
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return strings.Trim(buf.String(), "\r\n")
}
