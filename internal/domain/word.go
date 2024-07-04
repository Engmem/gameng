package domain

import (
	"math/rand"
	"time"
)

type WordCard struct {
	Word        string
	Translation string
	Audio       []byte // audio file
	Sentence    string // sentence with the word
}

func GenerateWordCard() *WordCard {
	const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789"

	seededRand := rand.New(
		rand.NewSource(time.Now().UnixNano()),
	)

	randBytes := func(length int) []byte {
		b := make([]byte, length)
		for i := range b {
			b[i] = charset[seededRand.Intn(len(charset))]
		}
		return b
	}

	return &WordCard{
		Word:        string(randBytes(10)),
		Translation: string(randBytes(10)),
		Audio:       randBytes(10),
		Sentence:    string(randBytes(10)),
	}
}
