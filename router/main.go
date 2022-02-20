package router

import (
	"github.com/Rafipratama22/bank_test/config"
	"github.com/Rafipratama22/bank_test/controller"
	"github.com/Rafipratama22/bank_test/middleware"
	"github.com/Rafipratama22/bank_test/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func Start() *gin.Engine{
	// Load .env file
	route := gin.Default()
	
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
	}
	return route
}