package main

import (
	"fmt"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)


var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/MobileOA?charset=utf8")

	database.SetMaxOpenConns(20)
	database.SetMaxIdleConns(10)
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
}

func checkErr(err error) {
	if err != nil {
			panic(err)
	}
}


//插入操作
func insert() {
	stmt, err := Db.Prepare(`insert into person(username, sex, email)values(?, ?, ?)`)
	checkErr(err)
	res, err := stmt.Exec("stu002", "man", "stu02@qq.com")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}

func query1() {
	rows, err := Db.Query("SELECT * FROM person")
	checkErr(err)
	//普通demo
	for rows.Next() {
		var user_id int
		var username string
		var sex string
		var email string

		rows.Columns()
		err = rows.Scan(&user_id, &username, &sex, &email)
		checkErr(err)

		//fmt.Println(user_id, username, sex, email)
	}
}

func query2() {
	//rows, err := Db.Query("SELECT * FROM user")
	rows, err := Db.Query("SELECT * FROM person")
	checkErr(err)

	//字典类型
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
}

//更新数据
func update() {
	stmt, err := Db.Prepare(`UPDATE person SET username = ?, sex = ?, email = ? WHERE username=?`)
	checkErr(err)
	res, err := stmt.Exec("stu003", "man", "stu03@qq.com", "stu002")
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}

//删除数据
func remove() {
	stmt, err := Db.Prepare(`DELETE FROM person WHERE username=?`)
	checkErr(err)
	res, err := stmt.Exec("stu002")
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}

func getJSON(sqlString string) (string, error) {
    stmt, err := Db.Prepare(sqlString)
    if err != nil {
        return "", err
    }
    defer stmt.Close()
    rows, err := stmt.Query()
    if err != nil {
        return "", err
    }
    defer rows.Close()
    columns, err := rows.Columns()
    if err != nil {
      return "", err
    }
    count := len(columns)
    tableData := make([]map[string]interface{}, 0)
    values := make([]interface{}, count)
    valuePtrs := make([]interface{}, count)
    for rows.Next() {
      for i := 0; i < count; i++ {
          valuePtrs[i] = &values[i]
      }
      rows.Scan(valuePtrs...)
      entry := make(map[string]interface{})
      for i, col := range columns {
          var v interface{}
          val := values[i]
          b, ok := val.([]byte)
          if ok {
              v = string(b)
          } else {
              v = val
          }
          entry[col] = v
      }
      tableData = append(tableData, entry)
    }
    jsonData, err := json.Marshal(tableData)
    if err != nil {
      return "", err
    }
    fmt.Println(string(jsonData))
    return string(jsonData), nil
}

func ExecSql(sql string) string {
    stmt,_ := Db.Prepare(sql)
    rst, err := stmt.Exec()
    if err != nil {
        return "{\"ERROR\":\"Exec sql ERROR\"}"
    }
    num, err := rst.RowsAffected()
    if err != nil {
        return "{\"ERROR\":\"rst.RowsAffected ERROR\"}"
    }
	fmt.Println("num:" , num)
    //return "{\"LINENUM\":\""+Itoa(num)+"\"}"
	return ""
}


func main() {
	insert()
	fmt.Println("========================================")
	query1()
		/*
	query2()
	update()
	remove()
	query2()
*/
//	ExecSql("update xxxtestxxx set mobile='15088889999' where id ='ce609f4077f04169a9211bc0'")
	rst, _ := getJSON("select * from person")
	fmt.Println(rst)
}



