package middleware

import (
	"portal/common"
	"github.com/dgrijalva/jwt-go"
)
// Parse token
func ParseToken(tokenString string) (*common.MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &common.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			return common.SECRET, nil
	})
	// validate
	if claims, ok := token.Claims.(*common.MyClaims); ok && token.Valid {
			// fmt.Printf("%v %v", claims.UserId, claims.StandardClaims.ExpiresAt)
			return claims, nil
	} else {
			// fmt.Println("valid fail: ", err)
			return nil, err
	}
}