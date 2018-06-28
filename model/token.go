// Define token
package model

import (
	"github.com/dgrijalva/jwt-go"
)
// claims
type MyClaims struct {
	UserId string `json:"userId,omitempty"` // 用户id
	RoleId string `json:"roleId,omitempty"` // 角色id
	jwt.StandardClaims                      // jwt标准对象
}