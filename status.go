package main

import (
	"log/slog"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

var PluginName = "gamerstatus"
var PluginVersion = ""
var PluginBuilddate = ""

func Start(discord *discordgo.Session) {
	go setIdleStatus(discord)
}

func setIdleStatus(discord *discordgo.Session) {
	games := []string{
		"Terranigma",
		"Secret of Mana",
		"Quake III Arena",
		"Duke Nukem 3D",
		"Monkey Island 2: LeChuck's Revenge",
		"Turtles in Time",
		"Unreal Tournament",
		"Half-Life",
		"Half-Life 2",
		"Warcraft II",
		"Starcraft",
		"Diablo",
		"Diablo II",
		"A Link to the Past",
		"Ocarina of Time",
		"Star Fox",
		"Tetris",
		"Pokémon Red",
		"Pokémon Blue",
		"Die Siedler II",
		"Day of the Tentacle",
		"Maniac Mansion",
		"Prince of Persia",
		"Super Mario Kart",
		"Pac-Man",
		"Frogger",
		"Donkey Kong",
		"Donkey Kong Country",
		"Asteroids",
		"Doom",
		"Breakout",
		"Street Fighter II",
		"Wolfenstein 3D",
		"Mega Man",
		"Myst",
		"R-Type",
	}
	for {
		err := discord.UpdateStreamingStatus(1, "", "")
		if err != nil {
			slog.Error("Could not set streaming status", "error", err)
		}
		err = discord.UpdateGameStatus(0, games[randomRange(0, len(games))])
		if err != nil {
			slog.Error("Could not set game status", "error", err)
		}
		time.Sleep(time.Duration(randomRange(5, 15)) * time.Minute)
	}
}

func randomRange(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}
