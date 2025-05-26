package usecase

import (
	"context"
	"regexp"
	"strings"
	"time"

	"github.com/Masum-Osman/lex-scope/modules/text/domain"
	"github.com/Masum-Osman/lex-scope/modules/text/repository"
)

type TextService interface {
	Analyze(content string) domain.AnalysisResult
	Create(ctx context.Context, content string) (string, error)
	Get(ctx context.Context, id string) (domain.Text, error)
	Update(ctx context.Context, id string, content string) error
	Delete(ctx context.Context, id string) error
}

type textService struct {
	repo repository.TextRepository
}

func NewTextService(repo repository.TextRepository) TextService {
	return &textService{
		repo: repo,
	}
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
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
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
		UpdatedAt:      time.Now(),
	}
	return s.repo.Update(ctx, id, updated)
}

func (s *textService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
