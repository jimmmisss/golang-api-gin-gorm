package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jimmmisss/go-api-gin-gorm/controllers"
	"github.com/jimmmisss/go-api-gin-gorm/database"
	"github.com/jimmmisss/go-api-gin-gorm/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRotasTest() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{
		Nome: "Wesley Pereira",
		CPF:  "12345678901",
		RG:   "123456789",
	}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestVerificaStatusCodeSaudacaoComParametro(t *testing.T) {
	r := SetupRotasTest()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/isa", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "deveriam ser iguais")
	mockResposta := `{"API diz":"E ai isa, Tudo blz"}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)
	assert.Equal(t, mockResposta, respostaBody)
}

func TestListandoTodosAlunosHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunMock()
	r := SetupRotasTest()
	r.POST("/:criar", controllers.CriarNovoAluno)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaAlunoPorCPFHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunMock()
	r := SetupRotasTest()
	r.GET("/alunos/:cpf", controllers.BuscaAlunoPorCPF)
	req, _ := http.NewRequest("GET", "/alunos/12345678901", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}
