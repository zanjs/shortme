package conf

import (
	"log"
	"runtime"
)

var Version string

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
