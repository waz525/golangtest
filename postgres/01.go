package main

import (
	//"database/sql" //通用的接口
	"github.com/jmoiron/sqlx"
	"fmt"
	_ "github.com/bmizerany/pq" //必须要有相应的驱动
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""			//你自己数据库的密码
	dbname   = "genericdb" //创建的数据库
)


var db *sqlx.DB // 连接池对象
var err error

type product struct {
	ProductNo string
	Name      string
	Price     float64
}

type productDao struct{}


func main() {
	pd := new(productDao)
	err = pd.initDB()
	if err != nil {
		fmt.Println("initDB() failed. ")
	}
	defer pd.closeDB()
	pd.doQueryAll()
	//pd.doPreQueryByName("apple")
}



func (pd *productDao) initDB() (err error) {
	pdqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	//  pdqlInfo := "postgres:密码@tcp(127.0.0.1:5432)/test" // 用户名:密码@tcp(ip端口)/数据库名字，暂时出错
	db, err = sqlx.Open("postgres", pdqlInfo) //Open(driverName 驱动名字, dataSourceName string 数据库信息)
	// DB 代表一个具有零到多个底层连接的连接池，可以安全的被多个go程序同时使用
	//这里的open函数只是验证参数是否合法，而不会创建和数据库的连接,也不会检查账号密码是否正确
	if err != nil {
		fmt.Println("Wrong args.Connected failed.")
		return err
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Connected failed.")
		return err
	}
	db.SetMaxOpenConns(20) //设置数据库连接池最大连接数
	db.SetMaxIdleConns(10) //设置最大空闲连接数
	fmt.Println("Successfully connected!")
	return nil
}


func (pd *productDao) closeDB() (err error) {
	return db.Close()
}

func (pd *productDao) doQueryAll() (error, []product) {
	rows, err := db.Query(`Select * from product`)
	if err != nil {
		fmt.Println("Some amazing wrong happens in the process of Query.", err)
		return err, []product{}
	}
	products := make([]product, 0)
	defer rows.Close() //关闭连接
	index := 0
	var p product
	for rows.Next() {
		err := rows.Scan(&p.ProductNo, &p.Name, &p.Price)
		products = append(products, p)
		if err != nil { // 获得的都是字符串
			fmt.Println("Some amazing wrong happens in the process of queryAll.", err)
			return err, products
		}
		index++
	}
	if index > 0 {
		fmt.Println("The data of table is as follow.")
		for _, p := range products {
			fmt.Printf("%v %s %v\n", p.ProductNo, p.Name, p.Price)
		}
		fmt.Println("Successfully query ", len(products))
		return nil, products
	} else {
		fmt.Println("No such data exists in database. ")
		return fmt.Errorf("No such data exists in database. "), products
	}
}



