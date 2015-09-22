package mmsg

import (
	. "github.com/mabetle/mcore/mtest"
	"testing"
)

func TestPutMsg(t *testing.T) {
	PutMsg("en", "hello", "hello demo")
	AssertEqual(Message("en", "hello"), "hello demo")
}

func TestPutText(t *testing.T) {
	text := `
# this is a demo
hello2=hello demo
name=demo
	`
	PutMsgText("en", text)
	AssertEqual(Message("en", "hello2"), "hello demo")
	AssertEqual(Message("en", "name"), "demo")
}

func TestFile(t *testing.T) {
	LoadMessages("testdata")
	AssertEqual(Message("en", "hello"), "Hello from English")
	AssertEqual(Message("zh", "hello"), "Hello from Chinese, 你好")
}
