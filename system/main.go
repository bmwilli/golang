package main
import "fmt"
import "os"
func main() {
   h, _ := os.Hostname()
   fmt.Println("hostname: %s",h)
}
