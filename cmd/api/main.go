package main

import (
	"context"

	// "github.com/soheilhy/cmux" // Will run Article and Vote services on HTTP and GRPC ports using cmux

	"github.com/gorilla/mux"
	"github.com/pronuu/lincoln/internal/config"
	"github.com/pronuu/lincoln/internal/database"
	"github.com/pronuu/lincoln/internal/server"
	"github.com/pronuu/lincoln/internal/utils"
	"github.com/pronuu/lincoln/pkg/article"
	"github.com/pronuu/lincoln/pkg/vote"

	_ "github.com/joho/godotenv/autoload" // Autoload environment variables from file
)

func main() {
	// Create context, done channel, and logger
	ctx, cancel := context.WithCancel(context.Background())
	done := utils.GetDoneChannel()
	lgr := utils.GetLogger("lincoln")

	// Create store
	str, err := database.NewStore()
	if err != nil {
		lgr.Log("err", "Database connection failed")
		cancel()
		return
	}

	// Create services
	article := article.NewArticleService(str, lgr)
	vote := vote.NewVoteService(str, lgr)

	// Create router
	rtr := mux.NewRouter()
	api := rtr.PathPrefix("/api/v1").Subrouter()
	article.Route(api)
	vote.Route(api)

	// Create config and server
	cfg := config.NewConfig()
	srv := server.NewHTTP(cfg, lgr, rtr)
	defer srv.Stop(ctx)

	// Start server and bind until exit
	go srv.Start(ctx)
	<-*done
	cancel()
}
