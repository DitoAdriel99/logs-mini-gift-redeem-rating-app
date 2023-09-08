package logs_user

import (
	"encoding/json"
	"go-learn/entities"
	"go-learn/library/response"
	"log"
	"net/http"
)

func (c *_Controller) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var (
		payload     entities.LogsPayload
		errResponse = response.NewResponse().
				WithCode(http.StatusUnprocessableEntity).
				WithStatus("Failed").
				WithMessage("Failed")
		succResponse = response.NewResponse().
				WithStatus("Success").
				WithMessage("Success")
	)

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response := *errResponse.WithError(err.Error())
		output, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(output)
		return
	}

	if err := c.service.LogsService.Create(&payload); err != nil {
		response := *errResponse.WithError(err)
		output, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(output)
		return
	}

	response := *succResponse.WithData(payload)
	object, err := json.Marshal(response)
	if err != nil {
		response := *errResponse.WithError(err)
		output, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(output)
		return
	}
	log.Println("success save logs...")
	w.WriteHeader(http.StatusCreated)
	w.Write(object)

}
