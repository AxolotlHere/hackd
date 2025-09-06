package main

import (
	"context"
	"fmt"
	"github.com/AxolotlHere/hackd/internal/env"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

var pool *pgxpool.Pool
var ctx = context.Background()

type User struct {
	Uid         int    `json:"uid"`
	Username    string `json:"username"`
	Password    string `json:"Password"`
	Points      int    `json:"Points"`
	Rounds      int    `json:"Rounds"`
	Description string `json:"Description"`
	Email       string `json:"Email"`
}

func initDbController() {
	var err error
	var pg_param env.DatabaseCreds = env.GetDatabaseCreds()
	var conn_string = fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=require&channel_binding=require", pg_param.Db_user, pg_param.Db_pass, pg_param.Db_url, pg_param.Db_name)
	pool, err = pgxpool.New(ctx, conn_string)
	if err != nil {

		log.Fatal("Unable to connect to database")
	}
	if err := pool.Ping(ctx); err != nil {
		log.Fatal("Unable to ping database")
	}
}
