package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"userInfoService/routers"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := routers.InitRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	type r struct {
		Message string
	}
	var returnData r
	err := json.Unmarshal(w.Body.Bytes(), &returnData)
	if err != nil {
		panic(err)
	}
	fmt.Println(returnData.Message)
	assert.Equal(t, "pong", returnData.Message)
}

func TestAuth(t *testing.T) {
	router := routers.InitRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/auth?username=cjb&password=123", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	type r struct {
		Code int
		Msg  string
		Data map[string]interface{}
	}
	var returnData r
	err := json.Unmarshal(w.Body.Bytes(), &returnData)
	if err != nil {
		t.Errorf("Unmarshal error : %s", err)
	}
	fmt.Printf("returnData.Code : %d\n", returnData.Code)
	fmt.Printf("returnData.Msg : %s\n", returnData.Msg)
	fmt.Printf("returnData.Data : %s\n", returnData.Data)
	// assert.Equal(t, 200, returnData.Code)
	// assert.Equal(t, "ok", returnData.Msg)
	// t.Errorf("returnData.Code : %d", returnData.Code)
	// t.Errorf("returnData.Msg : %s", returnData.Msg)
	assert.Equal(t, 200, returnData.Code)
	assert.Equal(t, "ok", returnData.Msg)
}
