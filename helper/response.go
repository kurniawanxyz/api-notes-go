package helper

import "github.com/gin-gonic/gin"

type Response struct {
	Status  int         `json:"status"`
	Success bool		`json:"success"`
	Message string      `json:"message"`
	Data    any `json:"data"`
}


func HandleResponse(c *gin.Context , Status int, Data any) {
	var response Response
	response.Status = Status

	if Status >= 200 && Status < 300 {
		response.Success = true
		response.Message = "Success"
	} else if Status >= 400 && Status < 500 {
		response.Success = false
		response.Message = "Bad Request Error"
	}else{
		response.Success = false
		response.Message = "Internal Server Error"
	}
	response.Data = Data
	c.JSON(Status, response)
}