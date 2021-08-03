package cli

import (
	"flag"
	"log"
	"skeleton.service/database"
	"skeleton.service/env"
	"skeleton.service/routing"
	validators2 "skeleton.service/validators"
)

type serverRunCommand struct {
	flagSet *flag.FlagSet
}

func MakeServerRunCommand() *serverRunCommand {
	return &serverRunCommand{
		flagSet: flag.NewFlagSet("server-run", flag.ContinueOnError),
	}
}

func (cmd *serverRunCommand) Name() string {
	return cmd.flagSet.Name()
}

func (cmd *serverRunCommand) Init(args []string) error {
	return cmd.flagSet.Parse(args)
}

func (cmd *serverRunCommand) Run() error {
	log.Println("Starting system")

	env.Initialize()

	database.Init()

	validators2.Init()

	return routing.InitHttpServer(env.Port)
}
