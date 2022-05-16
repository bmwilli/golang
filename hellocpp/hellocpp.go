package main

// #cgo CXXFLAGS: -I/usr/lib/
// #cgo LDFLAGS: -L/usr/lib/ -lstdc++
// #include "helloworld.hpp"
import "C"

import "fmt"

func main() {
  C.hello_world()
}

