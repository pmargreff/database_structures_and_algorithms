package main

import "testing"

func TestScan(t *testing.T) {
  var expected string = "Hello"

  returned := Scan()

  if returned != expected {
    t.Errorf("Incorrect scan, got: %s, expected: %s.", returned, expected)
  }
}
