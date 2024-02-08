package main

//Attibute struct
type attribute struct {
	Value string `json:"value"`
	Type string `json:"type"`
} 

//Attibute struct
type OutgoingData struct {
	Event string `json:"event"`
	EventType string `json:"event_type"`
	AppID string `json:"app_id"`
	UserID string `json:"user_id"`
	MessageID string `json:"message_id"`
	PageTitle string `json:"page_title"`
	PageURL string `json:"page_url"`
	BrowserLanguage string `json:"browser_language"`
	ScreenSize string `json:"screen_size"`
	Attributes map[string]attribute `json:"attributes"`
	UserAttributes map[string]attribute `json:"traits"`
}