package main 

import "fmt"

func sayHello(name string) string {
	if len(name) == 0 {
		return "hello Anonymous"
	} 
	return fmt.Sprintf("hello %s", name)
}

func sayGoodBye(name string) string {
	if len(name) == 0 {
		return "Bye Bye Anonymous!"
	}
	return fmt.Sprintf("Bye Bye %s!", name)
}