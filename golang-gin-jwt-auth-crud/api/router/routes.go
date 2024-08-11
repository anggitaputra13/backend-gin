package router

import (
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/api/controllers"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/api/middleware"
	"github.com/gin-gonic/gin"
)

func GetRoute(r *gin.Engine) {
	// User routes
	r.POST("/api/signup", controllers.Signup)
	r.POST("/api/login", controllers.Login)

	r.Use(middleware.RequireAuth)
	r.POST("/api/logout", controllers.Logout)
	userRouter := r.Group("/api/users")
	{
		userRouter.GET("/", controllers.ListUsers)
		userRouter.GET("/:id", controllers.GetUser)
		userRouter.PUT("/update/:id", controllers.UpdateUser)
		userRouter.DELETE("/:id", controllers.DeleteUser)
		userRouter.DELETE("/delete-permanent/:id", controllers.PermanentlyDeleteUser)
	}

	// Book routes
	bookRouter := r.Group("/api/books")
	{
		bookRouter.GET("/", controllers.ListBook)
		bookRouter.POST("/create", controllers.CreateBook)
		bookRouter.GET("/:id", controllers.GetBook)
		bookRouter.PUT("/update/:id", controllers.UpdateBook)
		bookRouter.DELETE("/:id", controllers.DeleteBook)
		bookRouter.DELETE("/delete-permanent/:id", controllers.DeleteBookPermanent)
	}

	// Employee routes
	customerRouter := r.Group("/api/customers")
	{
		customerRouter.GET("/", controllers.ListEmployee)
		customerRouter.POST("/create", controllers.CreateEmployee)
		customerRouter.GET("/:id", controllers.GetEmployee)
		customerRouter.PUT("/update/:id", controllers.UpdateEmployee)
		customerRouter.DELETE("/:id", controllers.DeleteEmployee)
		customerRouter.DELETE("/delete-permanent/:id", controllers.DeleteEmployeePermanent)
	}

	// Order routes
	orderRouter := r.Group("/api/orders")
	{
		orderRouter.GET("/", controllers.ListOrders)
		orderRouter.POST("/create", controllers.CreateOrder)
		orderRouter.PUT("/update/:id", controllers.UpdateOrder)
		orderRouter.DELETE("/:id", controllers.DeleteOrder)
		orderRouter.DELETE("/delete-permanent/:id", controllers.PermanentlyDeleteOrder)
	}
}
