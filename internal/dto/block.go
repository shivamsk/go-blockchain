package dto

type Block struct {
	Index     int           `json:"index" `
	Timestamp string        `json:"timestamp"`
	Content   *BlockRequest `json:"content"`
	Hash      string        `json:"hash"`
	PrevHash  string        `json:"prevHash"`
}
