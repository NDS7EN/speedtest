package main

import (
	"github.com/maddie/speedtest/config"
	"github.com/maddie/speedtest/database"
	"github.com/maddie/speedtest/web"

	log "github.com/sirupsen/logrus"
)

func main() {
	conf := config.Load()

	database.SetDBInfo(&conf)
	log.Fatal(web.ListenAndServe(&conf))
}
