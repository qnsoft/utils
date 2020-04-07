package DbHelper

import (
	"fmt"
	//"log"

	"github.com/astaxie/beego"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/qnsoft/utils/ErrorHelper"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
	"xorm.io/xorm/names"
)

func MsSqlDb() *xorm.Engine {
	var Engine *xorm.Engine
	var dbError error
	//数据库类型
	_type := beego.AppConfig.String("database_mssql::db_type")

	//数据库IP
	//数据库IP
	_server := beego.AppConfig.String("database_mssql::db_server")
	//数据库端口
	_port := beego.AppConfig.String("database_mssql::db_port")
	////数据库
	_database := beego.AppConfig.String("database_mssql::db_database")
	//数据库用户名
	_user := beego.AppConfig.String("database_mssql::db_user")
	//数据库密码
	_password := beego.AppConfig.String("database_mssql::db_password")
	//数据库表前缀
	_prefix := beego.AppConfig.String("database_mysql::db_prefix")
	if Engine != nil {
		return Engine
	}
	//连接字符串
	_connString := fmt.Sprintf("server=%s;port=%s;database=%s;user id=%s;password=%s", _server, _port, _database, _user, _password)
	Engine, dbError = xorm.NewEngine(_type, _connString)
	tbMapper := names.NewPrefixMapper(names.SnakeMapper{}, _prefix)
	Engine.SetTableMapper(tbMapper)
	//Engine.SetMaxIdleConns(50)
	//Engine.SetMaxOpenConns(200)
	Engine.ShowSQL(true)
	Engine.Logger().SetLevel(log.LOG_DEBUG)
	if dbError != nil {
		ErrorHelper.CheckErr(dbError)
		panic(dbError)
	}
	return Engine
}
