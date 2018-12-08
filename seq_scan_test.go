package main

import "testing"

func TestSeqScanWithoutTargetValue(t *testing.T) {
  var expected_index int = -1
  array := []int{2, 3, 5, 7, 11, 13}

  returned_index := SeqScan(array, 100)

  if returned_index != expected_index {
    t.Errorf("Incorrect SeqScan, got: %d, expected: %d.", returned_index, expected_index)
  }

}

func TestSeqScanWithSmallSet(t *testing.T) {
  var expected_index int = 2
  array := []int{2, 3, 5, 7, 11, 13}

  returned_index := SeqScan(array, 5)

  if returned_index != expected_index {
    t.Errorf("Incorrect SeqScan, got: %d, expected: %d.", returned_index, expected_index)
  }
}

func TestSeqScanWithDuplicatedValue(t *testing.T) {
  var expected_index int = 1
  array := []int{2, 3, 5, 7, 11, 3}

  returned_index := SeqScan(array, 3)

  if returned_index != expected_index {
    t.Errorf("Incorrect SeqScan, got: %d, expected: %d.", returned_index, expected_index)
  }
}

func TestSeqScanOnLargeArray(t *testing.T) {
  var expected_index int = 100
  var array_size int = 100000

  array := make([]int, array_size)
  for i := range array {
    array[i] = i
  }

  returned_index := SeqScan(array, 100)

  if returned_index != expected_index {
    t.Errorf("Incorrect SeqScan, got: %d, expected: %d.", returned_index, expected_index)
  }
}
