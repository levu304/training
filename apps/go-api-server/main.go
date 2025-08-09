package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

func connectWebSocket(addr *string, symbol *string, wg *sync.WaitGroup) {
	defer wg.Done()

	u := url.URL{Scheme: "wss", Host: *addr, Path: "/ws/" + *symbol + "@aggTrade"}
	log.Printf("Connecting to %s", u.String())

	done := make(chan struct{})
	reconnectDelay := 5 * time.Second

	for {
		c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

		if err != nil {
			log.Printf("Dial error for %s: %v, retrying in %v", *symbol, err, reconnectDelay)
			time.Sleep(reconnectDelay)
			continue
		}

		go func() {
			defer close(done)
			for {
				_, message, err := c.ReadMessage()
				if err != nil {
					log.Println("Read:", err)
					return
				}
				log.Printf("recv: %s", message)
			}
		}()

		// Wait for connection to close before attempting reconnect
		<-done
		c.Close()
	}
}

func main() {
	addr := flag.String("addr", "fstream.binance.com", "http service address")
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	symbols := []string{"btcusdt", "ethusdt", "bnbusdt", "xrpusdt", "adausdt"}

	var wg sync.WaitGroup

	// Start connections
	for _, symbol := range symbols {
		wg.Add(1)
		s := symbol // Create new variable to avoid closure issues
		go connectWebSocket(addr, &s, &wg)
	}

	// Wait for interrupt signal
	<-interrupt
	log.Println("Received interrupt signal, shutting down...")

	// Wait for all goroutines to finish
	wg.Wait()
}
