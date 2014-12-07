package TLog

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

type Logger struct {
	traceID int32
	inTrace bool
}

func init() {
	rand.Seed(time.Now().Unix())
}

func (l *Logger) setTraceID() {
	l.traceID = rand.Int31()
}

func (l *Logger) getLogHead(skip int) string {
	pc, file, line, _ := runtime.Caller(skip + 1)
	funcName := runtime.FuncForPC(pc).Name()
	timeStr := time.Now().UTC().Format(time.RFC822)
	str := fmt.Sprintf("[%d] %s [%s:%d] %s: ", l.traceID, timeStr, file, line, funcName)
	return str
}

func (l *Logger) getLogContent(args ...interface{}) string {
	str := ""
	for _, arg := range args {
		str += fmt.Sprintf("%+v ", arg)
	}
	return str
}

func (l *Logger) StartTrace(args ...interface{}) {
	l.inTrace = true
	l.setTraceID()
	//Caller -> StartTrace
	str := l.getLogHead(1)
	str += "[Start Tracing] "
	str += l.getLogContent(args...)
	fmt.Println(str)
}

func (l *Logger) Start(args ...interface{}) {
	if l.inTrace {
		id := l.traceID
		l.setTraceID()
		str := l.getLogHead(1)
		str += fmt.Sprintf("[Start From %d] ", id)
		str += l.getLogContent(args...)
		fmt.Println(str)
	}
}

func (l *Logger) End(args ...interface{}) {
	if l.inTrace {
		str := l.getLogHead(1)
		str += "[End] "
		str += l.getLogContent(args...)
		fmt.Println(str)
	}
}

func (l *Logger) Log(args ...interface{}) {
	if l.inTrace {
		//Caller -> LogFunc -> logWithSkip
		str := l.getLogHead(1)
		str += "[Log] "
		str += l.getLogContent(args...)
		fmt.Println(str)
	}
}
