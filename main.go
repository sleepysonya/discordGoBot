package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
	"github.com/scraly/learning-go-by-examples/go-gopher-bot-discord/birthday"
	storage "github.com/scraly/learning-go-by-examples/go-gopher-bot-discord/util"

	register "github.com/scraly/learning-go-by-examples/go-gopher-bot-discord/test"
)

var (
	GuildID        string = storage.GuildID
	BotToken       string = storage.Token
	RemoveCommands bool   = storage.RemoveCommands
)

const KuteGoAPIURL = "https://ibb.co/tMqStgz"

// reference to discord session
var s *discordgo.Session

func init() {
	schedule := gocron.NewScheduler(time.UTC)

	_, err := schedule.CronWithSeconds("0 0 14 * * *").Do(func() {
		data := birthday.GetAllCols()
		for _, record := range data {
			fmt.Println(record)
			day := record.Day
			month := record.Month
			today := time.Now().Day()
			todayMonth := int(time.Now().Month())
			if day == strconv.Itoa(today) && month == strconv.Itoa(todayMonth) {
				var err error
				s, err = discordgo.New("Bot " + BotToken)
				if err != nil {
					fmt.Println("error creating Discord session,", err)
					return
				}
				_, err = s.ChannelMessageSend("1052656598046220321", "Meow, Today is <@"+record.Id+"> birthday! Time to nom <:BlobCatHeart:1057001702072537159>")
				if err != nil {
					fmt.Println("error sending message,", err)
				}
				err = s.Close()
				if err != nil {
					fmt.Println("error closing session,", err)
				}
			}

		}
	})
	if err != nil {
		panic(err)
	}
	schedule.StartAsync()
}

func main() {
	//flag.BoolVar(&FlagValue, "rmcmd", true, "description of the flag")
	// Create a new Discord session using the provided bot token.
	fmt.Println("Starting bot...")
	register.P()
}
