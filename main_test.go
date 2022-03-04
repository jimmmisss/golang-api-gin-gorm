package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
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

func TestBuscaAlunoPorIdFHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunMock()
	r := SetupRotasTest()
	r.GET("/alunos/:id", controllers.BucaAlunoPorId)
	pathDaBusca := "/aluno" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathDaBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)
	assert.Equal(t, "Wesley Pereira", alunoMock.Nome, "Os nomes devem ser iguais")
	assert.Equal(t, "12345678901", alunoMock.CPF)
	assert.Equal(t, "123456789", alunoMock.RG)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestDeletaAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	r := SetupRotasTest()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	pathDaBusca := "/aluno" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathDaBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestAtualizaAlunoFHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunMock()
	r := SetupRotasTest()
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	aluno := models.Aluno{
		Nome: "Wesley Pereira",
		CPF:  "12345678901",
		RG:   "123456789",
	}
	valorJson, _ := json.Marshal(aluno)
	pathDaEditar := "/aluno/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathDaEditar, bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMockAtializado models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtializado)
	assert.Equal(t, "Wesley Pereira", alunoMockAtializado.Nome, "Os nomes devem ser iguais")
	assert.Equal(t, "12345678901", alunoMockAtializado.CPF)
	assert.Equal(t, "123456789", alunoMockAtializado.RG)
	assert.Equal(t, http.StatusOK, resposta.Code)
}
