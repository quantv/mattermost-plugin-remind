package main

import (
	"time"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/nicksnyder/go-i18n/i18n"
)

func (p *Plugin) translation(user *model.User) (i18n.TranslateFunc, string) {
	locale := "en"
	for l := range p.locales {
		if user.Locale == l {
			locale = user.Locale
			break
		}
	}
	return p.GetUserTranslations(locale), locale
}

func (p *Plugin) location(user *model.User) *time.Location {
	tz := user.GetPreferredTimezone()
	//Fixed deprecated timezone
	if tz == "Asia/Saigon" {
		tz = "Asia/Ho_Chi_Minh"
	}
	if tz == "" {
		// Use server's timezone
		return time.Now().Location()
	} else {
		location, err := time.LoadLocation(tz)
		if err != nil {
			p.API.LogWarn("Failed to load user timezone, using server timezone",
				"user_id", user.Id, "tz", tz, "error", err)
			return time.Now().Location() // Fallback to server TZ
		}
		return location
	}

}
