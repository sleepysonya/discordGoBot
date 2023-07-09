package register

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/scraly/learning-go-by-examples/go-gopher-bot-discord/birthday"
	randomCommand "github.com/scraly/learning-go-by-examples/go-gopher-bot-discord/random"
	"github.com/scraly/learning-go-by-examples/go-gopher-bot-discord/reminder"
	storage "github.com/scraly/learning-go-by-examples/go-gopher-bot-discord/util"
)

var (
	GuildID        string = storage.GuildID
	BotToken       string = storage.Token
	RemoveCommands bool   = storage.RemoveCommands
	rolls          int64
	sides          int64
	YearNow        float64 = float64(time.Now().Year())
)

var s *discordgo.Session

func init() {
	fmt.Println(YearNow)

	var err error
	s, err = discordgo.New("Bot " + BotToken)
	println("Bot Token: " + BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

var (
	integerOptionMin               = 1.0
	integerOptionMinYear           = 1900.0
	dmPermission                   = false
	defaultMemberPermissions int64 = discordgo.PermissionManageServer

	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "random",
			Description: "Random Roll",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "times",
					Description: "Number of times",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "sides",
					Description: "Number of sides",
					Required:    true,
				},
			},
		},
		{
			Name:        "remind-in",
			Description: "Remind command",
			Type:        discordgo.ChatApplicationCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "message",
					Description: "Message to remind",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "in",
					MinValue:    &integerOptionMin,
					Description: "Time",
					Required:    true,
				},
				{
					Name:        "time",
					Description: "Time",
					Type:        discordgo.ApplicationCommandOptionInteger,
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{

							Name:  "minute(s)",
							Value: 10,
						},
						{
							Name:  "hour(s)",
							Value: 3600,
						},
						{
							Name:  "days",
							Value: 86400,
						},
					},
				},
			},
		},
		{
			Name:        "birthday",
			Description: "Birthday command",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "day",
					Description: "Day of your birthday",
					MinValue:    &integerOptionMin,
					MaxValue:    31,
					Required:    true,
				},

				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "month",
					Description: "Month of your birthday",
					MinValue:    &integerOptionMin,
					MaxValue:    12,
					Required:    true,
				},
			},
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"birthday": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			fmt.Println("Birthday command called")
			day := fmt.Sprint(i.ApplicationCommandData().Options[0].Value)
			month := fmt.Sprint(i.ApplicationCommandData().Options[1].Value)

			message_response := birthday.AddCol(string(i.Member.User.ID), day, month)
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: message_response,
				},
			})
		},
		"remind-in": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			fmt.Println("Remind command called")
			var message string
			var futureTime, secValue int
			for _, option := range i.ApplicationCommandData().Options {
				switch option.Name {
				case "time":
					vi, err := strconv.Atoi(fmt.Sprint(option.Value))
					if err != nil {
						fmt.Println(err)
					}
					futureTime = vi
					break
				case "message":
					message = fmt.Sprint(option.Value)
					break
				case "in":
					vi, err := strconv.Atoi(fmt.Sprint(option.Value))
					if err != nil {
						fmt.Println(err)
					}
					secValue = vi
					break
				}
			}
			userId := string(i.Member.User.ID)
			channel := string(i.ChannelID)
			secondsNow := int64(time.Now().Unix())
			newTime := secondsNow + int64(futureTime*secValue)
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Reminder added",
				},
			})
			reminder.StartReminder(newTime, message, userId, channel)
		},
		"random": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			rolls = i.ApplicationCommandData().Options[0].IntValue()
			sides = i.ApplicationCommandData().Options[1].IntValue()
			fmt.Println("Random command called")
			// Roll dice
			var total, eachRollsString = randomCommand.Random(rolls, sides)
			// Send response
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("You rolled %v times with %v max. Total: %v\n %v", rolls, sides, total, eachRollsString),
				},
			})
		},
	}
)

func init() {

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}

		fmt.Println("Received interaction")
	})
}

func P() {
	fmt.Println("Starting bot...")
	s, er := discordgo.New("Bot " + BotToken)
	fmt.Println(BotToken)

	if er != nil {
		log.Fatalf("Cannot create a session: %v", er)
	}
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
		return
	}

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	if RemoveCommands {
		log.Println("Removing commands...")
		// // We need to fetch the commands, since deleting requires the command ID.
		// // We are doing this from the returned commands on line 375, because using
		// // this will delete all the commands, which might not be desirable, so we
		// // are deleting only the commands that we added.
		// registeredCommands, err := s.ApplicationCommands(s.State.User.ID, GuildID)
		// if err != nil {
		// 	log.Fatalf("Could not fetch registered commands: %v", err)
		// }

		for _, v := range registeredCommands {
			err := s.ApplicationCommandDelete(s.State.User.ID, GuildID, v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}

	log.Println("Gracefully shutting down.")
}
