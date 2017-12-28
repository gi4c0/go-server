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

func Change(oldName, newName string) error {
  _, updateCategoryErr := db.Con.Exec("UPDATE test.Categories SET Name = ? WHERE Name = ?", newName, oldName)
  if updateCategoryErr != nil {
    fmt.Println(updateCategoryErr)
    return updateCategoryErr
  }

  _, updateArticleErr := db.Con.Exec("UPDATE test.Articles SET Category = ? WHERE Category = ?", newName, oldName)
  if updateArticleErr != nil {
    fmt.Println(updateArticleErr)
    return updateArticleErr
  }

  return nil
}
