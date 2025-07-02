package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GamesRouter struct{}

// InitGamesRouter  games
func (s *GamesRouter) InitGamesRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	gamesRouter := Router.Group("games").Use(middleware.OperationRecord())
	gamesRouterWithoutRecord := Router.Group("games")
	gamesRouterWithoutAuth := PublicRouter.Group("games")
	{
		gamesRouter.POST("createGames", gamesApi.CreateGames)             // games
		gamesRouter.DELETE("deleteGames", gamesApi.DeleteGames)           // games
		gamesRouter.DELETE("deleteGamesByIds", gamesApi.DeleteGamesByIds) // games
		gamesRouter.PUT("updateGames", gamesApi.UpdateGames)              // games
	}
	{
		gamesRouterWithoutRecord.GET("findGames", gamesApi.FindGames)       // IDgames
		gamesRouterWithoutRecord.GET("getGamesList", gamesApi.GetGamesList) // games
	}
	{
		gamesRouterWithoutAuth.GET("getGamesPublic", gamesApi.GetGamesPublic) // games
		gamesRouterWithoutAuth.GET("list", gamesApi.GetList)                  // games
	}
}
