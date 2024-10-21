package accounts

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	Username string `example:"username"`
	Password string `example:"password"`
	Email    string `example:"email"`
}

func CreateAccount(acc Account) {

}

func DeleteAccount(acc Account) {

}

func AccessAccount(c *gin.Context) {
	var acc Account
	if err := c.BindJSON(&acc); err != nil {
		fmt.Println("Illegal access attempt")
	}

	if nil != bcrypt.CompareHashAndPassword(nil, []byte(acc.password)) {
		fmt.Println("Failed login attempt")
	}
}

func UpdateAccount(acc Account) {

}
