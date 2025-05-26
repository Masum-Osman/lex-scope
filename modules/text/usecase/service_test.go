package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnalyze_SimpleTRext(t *testing.T) {
	service := NewTextService()

	text := "The quick brown fox jumps over the lazy dog. The lazy dog slept in the sun."
	result := service.Analyze(text)

	assert.Equal(t, 16, result.WordCount)
	assert.Equal(t, 60, result.CharacterCount)
	assert.Equal(t, 2, result.SentenceCount)
	assert.Equal(t, 1, result.ParagraphCount)
	assert.ElementsMatch(t, []string{"jumps"}, result.LongestWords)
}

func TestAnalyze_MultiParagraph(t *testing.T) {
	service := NewTextService()

	text := `This is the first paragraph.
And still part of it.

This is a new paragraph.`
	result := service.Analyze(text)

	assert.Equal(t, 15, result.WordCount)
	assert.Equal(t, 61, result.CharacterCount)
	assert.Equal(t, 3, result.SentenceCount)
	assert.Equal(t, 3, result.ParagraphCount)
	assert.ElementsMatch(t, []string{"paragraph"}, result.LongestWords)
}

func TestAnalyze_EmptyText(t *testing.T) {
	service := NewTextService()
	result := service.Analyze("")

	assert.Equal(t, 0, result.WordCount)
	assert.Equal(t, 0, result.CharacterCount)
	assert.Equal(t, 0, result.SentenceCount)
	assert.Equal(t, 0, result.ParagraphCount)
	assert.Empty(t, result.LongestWords)
}
