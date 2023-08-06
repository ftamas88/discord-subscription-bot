// Package main Entry point for the bot, bootstraps the application
package main

import (
	"github.com/ftamas88/discord-subscription-bot/internal/bot"
	"github.com/ftamas88/discord-subscription-bot/internal/config"
	"github.com/ftamas88/discord-subscription-bot/internal/logger"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	l := logger.NewLogger()
	cfg := config.NewConfig()

	b := bot.NewBot(cfg, l)
	b.Run()
}
