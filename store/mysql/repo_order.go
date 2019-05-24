package mysql

import (
	"fmt"
	"gopkg.in/doug-martin/goqu.v5"
	"through/env"
	"time"
)

type sequenceRow struct {
	ID   int64     `db:"id"`
	Date time.Time `db:"date"`
}

type orderRow struct {
	ID   int64     `db:"id"`
	Date time.Time `db:"created_at"`
}

type OrderRepo struct {
	db *goqu.Database
}

func NewOrderRepo(db *goqu.Database) *OrderRepo {
	return &OrderRepo{
		db: db,
	}
}

func (r OrderRepo) AddThrough() bool {
	var table int64 = 1

	for i := 1; i <= env.Config.InsertCount; i++ {

		result, err := r.db.
			From(goqu.I("through")).
			Insert(sequenceRow{Date: time.Now()}).
			Exec()

		if err != nil {
			fmt.Println(err.Error())
		}

		id, _ := result.LastInsertId()

		_, err = r.db.
			From(goqu.I(fmt.Sprintf("orders_%d", table))).
			Insert(orderRow{
				ID:   id,
				Date: time.Now(),
			}).
			Exec()

		if err != nil {
			fmt.Println(err.Error())
		}

		if table == 3 {
			table = 1
		} else {
			table++
		}
	}

	return true
}

func (r OrderRepo) AddAutoincrement() bool {
	for i := 1; i <= 1000; i++ {

		_, err := r.db.
			From(goqu.I("auto_orders")).
			Insert(orderRow{Date: time.Now()}).
			Exec()

		if err != nil {
			fmt.Println(err.Error())
		}
	}

	return true
}
