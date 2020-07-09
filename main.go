package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

const token string = "NzMwNzc3MzE1MzUyNjQxNTY2.Xwdglg.Asw2As6uEoJJHPqnKgK4BjKExpY"

var BotID string

func main() {
	fmt.Println("Test")
	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println(err)
		return
	}
	u, err := dg.User("@me")
	if err != nil {
		fmt.Println(err)
		return
	}
	BotID = u.ID

	err = dg.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot Is rUnNiNg!")

	<-make(chan struct{})
	return
}
