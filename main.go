package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Age int
}

func main() {
	// config file - https://blog.gopheracademy.com/advent-2014/reading-config-files-the-go-way/
	var conf Config
	//if _, err := toml.DecodeFile("something.toml", &conf); err != nil {
	//	log.Fatal("Error: totallybroken config file")
	//}
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("Age: %s \n", conf.Age)

	db := dbConnect()

	fmt.Printf("Postgres Version: %s \n", selectVersion(db))
	fmt.Printf("Dead Rows:        %s \n", selectDeadrowCount(db))

	connRow := connectionCounter(db)
	fmt.Printf("Total: %d  Active: %v  Idle: %v \n", connRow.total, connRow.active, connRow.idle)

	fmt.Printf("CacheHit: %f  IndexHit: %f \n", cachehitCounter(db), indexhitCounter(db))
	fmt.Printf("# of missing indexes: %v \n", missingindexCounter(db))

	fmt.Printf("\nSYSTEM INFO \n")
	fmt.Print("cores: ", system_cpucores(), "\n")
}

// query
// exec

//if pg_stat_statements instaled,  get longest query duration

/*
	mem := sigar.Mem{}
	mem.Get()

	fmt.Fprintf(os.Stdout, "%18s %10s %10s\n",
		"total", "used", "free")

	fmt.Fprintf(os.Stdout, "Mem:    %10d %10d %10d\n",
		format(mem.Total), format(mem.Used), format(mem.Free))

	fmt.Fprintf(os.Stdout, "-/+ buffers/cache: %10d %10d\n",
		format(mem.ActualUsed), format(mem.ActualFree))
*/
