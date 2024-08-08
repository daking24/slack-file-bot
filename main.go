package main

import (
	"fmt"
	"os"
	"github.com/slack-go/slack"
)

func main()  {
	
	os.Setenv("SLACK_BOT_TOKEN", "")
	os.Setenv("CHANNEL_ID", "")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channel := os.Getenv("CHANNEL_ID")
	fileArr := []string{"faith.pdf"}

	for i := 0; i<len(fileArr); i++{
		params := slack.UploadFileV2Parameters{
			Channel: channel,
			File: fileArr[i],
			Filename: "faith.pdf",
			FileSize: 229090,
		}

		file, err := api.UploadFileV2(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Title: %s, ID: %s\n", file.Title, file.ID)
		
	}
}