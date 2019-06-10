package mysql

import (
	"database/sql"
	"fmt"
	"time"

	goqu "gopkg.in/doug-martin/goqu.v5"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gobuffalo/packr"
	"github.com/rubenv/sql-migrate"
	// Регистрация адаптера mysql
	_ "gopkg.in/doug-martin/goqu.v5/adapters/mysql"
)

// ConnectionOptions параметры соединения с mysql бд
type ConnectionOptions struct {
	Host  string
	Port  int
	User  string
	Pass  string
	Name  string
	Debug bool
}

// Addr адрес соединеия
func (o *ConnectionOptions) Addr() string {
	return fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", o.User, o.Pass, o.Host, o.Port, o.Name)
}

// CensoredAddr адрес соединеия со скрытым паролем
func (o *ConnectionOptions) CensoredAddr() string {
	return fmt.Sprintf("%s:*****@(%s:%d)/%s", o.User, o.Host, o.Port, o.Name)
}

// NewConnection открыает новое соединение с бд
func NewConnection(o *ConnectionOptions) (*goqu.Database, error) {
	// создает экземляр DB используя mysql драйвер
	// только проверяет аргементы  и не производит реального соединения
	con, err := sql.Open("mysql", o.Addr())
	if err != nil {
		msg := "Cannot open mysql connection %v"
		return nil, fmt.Errorf(msg, err)
	}

	// установит соединение и проверит доступность базы
	err = con.Ping()
	if err != nil {
		msg := "Mysql ping failed: [%s]: %v"
		return nil, fmt.Errorf(msg, o.CensoredAddr(), err)
	}

	// открытое соединение держится не более часа
	con.SetConnMaxLifetime(time.Hour)

	// инициализация конструктора запросов
	// если параметр debug установлен в true
	// будет логировать все запросы в консоль
	db := goqu.New("mysql", con)

	// папка с файлами миграций
	migrations := &migrate.PackrMigrationSource{
		Box: packr.NewBox("./migrations/"),
	}

	// применение миграций
	n, err := migrate.Exec(con, "mysql", migrations, migrate.Up)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Applied %d migrations!\n", n)
	return db, nil
}
