package delivery

import (
	"menu-manage/apperror"
	"menu-manage/delivery/appresponse"
	"menu-manage/dto"
	"menu-manage/usecase"

	"github.com/gin-gonic/gin"
)

type CustomerTableApi struct {
	usecase     usecase.CustomerTableUseCase
	publicRoute *gin.RouterGroup
}

func NewCustomerTableApi(publicRoute *gin.RouterGroup, usecase usecase.CustomerTableUseCase) *CustomerTableApi {
	customerTableApi := CustomerTableApi{
		usecase:     usecase,
		publicRoute: publicRoute,
	}
	customerTableApi.initRouter()
	return &customerTableApi
}

func (a *CustomerTableApi) initRouter() {
	tableRoute := a.publicRoute.Group("/table")
	tableRoute.GET("", a.getTableList)
	tableRoute.POST("/checkin", a.tableCheckIn)
	tableRoute.PUT("/checkout", a.tableCheckOut)
}

func (a *CustomerTableApi) getTableList(c *gin.Context) {
	tables, err := a.usecase.GetTodayListCustomerTable()
	response := appresponse.NewJsonResponse(c)
	if err != nil {
		response.SendError(appresponse.NewInternalServerError(err, "Failed to get table list"))
		return
	}
	response.SendData(appresponse.NewResponseMessage("SUCCESS", "List table", tables))
}

func (a *CustomerTableApi) tableCheckIn(c *gin.Context) {
	var checkInRequest dto.Request
	response := appresponse.NewJsonResponse(c)
	if err := c.ShouldBindJSON(&checkInRequest); err != nil {
		response.SendError(appresponse.NewBadRequestError(err, "Failed to check in"))
		return
	}

	data, err := a.usecase.TableCheckIn(checkInRequest)
	if err != nil {
		if err == apperror.TableOccupiedErr {
			response.SendData(appresponse.NewResponseMessage("FAILED", err.Error(), data))
			return
		}
		response.SendError(appresponse.NewInternalServerError(err, err.Error()))
		return
	}
	response.SendData(appresponse.NewResponseMessage("SUCCESS", "Check in", data))
}

func (a *CustomerTableApi) tableCheckOut(c *gin.Context) {
	billNo := c.Query("billNo")
	response := appresponse.NewJsonResponse(c)
	if billNo == "" {
		response.SendError(appresponse.NewBadRequestError(apperror.FieldRequiredErr, "Failed"))
		return
	}
	err := a.usecase.TableCheckOut(billNo)
	if err != nil {
		if err == apperror.NoRecordFoundErr {
			response.SendError(appresponse.NewBadRequestError(err, "Failed to check out"))
			return
		}
		response.SendError(appresponse.NewInternalServerError(err, "Failed to check out"))
		return
	}
	response.SendData(appresponse.NewResponseMessage("SUCCESS", "check out", nil))

}
