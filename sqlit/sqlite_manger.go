package main

import (
    "database/sql"
	"fmt"
	"time"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
	// 打开/创建
	db, err := sql.Open("sqlite3", "./my.db")

	// 关闭
	defer db.Close()


	table := `
	    CREATE TABLE IF NOT EXISTS user (
		    uid INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(128) NULL,
	        created DATE NULL
		);
		`
	_, err = db.Exec(table)
	if err != nil {
		panic(err)
		return ;
	}

	fmt.Println("CREATE TABLE Success!");


    stmt, err := db.Prepare("INSERT INTO user(name,  created) values(?,?)")
    if err != nil {
        panic(err)
    }
    // res 为返回结果
    res, err := stmt.Exec("guoke", "2012-12-09")
    if err != nil {
        panic(err)
    }

    // 可以通过res取自动生成的id
    id, err := res.LastInsertId()
    if err != nil {
        panic(err)
    }
	fmt.Println(id);

	rows, err := db.Query("SELECT * FROM user")
    if err != nil {
          panic(err)
     }
    defer rows.Close()

    for rows.Next() {
        var uid int
        var name string
        var created time.Time
        err = rows.Scan(&uid, &name,  &created)
        if err != nil {
          panic(err)
        }

        fmt.Println(uid)
        fmt.Println(name)
        fmt.Println(created)
    }
}
