package storage

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Token          string
	GuildID        string
	RemoveCommands bool
	OpenAiToken    string
)

type BirthType struct {
	Id    string
	Day   string
	Month string
}

type ReminderOptions struct {
	Name  string
	Value any
	Type  int
}
type ReminderInBoundData struct {
	Id      string
	Name    string
	Options []ReminderOptions
}

type ReminderReturn struct {
	Id   string
	Data ReminderInBoundData
}

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.StringVar(&GuildID, "g", "", "Guild ID")
	flag.BoolVar(&RemoveCommands, "r", false, "RemoveCommands")
	flag.StringVar(&OpenAiToken, "o", "", "OpenAiToken")
	flag.Parse()
}

func GetEnvVar(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}
