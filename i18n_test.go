package mmsg

import (
	"testing"
	. "github.com/mabetle/mcore/mtest"
)

func TestPutMsg(t *testing.T) {
	PutMsg("en","hello", "hello demo")
	AssertEqual(LocaleMessage("en","hello"), "hello demo")
	//default to zh, but i donot set zh key
	AssertEqual(Message("hello"), "hello demo")
}

func TestPutText(t *testing.T){
	text:=`
# this is a demo
hello2=hello demo
name=demo
	`
	PutMsgText("en", text)
	AssertEqual(Message("hello2"), "hello demo")
	AssertEqual(Message("name"), "demo")
}


func TestFile(t *testing.T) {
	LoadMessages("testdata")
	AssertEqual(LocaleMessage("en", "hello"),"Hello from English")
	AssertEqual(LocaleMessage("zh", "hello"),"Hello from Chinese, 你好")
}

