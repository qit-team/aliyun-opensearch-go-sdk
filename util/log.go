package util

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

type Logger struct {
}

//Log - log.Logger
var Log *log.Logger

//LogWriter io.Writer
var LogWriter io.Writer

func initLog() {
	os.MkdirAll("logs", os.ModePerm)
	logName := "logs/app"
	logFile, err := os.OpenFile(logName+"-"+NowDate()+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("----Cannot open log file----")
	}
	writers := []io.Writer{
		logFile,
		os.Stdout,
	}
	LogWriter = io.MultiWriter(writers...)
	Log = log.New(LogWriter, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

// P Convenience function for printing to stdout
func P(a ...interface{}) {
	fmt.Println(NowTime(), a)
}

// LogPrint 输出并Log
func LogPrint(a ...interface{}) {
	aa, _ := json.Marshal(a)
	newstr := filterWarningInfo(string(aa[:]))
	var res interface{}
	err := json.Unmarshal([]byte(newstr), &res)
	if err != nil {
		Log.Println(a)
	}
	Log.Println(res)
}

// LogPrintln 输出并Log
func LogPrintln(a ...interface{}) {
	aa, _ := json.Marshal(a)
	newstr := filterWarningInfo(string(aa[:]))
	var res interface{}
	err := json.Unmarshal([]byte(newstr), &res)
	if err != nil {
		Log.Println(a)
	}
	Log.Println(res)
}

// LogPrintf 输出并Log
func LogPrintf(format string, a ...interface{}) {
	ss, _ := json.Marshal(format)
	newstr := filterWarningInfo(string(ss[:]))
	var res string
	err := json.Unmarshal([]byte(newstr), &res)
	if err != nil {
		res = format
	}
	aa, _ := json.Marshal(a)
	newstr = filterWarningInfo(string(aa[:]))
	var res1 []interface{}
	err = json.Unmarshal([]byte(newstr), &res1)
	if err != nil {
		res1 = a
	}
	Log.Printf(res, res1...)
}

func Error(message string, a ...interface{}) {
	aa, _ := json.Marshal(a)
	newStr := filterWarningInfo(string(aa[:]))
	var res interface{}
	err := json.Unmarshal([]byte(newStr), &res)
	if err != nil {
		Log.Println(a)
	}
	Log.Println(res)
}

func Warning(message string, a ...interface{}) {
	aa, _ := json.Marshal(a)
	newStr := filterWarningInfo(string(aa[:]))
	var res interface{}
	err := json.Unmarshal([]byte(newStr), &res)
	if err != nil {
		Log.Println(a)
	}
	Log.Println(res)
}

func Info() {

}

func filterWarningInfo(format string) (result string) {
	s := filterIdNum(format)
	format = filterPhoneNum(s)
	return format
}

// 身份证号脱敏
func filterIdNum(searchIn string) (result string) {
	pat := `([1-9])[0-7]\d{4}(19[0-9][0-9]|20[0-3][0-9])\d{7}(\d|X|x)` //正则
	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		re, _ := regexp.Compile(pat)
		//将匹配到的部分替换为"##.#"
		str := re.ReplaceAllString(searchIn, "${1}0000000000000000${3}")
		return str
	}
	return searchIn
}

// 手机号脱敏
func filterPhoneNum(searchIn string) (result string) {
	pat := `(1)(\d{9})(\d)` //正则
	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		re, _ := regexp.Compile(pat)
		//将匹配到的部分替换为"##.#"
		str := re.ReplaceAllString(searchIn, "${1}000000000${3}")
		return str
	}
	return searchIn
}
