package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	fmt.Println("start")
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for {
		v, _ := mem.VirtualMemory()
		c, _ := cpu.Times(false)
		fmt.Fprintf(f, "%v total: %v, free: %v, used: %v%%, cpu: %v\n",
			time.Now().Format("2006-02-01 15:04:05"),
			v.Total,
			v.Free,
			v.UsedPercent,
			c,
		)
		time.Sleep(5 * time.Second)
	}
}
