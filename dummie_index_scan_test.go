package main

import "testing"

func TestDummieIndexScanWithLowerInexistentTarget(t *testing.T) {
  var expected_index int = -1
  array := []int{2, 3, 5, 7, 11, 13}

  returned_index := DummieIndexScan(array, 1, 0, len(array))

  if returned_index != expected_index {
    t.Errorf("Incorrect DummieIndexScan, got: %d, expected: %d.", returned_index, expected_index)
  }
}

func TestDummieIndexScanWithHigherInexistentTarget(t *testing.T) {
  var expected_index int = -1
  array := []int{2, 3, 5, 7, 11, 13}

  returned_index := DummieIndexScan(array, 100, 0, len(array))

  if returned_index != expected_index {
    t.Errorf("Incorrect DummieIndexScan, got: %d, expected: %d.", returned_index, expected_index)
  }
}

func TestDummieIndexScanWithInRangeInexistentTarget(t *testing.T) {
  var expected_index int = -1
  array := []int{2, 3, 5, 7, 11, 13}

  returned_index := DummieIndexScan(array, 9, 0, len(array))

  if returned_index != expected_index {
    t.Errorf("Incorrect DummieIndexScan, got: %d, expected: %d.", returned_index, expected_index)
  }
}

func TestDummieIndexScanWithSmallSet(t *testing.T) {
  var expected_index int = 2
  array := []int{2, 3, 5, 7, 11, 13}

  returned_index := DummieIndexScan(array, 5, 0, len(array))

  if returned_index != expected_index {
    t.Errorf("Incorrect DummieIndexScan, got: %d, expected: %d.", returned_index, expected_index)
  }
}

func TestDummieIndexScanWithDuplicatedValue(t *testing.T) {
  var expected_index int = 1
  array := []int{2, 3, 5, 7, 11, 3}

  returned_index := DummieIndexScan(array, 3, 0, len(array))

  if returned_index != expected_index {
    t.Errorf("Incorrect DummieIndexScan, got: %d, expected: %d.", returned_index, expected_index)
  }
}

func TestDummieIndexScanOnLargeArray(t *testing.T) {
  var expected_index int = 100
  var array_size int = 1000

  array := make([]int, array_size)
  for i := range array {
    array[i] = i
  }

  returned_index := DummieIndexScan(array, 100, 0, len(array))

  if returned_index != expected_index {
    t.Errorf("Incorrect DummieIndexScan, got: %d, expected: %d.", returned_index, expected_index)
  }
}
