package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Logger struct {
	// contains filtered or unexported fields
}

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func Init(traceHandle io.Writer, infoHandle io.Writer,
	warningHandle io.Writer, errorHandle io.Writer) (Trace, Info, Warning, Error *log.Logger) {
	Trace = log.New(traceHandle, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(warningHandle, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	// Trace.Println("XXXXXXXXXXXX")
	return //Trace,Info,Warning,Error
}

func test_log2() {
	Trace, Info, Warning, Error := Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	Trace.Println("I have something standard to say")
	Info.Println("Special Information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
}

func test_log1() {
	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Lshortfile|log.Ldate)
	logger.Print("Hello, log file!")
	fmt.Print(&buf)
}

func main() {

	test_log1()
	test_log2()

}
