package main

import (
	"github.com/braintree/manners"
	"github.com/gin-gonic/gin"
	_ "log"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	routerMy := gin.Default()

	router.Register(routerMy)

	manners.ListenAndServe(":81", routerMy)
	manners.Close()
}
