package user

import (
	"go-server/db"
	"golang.org/x/crypto/bcrypt"
	"github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"database/sql"
)

type User struct {
	Username string
	Password string
	UserId int
	Token sql.NullString
}

type errorUser struct {
	err string
	code int
}

func (e *errorUser) Error() string {
	return fmt.Sprintf("Error %d: %s", e.code, e.err)
}

func (e *errorUser) GetError() string {
	return e.err
}

func (e *errorUser) GetCode() int {
	return e.code
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func GenerateToken(username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})

	tokenString, _ := token.SignedString([]byte("secret"))

	return tokenString
}

func VerifyToken (tokenString string) (bool, interface{}) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if err != nil {
		return false, ""
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, claims["username"]
	}

	return false,""
}

func CreateUser(user User) *errorUser {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	checkError(err)

	stmt, stmtErr := db.Con.Prepare("INSERT INTO test.Users (Username, Password) VALUES (?, ?)")
	checkError(stmtErr)

	_, resErr := stmt.Exec(user.Username, hashedPassword)

	if resErr != nil {
		me, ok := resErr.(*mysql.MySQLError)
		if !ok {
			return &errorUser{"Something went wrong", 500}
		}

		if me.Number == 1062 {
			return &errorUser{"This username is already exist", 400}
		}
	}

	return nil
}

func VerifyUser(user User) *errorUser {
	var existUser User

	err := db.Con.QueryRow("SELECT * FROM test.Users WHERE Username = ?", user.Username).Scan(&existUser.UserId, &existUser.Username, &existUser.Password, &existUser.Token)
	checkError(err)

	wrongPassword := bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(user.Password))
	if wrongPassword != nil { return &errorUser{"Wrong login or password", 401} }

	return nil
}