package models

type Command struct {
	X    float64 `json:"x,omitempty"`
	Y    float64 `json:"y,omitempty"`
	Key  string  `json:"key,omitempty"`
	Text string  `json:"text,omitempty"`
}

type WSMessage struct {
	Event string  `json:"event"`
	X     float64 `json:"x,omitempty"`
	Y     float64 `json:"y,omitempty"`
	Key   string  `json:"key,omitempty"`
	Text  string  `json:"text,omitempty"`
}
