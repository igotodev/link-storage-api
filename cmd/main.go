package main

import (
	"link-storage-api/internal/handler"
	"link-storage-api/internal/service"
	"link-storage-api/internal/storage/db"
	"link-storage-api/pkg/config"
	"link-storage-api/pkg/psql"
	"link-storage-api/pkg/router"
	"log"
	"net/http"
)

func main() {
	cfg := config.GetConfig()

	addr := cfg.Listen.BindIP + ":" + cfg.Listen.Port
	mux := router.Router(addr)

	myDB := psql.NewPSLQ(cfg)
	storage := db.NewStorage(myDB)
	appService := service.NewService(storage)

	appHandler := handler.NewHandler(mux, appService)
	appRouting := appHandler.RegisterRouting()

	log.Fatal(http.ListenAndServe(addr, appRouting))
}
