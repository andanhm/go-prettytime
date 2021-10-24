//Package prettytime provides utility functions for calculating and presenting
//in human readable and understandable form.
package prettytime

import (
	"strconv"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

const (
	// Unix epoch (or Unix time or POSIX time or Unix timestamp)  1 year (365.24 days)
	infinity                  float64 = 31556926 * 1000
	yearTimePeriodMessageID           = "year"
	monthTimePeriodMessageID          = "month"
	weekTimePeriodMessageID           = "week"
	dayTimePeriodMessageID            = "day"
	hourTimePeriodMessageID           = "hour"
	minuteTimePeriodMessageID         = "minute"
	secondTimePeriodMessageID         = "second"
	fromSinceMessageID                = "from_now"
	agoSinceMessageID                 = "ago"
	tomorrowMessageID                 = "tomorrow"
	justNowMessageID                  = "just_now"
	yesterdayMessageID                = "yesterday"
)

// Handler function which determines the time difference based on defined time spams
func handler(timeIntervalThreshold float64, timePeriodMessageID, sinceMessageID string, localizer *i18n.Localizer) func(float64) string {
	return func(difference float64) string {
		var str strings.Builder
		n := difference / timeIntervalThreshold
		nStr := strconv.FormatFloat(n, 'f', 0, 64)

		timePeriod := localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: timePeriodMessageID,
			TemplateData: map[string]interface{}{
				"Value": nStr,
			},
			PluralCount: nStr,
		})

		since := localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: sinceMessageID,
			TemplateData: map[string]string{
				"Time": timePeriod,
			},
		})
		str.WriteString(since)
		return str.String()
	}
}

// timeLapse condition struct
type timeLapse struct {
	// Time stamp threshold to handle the time lap condition
	Threshold float64
	// Handler function which determines the time lapse based on the condition
	Handler func(float64) string
}

func initTimeLapse(localizer *i18n.Localizer) []timeLapse {
	return []timeLapse{
		{Threshold: -31535999, Handler: handler(-31536000, yearTimePeriodMessageID, fromSinceMessageID, localizer)},
		{Threshold: -2591999, Handler: handler(-2592000, monthTimePeriodMessageID, fromSinceMessageID, localizer)},
		{Threshold: -604799, Handler: handler(-604800, weekTimePeriodMessageID, fromSinceMessageID, localizer)},
		{Threshold: -172799, Handler: handler(-86400, dayTimePeriodMessageID, fromSinceMessageID, localizer)},
		{Threshold: -86399, Handler: func(diff float64) string {
			return localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: tomorrowMessageID,
			})
		}},
		{Threshold: -3599, Handler: handler(-3600, hourTimePeriodMessageID, fromSinceMessageID, localizer)},
		{Threshold: -59, Handler: handler(-60, minuteTimePeriodMessageID, fromSinceMessageID, localizer)},
		{Threshold: -0.9999, Handler: handler(-1, secondTimePeriodMessageID, fromSinceMessageID, localizer)},
		{Threshold: 1, Handler: func(diff float64) string {
			return localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: justNowMessageID,
			})
		}},
		{Threshold: 60, Handler: handler(1, secondTimePeriodMessageID, agoSinceMessageID, localizer)},
		{Threshold: 3600, Handler: handler(60, minuteTimePeriodMessageID, agoSinceMessageID, localizer)},
		{Threshold: 86400, Handler: handler(3600, hourTimePeriodMessageID, agoSinceMessageID, localizer)},
		{Threshold: 172800, Handler: func(diff float64) string {
			return localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: yesterdayMessageID,
			})
		}},
		{Threshold: 604800, Handler: handler(86400, dayTimePeriodMessageID, agoSinceMessageID, localizer)},
		{Threshold: 2592000, Handler: handler(604800, weekTimePeriodMessageID, agoSinceMessageID, localizer)},
		{Threshold: 31536000, Handler: handler(2592000, monthTimePeriodMessageID, agoSinceMessageID, localizer)},
		{Threshold: infinity, Handler: handler(31536000, yearTimePeriodMessageID, agoSinceMessageID, localizer)},
	}
}

//Format returns a string describing how long it has been since
//the time argument passed int
func Format(t time.Time, lang string) (timeSince string) {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("de.toml")
	localizer := i18n.NewLocalizer(bundle, lang)

	timeLapses := initTimeLapse(localizer)
	timestamp := t.Unix()
	now := time.Now().Unix()

	if timestamp > now || timestamp <= 0 {
		timeSince = ""
	}

	timeElapsed := float64(now - timestamp)
	for _, formatter := range timeLapses {
		if timeElapsed < formatter.Threshold {
			timeSince = formatter.Handler(timeElapsed)
			break
		}
	}
	return timeSince
}
