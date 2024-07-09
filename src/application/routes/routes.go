package routes

import (
	"os"
	"typograph_back/src/application"
	"typograph_back/src/controller"
	"typograph_back/src/middleware"
	"typograph_back/src/repository"
	"typograph_back/src/service"
)

func RegisterRoute(urlPrefix string) {
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	routeGroup := application.GlobalApp.Group(urlPrefix)
	authM := middleware.NewJwtMiddleware([]byte(os.Getenv("JWT_SECRET")), "Bearer ", userService)
	permissionM := middleware.NewPermissionMiddleware(userService)

	routeGroup.GET("/users", userController.Index, authM.Run(), permissionM.Run("users-index"))
	routeGroup.POST("/users", userController.Store)
	routeGroup.GET("/users/:id", userController.Show, authM.Run())
	routeGroup.PATCH("/users/:id", userController.UpdatePassword, authM.Run())
	routeGroup.DELETE("/users/:id", userController.Delete, authM.Run())

	authService := service.NewAuthService(userService)
	authController := controller.NewAuthController(authService)
	routeGroup.POST("/login", authController.Login)
	routeGroup.POST("/register", authController.Register)
	routeGroup.GET("/me", authController.Me, authM.Run())

	roleController := controller.NewRoleController()
	routeGroup.GET("/roles", roleController.Index, authM.Run(), permissionM.Run("roles-index"))
	routeGroup.POST("/roles", roleController.Store, authM.Run(), permissionM.Run("roles-store"))
	routeGroup.GET("/roles/:id", roleController.Show, authM.Run(), permissionM.Run("roles-show"))
	routeGroup.PATCH("/roles/:id", roleController.Update, authM.Run(), permissionM.Run("roles-update"))
	routeGroup.DELETE("/roles/:id", roleController.Delete, authM.Run(), permissionM.Run("roles-delete"))
	routeGroup.POST("/roles/:id/permissions", roleController.AddPermissions, authM.Run(), permissionM.Run("roles-add-permissions"))
	routeGroup.DELETE("/roles/:id/permissions", roleController.RemovePermissions, authM.Run(), permissionM.Run("roles-remove-permissions"))

	permissionController := controller.NewPermissionController()
	routeGroup.GET("/permissions", permissionController.Index, authM.Run(), permissionM.Run("permissions-index"))
	routeGroup.POST("/permissions", permissionController.Store, authM.Run(), permissionM.Run("permissions-store"))
	routeGroup.GET("/permissions/:id", permissionController.Show, authM.Run(), permissionM.Run("permissions-show"))
	routeGroup.PATCH("/permissions/:id", permissionController.Update, authM.Run(), permissionM.Run("permissions-update"))
	routeGroup.DELETE("/permissions/:id", permissionController.Delete, authM.Run(), permissionM.Run("permissions-delete"))

	paragraphController := controller.NewParagraphController()
	routeGroup.GET("/random_paragraph", paragraphController.GetRandom)

	raceRepository := repository.NewRaceRepository()
	raceService := service.NewRaceService(raceRepository, userService)
	userRaceResultService := service.NewUserRaceResultService()
	raceController := controller.NewRaceController(raceService, userRaceResultService)
	routeGroup.GET("/races", raceController.Index, authM.Run())
	routeGroup.GET("/races/:id", raceController.Show, authM.Run())
	routeGroup.POST("/races", raceController.Store, authM.Run())
	routeGroup.PATCH("/races", raceController.Update, authM.Run())
	routeGroup.DELETE("/races", raceController.Delete, authM.Run())
	routeGroup.POST("/races/add_user_race_result", raceController.AddUserRaceResult, authM.Run())
	routeGroup.GET("/races/get_user_race_result/:user_id", raceController.GetUserRaceResults, authM.Run())
	routeGroup.GET("/races/leaderboard", raceController.Leaderboard)

	lobbyWsRepository := repository.NewLobbyWsRepository()
	lobbyWsService := service.NewLobbyWsService(lobbyWsRepository)
	lobbyWSController := controller.NewLobbyWSController(lobbyWsService)
	lobbyRepository := repository.NewLobbyRepository()
	lobbyService := service.NewLobbyService(lobbyRepository, userService, raceService, lobbyWsService)
	lobbyController := controller.NewLobbyController(lobbyService)
	routeGroup.GET("/lobbies", lobbyController.Index, authM.Run())
	routeGroup.GET("/lobbies/:id", lobbyController.Show, authM.Run())
	routeGroup.POST("/lobbies", lobbyController.Store, authM.Run())
	routeGroup.PATCH("/lobbies", lobbyController.Update, authM.Run())
	routeGroup.DELETE("/lobbies", lobbyController.Delete, authM.Run())
	routeGroup.POST("/lobbies/enter", lobbyController.EnterLobby, authM.Run())
	routeGroup.POST("/lobbies/leave", lobbyController.LeaveLobby, authM.Run())
	routeGroup.PATCH("/lobbues/:id/start", lobbyController.StartLobby, authM.Run())

	routeGroup.GET("/ws", lobbyWSController.Index)
}
