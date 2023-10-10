package storage

import (
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

func GetEnvVar(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

type Generation struct {
	Id            string `json:"id"`
	Text          string `json:"text"`
	Finish_reason string `json:"finish_reason"`
}
type CohereResponse struct {
	Id          string       `json:"id"`
	Generations []Generation `json:"generations"`
	Prompt      string       `json:"prompt"`
}
