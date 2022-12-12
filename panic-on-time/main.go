package main

import (
	"os"
	"strconv"
	"time"
)

func main() {
	n, _ := strconv.Atoi(os.Getenv("TIME_INTERVAL"))
	time.Sleep(time.Second * time.Duration(n))
	panic("system encountered error")
}
