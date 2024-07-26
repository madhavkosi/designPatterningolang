package main

import (
	"fmt"

	"gopkg.in/guregu/null.v4"
)

type Channels struct {
	ChannelName string `json:"channelName" binding:"required,ne=,min=1"`
	WebhookURL  string `json:"webhookUrl" binding:"required,ne=,min=1"`
}

type AlertRule struct {
	RuleId null.String `json:"alertRuleId,omitempty" db:"alertruleid" swaggertype:"string"`
}

type MSTeamsConfiguration struct {
	Channels []Channels `json:"channels" validate:"required,dive,required"`
}

type Email struct {
	Guest []string `json:"guest"`
}

type TeamsChannel struct {
	ChannelName string `json:"channelName"`
	WebhookURL  string `json:"webhookUrl"`
}

type Recipient struct {
	Emails        Email          `json:"emails"`
	TeamsChannels []TeamsChannel `json:"teamsChannels"`
}

type SomeAlertRule struct {
	AlertRule AlertRule
	Recipient Recipient
}

func main() {
	config := MSTeamsConfiguration{
		Channels: []Channels{
			{ChannelName: "General", WebhookURL: "https://example.com/webhook1"},
			{ChannelName: "Alerts", WebhookURL: "https://example.com/webhook2"},
		},
	}

	channelMap := make(map[string]string)
	for _, channel := range config.Channels {
		channelMap[channel.ChannelName] = channel.WebhookURL
	}
	channelWebHookUrl := make(map[string]string)
	for _, channel := range config.Channels {
		channelMap[channel.ChannelName] = channel.WebhookURL
	}

	// Printing the map for demonstration
	for name, url := range channelMap {
		fmt.Printf("Channel Name: %s, Webhook URL: %s\n", name, url)
	}

	// Get and print the rules
	someAlertRules := getRules()
	for _, sar := range someAlertRules {
		fmt.Printf("AlertRule ID: %s, Recipients: %+v\n", sar.AlertRule.RuleId.String, sar.Recipient)
	}
	rules := getRules()

}

func checkMissingChannels(rules []SomeAlertRule, channelMap, channelWebHookUrl map[string]string) []SomeAlertRule {
	var rulesToBeUpdated []SomeAlertRule

	for _, rule := range rules {
		teamChannel := make([]TeamsChannel, 0)
		changed := false
		for _, channel := range rule.Recipient.TeamsChannels {
			storedWebHookUrl, exists1 := channelMap[channel.ChannelName]
			storedChannelName, exists2 := channelWebHookUrl[channel.WebhookURL]
			if !exists1 && !exists2 {
				changed = true
				continue
			} else if exists1 && storedWebHookUrl != channel.WebhookURL {
				changed = true
				teamChannel = append(teamChannel, TeamsChannel{ChannelName: channel.ChannelName, WebhookURL: storedWebHookUrl})
			} else if exists2 && storedChannelName != channel.ChannelName {
				changed = true
				teamChannel = append(teamChannel, TeamsChannel{ChannelName: storedChannelName, WebhookURL: storedWebHookUrl})
			} else {
				teamChannel = append(teamChannel, channel)
			}
		}

		if changed {
			rulesToBeUpdated = append(rulesToBeUpdated, SomeAlertRule{AlertRule: rule.AlertRule, Recipient: Recipient{Emails: rule.Recipient.Emails, TeamsChannels: rule.Recipient.TeamsChannels}})
		}
	}
	return rulesToBeUpdated
}

func getRules() []SomeAlertRule {
	// Sample data for AlertRules
	alertRule1 := AlertRule{RuleId: null.StringFrom("rule1")}
	alertRule2 := AlertRule{RuleId: null.StringFrom("rule2")}

	// Sample data for Recipients
	recipient1 := Recipient{
		Emails: Email{Guest: []string{"alice@example.com", "bob@example.com"}},
		TeamsChannels: []TeamsChannel{
			{ChannelName: "General", WebhookURL: "https://example.com/webhook/general"},
		},
	}
	recipient2 := Recipient{
		Emails: Email{Guest: []string{"charlie@example.com", "dave@example.com"}},
		TeamsChannels: []TeamsChannel{
			{ChannelName: "Alerts", WebhookURL: "https://example.com/webhook/alerts"},
		},
	}

	// Create a slice of SomeAlertRule
	someAlertRules := []SomeAlertRule{
		{AlertRule: alertRule1, Recipient: recipient1},
		{AlertRule: alertRule2, Recipient: recipient2},
	}

	return someAlertRules
}
