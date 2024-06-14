package task

import "github.com/Galdoba/tabletoptools/pkg/gametime"

type Task struct {
	Phrase     string `json:"Phrase"`
	Assets     []Asset
	Difficulty int    `json:"Difficulty,omitempty"`
	Comments   string `json:"Comments"`
	StartTime  string `json:"Start At,omitempty"`
	Duration   string `json:"Duration,omitempty"`
	EndTime    string `json:"End By,omitempty"`
	Complete   bool   `json:"Complete,omitempty"`
}

type Asset struct {
	AType    string `json:"Type"`
	Key      string `json:"Description,omitempty"`
	Mod      int    `json:"Mod,omitempty"`
	Required bool   `json:"Required,omitempty"`
}

func NewTask() *Task {
	t := Task{}
	return &t
}

func (t *Task) WithPhrase(phrase string) *Task {
	t.Phrase = phrase
	return t
}

func (t *Task) WithDuration(dur gametime.Duration) *Task {
	t.Duration = dur.String()
	return t
}

type Resolution struct {
	Outcome string //Success=1, Failure=-1, none=0
	// Success                  bool
	// Failure                  bool
	// SpectacularSuccess       bool
	// SpectacularFailure       bool
	SpectacularlyStupid bool
	// SpectacularlyInteresting bool
}

const (
	SUCCESS                  = "Success"
	FAILURE                  = "Failure"
	SPECTACULARLY_INTERESING = "Spectacularly Interesting"
	SPECTACULAR_SUCCESS      = "Spectacular Success"
	SPECTACULAR_FAILURE      = "Spectacular Failure"
)

func newResolution(tn int, result []int) Resolution {
	r := Resolution{}
	if len(result) < tn {
		r.SpectacularlyStupid = true
	}
	sSuccess := ifScored(1, result) >= 3
	sFailure := ifScored(6, result) >= 3
	switch {
	case sSuccess && sFailure:
		r.Outcome = SPECTACULARLY_INTERESING
	case sSuccess:
		r.Outcome = SPECTACULAR_SUCCESS
	case sFailure:
		r.Outcome = SPECTACULAR_FAILURE
	default:
		r.Outcome = FAILURE
		if sum(result) <= tn {
			r.Outcome = SUCCESS
		}
	}
	// return r
	// if sSuccess && sFailure {
	// 	r.Outcome = SPECTACULARLY_INTERESING
	// 	return r
	// }
	// if sSuccess {
	// 	r.Outcome = SPECTACULAR_SUCCESS
	// 	return r
	// }
	// if sFailure {
	// 	r.Outcome = SPECTACULAR_FAILURE
	// 	return r
	// }
	// switch sum(result) <= tn {
	// case true:
	// 	r.Outcome = SUCCESS
	// case false:
	// 	r.Outcome = FAILURE
	// }
	return r
}

func sum(result []int) int {
	s := 0
	for _, r := range result {
		s += r
	}
	return s
}

func ifScored(i int, result []int) int {
	s := 0
	for _, r := range result {
		if r == i {
			s++
		}
	}
	return s
}
