package helpers

import (
	"dh/api"
	"fmt"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/gotgbot/keyboard"
)

func GetAccessKey(message string) string {
	fields := strings.Fields(message)
	if len(fields) == 2 {
		kv := fields[1]
		if len(kv) > 13 {
			return kv[13:]
		}
	}
	return ""
}

func EnsureIndent(indent int, len int) int {
	if !(len > indent) {
		indent = 0
	}
	return indent
}

func GetReplyMarkup(track api.Track) *gotgbot.InlineKeyboardMarkup {
	return new(
		keyboard.InlineKeyboard,
	).Url(
		"Play on Deezer",
		track.Link,
	).Url(
		"Album",
		track.Album.Link,
	).Row().Url(
		"Share",
		fmt.Sprintf("https://t.me/share/url?url=%s", track.Link),
	).Build()
}