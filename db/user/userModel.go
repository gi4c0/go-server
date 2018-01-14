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
	Permission string
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

func GenerateToken(username, permission string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"permission": permission,
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
		fmt.Println(err)
		return false, ""
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, claims["username"]
	}

	return false,""
}

func CreateUser(user User) *errorUser {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return &errorUser{err.Error(), 500}
	}

	stmt, stmtErr := db.Con.Prepare("INSERT INTO test.Users (Username, Password) VALUES (?, ?)")
	if stmtErr != nil {
		return &errorUser{stmtErr.Error(), 500}
	}

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

func VerifyUser(user *User) *errorUser {
	var dbUser User

	err := db.Con.QueryRow("SELECT * FROM test.Users WHERE Username = ?", user.Username).Scan(&dbUser.UserId, &dbUser.Username, &dbUser.Password, &dbUser.Token, &user.Permission)
	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			return &errorUser{"Wrong username", 401}
		}
		return &errorUser{err.Error(), 500}
	}

	wrongPassword := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if wrongPassword != nil { return &errorUser{"Wrong login or password", 401} }

	return nil
}

func GetUserId (token string) (int, string) {
	validToken, username := VerifyToken(token)
	if !validToken {
		return -1, ""
	}

	var userId int
  var permission string
	err := db.Con.QueryRow("SELECT UserId, Permission From test.Users WHERE Username = ?", username).Scan(&userId, &permission)
	if err != nil {
		fmt.Println(err)
		return -1, ""
	}

	return userId, permission
}

func VerifyPermission(token string, perm string) (int, string) {
	validToken, username := VerifyToken(token)
	if !validToken {
		return -1, ""
	}

	var userId int
	var permission string
	err := db.Con.QueryRow("SELECT UserId, Permission From test.Users WHERE Username = ?", username).Scan(&userId, &permission)
	if err != nil {
		fmt.Println(err)
		return -1, ""
	}

  if permission != perm && permission != "admin" {
    return -1, ""
  }

	return userId, permission
}

func CheckUsername (username string) bool {
	var exist int
	err := db.Con.QueryRow("SELECT COUNT(UserId) FROM test.Users WHERE Username = ?", username).Scan(&exist)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if exist == 1 {
		return false
	}

	return true
}

