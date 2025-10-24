package data

type AddGameParamsJSON struct {
	Title       string   `json:"title"`
	Developer   string   `json:"developer"`
	Publisher   string   `json:"publisher"`
	ReleaseYear int      `json:"release_year"`
	Platforms   []string `json:"platforms"`
	Description string   `json:"description"`
}
