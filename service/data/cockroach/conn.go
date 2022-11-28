package cockroach

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type ConnOption struct {
	Host     string
	Port     string
	DBName   string
	User     string
	Password string
}

type Conn struct {
	conn *pgx.Conn
}

// urlExample := "postgres://username:password@localhost:5432/database_name"
func makeBaseUrl(option *ConnOption) string {
	str := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		option.User,
		option.Password,
		option.Host,
		option.Port,
		option.DBName,
	)

	return str
}

func (cp *Conn) Init(option *ConnOption) error {
	dataBaseUrl := makeBaseUrl(option)
	conn, err := pgx.Connect(context.Background(), dataBaseUrl)
	if err != nil {
		return err
	}

	cp.conn = conn
	return nil
}

func (cp *Conn) Release() error {
	return cp.conn.Close(context.Background())
}

func (cp *Conn) QueryRow(sql string, result interface{}) error {
	err := cp.conn.QueryRow(context.Background(), sql).Scan(result)
	if err != nil {
		return err
	}
	return nil
}

func NewConn(option *ConnOption) (*Conn, error) {
	pool := new(Conn)
	err := pool.Init(option)
	return pool, err
}

type DBHandler struct {
	db *sql.DB
}

func NewDBHandler(option *ConnOption) (*DBHandler, error) {
	h := new(DBHandler)
	err := h.Init(option)
	if err != nil {
		return nil, err
	}
	return h, nil
}

func (h *DBHandler) Init(option *ConnOption) error {
	dataBaseUrl := makeBaseUrl(option)
	db, err := sql.Open("pgx", dataBaseUrl)
	if err != nil {
		return err
	}
	h.db = db
	return nil
}

func (h *DBHandler) GetDB() *sql.DB {
	return h.db
}
