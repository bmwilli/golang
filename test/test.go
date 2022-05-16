package main

// #cgo CXXFLAGS: -I/usr/lib/
// #cgo LDFLAGS: -L/usr/lib/ -lstdc++
// #include "test.hpp"
import "C"

func main() {
  C.test()
}
