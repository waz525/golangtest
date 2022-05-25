/*
CREATE TABLE `person` (
    `user_id` int(11) NOT NULL AUTO_INCREMENT,
    `username` varchar(260) DEFAULT NULL,
    `sex` varchar(260) DEFAULT NULL,
    `email` varchar(260) DEFAULT NULL,
    PRIMARY KEY (`user_id`)
  ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

CREATE TABLE place (
    country varchar(200),
    city varchar(200),
    telcode int
)ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

insert into person(username, sex, email)values("stu001", "man", "stu01@qq.com") ;
*/

package main

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
	"encoding/json"
)

type Person struct {
    UserId   int    `db:"user_id"`
    Username string `db:"username"`
    Sex      string `db:"sex"`
    Email    string `db:"email"`
}

type Place struct {
    Country string `db:"country"`
    City    string `db:"city"`
    TelCode int    `db:"telcode"`
}


func conn_mysql(username string, password string, host string, port string, tablespace string) *sqlx.DB {

	//database, err := sqlx.Open("mysql", "root:Sq_123456@tcp(127.0.0.1:3306)/sq_iptables")
    database, err := sqlx.Open("mysql", ""+username+":"+password+"@tcp("+host+":"+port+")/"+tablespace+"")
    if err != nil {
        fmt.Println("open mysql failed,", err)
        return nil
    }
    //defer database.Close()
    return database
    //defer Db.Close()  // 注意这行代码要写在上面err判断的下面
}

func main() {

	var Db =  conn_mysql("root","123456","127.0.0.1","3306","MobileOA")
	defer Db.Close()
    var person []Person
    //err := Db.Select(&person, "select user_id, username, sex, email from person where user_id = ?", 2)
    err := Db.Select(&person, "select user_id, username, sex, email from person")
    if err != nil {
        fmt.Println("exec failed, ", err)
        return
    }
    fmt.Println("select succ:", person)

	//转成json
	jsonBytes, err := json.Marshal(person)
        if err != nil {
                fmt.Println(err)
        }
        fmt.Println(string(jsonBytes))



}
