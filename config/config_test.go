package config

import (
	"reflect"
	"testing"
)

func TestLoadConfiguration(t *testing.T) {
	var config Config
	err := LoadConfiguration("../config.json", &config)
	if err != nil {
		t.Error(err)
	}
	defaultConfig := Config{
		TelegramToken:          "write-your-telegram-token-here",
		TelegramBotName:        "golosovalochka_bot",
		Account:                "golosovalochka",
		PostingKey:             "5...",
		ActiveKey:              "5...",
		TextRuToken:            "",
		ReferralFee:            5.0,
		RequiredVotes:          4,
		InitialUserRating:      10,
		MaximumOpenedVotes:     3,
		MaximumUserVotesPerDay: 4,
		MinimumPostLength:      500,
		Developer:              "@babin",
		GroupID:                -1001143551951,
		GroupLink:              "https://t.me/joinchat/AlKeQUQpN8-9oShtaTcY7Q",
		DatabasePath:           "./db/database.db",
		Domains:                []string{"golos.io", "golos.blog", "goldvoice.club", "golosd.com", "golosdb.com", "mapala.net", "newbie.goloses.ru"},
		Chain:                  "golos",
		Rpc:                    []string{"wss://ws.golos.io", "wss://api.golos.cf"},
	}
	if !reflect.DeepEqual(defaultConfig, config) {
		t.Error("Конфиги не совпадают")
	}
}
