package roomba

import "time"

type (
	// ReportPayload is the payload Roomba will post to the report callback URL (see config.go for details)
	// It contains a mix of Config{}, Record{} and runtime information
	ReportPayload struct {
		ChannelID string     `json:"channel_id"`
		Datetime  time.Time  `json:"datetime"`
		PRs       []Entry    `json:"prs"`
		Reminders []Reminder `json:"reminders"`
	}

	Reminder struct {
		Date time.Time `json:"date"`
		Text string    `json:"text"`
	}
)
