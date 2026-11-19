package domain

type Banner struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Subtitle   string `json:"subtitle"`
	ImageURL   string `json:"imageUrl"`
	TargetSlug string `json:"targetSlu,omitempty"`
	BadgeText  string `json:"badgeText,omitempty"`
	Source     string `json:"source"`
}