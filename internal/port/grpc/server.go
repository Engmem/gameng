package wordgrpc

import (
	"context"
	gamengv02 "github.com/Engmem/wordbox-api/gen/go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"wordbox/internal/domain"
)

type WordService interface {
	AddNewWords(ctx context.Context, UUID string) ([]*domain.WordCard, error)
	GetWordsToRepeat(ctx context.Context, UUID string) ([]*domain.WordCard, error)
}

type WordServer struct {
	WordService WordService
	gamengv02.UnimplementedWordboxServer
}

func Register(s *grpc.Server, ws WordService) {
	gamengv02.RegisterWordboxServer(s, &WordServer{
		WordService: ws,
	})
}

func (ws *WordServer) AddNewWords(ctx context.Context, req *gamengv02.AddNewWordsRequest) (*gamengv02.AddNewWordsResponse, error) {
	if req.Uuid == "" {
		return nil, status.Error(codes.InvalidArgument, "user id must not be empty")
	}

	words, err := ws.WordService.AddNewWords(ctx, req.Uuid)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	resp := &gamengv02.AddNewWordsResponse{
		WordsTranslations: make(map[string]string, len(words)),
	}
	for _, word := range words {
		resp.WordsTranslations[word.Word] = word.Translation
	}

	return resp, nil
}

func (ws *WordServer) GetWordsToRepeat(ctx context.Context, req *gamengv02.GetWordsToRepeatRequest) (*gamengv02.GetWordsToRepeatResponse, error) {
	if req.Uuid == "" {
		return nil, status.Error(codes.InvalidArgument, "user id must not be empty")
	}

	words, err := ws.WordService.GetWordsToRepeat(ctx, req.Uuid)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	resp := &gamengv02.GetWordsToRepeatResponse{
		WordsTranslations: make(map[string]string, len(words)),
	}
	for _, word := range words {
		resp.WordsTranslations[word.Word] = word.Translation
	}

	return resp, nil
}
