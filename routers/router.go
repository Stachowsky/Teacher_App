package routers

import (
	"Stachowsky/Teacher_App/controllers"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func CreateUrlMappings() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("sites/html/*.html")
	r.Use(static.Serve("/sites", static.LocalFile("./sites/js", true)))
	r.GET("/", controllers.Page)
	r.POST("/student", controllers.IsAdmin, controllers.CreateStudent)
	r.GET("/student", controllers.ReadStudents)
	r.GET("/student/:id", controllers.ReadStudentById)
	r.PUT("/student/:id", controllers.IsDev, controllers.UpdateStudentById)
	r.DELETE("/student/:id", controllers.IsAdmin, controllers.DeleteStudentById)
	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "error.html", gin.H{"title": "Page not found!"})
	})
	r.POST("/login", controllers.LoginUser)
	r.POST("/register", controllers.RegisterUser)
	r.POST("/refresh", controllers.RefreshToken)
	return r
}
