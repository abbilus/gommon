package main

import (
	"fmt"
	"github.com/abbilus/gommon/telegram"
)

func main(){
	botToken := "444557213:AAG0T1F-N3UxxcTJKNrO2GVDAqv_uLf_BSc"
	b := telegram.New(botToken, "-375723579", true)
	err := b.Send("**text** test message\nnew line","")

	fmt.Printf("err: %+v\n", err)
}