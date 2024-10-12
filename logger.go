package main

import (
	"fmt"
	"sync"
)

type level int

const (
	debug level = iota
	info
)

func (l level)GetString()string{
	if l == debug{
		return "debug"
	}else if  l ==info{
		return "info"
	}
	return ""
}

type LogSink interface {
	Log(message string)
}

type ConsoleSink struct {
}

func (c ConsoleSink) Log(message string) {
	fmt.Println(message)
}

type FileSink struct {
}

func (c FileSink) Log(message string) {
	fmt.Println("writing to file")
}

var logger *Logger
var mu = &sync.Mutex{}

type Logger struct {
	LogLevel level
	Sink LogSink
}
func NewLogger(level level)*Logger{
	if logger == nil{
		mu.Lock()
		defer mu.Unlock()
		if logger == nil{
			logger = &Logger{LogLevel:level,Sink:ConsoleSink{}}
			return logger
		}
	}
	return logger	
}
func(l Logger)log(level level,message string){
	mess:=fmt.Sprintf("[level:%s,msg:%s",level.GetString(),message)
	l.Sink.Log(mess)
}

func main(){
	logger:= NewLogger(info)
	logger.log(debug,"hello how are you")
	logger.log(debug,"important issue arising because of this")
}

