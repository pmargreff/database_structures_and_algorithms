package main

const NOT_FOUND int = -1

func SeqScan(array []int, find int) int {
  for index, element := range array {
    if element == find {
      return index
    }
  }
  return NOT_FOUND
}
