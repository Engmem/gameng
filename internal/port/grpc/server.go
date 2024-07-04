package wordgrpc

import (
	"context"
	"gameng/internal/domain"
	gamengv1 "github.com/Engmem/gameng-api/gen/go/proto/gameng"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WordService interface {
	GetNewWords(ctx context.Context, userID uint64) ([]*domain.WordCard, error)
	GetWords(ctx context.Context, offset uint64) ([]*domain.WordCard, error)
}

type WordServer struct {
	WordService WordService
	gamengv1.UnimplementedGamengServer
}

func Register(s *grpc.Server, ws WordService) {
	gamengv1.RegisterGamengServer(s, &WordServer{
		WordService: ws,
	})
}

func (ws *WordServer) GetNewWords(ctx context.Context, req *gamengv1.GetNewWordsRequest) (*gamengv1.GetNewWordsResponse, error) {
	if req.UserId == 0 {
		return nil, status.Error(codes.InvalidArgument, "user id must not be empty")
	}

	words, err := ws.WordService.GetNewWords(ctx, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	resp := &gamengv1.GetNewWordsResponse{
		WordsTranslations: make(map[string]string, len(words)),
	}
	for _, word := range words {
		resp.WordsTranslations[word.Word] = word.Translation
	}

	return resp, nil
}

func (ws *WordServer) GetWordWithTranslations(ctx context.Context, req *gamengv1.GetWordsWithTranslationsRequest) (*gamengv1.GetWordsWithTranslationsResponse, error) {
	words, err := ws.WordService.GetWords(ctx, req.Offset)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	resp := &gamengv1.GetWordsWithTranslationsResponse{
		WordsTranslations: make(map[string]string, len(words)),
	}

	for _, word := range words {
		resp.WordsTranslations[word.Word] = word.Translation
	}

	return resp, nil
}

func (ws *WordServer) GetWordByAudio(ctx context.Context, req *gamengv1.GetWordsByAudioRequest) (*gamengv1.GetWordsByAudioResponse, error) {
	words, err := ws.WordService.GetWords(ctx, req.Offset)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	resp := &gamengv1.GetWordsByAudioResponse{
		AudioTranslationPairs: make([]*gamengv1.AudioTranslationPair, 0, len(words)),
	}

	for _, word := range words {
		resp.AudioTranslationPairs = append(resp.AudioTranslationPairs, &gamengv1.AudioTranslationPair{
			Audio:       word.Audio,
			Translation: word.Translation,
		})
	}

	return resp, nil
}
func (ws *WordServer) GetSentencesWithMissingWord(ctx context.Context, req *gamengv1.GetSentencesWithMissingWordRequest) (*gamengv1.GetSentencesWithMissingWordResponse, error) {
	words, err := ws.WordService.GetWords(ctx, req.Offset)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	resp := &gamengv1.GetSentencesWithMissingWordResponse{
		SentencesWords: make(map[string]string, len(words)),
	}

	for _, word := range words {
		resp.SentencesWords[word.Sentence] = word.Word
	}

	return resp, nil
}
func (ws *WordServer) GetWordForTranslation(ctx context.Context, req *gamengv1.GetWordForTranslationRequest) (*gamengv1.GetWordForTranslationResponse, error) {
	words, err := ws.WordService.GetWords(ctx, req.Offset)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	resp := &gamengv1.GetWordForTranslationResponse{
		WordsTranslations: make(map[string]string, len(words)),
	}

	for _, word := range words {
		// for translation game, we need to swap word and translation
		resp.WordsTranslations[word.Translation] = word.Word
	}

	return resp, nil
}
