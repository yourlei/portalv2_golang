package common
import (
	"time"
	// "fmt"

	"github.com/dgrijalva/jwt-go"
)

var SECRET = []byte("dcfcd07e645d245babe887e5e2daa016")
type MyClaims struct {
    UserId string `json:"userId,omitempty"`
    RoleId string `json:"roleId,omitempty"`
    jwt.StandardClaims
}
// Generate Token
func GenerateToken(userId, roleId string) (string, error) {
	// 失效时间
  expireTime := time.Now().Add(1 * time.Hour)
	// Create the Claims
	claims := MyClaims{
		userId,
		roleId,
		jwt.StandardClaims {
			ExpiresAt: expireTime.Unix(),
			Issuer:    "yourlin127@gmail.com",
		},
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(SECRET)
	// IF error
	if err != nil {
		return "",err
	}
	return ss, nil
}