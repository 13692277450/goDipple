package main

func PrettyShutDownCfg() {

	if Init_PrettyShutDown {

		FolderCheck("util/prettyshutdown", "util/prettyshutdown", "[PRETTY SHUTDOWN] ")
		WriteContentToConfigYaml(PrettyShutDown_Init_Content, "util/prettyshutdown/prettyshutdown.go", "[PRETTY SHUTDOWN] ")
	}
}

var PrettyShutDown_Init_Content = `
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	r := SetupRouters()
	//Pretty shutdown...
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Info("Listen: %v\n", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

}

func SetupRouters() *gin.Engine {
	// Initiallize routes
	r := gin.New()
	return r
}
`
