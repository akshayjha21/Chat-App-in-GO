package main

import (
	"context"
	// "fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	// "log/slog"
	"github.com/akshayjha21/Chat-App-in-GO/Backend/internal/chat"
	"github.com/akshayjha21/Chat-App-in-GO/Backend/internal/config"
	"github.com/akshayjha21/Chat-App-in-GO/Backend/internal/storage/postgres"
)

func main() {

	cfg := config.MustLoad()
	//db connection

	log.Printf("Conecting to the database")
	db,err:=postgres.New(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	//1. hub ko initialize krenge
	hub := chat.NewHub(db)
	// Get the generic sql.DB object from GORM to close it
	

	//2. hub ko background me execute krenge
	go hub.Run()
	router := http.NewServeMux()
	// 3. Route Handler Update karenge
	// Note: Hum seedha chat.ServeWs nahi de sakte kyunki usse 'hub' chahiye.
	// Isliye hum ek "Closure" (anonymous function) use karenge.

	router.HandleFunc("/ws",
		func(w http.ResponseWriter, r *http.Request) {
			chat.ServeWs(hub, w, r)
		})

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	go func() {
		log.Printf("Server starting on %s", cfg.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done
	log.Println("Server stopping.....")
	//  Graceful Shutdown (Give active connections 5 seconds to finish)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	if sqlDB, err := db.Db.DB(); err == nil {
		log.Println("Closing database connection...")
		sqlDB.Close()
	}

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exited properly")
}
