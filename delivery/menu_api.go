package delivery

import (
	"menu-manage/delivery/appresponse"
	"menu-manage/usecase"

	"github.com/gin-gonic/gin"
)

type MenuApi struct {
	usecase     usecase.MenuUseCase
	publicRoute *gin.RouterGroup
}

func NewMenuApi(publicRoute *gin.RouterGroup, usecase usecase.MenuUseCase) *MenuApi {
	api := MenuApi{
		usecase:     usecase,
		publicRoute: publicRoute,
	}
	api.initRouter()
	return &api
}

func (a *MenuApi) initRouter() {
	menuRoute := a.publicRoute.Group("/menu")
	menuRoute.GET("", a.getMenu)
}

func (a *MenuApi) getMenu(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")
	response := appresponse.NewJsonResponse(c)

	if id != "" {
		menu, err := a.usecase.SearchMenuById(id)
		if err != nil {
			response.SendError(appresponse.NewInternalServerError(err, "FAILED"))
			return
		}
		response.
			SendData(appresponse.NewResponseMessage("SUCCESS", "Menu By Id", menu))
		return
	}

	if name != "" {
		menuList, err := a.usecase.SearchMenuByName(name)
		if err != nil {
			response.SendError(appresponse.NewInternalServerError(err, "FAILED"))
			return
		}
		response.SendData(appresponse.NewResponseMessage("SUCCESS", "Menu By Name", menuList))
		return
	}

	menuList, err := a.usecase.GetAll()
	if err != nil {
		response.SendError(appresponse.NewInternalServerError(err, "FAILED"))
		return
	}
	response.SendData(appresponse.NewResponseMessage("SUCCESS", "Menu All", menuList))
}
