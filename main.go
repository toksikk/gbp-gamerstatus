package main

import (
	"github.com/bwmarrin/discordgo"
	plugin "github.com/toksikk/gbp-gamerstatus/plugin"
)

var PluginName = "gamerstatus"
var PluginVersion = ""
var PluginBuilddate = ""

func Start(discord *discordgo.Session) {
	plugin.Start(discord)
}
