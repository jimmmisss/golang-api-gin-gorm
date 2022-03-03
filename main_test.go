package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jimmmisss/go-api-gin-gorm/controllers"
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
	if resposta.Code != http.StatusOK {
		t.Fatalf("Status error: valor recebido foi %d e o valor esperado era %d", resposta.Code, http.StatusOk)
	}
}
