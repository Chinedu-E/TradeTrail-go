package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Chinedu-E/TradeTrail-go/internal/leaderboards"
	"github.com/Chinedu-E/TradeTrail-go/internal/middleware"
	"github.com/Chinedu-E/TradeTrail-go/internal/portfolios"
	"github.com/Chinedu-E/TradeTrail-go/internal/sessions"
	"github.com/Chinedu-E/TradeTrail-go/internal/storage"
	"github.com/Chinedu-E/TradeTrail-go/internal/transactions"
	"github.com/Chinedu-E/TradeTrail-go/internal/users"
	"github.com/Chinedu-E/TradeTrail-go/internal/watchlists"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := storage.BootstrapPostgres()

	gin.ForceConsoleColor()

	router := gin.Default()

	config := cors.Config{
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Access-Control-Allow-Origin"},
	}

	router.Use(cors.New(config))
	router.Use(middleware.VerifyToken())

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome App Server")
	})
	// Users
	userStore := users.NewUserStorage(db)
	userController := users.NewUserController(userStore)
	users.AddUserRoutes(router, userController)

	//Watch lists
	watchlistStore := watchlists.NewWatchListStorage(db)
	watchlistController := watchlists.NewWatchListController(watchlistStore)
	watchlists.AddWatchListRoutes(router, watchlistController)

	// Trading  Session
	sessionStore := sessions.NewSessionStorage(db)
	sessionController := sessions.NewSessionController(sessionStore)
	sessions.AddSessionRoutes(router, sessionController)

	// Leaderboards
	leaderStore := leaderboards.NewLeaderBoardStorage(db)
	leaderController := leaderboards.NewLeaderBoardController(leaderStore)
	leaderboards.AddLeaderboardRoutes(router, leaderController)

	// Portfolios
	portfolioStore := portfolios.NewPortfolioStorage(db)
	portfolioController := portfolios.NewPortfolioController(portfolioStore)
	portfolios.AddPortfolioRoutes(router, portfolioController)

	// Transactions
	transactionStore := transactions.NewTransactionStorage(db)
	transactionController := transactions.NewTransactionController(transactionStore)
	transactions.AddTransactionRoutes(router, transactionController)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
