package handlers

type DocumentHandler struct{}

func NewDocumentHandler() *DocumentHandler {
	return &DocumentHandler{}
}

func (documentHandler *DocumentHandler) ReadDocument(document Document)
