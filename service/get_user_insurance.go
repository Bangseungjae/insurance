package service

import (
	"Bangseungjae/insurance/store"
	"context"
	"errors"
)

type ListUserInsurance struct {
}

var (
	ErrNotFound = errors.New("not found")
)

func (ig *ListUserInsurance) ListUserInsurance(ctx context.Context, id int) (*store.UserInsurance, error) {
	ctx.Done()
	select {
	case <-ctx.Done():
		return nil, nil
	default:
		rs, ok := store.UserInsurances[id]
		if !ok {
			return nil, ErrNotFound
		}
		return &rs, nil
	}
}
