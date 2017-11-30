package user

import (
	"test/db"
	"golang.org/x/crypto/bcrypt"
	"github.com/go-sql-driver/mysql"
	"fmt"
)

type User struct {
	Username string `db:"username",json:"username"`
	Password string `db:"password",json:"password"`
	UserId int `db:"userId"`
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

func CreateUser(user User) *errorUser {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	query := "INSERT INTO test.users (username, password) VALUES (:name,:pass)"
	res, err := db.Con.NamedExec(query, map[string]interface{}{
		"name": user.Username,
		"pass": string(hashedPassword),
	})
	fmt.Print(res)
	if err != nil {
		me, ok := err.(*mysql.MySQLError)
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
	query := "SELECT * FROM test.users WHERE username = \"" + user.Username + "\""
	err := db.Con.Get(&existUser, query)
	if err != nil {
		fmt.Println(err)
		return &errorUser{"ERROR", 500}
	}

	wrongPassword := bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(user.Password))
	if wrongPassword != nil { return &errorUser{"Wrong login or password", 401} }

	return nil
}
