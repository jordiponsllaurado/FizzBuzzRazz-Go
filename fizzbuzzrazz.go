package main

import (
  "strconv"
  "fmt"
  "os"
  "bufio"
  "strings"
  )


func serial_transformer(dictionary map[int]string, number int) string {
  var result string

  for k, v := range dictionary {
    if number % k == 0 {
      result += v
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
