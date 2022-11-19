package main

import (
	"fmt"

	goodbye "github.com/akilans/go-module-demo/bye"
	"github.com/akilans/go-module-demo/hello"
)

func main() {

	// sayHi() , myName are directly consumed here
	sayHi()
	fmt.Println(myName)

	// SayBye() exported because it follows camelcase pattern
	goodbye.SayBye()
	// Throws error as byeMessage variable not exported
	// fmt.Println(goodbye.byeMessage)

	// SayHello() exported because it follows camelcase pattern
	hello.SayHello()

}
