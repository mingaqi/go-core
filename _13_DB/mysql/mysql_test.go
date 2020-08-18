package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // mysql init()函数向sql规范接口中注册了MySQL驱动
	"testing"
	"time"
)

// Go语言中的database/sql包提供了保证SQL或类SQL数据库的泛用接口，并不提供具体的数据库驱动。
// db.query 有SQL注入风险
// db.prepare() stmt.query 可以避免注入
// db.Begin() tx.commit() tx.rollback() 使用事物
// DB内部实现连接池, 并且线程安全所以使用全局变量
var db *sql.DB

func TestConnect(t *testing.T) {
	err := initDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("数据库链接成功")

	// 查询所有表 包括DML操作也使用query Exec 增删改
	rows, _ := db.Query("show tables")
	var tableName string
	for rows.Next() {
		// 非常重要：确保Query之后调用Scan方法，否则持有的数据库链接不会被释放
		rows.Scan(&tableName)
		fmt.Println(tableName)
	}
	defer db.Close()
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

}

func TestUpdate(t *testing.T) {
	initDB()
	p := Dept{3, "品管部", "couldDB1"}
	affect, err := update(&p)
	if err != nil {
		fmt.Println("update fialed: error:", err)
	}
	fmt.Println("受影响行数:", affect)
	defer db.Close()
}

// 预处理
// 1. 把SQL语句分成两部分，命令部分与数据部分。
// 2. 先把命令部分发送给MySQL服务端，MySQL服务端进行SQL预处理。
// 3. 然后把数据部分发送给MySQL服务端，MySQL服务端对SQL语句进行占位符替换。
// 4. MySQL服务端执行完整的SQL语句并将结果返回给客户端。
func TestStmt(t *testing.T) {
	initDB()

	stmt()
}

func initDB() (err error) {
	// open函数不校验数据库账号和密码, 只做连接串的格式校验
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/cloudDB01")
	if err != nil {
		return
	}
	err = db.Ping() //尝试与数据库建立链接
	if err != nil {
		return
	}
	db.SetMaxOpenConns(10) // 最大连接数
	db.SetMaxIdleConns(5)  // 最大活跃连接数
	db.SetConnMaxLifetime(30 * time.Second)
	return
}

func update(dept *Dept) (affect int64, err error) {
	sql := "update dept set deptName = ? where deptNo=?"
	result, err := db.Exec(sql, dept.DeptName, dept.Id)
	if err != nil {
		return 0, err
	}
	affect, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return
}

func stmt() (err error) {
	db.Begin()
	stmt, err := db.Prepare("select * from dept where deptNo>?")
	if err != nil {
		return err
	}
	rows, err := stmt.Query(0)
	if err != nil {
		return err
	}
	for rows.Next() {
		dept := Dept{}
		rows.Scan(&dept.Id, &dept.DeptName, &dept.DeptSource)
		fmt.Println(dept)
	}
	return nil
}

type Dept struct {
	Id         int
	DeptName   string
	DeptSource string
}
