package xlog

import "testing"

func TestDebugLogger(t *testing.T) {

	l, _ := New(Config{
		Color: true,
	})
	SetGlobal(l)

	Debugln("hello")
	Errorln("should be red")
	Warnln("should be yellow")
}

func TestFileLogger(t *testing.T) {

	l, _ := New(Config{
		LogFile: "log.log",
	})
	SetGlobal(l)

	Debugln("hello")
	Errorln("err")
}

func TestJsonLogger(t *testing.T) {

	l, _ := New(Config{
		JsonLog: true,
	})
	SetGlobal(l)

	Debugln("hello")
	Errorln("should be red")
	Warnln("should be yellow")
}
