package main

import (
  "flag"
  "runtime"
  "fmt"
)

func main() {
  cpunum := flag.Int("cpunum", 0, "The number of cores")
  flag.Parse()
  fmt.Println("The number of cores:", *cpunum)
  runtime.GOMAXPROCS(*cpunum)
  for i := 0; i < *cpunum - 1; i++ {
    go func() {
      for {
      }
    }()
  }
  for {
  }
}
