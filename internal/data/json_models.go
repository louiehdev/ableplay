package data

type GameParamsJSON struct {
	Title       string   `json:"title"`
	Developer   string   `json:"developer"`
	Publisher   string   `json:"publisher"`
	ReleaseYear string   `json:"release_year"`
	Platforms   []string `json:"platforms"`
	Description string   `json:"description"`
}
