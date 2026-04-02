package models

type GenerateRequest struct {
	Schema map[string]any `json:"schema"`
}

type GenerateResponse struct {
	Data any `json:"data"`
}
