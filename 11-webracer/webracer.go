package webracer

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func pingServer(u string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(u)
		close(ch)
	}()
	return ch
}

func WebsiteRacer(urlOne, urlTwo string) (string, error) {
	return ConfigurableWebsiteRacer(urlOne, urlTwo, tenSecondTimeout)
}

func ConfigurableWebsiteRacer(urlOne, urlTwo string, dt time.Duration) (string, error) {
	select {
	case <-pingServer(urlOne):
		return urlOne, nil
	case <-pingServer(urlTwo):
		return urlTwo, nil
	case <-time.After(dt):
		return "", fmt.Errorf("error timeout while waiting for %s and %s", urlOne, urlTwo)
	}
}
