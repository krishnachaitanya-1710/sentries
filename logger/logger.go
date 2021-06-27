package logger

import (
	"fmt"
	"os"
)

func main() {
	InfoS("INFO")
	ErrorS("ERROR")
	DebugS("DEBUG")
	WarnS("WARN")
	NoticeS("NOTICE")
	StreamRed("INFO")
	StreamBlue("ERROR")
	StreamGreen("DEBUG")
	StreamYellow("WARN")
	NoticeS("NOTICE")
}

/*
Simple logging utility with ANSI color support which formats the given message and prints it
*/
const (
	colorReset   = "\033[0m"
	colorRed     = "\033[31m"
	colorGreen   = "\033[32m"
	colorYellow  = "\033[33m"
	colorBlue    = "\033[34m"
	colorPurple  = "\033[35m"
	colorCyan    = "\033[36m"
	colorWhite   = "\033[37m"
	InfoColor    = "\033[1;32m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;34m%s\033[0m"
)

func Info(msg string) {
	fmt.Println("------------------------ \n  INFO :: " + msg + "\n ------------------------\n")
}

func Debug(msg string) {
	if os.Getenv("LOG_LEVEL") == "debug" {
		fmt.Println("------------------------ \n  DEBUG :: " + msg + "\n------------------------ \n")
	}
}

func Warn(msg string) {
	fmt.Println("------------------------ \n  WARN :: " + msg + "\n------------------------\n")
}

func Error(msg string) {
	fmt.Println("------------------------ \n  ERROR :: " + msg + "\n------------------------\n")
}

func Fatal(msg string) {
	fmt.Println("------------------------ \n  FATAL :: " + msg + "\n------------------------\n")
}

func InfoS(msg string) {
	fmt.Printf(InfoColor, "------------------------ \n  INFO :: "+msg+"\n------------------------\n")
}

func WarnS(msg string) {
	fmt.Printf(WarningColor, "------------------------ \n  WARN :: "+msg+"\n------------------------\n")
}

func DebugS(msg string) {
	if os.Getenv("LOG_LEVEL") == "debug" {
		fmt.Printf(DebugColor, "------------------------ \n  INFO :: "+msg+"\n------------------------\n")
	}
}

func ErrorS(msg string) {
	fmt.Printf(ErrorColor, "------------------------ \n  WARN :: "+msg+"\n------------------------\n")
}

func NoticeS(msg string) {
	fmt.Printf(NoticeColor, "------------------------ \n  WARN :: "+msg+"\n------------------------\n")
}

func Stream(msg string) {
	fmt.Println(msg)
}

func StreamRed(msg string) {
	fmt.Println(string(colorRed), msg)
}

func StreamGreen(msg string) {
	fmt.Println(string(colorGreen), msg)
}

func StreamBlue(msg string) {
	fmt.Println(string(colorBlue), msg)
}

func StreamPurple(msg string) {
	fmt.Println(string(colorPurple), msg)
}

func StreamCyan(msg string) {
	fmt.Println(string(colorCyan), msg)
}

func StreamYellow(msg string) {
	fmt.Println(string(colorYellow), msg)
}

func StreamWhite(msg string) {
	fmt.Println(string(colorWhite), msg)
}

func StreamReset(msg string) {
	fmt.Println(string(colorReset), msg)
}
