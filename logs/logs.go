package logs

import (
	"fmt"
	"runtime"
	"time"
)

var dateFormat = "2006-01-02 15:04:05"
var reset = "\033[0m"
var red = "\033[31m"
var blue = "\033[34m"

func init() {
	if runtime.GOOS == "windows" {
		reset = ""
		red = ""
		blue = ""

	}
}

func D(message ...interface{}) {
	var dt = time.Now()
	var currentTime = dt.Format(dateFormat)
	for i := 0; i < len(message); i++ {
		fmt.Println(currentTime, ":", message[i])
	}
}

func E(message ...interface{}) {
	var dt = time.Now()
	var currentTime = dt.Format(dateFormat)
	for i := 0; i < len(message); i++ {
		fmt.Println(red, currentTime, ": ", message[i], reset)
	}
}

func W(message ...interface{}) {
	var dt = time.Now()
	var currentTime = dt.Format(dateFormat)
	for i := 0; i < len(message); i++ {
		fmt.Println(blue+currentTime, ": ", message[i], reset)
	}
}
