package lib

import "github.com/bwmarrin/discordgo"

type Command interface {
	Name() string
	Description() string
	Options() []*discordgo.ApplicationCommandInteractionDataOption
	Version() string
	Exec(ctx CommandContext) error
}

type CommandContext struct {
}