package addons

import (
	"regexp"
)

func IsLinkValid(link string) bool {
	re := "^((?:https?:)?\\/\\/)?((?:www|m)\\.)?((?:youtube\\.com|youtu.be))(\\/(?:[\\w\\-]+\\?v=|embed\\/|v\\/)?)([\\w\\-]+)(\\S+)?$"
	valid, _ := regexp.MatchString(re, link)
	return valid
}