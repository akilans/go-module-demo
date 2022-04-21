package main

import (
	goodbye "github.com/akilans/go-module-demo/bye"
	"github.com/akilans/go-module-demo/hello"
)

func main() {
	hello.SayHello()
	goodbye.SayBye()
	sayHi()
	//sayBye()
	//sayHello()
}
