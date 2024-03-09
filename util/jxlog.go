package util

import "fmt"

func Log(a ...any) (n int, err error) {
	return fmt.Println(a...)
}
