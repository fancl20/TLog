package TLog

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

type Logger interface {
	StartTrace(args ...interface{})
	Start(args ...interface{})
	End(args ...interface{})
	Log(args ...interface{})
}

type LoggerImpl struct {
	traceID int32
	inTrace bool
}

func init() {
	rand.Seed(time.Now().Unix())
}

func NewLogger() *LoggerImpl {
	return new(LoggerImpl)
}

func (l *LoggerImpl) setTraceID() {
	l.traceID = rand.Int31()
}

func (l *LoggerImpl) getLogHead(skip int) string {
	pc, file, line, _ := runtime.Caller(skip + 1)
	funcName := runtime.FuncForPC(pc).Name()
	timeStr := time.Now().UTC().Format(time.RFC822)
	str := fmt.Sprintf("[%d] %s [%s:%d] %s: ", l.traceID, timeStr, file, line, funcName)
	return str
}

func (l *LoggerImpl) getLogContent(args ...interface{}) string {
	str := ""
	for _, arg := range args {
		str += fmt.Sprintf("%+v ", arg)
	}
	return str
}

func (l *LoggerImpl) StartTrace(args ...interface{}) {
	l.inTrace = true
	l.setTraceID()
	//Caller -> StartTrace
	str := l.getLogHead(1)
	str += "[Start Tracing] "
	str += l.getLogContent(args...)
	fmt.Println(str)
}

func (l *LoggerImpl) Start(args ...interface{}) {
	if l.inTrace {
		id := l.traceID
		l.setTraceID()
		//Caller -> Start
		str := l.getLogHead(1)
		str += fmt.Sprintf("[Start From %d] ", id)
		str += l.getLogContent(args...)
		fmt.Println(str)
	}
}

func (l *LoggerImpl) End(args ...interface{}) {
	if l.inTrace {
		//Caller -> End
		str := l.getLogHead(1)
		str += "[End] "
		str += l.getLogContent(args...)
		fmt.Println(str)
	}
}

func (l *LoggerImpl) Log(args ...interface{}) {
	if l.inTrace {
		//Caller -> Log
		str := l.getLogHead(1)
		str += "[Log] "
		str += l.getLogContent(args...)
		fmt.Println(str)
	}
}
