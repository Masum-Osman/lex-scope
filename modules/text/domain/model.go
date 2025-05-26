package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Text struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Content        string             `json:"content" bson:"content"`
	UserID         string             `json:"user_id,omitempty" bson:"user_id,omitempty"`
	AnalysisResult AnalysisResult     `json:"analysis_result" bson:"analysis_result"`
	CreatedAt      time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at" bson:"updated_at"`
}

type AnalysisResult struct {
	WordCount      int      `json:"word_count"`
	CharacterCount int      `json:"character_count"`
	SentenceCount  int      `json:"sentence_count"`
	ParagraphCount int      `json:"paragraph_count"`
	LongestWords   []string `json:"longest_words"`
}

type TextRequest struct {
	Content string `json:"content"`
}
