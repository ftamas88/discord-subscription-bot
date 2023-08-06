// Package bot Discord subscription bot
package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/ftamas88/discord-subscription-bot/internal/config"
	"github.com/ftamas88/discord-subscription-bot/internal/model"
	"github.com/ftamas88/discord-subscription-bot/internal/storage"
	"go.uber.org/zap"
)

type Bot struct {
	cfg        *config.AppConfig
	logger     *zap.SugaredLogger
	repository storage.Repository
	roles      map[string]string
	users      map[string]model.Member
}

func NewBot(cfg *config.AppConfig, logger *zap.SugaredLogger) *Bot {
	rep := storage.NewSqlite(logger)

	return &Bot{
		cfg:        cfg,
		logger:     logger,
		repository: rep,
		roles:      make(map[string]string),
		users:      make(map[string]model.Member),
	}
}

func (b *Bot) Run() {
	discord, err := discordgo.New(
		fmt.Sprint("Bot ", b.cfg.Token),
	)
	if err != nil {
		b.logger.Fatalf("error creating Discord session: %v", err)
	}

	defer func() {
		if err := discord.Close(); err != nil {
			b.logger.Fatalf("error closing Discord session: %v", err)
		}
	}()

	discord.AddHandler(b.startup)
	discord.AddHandler(b.roleChange)
	discord.AddHandler(b.messageReceive)

	discord.Identify.Intents = discordgo.IntentsAll

	// Open a websocket connection to Discord and begin listening.
	err = discord.Open()
	if err != nil {
		b.logger.Fatalf("error opening connection %v", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	b.logger.Infoln("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
