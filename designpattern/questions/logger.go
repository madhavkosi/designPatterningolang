//The logging framework should support different log levels, such as DEBUG, INFO, WARNING, ERROR, and FATAL.
//It should allow logging messages with a timestamp, log level, and message content.
//The framework should support multiple output destinations, such as console, file, and database.
//It should provide a configuration mechanism to set the log level and output destination.
//The logging framework should be thread-safe to handle concurrent logging from multiple threads.
//It should be extensible to accommodate new log levels and output destinations in the future.

package main
import(
	"fmt"
	"sync"
)
type Level int
const (
	DEBUG Level = iota + 1
	INFO 
)

type Log struct{
	Message string
	Level Level
}
type LogBuilder struct{
	l *Log
}
func NewLogBuilder()*LogBuilder{
	return &LogBuilder{l:&Log{}}
}
func(lb *LogBuilder)SetMessage(msg string)*LogBuilder{
	lb.l.Message = msg
	return lb
}
func(lb *LogBuilder)SetLevel(level Level)*LogBuilder{
	lb.l.Level = level
	return lb
}
func(lb *LogBuilder)Build()Log{
	return *(lb.l)
}

func NewLog(level Level,msg string)Log{
	return Log{Level:level,Message:msg}
}
type OutputSink interface{
	PublishMsg(log Log)
}

type Filestruct struct{

}

func (f Filestruct)PublishMsg(log Log){
	fmt.Println("%v",log)
}
func BuilderPattern(sinktype string)OutputSink{
	switch sinktype{
	case "file" :
		return &Filestruct{}
	default :
		return nil
	}
}
var logger *Logger
var mu sync.Mutex

type Logger struct{
	LogLevel Level
	OutputDestination  OutputSink
}
func GetLogger()*Logger{
	if logger == nil{
		mu.Lock()
		defer mu.Unlock()
		if logger == nil{
			logger =  &Logger{LogLevel:INFO,OutputDestination:&Filestruct{}}
		}
	}
	return logger
}
func(c *Logger)SetLogLevel(level Level){
	c.LogLevel = level
}

func(c *Logger)SetOutputDestination(sink OutputSink){
	c.OutputDestination = sink
}

func(l *Logger)Info(msg string){
	log := NewLog(INFO,msg)
	l.OutputDestination.PublishMsg(log)
}
func(l *Logger)Debug(msg string){
	log := NewLog(DEBUG,msg)
	l.OutputDestination.PublishMsg(log)
}
func main(){
	logger:=GetLogger()
	logger.Info("hello how re you")
	log :=NewLogBuilder().SetMessage("he").SetLevel(INFO).Build()
	something := BuilderPattern("file")
	something.PublishMsg(log)
}