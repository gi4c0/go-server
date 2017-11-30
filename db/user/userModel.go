package user

import (
	"test/db"
	"golang.org/x/crypto/bcrypt"
	"github.com/go-sql-driver/mysql"
	"fmt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
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
