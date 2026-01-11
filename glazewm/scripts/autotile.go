package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/websocket"
)

type GwmEvent struct {
	Data struct {
		EventType        string          `json:"eventType"`
		FocusedContainer json.RawMessage `json:"focusedContainer"`
	} `json:"data"`
}

type Container struct {
	Type     string      `json:"type"`
	HasFocus bool        `json:"hasFocus"`
	Width    float64     `json:"width"`
	Height   float64     `json:"height"`
	Children []Container `json:"children"`
}

func main() {
	exePath, err := os.Executable()
	if err != nil {
		os.Exit(1)
	}

	logPath := filepath.Join(filepath.Dir(exePath), "autotile.log")
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(f)
		defer f.Close()
	}

	uri := "ws://127.0.0.1:6123"

	for {
		if err := run(uri); err != nil {
			log.Printf("Socket error: %v. Retrying...", err)
			time.Sleep(5 * time.Second)
		}
	}
}

func run(uri string) error {
	conn, _, err := websocket.DefaultDialer.Dial(uri, nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	subs := []string{"focus_changed", "focused_container_moved", "application_exiting"}
	for _, s := range subs {
		if err := conn.WriteMessage(websocket.TextMessage, []byte("sub -e "+s)); err != nil {
			return err
		}
	}

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return err
		}

		var ev GwmEvent
		if err := json.Unmarshal(msg, &ev); err != nil {
			continue
		}

		if ev.Data.EventType == "application_exiting" {
			os.Exit(0)
		}

		if ev.Data.EventType == "focus_changed" || ev.Data.EventType == "focused_container_moved" {
			var root Container
			if err := json.Unmarshal(ev.Data.FocusedContainer, &root); err != nil {
				continue
			}

			if focused := findFocused(&root); focused != nil {
				if focused.Width > 0 && focused.Height > 0 {
					dir := "vertical"
					if focused.Width > focused.Height {
						dir = "horizontal"
					}
					cmd := fmt.Sprintf("command set-tiling-direction %s", dir)
					_ = conn.WriteMessage(websocket.TextMessage, []byte(cmd))
				}
			}
		}
	}
}

func findFocused(c *Container) *Container {
	if c.Type == "window" && c.HasFocus {
		return c
	}
	for _, child := range c.Children {
		if res := findFocused(&child); res != nil {
			return res
		}
	}
	return nil
}
