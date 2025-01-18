package lib

import (
	"context"
	"sync"

	"github.com/bwmarrin/discordgo"
)

var logger = NewLogger("commander")

// Commander manages all the interaction and other types of commands
type Commander struct {
	commandsMap map[string]Command
	mu sync.RWMutex

	Ctx context.Context
	Cancel context.CancelFunc
}

/*
* Create a new instance of Commander
*/
func NewCommander() *Commander {
	ctx, cancel := context.WithCancel(context.Background())

	return &Commander{
		commandsMap: make(map[string]Command),
		Ctx: ctx,
		Cancel: cancel,
	}
}

/*
* Pass in a map of commands to add
*/
func (c *Commander) Add(cmds map[string]Command) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range cmds {
		c.commandsMap[k] = v
	} 
}

func (c *Commander) Delete(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.commandsMap, name)
}

func (c *Commander) Get(name string) Command {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if cmd, ok := c.commandsMap[name]; !ok {
		return cmd
	}
	return nil
}

func (c *Commander) HandleInteraction(i *discordgo.Interaction) error {
	cmd , exists := c.commandsMap[i.ApplicationCommandData().Name]
	if !exists {
		return nil
	}
	ctx := CommandContext{}
	if err := cmd.Exec(ctx); err != nil {
		return err
	}
	logger.Print("Command ", cmd.Name(), " executed by @", i.Message.Author.Username, "in <c:", i.Message.ChannelID, "> <g:", i.Message.GuildID, ">")
	return nil
}

