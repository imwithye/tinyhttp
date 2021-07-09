package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	log "github.com/sirupsen/logrus"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	verbose  = kingpin.Flag("verbose", "Verbose mode.").Short('v').Default("false").Bool()
	headless = kingpin.Flag("headless", "Headless(Cli) mode.").Default("false").Bool()
	host     = kingpin.Flag("host", "IP address to listen.").Short('h').Default("127.0.0.1").IP()
	port     = kingpin.Flag("port", "Port to listen.").Short('p').Default("0").Int()
	dir      = kingpin.Flag("dir", "Directory to server.").Short('d').Default(".").String()
	open     = kingpin.Flag("open", "Open in browser. Default true, --no-open to disable.").Default("true").Bool()
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func ParseArgs() {
	kingpin.Parse()
	if *verbose {
		log.SetLevel(log.TraceLevel)
	}
	absDir, err := filepath.Abs(*dir)
	if err != nil {
		log.Fatal(err.Error())
	}
	*dir = absDir
	log.Trace(fmt.Sprintf("verbose: %v", *verbose))
	log.Trace(fmt.Sprintf("headless: %v", *headless))
	log.Trace(fmt.Sprintf("host: %s", *host))
	log.Trace(fmt.Sprintf("port: %d", *port))
	log.Trace(fmt.Sprintf("dir: %s", *dir))
	log.Trace(fmt.Sprintf("open: %v", *open))
	log.Trace()
}

func NewTCPListener() net.Listener {
	log.Trace("create tcp listener, looking for available port")
	tcp, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatal(err.Error())
	}
	*port = tcp.Addr().(*net.TCPAddr).Port
	log.Trace(fmt.Sprintf("using port %d", *port))
	return tcp
}

func OpenBrowser(url string) {
	switch runtime.GOOS {
	case "linux":
		exec.Command("xdg-open", url).Start()
	case "windows":
		exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		exec.Command("open", url).Start()
	}
}

func main() {
	ParseArgs()
	tcp := NewTCPListener()
	url := fmt.Sprintf("http://%s:%d", *host, *port)
	log.Info(fmt.Sprintf("tinyhttp serving %s on HTTP port %s", *dir, url))
	http.Handle("/", http.FileServer(http.Dir(*dir)))
	if *open {
		go func() {
			<-time.After(100 * time.Millisecond)
			OpenBrowser(url)
		}()
	}
	err := http.Serve(tcp, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
