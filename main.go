package main

import (
	"fmt"
	"github.com/df-mc/dragonfly/server"
	srvcmd "github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
	"os"
	"uno/cmd"
	"uno/game"
)

func main() {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{ForceColors: true}
	log.Level = logrus.DebugLevel

	chat.Global.Subscribe(chat.StdoutSubscriber{})

	conf, err := readConfig(log)
	if err != nil {
		log.Fatalln(err)
	}
	registerCommands()
	srv := conf.New()
	srv.CloseOnProgramEnd()
	srv.Listen()
	for srv.Accept(func(p *player.Player) {
		p.Handle(game.AddToHandler(p))
		game.CreateSession(p)
	}) {
	}
}

// readConfig reads the configuration from the config.toml file, or creates the
// file if it does not yet exist.
func readConfig(log server.Logger) (server.Config, error) {
	c := server.DefaultConfig()
	var zero server.Config
	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		data, err := toml.Marshal(c)
		if err != nil {
			return zero, fmt.Errorf("encode default config: %v", err)
		}
		if err := os.WriteFile("config.toml", data, 0644); err != nil {
			return zero, fmt.Errorf("create default config: %v", err)
		}
		return zero, nil
	}
	data, err := os.ReadFile("config.toml")
	if err != nil {
		return zero, fmt.Errorf("read config: %v", err)
	}
	if err := toml.Unmarshal(data, &c); err != nil {
		return zero, fmt.Errorf("decode config: %v", err)
	}
	return c.Config(log)
}

func registerCommands() {
	srvcmd.Register(srvcmd.New("startgame", "", nil, cmd.StartCommand{}))
	srvcmd.Register(srvcmd.New("joingame", "", nil, cmd.JoinGame{}))
	srvcmd.Register(srvcmd.New("creategame", "", nil, cmd.CreateGame{}))
}
