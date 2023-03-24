package config

import (
	"testing"
)

func TestNew(t *testing.T) {
	con := New()
	if con == nil {
		t.Error("config is nil")
	} else if con.LineMessageAPI.ChannelID == "" {
		t.Error("ChannelID is empty")
	} else if con.MongoDB.Url == "" {
		t.Error("MongoDB url is empty")
	}
}
