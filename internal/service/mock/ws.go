package wsmock

import (
	"context"
	"gameng/internal/domain"
)

type WSMock struct {
}

func New() *WSMock {
	return &WSMock{}
}

func (m *WSMock) GetNewWords(ctx context.Context, _ uint64) ([]*domain.WordCard, error) {
	words := make([]*domain.WordCard, 0, 4)
	for i := 0; i < 4; i++ {
		words = append(words, domain.GenerateWordCard())

	}
	return words, nil
}
func (m *WSMock) GetWords(ctx context.Context, _ uint64) ([]*domain.WordCard, error) {
	words := make([]*domain.WordCard, 0, 4)
	for i := 0; i < 4; i++ {
		words = append(words, domain.GenerateWordCard())

	}
	return words, nil
}
