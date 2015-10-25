package requests

import "github.com/etcinit/gonduit/constants"

// ManiphestQueryRequest represents a request to maniphest.query.
type ManiphestQueryRequest struct {
	IDs          []string                      `json:"ids"`
	PHIDs        []string                      `json:"PHIDs"`
	OwnerPHIDs   []string                      `json:"ownerPHIDs"`
	AuthorPHIDs  []string                      `json:"authorPHIDs"`
	ProjectPHIDs []string                      `json:"projectPHIDs"`
	CCPHIDs      []string                      `json:"ccPHIDs"`
	FullText     string                        `json:"fullText"`
	Status       constants.ManiphestTaskStatus `json:"status"`
	Order        constants.ManiphestQueryOrder `json:"order"`
	Limit        uint64                        `json:"offset"`
	Offset       uint64                        `json:"offset"`
	Request
}