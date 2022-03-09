package router

import (
	"atro/internal/handler"
	"atro/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

//RunAPI ->route setup
func RunAPI(address string) error {

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to Our Mini Ecommerce")
	})

	apiRoutes := r.Group("/api/v1")

	adminRoutes := apiRoutes.Group("/admin")
	adminHandler := handler.NewAdminHandler()
	productHandler := handler.NewProductHandler()
	{
		// unauthorize api 
		adminRoutes.POST("/login", adminHandler.SignInUser) // /admin/login
		adminRoutes.POST("/register", adminHandler.AddUser) // cứ cho đăng kí để có data trong db đã
		adminRoutes.POST("/logout", nil)

		// auth api 
		adminAuth := adminRoutes.Group("/auth", middleware.AuthorizeJWT())
		{

			// product
			adminAuth.PUT("/products/:id", productHandler.UpdateProduct) // /admin/auth/products
			adminAuth.POST("/products/", productHandler.AddProduct)
			adminAuth.DELETE("/products/:id", productHandler.DeleteProduct)

			// category

			// order info
		}

	}

	userRoutes := apiRoutes.Group("/user")
	userHandler := handler.NewUserHandler()

	{
		// unauthorize api 
		userRoutes.POST("/login", nil) 
		userRoutes.POST("/register", nil) 
		userRoutes.POST("/logout", nil)
		userRoutes.GET("/products/", productHandler.GetAllProduct)
		userRoutes.GET("/products/:id", productHandler.GetProduct)

		// auth api 
		userRoutes.GET("", userHandler.GetUser) // api/user?ip=1
	}

	return r.Run(address)

}
