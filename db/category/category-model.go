package category

import (
  "go-server/db"
  "fmt"
)

type Category struct {
	Name string
}

func GetAll() ([]string, error) {
	var categories []string
	res, err := db.Con.Query("SELECT * FROM test.Categories")
	if err != nil {
		fmt.Println(err)
		return categories, err
	}

	for res.Next() {
		var category string

		scanErr := res.Scan(&category)
		if scanErr != nil {
			fmt.Println(scanErr)
			return categories, scanErr
		}

		categories = append(categories, category)
	}

	return categories, nil
}
