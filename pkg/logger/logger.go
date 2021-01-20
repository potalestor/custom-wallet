package logger

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/potalestor/custom-wallet/pkg/cfg"
)

const fmtAppInfo = `
	%s starting...

`

func createLogfile() string {
	file := os.Args[0]

	return strings.TrimSuffix(file, filepath.Ext(file)) + ".log"
}

// Initialize logger using cfg.
func Initialize(config *cfg.Config) {
	if config.Logger.File {
		logfile := createLogfile()

		file, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Failed to open/create log file '%s'", logfile)
		}

		multi := io.MultiWriter(file, os.Stdout)
		log.SetOutput(multi)
	}

	log.SetFlags(0)

	log.Printf(fmtAppInfo, strings.ToUpper(filepath.Base(os.Args[0])))
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
}
