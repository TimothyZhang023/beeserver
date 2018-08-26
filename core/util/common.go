package util


import (
	"fmt"
	"os"
	"runtime"
)

func Zeus() {
	if e := recover(); e != nil {

		const size = 64 << 10
		buf := make([]byte, size)
		buf = buf[:runtime.Stack(buf, false)]

		fmt.Printf("[panic] err=%v \n %s \n ", e, buf)
		os.Exit(-1)
	}
}