package handler

import (
	"Bangseungjae/insurance/service"
	"net/http"
	"strconv"
)

type GetUserInsurance struct {
	Service service.GetUserInsuranceService
}

func (gui *GetUserInsurance) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}
	insurance, err := gui.Service.ListUserInsurance(ctx, id)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusNotFound)
		return
	}
	RespondJSON(ctx, w, insurance, http.StatusOK)
}
