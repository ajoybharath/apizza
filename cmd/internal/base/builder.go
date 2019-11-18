package base

import (
	"io"
	"os"

	"github.com/harrybrwn/apizza/pkg/cache"
	"github.com/harrybrwn/apizza/pkg/config"
)

// Builder defines an interface for an object that builds commands.
type Builder interface {
	Config() config.Config
	DB() *cache.DataBase
	Output() io.Writer
	Build(use, short string, r Runner) *Command
}

// BasicBuilder is a build that has no database, no config, no special output,
// only the ablility to build commands.
type BasicBuilder struct {
}

// Build a command.
func (bb *BasicBuilder) Build(use, short string, r Runner) *Command {
	return NewCommand(use, short, r.Run)
}

// Config does nothing.
func (bb *BasicBuilder) Config() config.Config { return nil }

// DB does nothing
func (bb *BasicBuilder) DB() *cache.DataBase { return nil }

// Output returns stdout
func (bb *BasicBuilder) Output() io.Writer { return os.Stdout }