package routes

import (
	"github.com/kataras/iris"
	"test/db/user"
	"github.com/dgrijalva/jwt-go"
)

type UserController struct {
	iris.C
}

type Response struct {
	Token string `json:"token"`
}

func generateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})

	tokenString, err := token.SignedString([]byte("secret"))
	return tokenString, err
}

func (u *UserController) PostRegister() (interface{}, int) {
	var newUser user.User
	u.Ctx.ReadJSON(&newUser)

	newUserError := user.CreateUser(newUser)
	if newUserError != nil {
		return newUserError.GetError(), newUserError.GetCode()
	}

	token, err := generateToken(newUser.Username)
	if err != nil { return err, 500 }

	response := Response{Token: token}

	return response, 200
}

func (u *UserController) PostLogin() (interface{}, int) {
	var existUser user.User
	u.Ctx.ReadJSON(&existUser)

	err := user.VerifyUser(existUser)

	if err != nil {
		return err.GetError(), err.GetCode()
	}

	token, tokenErr := generateToken(existUser.Username)
	if tokenErr != nil { return tokenErr, 500 }

	response := Response{Token: token}

	return response, 200
}