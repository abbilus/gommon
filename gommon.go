package main

import (
	"github.com/abbilus/gommon/log"
	"github.com/abbilus/gommon/telegram"
)

func main() {
	botToken := "444557213:AAG0T1F-N3UxxcTJKNrO2GVDAqv_uLf_BSc"
	b := telegram.New(botToken, "-562352213", true)
	log.SetNotifier(&b)
	log.EnableColor()
	log.Info("test info")
	log.Errorf("test error %v", "message")
	log.Warn("test warn message")
}

// func newNotifier(chatid, bottoken string) (notifier Notifier) {
// 	notifier.
// 	return
// }

// type Notifier struct {
// 	chatid   string
// 	bottoken string
// 	bot      telegram.Bot
// }

// func (n *Notifier) Init(chatid)

// func (n *Notifier) Send(message string) {

// }
