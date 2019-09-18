package db

import (
	"database/sql"
	"fmt"
	"github.com/aasumitro/go-learn/entity"
	_ "github.com/go-sql-driver/mysql"
)

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/example_db")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func SqlQuerySelect() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from examples")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []entity.Example

	for rows.Next() {
		var each = entity.Example{}
		var err = rows.Scan(&each.Id, &each.Title, &each.Description)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.Title)
	}
}

func SqlQuerySelectRow() {
	var db, err = connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var result = entity.Example{}
	var id = 1
	err = db.
		QueryRow("select title, description from examples where id = ?", id).
		Scan(&result.Title, &result.Description)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("title: %s\ndescription: %s\n", result.Title, result.Description)
}

func SqlQueryPrepare() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("select title, description from examples where id = ?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result1 = entity.Example{}
	stmt.QueryRow(1).Scan(&result1.Title, &result1.Description)
	fmt.Printf("title: %s\ndescription: %s\n", result1.Title, result1.Description)

	var result2 = entity.Example{}
	stmt.QueryRow(2).Scan(&result2.Title, &result2.Description)
	fmt.Printf("title: %s\ndescription: %s\n", result2.Title, result2.Description)

}

func SqlExec() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("insert into examples values (?, ?, ?)", nil, "ASd", "qwe")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("insert success!")

	_, err = db.Exec("update examples set description = ? where id = ?", 2, "asd qwe")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("update success!")

	_, err = db.Exec("delete from examples where id = ?", 1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("delete success!")
}

func FetchExampleData() []entity.Example {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer db.Close()

	rows, err := db.Query("select * from examples")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer rows.Close()

	var result []entity.Example

	for rows.Next() {
		var each = entity.Example{}
		var err = rows.Scan(&each.Id, &each.Title, &each.Description)

		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return result

}
