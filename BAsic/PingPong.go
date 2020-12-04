package main
import (
	"fmt"
)
func main() {
	pingChan := make(chan string)
	pongChan := make(chan string)
	go printer(pongChan)
	go pinger(pingChan)
	go ponger(pingChan, pongChan)
	var input string
	fmt.Scanln(&input)
}
func pinger(pingChan chan<- string) {
	fmt.Println("pinger sending \"ping\"")
	pingChan <- "ping"
}
func ponger(pingChan <-chan string, pongChan chan<- string) {
	fmt.Println("ponger received", <-pingChan)
	fmt.Println("ponger replying with \"pong\"")
	pongChan <- "pong"
}
func printer(pongChan <-chan string) {
	fmt.Println("printer received", <-pongChan)
}