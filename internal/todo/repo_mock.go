package todo

import (
	"context"
	"time"
)

type MockRepo struct {
	opt MockRepoOpt
}

type MockRepoOpt struct {
	Err error
}

var _ Repository = (*MockRepo)(nil)

func NewMockRepo(opt MockRepoOpt) *MockRepo {
	return &MockRepo{opt}
}

func (r *MockRepo) Create(_ context.Context, description string, dueDate time.Time) (Todo, error) {
	return Todo{
		ID:          1,
		Description: description,
		DueDate:     dueDate,
	}, r.opt.Err
}

func (r *MockRepo) FindLast(_ context.Context) (Todo, error) {
	return Todo{
		ID:          1,
		Description: "mock todo",
		DueDate:     time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC),
	}, r.opt.Err
}
