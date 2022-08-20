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
	r.{{.Method}}("{{.Path}}", WooAuthMiddleware, _{{$svrType}}_{{.Name}}{{.Num}}_HTTP_Handler)
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
		return
	}
	uv := c.Request.Form
	var req *{{.Request}} = new({{.Request}})
	{{- range .Vars}}
	uv.Add("{{.ProtoName}}", c.Param("{{.ProtoName}}"))
	{{- end}}
	if len(uv) > 0 {
		if err := runtime.BindValues(req, uv); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"reason": "request data format error",
				"msg":    "",
			})
			return
		}
	}
	reply, err := {{$svrType}}{{.Name}}BusinessHandler{{.Num}}(req, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": "BusinessHandler error",
			"msg":    "",
		})
		return
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
	{{- if .RequireToken}}
	//u, _ := c.Get("username")	// username 
	//t, _ := c.Get("domain")	// tenant {{end}}
	// Here can put your business logic, can use ORM:github.com/go-woo/protoc-gen-ent
	// INSERT_POINT: DO NOT DELETE THIS LINE!

	return {{.Reply}}{}, nil{{end}}
}
{{end}}
`
var authTypeTemplate = `

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
	Domain   int    ` + "`json:\"domain\"`" + `
	Username string ` + "`json:\"username\"`" + `
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
	c.Set("domain", myclaims.Domain)
	c.Set("username", myclaims.Username)
	if has, err := enforcer.Enforce(myclaims.Username, myclaims.Domain, c.Request.RequestURI, c.Request.Method); err != nil || !has {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"reason": "casbin did not permission",
			"msg":    "forbidden",
		})
		return
	}

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
	Vars         []*PathVar
	LoginUrl     string
	RequireToken bool
	IsLogin      bool
}

type PathVar struct {
	ProtoName string
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
