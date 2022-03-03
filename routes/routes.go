package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jimmmisss/go-api-gin-gorm/controllers"
)

func HandleRequests() {
	addr := "127.0.0.1:8080"
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/:criar", controllers.CriarNovoAluno)
	r.GET("/alunos/:id", controllers.BucaAlunoPorId)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/:cpf", controllers.BuscaAlunoPorCPF)
	r.Run(addr)
}
