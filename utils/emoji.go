package utils

import "regexp"

func RemoveEmojis(input string) string {
	emojiRegex := regexp.MustCompile(`[\x{1F1E6}-\x{1F1FF}🏴󠁧󠁢󠁥󠁮󠁧󠁿]+`)
	return emojiRegex.ReplaceAllString(input, "")
}
