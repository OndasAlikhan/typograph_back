package dto

import "typograph_back/src/model"

type ParagraphResponse struct {
	ID     uint   `json:"id"`
	Source string `json:"source"`
	Text   string `json:"text"`
	Length uint   `json:"length"`
}

func NewParagraphResponse(paragraph *model.Paragraph) *ParagraphResponse {
	return &ParagraphResponse{
		ID:     paragraph.ID,
		Source: paragraph.Source,
		Length: paragraph.Length,
		Text:   paragraph.Text,
	}
}
