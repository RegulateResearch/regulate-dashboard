package prep

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func PrepFile() (warnLogFile *os.File, errLogFile *os.File) {
	const warnFileName = "warn.log"
	const errFileName = "error.log"
	warnLogFile = openFile(warnFileName)
	errLogFile = openFile(errFileName)

	return
}

func openFile(filename string) *os.File {
	flag := os.O_CREATE | os.O_WRONLY | os.O_APPEND
	var perm fs.FileMode = 0666 // read+write to all group (should be)
	logFile, err := os.OpenFile(filename, flag, perm)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalf("cannot open file for log: %s", filename)
	}

	return logFile
}
