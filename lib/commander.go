package lib

import (
	"context"
	"sync"
)

// Commander manages all the interaction and other types of commands
type Commander struct {
	commandsMap map[string]Command
	mu sync.RWMutex

	Ctx context.Context
	Cancel context.CancelFunc
}

func NewCommander() *Commander {
	ctx, cancel := context.WithCancel(context.Background())

	return &Commander{
		commandsMap: make(map[string]Command),
		Ctx: ctx,
		Cancel: cancel,
	}
}
