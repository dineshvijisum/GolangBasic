package main

import (
	"fmt"
	"time"
)

type MyError struct {
	d time.Time
	s string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.d, e.s)
}
func oop() error {
	return MyError{
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"the file system has gone away",
	}
}
func main() {
	if err := oop(); err != nil {
		fmt.Println(err)
	}
}
