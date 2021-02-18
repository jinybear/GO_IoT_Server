package main

import (
	"GO_IoT_Server/router"
	"bufio"
	"clientTester/util"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/getsentry/sentry-go"
	"github.com/valyala/fasthttp"
)

var (
	config   = util.GetConfigInstance()
	addr     = flag.String("addr", config.Port.Http, "TCP address to listen to")
	compress = flag.Bool("compress", false, "Whether to enable transparent response compression")
)

func main() {
	defer func() {
		log.Println("Server teardown")
	}()

	//InitSentry()

	h := router.RequestHandler

	if *compress {
		h = fasthttp.CompressHandler(h)
	}
	go func() {
		if err := fasthttp.ListenAndServe(*addr, h); err != nil {
			log.Fatalf("Error in ListenAndServe: %s", err)
		}
	}()

	// 외부 입력 받는 기능 실행
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("quit", text) == 0 {
			break
		}
	}
}

// sentry 초기화
func InitSentry() {
	// err := sentry.Init(sentry.ClientOptions{
	// 	Dsn: "https://21f2f756a7724954a37364e254f0a4d2@o473756.ingest.sentry.io/5509103",
	// })
	// if err != nil {
	// 	log.Fatalf("sentry.Init: %s", err)
	// }
}
