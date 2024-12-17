package event

type Event struct {
	Type     string //
	Name     string //
	ID       int    //
	Text     string
	Start    string
	Decidion []string
	Code     string //Common-Life_event-2.json
}

type Action struct {
	Check       string //Gabbler 8+
	Description string //risk using Gambler
	Success     Reward
}

type Consequances struct {
	takeAll []string `json:"take all,omitempty"`
	takeOne []string `json:"take one of,omitempty"`
}

type Reward struct {
}
