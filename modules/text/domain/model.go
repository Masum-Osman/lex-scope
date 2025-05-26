package domain

type Text struct {
	ID      string `json:"id" bson:"_id,omitempty"`
	Content string `json:"content" bson:"content"`
	UserID  string `json:"user_id" bson:"user_id"`
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
