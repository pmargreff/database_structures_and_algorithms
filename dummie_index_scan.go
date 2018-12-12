package main

func DummieIndexScan(array []int, target int, min int, max int) int {
  mid := min + (len(array[min:max]) / 2)

  if len(array[min:max]) == 0 {
    return NOT_FOUND
  } else if array[mid] == target {
    return mid
  } else if array[mid] > target {
    return DummieIndexScan(array, target, min, mid)
  } else if array[mid] < target {
    return DummieIndexScan(array, target, mid + 1, max)
  }

  return NOT_FOUND
}
