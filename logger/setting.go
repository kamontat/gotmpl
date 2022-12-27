package logger

import (
	"io"
	"log"
	"os"
)

func Setup(debug bool) {
	log.SetFlags(log.Ltime | log.Lshortfile | log.Lmsgprefix)
	log.SetPrefix("[DEBUG] ")
	if debug {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(io.Discard)
	}
}
