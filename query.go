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

func cachehitCounter(db *sql.DB) float32 {
	var cache_hit float32
	db.QueryRow("SELECT round(sum(heap_blks_hit) / (sum(heap_blks_hit) + sum(heap_blks_read)), 4) as ratio FROM pg_statio_user_tables;").Scan(&cache_hit)
	return cache_hit
}
func indexhitCounter(db *sql.DB) float32 {
	var index_hit float32
	db.QueryRow("SELECT round(sum(idx_blks_hit) / sum(idx_blks_hit + idx_blks_read), 4) as ratio FROM pg_statio_user_indexes;").Scan(&index_hit)
	return index_hit
}
func missingindexCounter(db *sql.DB) int {
	var missing int
	db.QueryRow("SELECT COALESCE(sum(case when seq_scan-idx_scan > 0 THEN 1 ELSE 0 END),0) as counter FROM pg_stat_all_tables WHERE schemaname='public' AND pg_relation_size(relname::regclass) > 80000;").Scan(&missing)
	return missing
}

// DB HELPERs
func dbConnect() *sql.DB {
	conn, err := sql.Open("postgres", "user=rob dbname=rob sslmode=disable")
	if err != nil {
		log.Fatal("Error: The data source arguments are not valid")
	}
	return conn
}
