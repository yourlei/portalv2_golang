package test

import (
	// "fmt"
	// "time"
	"testing"
	// "fmt"
	// "portal/middleware"
	// "github.com/dgrijalva/jwt-go"
)

func TestToken(t *testing.T) {
	// var (
	// 	userId = "1"
	// 	roleId = "1"
	// )
	// ss, err := common.GenerateToken(userId, roleId)
	// if (err != nil ) {
	// 	t.Error("Error: ", err)
	// } else {
	// 	t.Log("success: ", ss)
	// }
	// time.Sleep(11 * time.Second)
	// ss := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxIiwicm9sZUlkIjoiMSIsImV4cCI6MTUyODQ1NDI1OSwiaXNzIjoieW91cmxpbjEyN0BnbWFpbC5jb20ifQ.q1v2NsuPdLtaj2uHruMjDmpPJp0iaszepNm7aQWpE6s"
	// ss := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Mjg3MDI2NTcsImlzcyI6InlvdXJsaW4xMjdAZ21haWwuY29tIn0.r4a5pPovjlbjARMIEZ7Js3P0yuUe3SbyrYH8h8sHgPY"
	// Claims, err1 := middleware.ParseToken(ss)

	// if (err1 != nil ) {
	// 	t.Error("Error: ", err1)
	// }
	// fmt.Println(Claims.UserId)

	// sample token is expired.  override time so it parses as valid
	// at(time.Unix(0, 0), func() {
	// token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
	// 		return []byte("AllYourBase"), nil
	// })

	// if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
	// 		fmt.Printf("%v %v", claims.Foo, claims.StandardClaims.ExpiresAt)
	// } else {
	// 		fmt.Println(err)
	// }
	// })
}