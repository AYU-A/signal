package signal

import (
	"context"
	"log"
	"net/http"

	"github.com/AYU-A/signal/config"
	"github.com/AYU-A/signal/db"
	"github.com/AYU-A/signal/db/migrations"
	"github.com/AYU-A/signal/scheduler"
	"github.com/AYU-A/signal/server"
)

var s *server.Server

func init() {
	ctx := context.Background()

	// config
	cfg := config.Get()

	pgDB, err := db.Dial(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Run Postgres migrations
	if err := migrations.Run(pgDB); err != nil {
		log.Fatal(err)
	}

	// create google cloud scheduler client
	sch, err := scheduler.Init(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// create new server instance
	s = server.Init(ctx, cfg, pgDB, sch)
}

// ScheduleCall
func ScheduleCall(w http.ResponseWriter, r *http.Request) {
	s.ScheduleCall(w, r)
}

// CallRoom
func CallRoom(w http.ResponseWriter, r *http.Request) {
	s.CallRoom(w, r)
}
