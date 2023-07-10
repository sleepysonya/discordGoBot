module github.com/sleepysonya/discordGoBot

go 1.20

require (
	github.com/bwmarrin/discordgo v0.27.1
	github.com/go-co-op/gocron v1.30.1
	github.com/joho/godotenv v1.5.1
)

replace (
	github.com/sleepysonya/discordGoBot => ../discordGoBot
)

require (
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/crypto v0.11.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
)
