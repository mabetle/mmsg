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
	// Win
	//goPath := "D:/devlab/gocodes/"
	// Linux
	goPath := "/home/korbenzhang/dev/gocodes/"
	//	/devlab/gocodes/src/mabetle/apps/web/messages/msg.en
	path := goPath + "src/mabetle/apps/web/messages"
	path = "/rundata/"
	prefix := "msg"
	mmsg.LoadPrefixMessages(path, prefix)
	DemoShow()

	//outPath := "/home/korbenzhang/"

	//mmsg.SaveMessageFile(outPath, prefix)
}

func main() {
	//Demo()
	DemoLoadPrefixMessage()
}
