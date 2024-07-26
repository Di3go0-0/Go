package main

import (
	"github.com/gin-gonic/gin" //Paquete gin

	"Gin/routes" //Se importa el paquete routes
)

func main() {
	// Usando gin.New()
	// r := gin.New()
	// Aquí puedes agregar tus propios middlewares si lo deseas
	// r.Use(myCustomMiddleware)

	// Usando gin.Default()
	r := gin.Default()
	// Incluye middlewares Logger y Recovery por defecto
	
	routes.MainRoutes(r)	

	r.Run(":8080")
}