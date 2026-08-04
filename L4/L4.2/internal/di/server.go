package di

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (a *App) Run() error {

	router := gin.Default()

	a.Server.Register(router)

	addr := fmt.Sprintf(
		"%s:%d",
		a.Config.Server.Host,
		a.Config.Server.Port,
	)

	return router.Run(addr)
}