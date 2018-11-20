package logger

import (
	"fmt"
	"time"
	"encoding/json"
	"bytes"
	"os"
)

var (
	level       = "debug"
	tranceLevel = 0
	debugLevel  = 1
	infoLevel   = 2
	warnLevel   = 3
	errorLevel  = 4
	fatalLevel  = 5
)

const (
	TranceLevel = "trance"
	DebugLevel  = "debug"
	ErrorLevel  = "error"
	WarnLevel   = "warn"
	FatalLevel  = "fatal"
	InfoLevel   = "info"
)

var isJson bool
var isAsync bool

type Logfile interface {
	WriteLog(msg string)
}

var f *os.File
var msgChan chan string

type Log struct {
	Level string `json:"level"`
	Time  string `json:"time"`
	Msg   string `json:"msg"`
}

func IsJson(b bool) {
	isJson = b
}

func SetLogLevel(logLevel string) {
	level = logLevel
}
func SetLogFile(logfile *os.File) {
	f = logfile
}
func SetIsAsync(b bool, size int, fun func(m string)) {
	isAsync = b
	if size == 0 {
		size = 5000
	}
	msgChan = make(chan string, size)
	go func() {
		for msg := range msgChan {
			if fun != nil {
				fun(msg)
			} else {
				f.WriteString(msg)
			}
		}
	}()
}
func buildMsg(logLevel string, format string, args interface{}) string {
	var msg string
	if args == nil {
		msg = format
	} else {
		msg = fmt.Sprintf(format, args)

	}
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	if isJson {
		b := make([][]byte, 2)
		logbyte, _ := json.Marshal(Log{Level: logLevel, Msg: msg, Time: timeStr})
		b[0] = logbyte
		b[1] = []byte("\n")
		seq := []byte("")
		return string(bytes.Join(b, seq))
	} else {
		logStr := fmt.Sprintf("level:%s time:%s msg:%s \n", logLevel, timeStr, msg)
		return logStr
	}
}
func writeLog(msg string) {
	if isAsync {
		msgChan <- msg
	} else {
		f.WriteString(msg)
	}
}
func TranceF(format string, args interface{}) {
	if getLevel(level) > tranceLevel {
		return
	}
	writeLog(buildMsg(TranceLevel, format, args))
}
func DebugF(format string, args interface{}) {
	if getLevel(level) > debugLevel {
		return
	}
	writeLog(buildMsg(DebugLevel, format, args))
}
func InfoF(format string, args interface{}) {
	if getLevel(level) > infoLevel {
		return
	}
	writeLog(buildMsg(InfoLevel, format, args))
}
func WarnF(format string, args interface{}) {
	if getLevel(level) > warnLevel {
		return
	}
	writeLog(buildMsg(WarnLevel, format, args))
}

func ErrorF(format string, args interface{}) {
	if getLevel(level) > errorLevel {
		return
	}
	writeLog(buildMsg(ErrorLevel, format, args))
}
func FatalF(format string, args interface{}) {
	if getLevel(level) > fatalLevel {
		return
	}
	writeLog(buildMsg(FatalLevel, format, args))
}

func Trance(msg string) {
	if getLevel(level) > tranceLevel {
		return
	}
	writeLog(buildMsg(TranceLevel, msg, nil))
}
func Debug(msg string) {
	if getLevel(level) > debugLevel {
		return
	}
	writeLog(buildMsg(DebugLevel, msg, nil))
}
func Info(msg string) {
	if getLevel(level) > infoLevel {
		return
	}
	writeLog(buildMsg(InfoLevel, msg, nil))
}
func Warn(msg string) {
	if getLevel(level) > warnLevel {
		return
	}
	writeLog(buildMsg(WarnLevel, msg, nil))
}

func Error(msg string) {
	if getLevel(level) > errorLevel {
		return
	}
	writeLog(buildMsg(ErrorLevel, msg, nil))
}
func Fatal(msg string) {
	if getLevel(level) > fatalLevel {
		return
	}
	writeLog(buildMsg(FatalLevel, msg, nil))
}
func getLevel(l string) int {
	switch l {
	case "trance":
		return tranceLevel
	case "debug":
		return debugLevel
	case "wran":
		return warnLevel
	case "info":
		return infoLevel
	case "error":
		return errorLevel
	case "fatal":
		return fatalLevel
	default:
		return debugLevel
	}
}

func init() {
	f = os.Stdout
}
