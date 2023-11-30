/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var translationMap = map[rune]rune{
	// translation map
	'q':  'й',
	'w':  'ц',
	'e':  'у',
	'r':  'к',
	't':  'е',
	'y':  'н',
	'u':  'г',
	'i':  'ш',
	'o':  'щ',
	'p':  'з',
	'[':  'х',
	'{':  'Х',
	']':  'ї',
	'}':  'Ї',
	'|':  '/',
	'`':  '\'',
	'~':  '₴',
	'a':  'ф',
	's':  'і',
	'd':  'в',
	'f':  'а',
	'g':  'п',
	'h':  'р',
	'j':  'о',
	'k':  'л',
	'l':  'д',
	';':  'ж',
	':':  'Ж',
	'\'': 'є',
	'\\': 'Є',
	'z':  'я',
	'x':  'ч',
	'c':  'с',
	'v':  'м',
	'b':  'и',
	'n':  'т',
	'm':  'ь',
	',':  'б',
	'<':  'Б',
	'.':  'ю',
	'>':  'Ю',
	'/':  '.',
	'?':  ',',
	'@':  '\\',
	'#':  '№',
	'$':  ';',
	'^':  ':',
	'&':  '?',
	'Q':  'Й',
	'W':  'Ц',
	'E':  'У',
	'R':  'К',
	'T':  'Е',
	'Y':  'Н',
	'U':  'Г',
	'I':  'Ш',
	'O':  'Щ',
	'P':  'З',
	'A':  'Ф',
	'S':  'І',
	'D':  'В',
	'F':  'А',
	'G':  'П',
	'H':  'Р',
	'J':  'О',
	'K':  'Л',
	'L':  'Д',
	'Z':  'Я',
	'X':  'Ч',
	'C':  'С',
	'V':  'М',
	'B':  'И',
	'N':  'Т',
	'M':  'Ь',
}

var (
	// TeleToken bot
	TeleToken = os.Getenv("TELE_TOKEN")
)

// Translation function
func translateToUkrainian(text string) string {
	translatedRunes := make([]rune, 0, len(text))

	for _, char := range text {
		// If the symbol is in the match map, we replace it
		if translatedChar, ok := translationMap[char]; ok {
			translatedRunes = append(translatedRunes, translatedChar)
		} else {
			// If the symbol is not found, we leave it unchanged
			translatedRunes = append(translatedRunes, char)
		}
	}

	return string(translatedRunes)
}

// kbotCmd represents the kbot command
var kbotCmd = &cobra.Command{
	Use:     "kbot",
	Aliases: []string{"start"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("kbot %s started", appVersion)
		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatalf("Please check TELE_TOKEN env variable %s", err)
			return
		}
		kbot.Handle(telebot.OnText, func(m telebot.Context) error {

			log.Print(m.Message().Payload, m.Text())
			//	payload := m.Message().Payload
			text := m.Text()

			switch text {
			case "hello":
				err = m.Send(fmt.Sprintf("Hello I'm Kbot %s!", appVersion))
			default:
				translatedText := translateToUkrainian(text)
				err = m.Send(translatedText)
			}

			return err

		})
		kbot.Start()
	},
}

func init() {
	rootCmd.AddCommand(kbotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kbotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kbotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
