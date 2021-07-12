package args

import (
	"fmt"
	"os"
	"path/filepath"
	"tinyhttp/log"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	Verbose = kingpin.Flag("verbose", "Verbose mode.").Short('v').Default("false").Bool()
	Host    = kingpin.Flag("host", "IP address to listen.").Short('h').Default("127.0.0.1").IP()
	Port    = kingpin.Flag("port", "Port to listen.").Short('p').Default("0").Int()
	Open    = kingpin.Flag("open", "Open in browser. Default true, --no-open to disable.").Default("true").Bool()
	Dir     = kingpin.Arg("dir", "Directory to server.").Default("").String()
)

func Parse() {
	kingpin.Parse()
	if *Verbose {
		log.TraceMode = true
	}
	if *Dir == "" {
		ex, err := os.Executable()
		if err != nil {
			log.Fatal(err.Error())
		}
		*Dir = filepath.Dir(ex)
	}
	absDir, err := filepath.Abs(*Dir)
	if err != nil {
		log.Fatal(err.Error())
	}
	*Dir = absDir
	log.Trace(fmt.Sprintf("verbose: %v", *Verbose))
	log.Trace(fmt.Sprintf("host: %s", *Host))
	log.Trace(fmt.Sprintf("port: %d", *Port))
	log.Trace(fmt.Sprintf("open: %v", *Open))
	log.Trace(fmt.Sprintf("dir: %s", *Dir))
}
