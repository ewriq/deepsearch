package main

import (
	"deepsearch/cmd/api"
	"deepsearch/pkg"
	"log"
	"os"
	"time"

	"github.com/getlantern/systray"
)

func onReady() {
	
	iconData, err := os.ReadFile("assets/favicon.ico")
	if err != nil {
		log.Fatal(err)
	}

	systray.SetIcon(iconData)
	systray.SetTitle("zaaa")
	go  api.Start()

	
	google, errs := pkg.Google("DeepSearch")
	if errs != nil {
		log.Printf("Error fetching Google results: %v", errs)
	} else {
		log.Printf("Google results: %v", google)
	}

	google, errss := pkg.Yandex("DeepSearch")
	if errss != nil {
		log.Printf("Error fetching Yandex results: %v", errss)
	} else {
		log.Printf("Yandex results: %v", google)
	}
	bing, er := pkg.Bing("DeepSearch")
	if er != nil {
		log.Printf("Error fetching Bing results: %v", err)
	} else {
		log.Printf("Bing results: %v", bing)
	}
	log.Printf("Loaded %d bytes", len(iconData))

	go func() {
		time.Sleep(100 * time.Millisecond)
		systray.SetIcon(iconData)
	}()

	quit := systray.AddMenuItem("Quit", "Quit the app")

	go func() {
		<-quit.ClickedCh
		systray.Quit()
	}()
}

func onExit() {

}

func main() {
	systray.Run(onReady, onExit)
}
