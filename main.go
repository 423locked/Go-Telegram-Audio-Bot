package main

import (
	"log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/423locked/mybot/ytdownloader"
	"github.com/423locked/mybot/addons"
	"strings"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5793099170:AAHq2SHYbSkgjbaaUa6gLaekhmw1SlBbUqY")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			go func() {
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
				
				videoID := update.Message.Text[strings.Index(update.Message.Text, "?v=")+3:]
				log.Printf(videoID)

				if (!addons.IsLinkValid(update.Message.Text)) {
					log.Printf("[ERR] Link invalid!")
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Invalid link!")
					msg.ReplyToMessageID = update.Message.MessageID
					bot.Send(msg)
				} else {
					bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Video downloading started!"))
					
					filepath := ytdownloader.DownloadAndGetPath()
					/*musicBytes, err := ioutil.ReadFile(filepath)
					if (err != nil) {
						panic(err)
					}
					musicFileBytes := tgbotapi.FileBytes {
						Name: "music",
						Bytes: musicBytes,
					}
					*/
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, filepath)
					msg.ReplyToMessageID = update.Message.MessageID
					bot.Send(msg)
					bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Video downloading ended!"))
					bot.Send(tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath(filepath)))
					
				}
			}()
		}
	}
}
