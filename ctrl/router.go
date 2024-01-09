package ctrl

import "github.com/gin-gonic/gin"

func Setup(en *gin.Engine) {
	en.Use(Cors())
	pubRouter := en.Group("/api")
	pubRouter.GET("/person", GetPerson)
	pubRouter.GET("/job-experience/:personID", GetJobExperience)
	pubRouter.GET("/work/:personID", GetWork)
}
