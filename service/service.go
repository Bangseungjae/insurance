package service

import (
	"Bangseungjae/insurance/store"
	"context"
)

type GetUserInsuranceService interface {
	ListUserInsurance(ctx context.Context, id int) (*store.UserInsurance, error)
}
