package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jimmmisss/go-api-gin-gorm/controllers"
	"github.com/stretchr/testify/assert"
)

func SetupRotasTest() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TesteVerificaStatusCodeSaudacaoComParametro(t *testing.T) {
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
