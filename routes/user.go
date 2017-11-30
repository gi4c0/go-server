package routes

import (
	"github.com/kataras/iris"
	"test/db/user"
	"github.com/dgrijalva/jwt-go"
	"fmt"
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

func ValidateToken(token string) (interface{}, error) {
	jwtDecoded, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if claims, ok := jwtDecoded.Claims.(jwt.MapClaims); ok && jwtDecoded.Valid {
		return claims["username"], nil
	} else {
		fmt.Println(err)
		return "", err
	}
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

func (u *UserController) PostAuth() (interface{}, int) {
	userTokenString := u.Ctx.GetHeader("authorization")
	token, err := ValidateToken(userTokenString)
	if err != nil {
		return "Your token wrong or expired", 401
	}

	return token, 200
}