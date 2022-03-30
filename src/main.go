package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"oraculo/config"
	"oraculo/remotes/mongo"
	"oraculo/web/router"
	"oraculo/web/server"

	"github.com/gorilla/mux"

	"github.com/joaopandolfi/blackwhale/handlers"
	"github.com/joaopandolfi/blackwhale/utils"
)

func configInit() {
	config.Load(os.Args[1:])

	if config.Get().SnakeByDefault {
		handlers.ActiveSnakeCase()
	}
}

func resilient() {
	utils.Info("[SERVER] - Shutdown")

	if err := recover(); err != nil {
		utils.CriticalError("[SERVER] - Returning from the dark", err)
		main()
	}
}

func gracefullShutdown() {
	mongo.GracefullShutdown()
}

func welcome() {
	// https://patorjk.com/software/taag/#p=display&f=Broadway%20KB&t=oraculo
	fmt.Println("")
	fmt.Println(`
	 ___   ___    __    __    _     _     ___  
	/ / \ | |_)  / /\  / /   | | | | |   / / \ 
	\_\_/ |_| \ /_/--\ \_\_, \_\_/ |_|__ \_\_/ 
O ======================================================> (I know your secrets...)
	`)
}

func main() {
	defer resilient()

	welcome()
	//Init
	configInit()

	// Initialize Mux Router
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))

	srv := server.New(r, config.Get())
	nr := router.New(srv)
	nr.Setup()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go srv.Start()

	<-done
	utils.Info("[SERVER] Gracefully shutdown")
	gracefullShutdown()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if err := srv.Shutdown(ctx); err != nil {
		utils.CriticalError("Server Shutdown Failed", err.Error())
	}

	cancel()
}
