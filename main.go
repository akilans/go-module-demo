package main

import (
	"fmt"

	goodbye "github.com/akilans/go-module-demo/bye"
	"github.com/akilans/go-module-demo/hello"
)

func main() {

	sayHi()
	fmt.Println(myName)

	goodbye.SayBye()
	// Throws error as byeMessage variable not exported
	// fmt.Println(goodbye.byeMessage)

	hello.SayHello()

	//sayBye()
	//sayHello()
}
