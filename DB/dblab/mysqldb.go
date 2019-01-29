package dblab

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golab")
	if err != nil {
		return nil, fmt.Errorf("connect mysql err: %+v", err.Error())
	}

	return db, nil
}

// InsertUserInfo ...
func InsertUserInfo() error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	res, err := db.Exec(
		"INSERT INTO la_user (name, age) VALUES (?, ?)",
		"gopher",
		27,
	)
	if err != nil {
		return fmt.Errorf("Insert User error: %+v", err.Error())
	}

	fmt.Printf("res: %+v", res)

	return nil
}

// UserInfo ...
type UserInfo struct {
	ID   string
	Name string
	Age  uint32
}

// SelectUserInfo ...
func SelectUserInfo() ([]*UserInfo, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select user_id, name, age from la_user")
	if err != nil {
		return nil, fmt.Errorf("selectUserInfo error: %+v", err.Error())
	}

	userInfos := make([]*UserInfo, 0)
	for rows.Next() {
		var info UserInfo
		if err := rows.Scan(
			&info.ID,
			&info.Name,
			&info.Age,
		); err != nil {
			return nil, fmt.Errorf("selectUserInfo error: %v", err.Error())
		}

		userInfos = append(userInfos, &info)
	}

	return userInfos, nil
}

// InsertIncludeThings ...
func InsertIncludeThings() error {
	db, err := connectDB()
	if err != nil {
		return fmt.Errorf("insertInclueTings err: %v", err.Error())
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("insertInclueTings err: %v", err.Error())
	}

	if _, err = tx.Exec("INSERT INTO la_user (name, age) VALUES (?, ?)", "gopher", 27); err != nil {
		tx.Rollback()
		return fmt.Errorf("insertInclueTings err: %v", err.Error())
	}

	if _, err = tx.Exec("INSERT INTO la_company (company_name, addr) VALUES (?, ?)", "usce", "sh"); err != nil {
		tx.Rollback()
		return fmt.Errorf("insertInclueTings err: %v", err.Error())
	}

	err = tx.Commit()

	return nil
}
