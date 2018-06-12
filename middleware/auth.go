package middleware

import (
	"net/http"
	"strconv"
	"fmt"

	"portal/config"
	"portal/util"

	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)

// claims
type MyClaims struct {
    UserId string `json:"userId,omitempty"`
    RoleId string `json:"roleId,omitempty"`
    jwt.StandardClaims
}
// Parse token
func parseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.AppConfig.TokenSecrect), nil
	})
	// validate
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			fmt.Printf("%v %v", claims.UserId, claims.StandardClaims.ExpiresAt)
			return claims, nil
	} else {
			fmt.Println("valid fail: ", err)
			return nil, err
	}
}
// require user sigin
func SigninRequired(c *gin.Context) {
	token, err := c.Request.Cookie("token")
	
	if err != nil {
		util.RequireSignin(c)
		return
	}
	claims, err := parseToken(token.Value)
	fmt.Println(claims)
	if err != nil {
		fmt.Println("error: ",err)
		util.RequireSignin(c)
		return
	}

	c.Next()
}
// 管理员登录
func AdminRequired(c *gin.Context) {
	var (
		token  *http.Cookie
		claims *MyClaims
		err    error
	)
	// Get cookie
	token, err = c.Request.Cookie("token")
	fmt.Println(token, err)
	if err != nil {
		util.RequireSignin(c)
		return
	}
	// Parse Token
	claims, err = parseToken(token.Value)
	if err != nil {
		fmt.Println("error: ",err)
		util.RequireSignin(c)
		return
	}
	role, decryptErr := util.Decrypt([]byte(config.AppConfig.AesKey), claims.RoleId)
	if decryptErr != nil {
		util.RequireSignin(c)
		return
	}
	roleID, _ := strconv.Atoi(role)
	// 非管理员角色
	if roleID != 1 {
		util.RequireSignin(c)
		return
	} 
	c.Next()
}