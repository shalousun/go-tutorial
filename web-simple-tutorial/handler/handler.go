package handler

import (
	"encoding/json"
	"net/http"
	"web-simple-tutorial/model"
)

type helloHandler struct {
}

func (m *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "A Go Web Server")
	w.Write([]byte("welcome go"))
}

func Json(w http.ResponseWriter, r *http.Request) {
	response := model.Response{}
	response.Code = "200"
	response.Message = "hello world"
	data := make(map[string]interface{})
	data["id"] = "2"
	data["name"] = "jack"
	response.Data = data

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func MysqlConf(w http.ResponseWriter, r *http.Request) {
	m := new(model.MysqlConf)
	mysqlConf := m.GetConf()
	jsonData, err := json.Marshal(mysqlConf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
