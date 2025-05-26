package usecase

import (
	"context"
	"regexp"
	"strings"

	"github.com/Masum-Osman/lex-scope/modules/text/domain"
)

type TextService interface {
	Analyze(content string) domain.AnalysisResult
}

func (t TextService) Create(ctx context.Context, s string) (any, any) {
	panic("unimplemented")
}

type textService struct{}

func NewTextService() TextService {
	return &textService{}
}

func (s *textService) Analyze(content string) domain.AnalysisResult {
	cleaned := strings.ToLower(content)
	cleaned = strings.ReplaceAll(cleaned, "\n", " ")

	words := strings.Fields(cleaned)
	sentences := regexp.MustCompile(`[.!?]+`).Split(cleaned, -1)
	paragraphs := strings.Split(cleaned, "\n")

	filter := func(ss []string) []string {
		out := []string{}
		for _, s := range ss {
			s = strings.TrimSpace(s)
			if len(s) > 0 {
				out = append(out, s)
			}
		}
		return out
	}

	words = filter(words)
	sentences = filter(sentences)
	paragraphs = filter(paragraphs)

	longest := findLongestWords(words)

	return domain.AnalysisResult{
		WordCount:      len(words),
		CharacterCount: len(strings.ReplaceAll(cleaned, " ", "")),
		SentenceCount:  len(sentences),
		ParagraphCount: len(paragraphs),
		LongestWords:   longest,
	}
}

func findLongestWords(words []string) []string {
	maxLen := 0
	longest := []string{}

	for _, word := range words {
		word = strings.Trim(word, ".,!?;:\"'")
		length := len(word)
		if length > maxLen {
			maxLen = length
			longest = []string{word}
		} else if length == maxLen {
			longest = append(longest, word)
		}
	}

	return longest
}

func (s *textService) Create(ctx context.Context, content string) (string, error) {
	result := s.Analyze(content)
	text := domain.Text{
		Content:        content,
		AnalysisResult: result,
	}

	return s.repo.Save(ctx, text)
}

func (s *textService) Get(ctx context.Context, id string) (domain.Text, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *textService) Update(ctx context.Context, id string, content string) error {
	result := s.Analyze(content)
	updated := domain.Text{
		Content:        content,
		AnalysisResult: result,
	}
	return s.repo.Update(ctx, id, updated)
}

func (s *textService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
