package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	"app/pkg/store/postgres"
	"app/pkg/tracing"
	"app/pkg/type/context"
	log "app/pkg/type/logger"
	deliveryGrpc "app/services/contact/internal/delivery/grpc"
	deliveryHttp "app/services/contact/internal/delivery/http"
	repositoryStorage "app/services/contact/internal/repository/storage/postgres"
	useCaseContact "app/services/contact/internal/useCase/contact"
	useCaseGroup "app/services/contact/internal/useCase/group"
)

func init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetDefault("SERVICE_NAME", "contactService")
}

func main() {
	conn, err := postgres.New(postgres.Settings{})
	if err != nil {
		panic(err)
	}
	defer conn.Pool.Close()

	closer, err := tracing.New(context.Empty())
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = closer.Close(); err != nil {
			log.Error(err)
		}
	}()

	// repoContact, err := repositoryContact.New(conn.Pool, repositoryContact.Options{})
	// if err != nil {
	// 	panic(err)
	// }
	// repoGroup, err := repositoryGroup.New(conn.Pool, repoContact)
	// if err != nil {
	// 	panic(err)
	// }

	var (
		repoStorage  = repositoryStorage.New(conn.Pool)
		ucContact    = useCaseContact.New(repoStorage)
		ucGroup      = useCaseGroup.New(repoStorage)
		_            = deliveryGrpc.New(ucContact, ucGroup)
		listenerHttp = deliveryHttp.New(ucContact, ucGroup)
	)
	go func() {
		fmt.Printf("service started successfully on http port: %d", viper.GetUint("HTTP_PORT"))
		if err := listenerHttp.Run(); err != nil {
			panic(err)
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	<-signalCh

}
