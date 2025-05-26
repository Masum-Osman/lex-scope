package handler

import (
	"net/http"

	"github.com/Masum-Osman/lex-scope/modules/text/domain"
	"github.com/Masum-Osman/lex-scope/modules/text/usecase"
	"github.com/gin-gonic/gin"
)

type TextHandler struct {
	textService usecase.TextService
}

func NewTextHandler(service usecase.TextService) *TextHandler {
	return &TextHandler{
		textService: service,
	}
}

func (h *TextHandler) RegisterRoutes(rg *gin.RouterGroup) {
	rg.POST("/text", h.CreateText)

	rg.GET("/texts/:id/word-count", h.GetWordCount)
	rg.GET("/texts/:id/character-count", h.GetCharacterCount)
	rg.GET("/texts/:id/sentence-count", h.GetSentenceCount)
	rg.GET("/texts/:id/paragraph-count", h.GetParagraphCount)
	rg.GET("/texts/:id/longest-words", h.GetLongestWords)
}

func (h *TextHandler) CreateText(c *gin.Context) {
	var req domain.TextRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	result := h.textService.Analyze(req.Content)
	c.JSON(http.StatusOK, result)
}

func (h *TextHandler) GetWordCount(c *gin.Context) {
	h.analyzeAndRespond(c, func(result domain.AnalysisResult) interface{} {
		return gin.H{"word_count": result.WordCount}
	})
}

func (h *TextHandler) GetCharacterCount(c *gin.Context) {
	h.analyzeAndRespond(c, func(result domain.AnalysisResult) interface{} {
		return gin.H{"character_count": result.CharacterCount}
	})
}

func (h *TextHandler) GetSentenceCount(c *gin.Context) {
	h.analyzeAndRespond(c, func(result domain.AnalysisResult) interface{} {
		return gin.H{"sentence_count": result.SentenceCount}
	})
}

func (h *TextHandler) GetParagraphCount(c *gin.Context) {
	h.analyzeAndRespond(c, func(result domain.AnalysisResult) interface{} {
		return gin.H{"paragraph_count": result.ParagraphCount}
	})
}

func (h *TextHandler) GetLongestWords(c *gin.Context) {
	h.analyzeAndRespond(c, func(result domain.AnalysisResult) interface{} {
		return gin.H{"longest_words": result.LongestWords}
	})
}

func (h *TextHandler) analyzeAndRespond(c *gin.Context, build func(result domain.AnalysisResult) interface{}) {
	text := "The quick brown fox jumps over the lazy dog. The lazy dog slept in the sun."

	result := h.textService.Analyze(text)

	c.JSON(http.StatusOK, build(result))
}
