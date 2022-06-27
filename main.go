package main

import (
	"context"
	"net"
	"strconv"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/webview/webview"
)

var (
	port         int
	portstr      string
	startup_succ bool = false
	syncch       chan int
	e            *echo.Echo
)

func init() {
	var err error
	port, err = GetFreePort()
	if err == nil {
		startup_succ = true
	}
	if startup_succ {
		go run_server()
	}
}

func run_server() {
	e = echo.New()
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "static",
		HTML5:  true,
		Index:  "CyberChef_v9.39.1.html",
		Browse: true,
	}))
	portstr = strconv.FormatInt(int64(port), 10)
	e.Start("127.0.0.1:" + portstr)
}

func main() {

	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("CyberChef")
	w.SetSize(1080, 768, webview.HintNone)
	if startup_succ {
		w.SetHtml(`<meta http-equiv="Refresh" content="0; url='http://127.0.0.1:` + portstr + `'" />`)
	} else {
		w.SetHtml(`<h1>Oh No!</h1>
		Could not connect to Localhost on port <b>` + portstr + `</b><br>
		Try relaunching the app.`)
	}
	w.Run()
	e.Shutdown(context.Background())
}

// src: https://github.com/phayes/freeport/blob/master/freeport.go
// GetFreePort asks the kernel for a free open port that is ready to use.
func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
