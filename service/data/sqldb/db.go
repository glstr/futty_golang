package sqldb

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	ErrNotFound = errors.New("not found")
)

type SqlDBType string

const (
	MysqlDBType     = "mysql"
	CockroachDBType = "cockroach"
)

type InitFunc func(option *DBOption) (*sql.DB, error)

var Type2InitFunc = map[SqlDBType]InitFunc{
	MysqlDBType:     InitMysql,
	CockroachDBType: InitCockroach,
}

type DBOption struct {
	Type SqlDBType

	User     string
	Password string
	Host     string
	Port     string
	DBName   string

	ConnMaxLifetimeMs int64
	MaxOpenConns      int64
	MaxIdleConns      int64
}

func GetDB(option *DBOption) (*sql.DB, error) {
	if initF, ok := Type2InitFunc[option.Type]; ok {
		return initF(option)
	}
	return nil, ErrNotFound
}

// example:"snow:301025@tcp(127.0.0.1:3306)/snow"
func makeMysqlUrl(option *DBOption) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		option.User,
		option.Password,
		option.Host,
		option.Port,
		option.DBName)
}

func InitMysql(option *DBOption) (*sql.DB, error) {
	url := makeMysqlUrl(option)
	db, err := sql.Open("mysql", url)
	if err != nil {
		return db, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, nil
}

// urlExample := "postgres://username:password@localhost:5432/database_name"
func makeCockroachUrl(option *DBOption) string {
	str := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		option.User,
		option.Password,
		option.Host,
		option.Port,
		option.DBName,
	)

	return str
}

func InitCockroach(option *DBOption) (*sql.DB, error) {
	url := makeCockroachUrl(option)
	return sql.Open("pgx", url)
}
