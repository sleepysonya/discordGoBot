package reminder

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	storage "github.com/sleepysonya/discordGoBot/util"
)

var (
	BotToken string = storage.GetEnvVar("BOT_TOKEN")
)
var s *discordgo.Session

// this logic will stay here, since we can try to use every minute check
// however, I since the number of requests is small, we can assume that
// the bot can keep more stuff in it's stack memory
// and do not run every second etc etc
func AddReminder(message string, seconds int, period int, id string, channel string) int64 {
	secs := time.Now().Unix()
	fmt.Println(secs)
	newTime := secs + int64(seconds*period)
	fmt.Println(newTime)
	f, err := os.OpenFile("reminders.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	f.WriteString(fmt.Sprintf("%s,%s. %s,%d\n", id, channel, message, newTime))
	f.Sync()
	f.Close()
	fmt.Printf("Reminder set for %s", time.Unix(newTime, 0))
	fmt.Println("Reminder added")
	return newTime
}

func StartReminder(newTime int64, message string, id string, channel string) {
	ctx := context.Background()
	until, _ := time.Parse(time.RFC3339, time.Unix(newTime+2, 0).Format(time.RFC3339))
	waitUntil(ctx, until)
	var err error
	s, err = discordgo.New("Bot " + BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = s.ChannelMessageSend(channel, fmt.Sprintf("Hey <@%s>, *This is a reminder about*:\n%s", id, message))
	if err != nil {
		fmt.Println(err)
		return
	}

}

func waitUntil(ctx context.Context, until time.Time) {
	timer := time.NewTimer(time.Until(until))
	defer timer.Stop()

	select {
	case <-timer.C:
		return
	case <-ctx.Done():
		return
	}
}
