package main

import (
	"fmt"
	"github.com/mabetle/mmsg"
)

func Demo() {
	mmsg.LoadMessages("messages")
	DemoShow()
}

func DemoShow() {
	key := "msg.hello"
	fmt.Printf("Contains: %s: %v\n", key, mmsg.Contains("en", key))
	fmt.Printf("%s\n", mmsg.Message("en", key))
	fmt.Printf("%s\n", mmsg.Message("zh", key))
}

func DemoLoadPrefixMessage() {
	path := "/rundata"
	prefix := "msg"
	mmsg.LoadPrefixMessages(path, prefix)
	DemoShow()
}

func main() {
	//Demo()
	DemoLoadPrefixMessage()
}
