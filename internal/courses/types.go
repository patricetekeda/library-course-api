package courses


type Course struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Category    string `json:"category"`
	Level	   string `json:"level"`
	Description string `json:"description"`
}
