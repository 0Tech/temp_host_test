package hello

import "fmt"

func hello(name string) string {
	return fmt.Sprintf("hello, %s!", name)
}