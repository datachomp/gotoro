package main

import (
	"github.com/cloudfoundry/gosigar"
	//"os"
	"runtime"
)

func format(val uint64) uint64 {
	return val / 1024
}

func system_cpucores() int {
	return runtime.NumCPU()
}

func system_memstats() int {
	mem := sigar.Mem{}
	mem.Get()
	return runtime.NumCPU()
}

/*
fmt.Printf("SYSTEM INFO: \n")
fmt.Print("cores: ", runtime.NumCPU(), "\n")

mem := sigar.Mem{}
mem.Get()

fmt.Fprintf(os.Stdout, "%18s %10s %10s\n",
  "total", "used", "free")

fmt.Fprintf(os.Stdout, "Mem:    %10d %10d %10d\n",
  format(mem.Total), format(mem.Used), format(mem.Free))

fmt.Fprintf(os.Stdout, "-/+ buffers/cache: %10d %10d\n",
  format(mem.ActualUsed), format(mem.ActualFree))
*/
