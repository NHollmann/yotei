package controller

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/NHollmann/yotei/model"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type YoteiServer struct {
	db     *gorm.DB
	router *gin.Engine
}

type DatabaseConfig struct {
	Driver   string
	User     string
	Password string
	Port     uint16
	Host     string
	Database string
}

func (server *YoteiServer) Start(addr string, dbConfig DatabaseConfig) {
	server.initDatabase(dbConfig)
	server.initHttpServer()
	server.run(addr)
}

func (server *YoteiServer) initDatabase(dbConfig DatabaseConfig) {
	var err error

	if dbConfig.Driver == "mysql" {
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			dbConfig.User,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.Database,
		)
		server.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else if dbConfig.Driver == "sqlite" {
		server.db, err = gorm.Open(sqlite.Open(dbConfig.Database), &gorm.Config{})
	} else {
		err = fmt.Errorf("Unknown db driver: %s", dbConfig.Driver)
	}
	if err != nil {
		log.Fatal("Error connecting db: ", err)
	}

	log.Printf("Connected to %s database: %s", dbConfig.Driver, dbConfig.Database)

	err = server.db.AutoMigrate(
		&model.Event{},
		&model.Participant{},
		&model.ParticipantChoice{},
		&model.Timeslot{},
		&model.User{},
	)
	if err != nil {
		log.Fatal("Error on migrating database:  ", err)
	}
}

func (server *YoteiServer) initHttpServer() {
	server.router = gin.New()
	server.router.Use(gin.Logger())
	server.router.Use(gin.Recovery())
	server.initRoutes()
}

func (server *YoteiServer) run(addr string) {
	httpServer := &http.Server{
		Addr:    addr,
		Handler: server.router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		log.Printf("Listening on %s", addr)
		if err := httpServer.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("Server: %s", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
