package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

func main() {
	fmt.Println("start")
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	for {
		mem := &runtime.MemStats{}
		cpu := runtime.NumCPU()
		rot := runtime.NumGoroutine()
		runtime.ReadMemStats(mem)
		fmt.Fprintf(f, "%v cpu %v gor %v mem %v %v\n",
			time.Now().Format("2006-02-01 15:04:05"),
			cpu,
			rot,
			mem.Alloc,
			mem.HeapAlloc,
		)
		time.Sleep(5 * time.Second)
	}
}
