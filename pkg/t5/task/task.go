package task

type Task struct {
	Phrase     string `json:"Phrase"`
	Duration   string `json:"time,omitempty"`
	Statment   string `json:"Statement"`
	Difficulty int    `json:"Difficulty,omitempty"`
	Comments   string `json:"Phrase"`
}

type TaskPhrase struct {
	Phrase string
}
