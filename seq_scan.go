package main

func SeqScan(array []int, target int) int {
  for index, element := range array {
    if element == target {
      return index
    }
  }
  return NOT_FOUND
}
