package di

import (
	"DistributedGrep/internal/config"
	"DistributedGrep/internal/network"
	"DistributedGrep/internal/quorum"
	"DistributedGrep/internal/service"
)

type App struct {
	Config  *config.Config
	Service *service.GrepService
	Server  *network.Server
	Client  *network.Client
	Quorum  *quorum.Quorum
}

func Build() (*App, error) {

	cfg := config.MustLoad("config/local.yaml")

	client := network.NewClient()

	svc := service.NewService(
		cfg.Workers,
		cfg.Server.Host,
	)

	server := network.NewServer(svc)

	q := quorum.New(client)

	return &App{
		Config:  cfg,
		Service: svc,
		Server:  server,
		Client:  client,
		Quorum:  q,
	}, nil
}
