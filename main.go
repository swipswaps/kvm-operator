package main

import (
	"os"

	"github.com/giantswarm/kvm-operator/flag"
	"github.com/giantswarm/microkit/command"
	"github.com/giantswarm/microkit/logger"
	microserver "github.com/giantswarm/microkit/server"
	"github.com/spf13/viper"

	"github.com/giantswarm/kvm-operator/server"
	"github.com/giantswarm/kvm-operator/service"
)

var (
	description string     = "The kvm-operator handles Kubernetes clusters running on a Kubernetes cluster."
	f           *flag.Flag = flag.New()
	gitCommit   string     = "n/a"
	name        string     = "kvm-operator"
	source      string     = "https://github.com/giantswarm/kvm-operator"
)

func main() {
	var err error

	// Create a new logger which is used by all packages.
	var newLogger logger.Logger
	{
		loggerConfig := logger.DefaultConfig()
		loggerConfig.IOWriter = os.Stdout
		newLogger, err = logger.New(loggerConfig)
		if err != nil {
			panic(err)
		}
	}

	// We define a server factory to create the custom server once all command
	// line flags are parsed and all microservice configuration is storted out.
	newServerFactory := func(v *viper.Viper) microserver.Server {
		// Create a new custom service which implements business logic.
		var newService *service.Service
		{
			serviceConfig := service.DefaultConfig()

			serviceConfig.Flag = f
			serviceConfig.Logger = newLogger
			serviceConfig.Viper = v

			serviceConfig.Description = description
			serviceConfig.GitCommit = gitCommit
			serviceConfig.Name = name
			serviceConfig.Source = source

			newService, err = service.New(serviceConfig)
			if err != nil {
				panic(err)
			}
			go newService.Boot()
		}

		// Create a new custom server which bundles our endpoints.
		var newServer microserver.Server
		{
			serverConfig := server.DefaultConfig()

			serverConfig.MicroServerConfig.Logger = newLogger
			serverConfig.MicroServerConfig.ServiceName = name
			serverConfig.MicroServerConfig.Viper = v
			serverConfig.Service = newService

			newServer, err = server.New(serverConfig)
			if err != nil {
				panic(err)
			}
		}

		return newServer
	}

	// Create a new microkit command which manages our custom microservice.
	var newCommand command.Command
	{
		commandConfig := command.DefaultConfig()

		commandConfig.Logger = newLogger
		commandConfig.ServerFactory = newServerFactory

		commandConfig.Description = description
		commandConfig.GitCommit = gitCommit
		commandConfig.Name = name
		commandConfig.Source = source

		newCommand, err = command.New(commandConfig)
		if err != nil {
			panic(err)
		}
	}

	daemonCommand := newCommand.DaemonCommand().CobraCommand()

	daemonCommand.PersistentFlags().String(f.Service.Kubernetes.Address, "http://127.0.0.1:6443", "Address used to connect to Kubernetes. When empty in-cluster config is created.")
	daemonCommand.PersistentFlags().Bool(f.Service.Kubernetes.InCluster, false, "Whether to use the in-cluster config to authenticate with Kubernetes.")
	daemonCommand.PersistentFlags().String(f.Service.Kubernetes.TLS.CaFile, "", "Certificate authority file path to use to authenticate with Kubernetes.")
	daemonCommand.PersistentFlags().String(f.Service.Kubernetes.TLS.CrtFile, "", "Certificate file path to use to authenticate with Kubernetes.")
	daemonCommand.PersistentFlags().String(f.Service.Kubernetes.TLS.KeyFile, "", "Key file path to use to authenticate with Kubernetes.")

	newCommand.CobraCommand().Execute()
}