package mysqlx

import (
	"fmt"
	"go_web_demo/48_final_work/settings"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

var db *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", settings.ViperConfig.Mysql.Username,
		settings.ViperConfig.Mysql.Password,
		settings.ViperConfig.Mysql.Host,
		settings.ViperConfig.Mysql.Port,
		settings.ViperConfig.Mysql.Dbname,
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping()
	db.SetMaxOpenConns(settings.ViperConfig.Mysql.Maxopenconns)
	db.SetMaxIdleConns(settings.ViperConfig.Mysql.Maxidleconns)
	return
}
