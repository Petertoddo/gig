package gig

import (
	"fmt"

	"github.com/jessevdk/go-flags"
	"github.com/pkg/errors"
)

type config struct {
	List    bool `short:"l" long:"list" description:"Show list of available language"`
	File    bool `short:"f" long:"File" description:"Output .gitignore file"`
	Quiet   bool `short:"q" long:"quiet" description:"Hide stdout"`
	Version bool `short:"v" long:"version" description:"Show version"`
	Help    bool `short:"h" long:"help" description:"Show this help message"`
	Args    struct {
		Language string
	} `positional-args:"yes"`
}

func (g *Gig) initConfig() error {
	p := flags.NewParser(&g.Config, flags.None)
	_, err := p.Parse()
	if err != nil {
		return errors.Wrapf(err, "failed to parse. Config: %s", &g.Config)
	}

	if g.Config.Version {
		return fmt.Errorf("gig version %s", g.Version)
	}

	if g.Config.Help ||
		(!g.Config.List && g.Config.Args.Language == "") {
		p.WriteHelp(g.ErrStream)
		return errors.New("")
	}

	return nil
}
