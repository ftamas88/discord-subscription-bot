# Discord Subscription Bot

[![Latest Release](https://github.com/ftamas88/discord-subscription-bot/actions/workflows/release.yml/badge.svg)](https://github.com/ftamas88/discord-subscription-bot/actions/workflows/release.yml)

## Table of Contents
- [About](#about)
- [Quickstart](#quickstart)
    + [Using local environment](#using-local-environment)


## About

This is an example Discord bot which can watch your server, check for certain ROLE updates and notify a user/channel if that role changes.

For example: Someone subscribed, gained a new role, and you want them to be notified.

### Environment variables

| Environment           | Description                                                 |
|-----------------------|-------------------------------------------------------------|
| `APPLICATION_ID`      | (OPTIONAL - NOT IN USE) Bot application Id                  |
| `PUBLIC_KEY`          | (OPTIONAL - NOT IN USE) Channel public key                  |
| `BOT_TOKEN`           | Insert your bot secret token here                           |
| `SUBSCRIBED_ROLE_IDS` | List of ids(numbers) for the role you want to watch         |
| `NOTIFICATION_TYPE`   | The target for the notification: either `user` or `channel` |
| `NOTIFICATION_ID`     | The target Id for the notification (user or channel)        |
| `GO_ENV`              | "development/production"                                    |
| `GO_LOG_LEVEL`        | "debug/info/warn/panic"                                     |


## Quickstart
### Using local environment

#### Setup: Downloads the linter and other go tools
```shell
task install-requirements
```

#### Run & build the application
```shell
task run
```
or

```shell
go run cmd/bot/bot.go
```
