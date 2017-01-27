package fizzbuzz

import "strconv"


func serial_transformer(dictionary map[int]string, number int) string {
  result := ""

  for k, v := range dictionary {
    if number % k == 0 {
      result += v
    }
  }
  return result
}

func fizzbuzz(number int) string {
  dictionary := make(map[int]string)
  dictionary[3] = "Fizz"
  dictionary[5] = "Buzz"
  dictionary[7] = "Razz"

  result := serial_transformer(dictionary, number)

  if result == "" {
    return strconv.Itoa(number)
  }
  return result
}
