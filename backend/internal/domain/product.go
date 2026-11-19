package domain

type Product struct {
	ID          string   `json:"id"`
	Slug        string   `json:"slug"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	CoverURL    string   `json:"coverUrl"`
	Platform    string   `json:"platform"`
	Region      string   `json:"region"`
	Tags        []string `json:"tags,omitempty"`
}