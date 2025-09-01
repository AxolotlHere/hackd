package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

var pool *pgxpool.Pool
var ctx = context.Background()

type User struct {
	uid         int
	username    string
	password    string
	points      int
	rounds      int
	description string
	email       string
}

func initDbController() {
	var err error
	pool, err = pgxpool.New(ctx, "postgres://postgres:Py2nb-0513@127.0.0.1:5432/base_1")
	if err != nil {
		fmt.Println("Enters but shouldnt")
		log.Fatal("Unable to connect to database")
	}
	if err := pool.Ping(ctx); err != nil {
		fmt.Println("Enters but shouldnt pt-2")
		log.Fatal("Unable to ping database")
	}
}
