package driver

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

// 声明一个全局变量db
var db *sql.DB

func TestMySQLDriver(t *testing.T) {

	dsn := "root:123456@tcp(127.0.0.1:3306)/whig"
	// 打开连接
	db, err := sql.Open("mysql", dsn) //这里边就是注册呢
	if err != nil {
		panic(err)
	}

	// 初始化连接，就是检查用户名密码等等是否正确
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	row := db.QueryRow("select * from user_info")

	fmt.Println(row)
	defer db.Close()

}

// 插入数据
func insertData() {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, "王五", 38)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}
