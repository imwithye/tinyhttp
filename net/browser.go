package net

import (
	"os/exec"
	"runtime"
	"time"
)

func OpenBrowser(url string) {
	go func() {
		<-time.After(100 * time.Millisecond)
		switch runtime.GOOS {
		case "linux":
			exec.Command("xdg-open", url).Start()
		case "windows":
			exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
		case "darwin":
			exec.Command("open", url).Start()
		}
	}()
}
