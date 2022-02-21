package router

import (
	"github.com/Rafipratama22/bank_test/config"
	"github.com/Rafipratama22/bank_test/controller"
	"github.com/Rafipratama22/bank_test/middleware"
	"github.com/Rafipratama22/bank_test/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/Rafipratama22/bank_test/docs"
)

type Server struct {
	Router *gin.Engine
}


var (
	db *gorm.DB = config.SetUpDatabase()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	userController  controller.UserController = controller.NewUserController(userRepository) 
	authMiddleware middleware.AuthMiddleware = middleware.NewAuthMiddleware(db)
	rekeningRepository repository.RekeningRepository = repository.NewRekeningRepository(db)
	rekeningController controller.RekeningController = controller.NewRekeningController(rekeningRepository)
	merchantRepository repository.MerchantRepository = repository.NewMerchantRepository(db)
	merchantController controller.MerchantController = controller.NewMerchantController(merchantRepository)
)

func MainServer() *Server {
	return &Server{
		Router: gin.Default(),
	}
}


// @title Gin Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func Start() *gin.Engine{
	// Load .env file
	route := gin.Default()
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	userRoute := route.Group("/user")
	{
		userRoute.POST("/register", userController.Register)
		userRoute.POST("/login", userController.Login)
		userRoute.POST("/logout", authMiddleware.ValidateTokenUser , userController.Logout)
	}

	rekeningRoute := route.Group("/rekening")
	rekeningRoute.Use(authMiddleware.ValidateTokenUser)
	{
		rekeningRoute.POST("/create", rekeningController.CreateRekening)
		rekeningRoute.GET("/list", rekeningController.GetRekeningAll)
		rekeningRoute.GET("/detail/:id", rekeningController.GetRekeningByID)
		rekeningRoute.PUT("/update/:id", rekeningController.UpdateRekening)
		rekeningRoute.POST("/transfer", rekeningController.TransferRekening)
	}

	merchantRoute := route.Group("/merchant")
	merchantRoute.Use(authMiddleware.ValidateTokenUser)
	{
		merchantRoute.POST("/create", merchantController.CreateMerchant)
		merchantRoute.GET("/list", merchantController.GetMerchantAll)
		merchantRoute.GET("/detail/:id", merchantController.DetailMerchant)
		merchantRoute.PUT("/update/:id", merchantController.UpdateMerchant)
		merchantRoute.DELETE("/delete/:id", merchantController.DeleteMerchant)
	}
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return route
}