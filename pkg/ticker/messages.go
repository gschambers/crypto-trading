package ticker

import "encoding/json"

const (
	SNAPSHOT      = "snapshot"
	LEVEL2_UPDATE = "l2update"
)

type Message struct {
	messageType string
	data        interface{}
}

type SnapshotMessage struct {
	Type      string     `json:"type"`
	ProductID string     `json:"product_id"`
	Bids      [][]string `json:"bids"`
	Asks      [][]string `json:"asks"`
}

func (m *SnapshotMessage) unmarshal(data []byte) bool {
	if err := json.Unmarshal(data, m); err != nil {
		return false
	}

	return m.Type == SNAPSHOT
}

type Level2UpdateMessage struct {
	Type      string     `json:"type"`
	ProductID string     `json:"product_id"`
	Changes   [][]string `json:"changes"`
}

func (m *Level2UpdateMessage) unmarshal(data []byte) bool {
	if err := json.Unmarshal(data, m); err != nil {
		return false
	}

	return m.Type == LEVEL2_UPDATE
}
