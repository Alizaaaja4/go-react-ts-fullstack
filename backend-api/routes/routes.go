package routes

import (
	"santrikoding/backend-api/controllers"
	"santrikoding/backend-api/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	//initialize gin
	router := gin.Default()

	// set up CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	// =============== Authentikasi ===================
	// route register
	router.POST("/api/register", controllers.Register)
	// route login
	router.POST("/api/login", controllers.Login)

	// ================ Users =========================
	// route users
	router.GET("/api/users", middlewares.AuthMiddleware(), controllers.FindUsers)
	// route user create
	router.POST("/api/users", middlewares.AuthMiddleware(), controllers.CreateUser)
	// route user by id
	router.GET("/api/users/:id", middlewares.AuthMiddleware(), controllers.FindUserById)
	// route user update
	router.PUT("/api/users/:id", middlewares.AuthMiddleware(), controllers.UpdateUser)
	// route user delete
	router.DELETE("/api/users/:id", middlewares.AuthMiddleware(), controllers.DeleteUser)

	// ================ Products =========================
	// route products
	router.GET("/api/products", middlewares.AuthMiddleware(), controllers.FindProducts)
	// route product create
	router.POST("/api/products", middlewares.AuthMiddleware(), controllers.CreateProduct)
	// route product by id
	router.GET("/api/products/:id", middlewares.AuthMiddleware(), controllers.FindProductById)
	// route product update
	router.PUT("/api/products/:id", middlewares.AuthMiddleware(), controllers.UpdateProduct)
	// route product delete
	router.DELETE("/api/products/:id", middlewares.AuthMiddleware(), controllers.DeleteProduct)

	return router
}
