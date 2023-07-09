﻿# Discord GO Bot for my server

 ## Why so weird?
 I am learning and testing some packages and approach in go, thus I am trying to implement as many technologies as possible to see their advantages and drawbacks

 ## How can I use this bot?

 Well, if you are a stranger that wants to play with it or use for your own server

 1. add a `Taskfile.yml`. it has to look like:
```yml
version: '3'

env: &env
  BOT_TOKEN: <Your Discord Bot Token>
  GUILDID: <Your Guild Id>
  OPEN_AI: <Your Open Ai Token>
  RMCM: <true or false if you want to clean commands after bot shuttdown>

tasks:
  build:
    desc: Build the app
    cmds:
      - GOFLAGS=-mod=mod go build -o bin/gopher-bot-discord main.go

  run:
    desc: Run the app
    cmds:
      - GOFLAGS=-mod=mod go run main.go -t $BOT_TOKEN -g $GUILDID -r $RMCM

  bot:
    desc: Execute the bot
    cmds:
      - ./bin/gopher-bot-discord -t $BOT_TOKEN -g $GUILDID -r $RMCM -o $OPEN_AI
```
2. Add .env file with a following var
```env
ORG_ID="<Open Ai Org Id>"
```
3. Open console with your bot and run `task run` it shall build a bot
