package wsmock

import (
	"context"
	"wordbox/internal/domain"
)

type WSMock struct{}

func New() *WSMock {
	return &WSMock{}
}

func (m *WSMock) AddNewWords(ctx context.Context, _ string) ([]*domain.WordCard, error) {
	words := make([]*domain.WordCard, 0, 4)
	for i := 0; i < 4; i++ {
		words = append(words, domain.GenerateWordCard())
	}
	return words, nil
}

func (m *WSMock) GetWordsToRepeat(ctx context.Context, _ string) ([]*domain.WordCard, error) {
	words := make([]*domain.WordCard, 0, 4)
	for i := 0; i < 4; i++ {
		words = append(words, domain.GenerateWordCard())
	}
	return words, nil
}
