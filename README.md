# Slack File Bot

## Table of Contents

- [About](#about)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [License](./LICENSE)

## About

The Slack File Bot is a simple bot that automates file uploads to a specified Slack channel. This project demonstrates how to use the Slack API to programmatically upload files to a Slack workspace, making it easier to manage and share documents within your team.

## Getting Started

These instructions will help you get the Slack File Bot up and running on your local machine for development and testing purposes.

### Prerequisites

You'll need the following to get started:

- Go (Golang) installed on your machine.
- A Slack workspace with a bot token.
- A Slack channel to upload files to.
- The Slack-Go SDK.

To install Go, follow the instructions [here](https://golang.org/doc/install).

Install the Slack-Go SDK by running:

```bash
#!/bin/bash
go get -u github.com/slack-go/slack
```

### Installing

1. Clone the repository:

   ```bash
   #!/bin/bash
   git clone https://github.com/daking24/slack-file-bot.git
   cd slack-file-bot
   ```

2. Set up environment variables:

   You need to set the Slack bot token and the channel ID where files will be uploaded.

   ```bash
   #!/bin/bash
   export SLACK_BOT_TOKEN="your-slack-bot-token"
   export CHANNEL_ID="your-channel-id"
   ```

3. Run the application:
  
   ```bash
   #!/bin/bash
   go run main.go
   ```

   This will upload the file specified in the code (`faith.pdf`) to the specified Slack channel.

## Usage

To use the Slack File Bot, ensure that you have configured your environment variables with the correct Slack bot token and channel ID. When you run the bot, it will automatically upload the files specified in the `fileArr` array to the designated Slack channel.

You can customize the files to be uploaded by modifying the `fileArr` variable in the `main.go` file. The bot will iterate over the files in this array and upload each one to your chosen channel.
