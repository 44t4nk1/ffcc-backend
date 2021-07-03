package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/load/csv", LoadCsv)
	router.GET("/load/courses", CourseList)
	router.GET("/load/faculty", FacultyList)
	router.GET("/get/courses", GetCourses)
	router.GET("/get/faculty", GetFaculty)
	router.POST("/user/signup", Signup)
	router.POST("/user/login", Login)
}
