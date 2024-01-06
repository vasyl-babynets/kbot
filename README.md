# kbot 

A telegram bot that replaces Cyrillic characters in the message with Latin characters.

## Requirements

[go v1.20](https://go.dev/doc/install)

## Installation

```bash
go get #download modules
go build -ldflags "-X="github.dev/vasyl-babynets/kbot/cmd.appVersion=1.06 #buid app
```

## Create a telegram bot

Go to `BotFarther` in telegram and push the `START` button.

Enter a `/newbot` command.

Type a name for your bot.

Choose a username for your bot. It must end in `bot`.

Copy a telegram API token.

## Start the application

```bash
read -s TELE_TOKEN # enter a telegram API token
export TELE_TOKEN
./kbot start
```
## Examples

`hello` - a telegram bot will respond with message of appVersion

`qwerty` - a telegram bot will respond `йцукен`

## Telegram bot URL

[kbot](https://t.me/vasylbabynets_bot)

## CI/CD
