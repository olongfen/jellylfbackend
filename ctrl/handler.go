package ctrl

import (
	"github.com/gin-gonic/gin"
	"jellylfbackend/services"
	"strconv"
)

// Response 返回体
type Response struct {
	// Code 0为成功，其他都是失败
	Code string `json:"code"`
	// Msg 返回提示
	Msg string `json:"msg"`
	// Data 返回数据
	Data any `json:"data"`
}

// successResponse 成功返回
func successResponse(c *gin.Context, data any) {
	c.JSON(200, Response{Code: "SUCCESS", Msg: "操作成功", Data: data})
}

func GetPerson(c *gin.Context) {
	email := c.Query("email")
	person, err := services.GetPersonOneWithEmail(c, email)
	if err != nil {
		return
	}
	successResponse(c, person)
}

func GetJobExperience(c *gin.Context) {
	personID := c.Param("personID")
	atoi, err := strconv.Atoi(personID)
	if err != nil {
		return
	}
	jobExperiences, err := services.GetJobExperienceWithPersonID(c, uint(atoi))
	if err != nil {
		return
	}
	successResponse(c, jobExperiences)
}

func GetWork(c *gin.Context) {
	personID := c.Param("personID")
	atoi, err := strconv.Atoi(personID)
	if err != nil {
		return
	}
	works, err := services.GetWorksWithPersonID(c, uint(atoi))
	if err != nil {
		return
	}
	successResponse(c, works)
}
