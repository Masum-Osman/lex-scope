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

	rg.GET("/texts/:id", h.GetText)
	rg.PUT("/texts/:id", h.UpdateText)
	rg.DELETE("/texts/:id", h.DeleteText)
}

func (h *TextHandler) CreateText(c *gin.Context) {
	var req domain.TextRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	id, err := h.textService.Create(c, req.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "create failed"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
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

func (h *TextHandler) GetText(c *gin.Context) {
	id := c.Param("id")
	text, err := h.textService.Get(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, text)
}

func (h *TextHandler) UpdateText(c *gin.Context) {
	id := c.Param("id")
	var req domain.TextRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	err := h.textService.Update(c, id, req.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (h *TextHandler) DeleteText(c *gin.Context) {
	id := c.Param("id")
	err := h.textService.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
