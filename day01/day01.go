package main

import (
  "os"
  "log"
  "bufio"
  "fmt"
  "strconv"
)

func main() {
  file, err := os.Open(os.Args[1])
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)
  total1 := 0
  total2 := 0
  for scanner.Scan() {
    i, _ := strconv.Atoi(scanner.Text())
    total1 = total1 + fuel_required(i)
  }
  fmt.Println("part 1", total1)
  fmt.Println("part 2", total2)
}

func fuel_required(mass int) int {
  return mass / 3 - 2
}
