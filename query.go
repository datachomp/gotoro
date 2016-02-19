package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func selectVersion(db *sql.DB) string {
	var pg_version string
	db.QueryRow("show server_version;").Scan(&pg_version)
	return pg_version
}

func selectDeadrowCount(db *sql.DB) string {
	row := db.QueryRow("select sum(n_dead_tup) as deadrows from pg_stat_user_tables;")
	var deadrows string
	err := row.Scan(&deadrows)
	if err != nil {
		panic(err)
	}
	return deadrows
}

func connectionCounter(db *sql.DB) *ConnectionState {
	var c ConnectionState
	db.QueryRow("select count(0) as total, count(0) FILTER(WHERE state = 'active') as active_count, count(0) FILTER(WHERE state = 'idle') as idle_count, count(0) FILTER(WHERE state = 'idle in transaction') as idle_transaction_count from pg_stat_activity;").Scan(&c.total, &c.active, &c.idle, &c.idle_transaction)
	return &c
}

// DB HELPERs
func dbConnect() *sql.DB {
	conn, err := sql.Open("postgres", "user=rob dbname=rob sslmode=disable")
	if err != nil {
		log.Fatal("Error: The data source arguments are not valid")
	}
	return conn
}