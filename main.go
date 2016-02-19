package main

import (
	"fmt"
)

func main() {
	db := dbConnect()

	pg_version := selectVersion(db)
	fmt.Printf("Postgres Version: %s \n", pg_version)

	deadRow := selectDeadrowCount(db)
	fmt.Printf("Dead Rows:        %s \n", deadRow)

	connRow := connectionCounter(db)
	fmt.Printf("Total: %d  Active: %v  Idle: %v \n", connRow.total, connRow.active, connRow.idle)
	// fmt.Printf("Active : %v \n", connRow.active)
	// fmt.Printf("Idle : %v \n", connRow.idle)

	// query
	// exec
}
