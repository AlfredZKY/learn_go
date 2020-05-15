package msyqls

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spaolacci/murmur3"
)

// 数据库配置
const (
	username = "root"
	password = "123456"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "loginserver" // 要连接的数据库名
)

// DB 数据库连接池
var db *gorm.DB

// 初始化init_driver
func init() {
	{
		// // 构建连接:用户名:密码@tcp(IP:端口)/数据库？chartset="utf8"
		// path := strings.Join([]string{username, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

		// // 打开数据库，前者是驱动名，所以要导入:_ "github.com/go-sql-driver/mysql"
		// DB, _ = sql.Open("mysql", path)

		// // 设置数据库的最大连接数
		// DB.SetConnMaxLifetime(100)

		// // 设置数据库最大闲置连接数
		// DB.SetMaxIdleConns(10)

		// // 验证连接
		// if err := DB.Ping(); err != nil {
		// 	fmt.Println("OPen database fail")
		// 	return
		// }
		// fmt.Println("connect success")
	}
	// db, err = gorm.Open("mysql", "hatlonely:keaiduo1@/hatlonely?charset=utf8&parseTime=True&loc=Local")
	path := strings.Join([]string{username, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	var err error
	db, err = gorm.Open("mysql", path)
	if err != nil {
		panic(err)
	}
	// if !db.HasTable(&View{}) {
	// 	if err := db.Set();err != nil {

	// 	}
	// }
}

// MysqlInsert insert data in to database
func MysqlInsert() bool {
	// 开启事务
	// tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return false
	}

	// 准备sql语句
	stmt, err := tx.Prepare("INSERT INTO nk_user (`name`,`password`) VALUES (?,?)")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	_ = stmt
	// // 设置参数以及执行sql语句
	// res,err := stmt.Exec()
	return false
}
