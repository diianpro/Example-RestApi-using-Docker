package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/diianpro/simpleApp/internal/models"
	"github.com/gofrs/uuid"
	"github.com/labstack/gommon/log"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestNoteHandler_CreateNote(t *testing.T) {
	body := models.TodoList{
		ID:          uuid.UUID{},
		Title:       "test",
		Description: "testing",
	}

	response, err := json.Marshal(body)
	if err != nil {
		return
	}

	request, err := http.NewRequest(http.MethodPost, "http://localhost:8080/note", bytes.NewReader(response))
	request.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient

	resp, err := client.Do(request)
	if err != nil {
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("error")
		}

	}(resp.Body)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	log.Info(string(data))
}

func TestNoteHandler_GetByID(t *testing.T) {
	q := make(url.Values)
	q.Set("id", "824c6e06-0155-4494-8090-5dc763997e20")

	request := httptest.NewRequest(http.MethodGet, "/node/"+q.Encode(), nil)

	request.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient

	resp, err := client.Do(request)
	if err != nil {
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("error")
		}

	}(resp.Body)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	log.Info(string(data))

}
