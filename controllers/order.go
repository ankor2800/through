package controllers

import (
	"fmt"
	"net/http"
	"sqnc/store/mysql"
)

type Controller struct {
	Order *mysql.OrderRepo
}

func (c Controller) AddSequence() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := c.Order.AddSequence()
		fmt.Println(res)
	}
}

func (c Controller) AddAutoincrement() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := c.Order.AddAutoincrement()
		fmt.Println(res)
	}
}
