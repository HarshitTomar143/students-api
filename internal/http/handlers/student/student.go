package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/HarshitTomar143/students-api/internal/types"
	"github.com/HarshitTomar143/students-api/internal/utils/response"
)

func New() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student
		error:= json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(error, io.EOF){
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(error))
			return
		}

		slog.Info("creating a student")

		response.WriteJson(w, http.StatusCreated, map[string]string{"status":"OK"})// making a splice of byte data.
	}
}