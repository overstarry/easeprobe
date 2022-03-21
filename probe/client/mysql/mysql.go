package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/megaease/easeprobe/probe/client/conf"
)

// Kind is the type of driver
const Kind string = "MySQL"

// MySQL is the MySQL client
type MySQL struct {
	conf.Options `yaml:",inline"`
	ConnStr      string `yaml:"conn_str"`
}

// New create a Redis client
func New(opt conf.Options) MySQL {

	var conn string
	if len(opt.Password) > 0 {
		conn = fmt.Sprintf("%s:%s@tcp(%s)/?%s",
			opt.Username, opt.Password, opt.Host, opt.Timeout.Round(time.Second))
	} else {
		conn = fmt.Sprintf("%s@tcp(%s)/?%s",
			opt.Username, opt.Host, opt.Timeout.Round(time.Second))
	}

	return MySQL{
		Options: opt,
		ConnStr: conn,
	}
}

// Kind return the name of client
func (r MySQL) Kind() string {
	return Kind
}

// Probe do the health check
func (r MySQL) Probe() (bool, string) {

	db, err := sql.Open("mysql", r.ConnStr)
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return false, err.Error()
	}
	_, err = db.Query("show status") // run a SQL to test
	if err != nil {
		return false, err.Error()
	}

	return true, "Check MySQL Server Successfully!"

}