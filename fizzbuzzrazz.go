package main

import (
  "strconv"
  "fmt"
  "os"
  "bufio"
  "strings"
  "sort"
  )


func serial_transformer(dictionary map[int]string, number int) string {
  var result string

  var keys []int
  for k := range dictionary {
      keys = append(keys, k)
  }
  sort.Ints(keys)

  for _, k := range keys {
      if number % k == 0 {
        result += dictionary[k]
      }
  }

  return result
}

func fizzbuzz(number int) string {
  dictionary := map[int]string {
    3: "Fizz",
    5: "Buzz",
    7: "Razz",
  }
  result := serial_transformer(dictionary, number)

  if result == "" {
    return strconv.Itoa(number)
  }
  return result
}

func main() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Printf("Write a number\n")
  text, _ := reader.ReadString('\n')
  i, err := strconv.Atoi(strings.TrimSpace(text))
  if err != nil {
    fmt.Println("This is not a number!", err)
  } else {
    fizzBuzzValue := fizzbuzz(i)
    fmt.Println(fizzBuzzValue)
  }
}
