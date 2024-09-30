package world

type WorldOption func(*worldOption)

type worldOption struct {
	name              string
	location          string
	profile           string
	profileKeys       map[string]string
	confirmedPresence map[string]bool
	confirmedAbsence  map[string]bool
	travelZone        string
}

func defaultOption() worldOption {
	return worldOption{
		name:              "",
		location:          "",
		profile:           "",
		profileKeys:       make(map[string]string),
		confirmedPresence: make(map[string]bool),
		confirmedAbsence:  make(map[string]bool),
		travelZone:        "",
	}
}

func WithName(name string) WorldOption {
	return func(wo *worldOption) {
		wo.name = name
	}
}

func WithLocation(location string) WorldOption {
	return func(wo *worldOption) {
		wo.location = location
	}
}

func WithUWP(uwp string) WorldOption {
	return func(wo *worldOption) {
		wo.profile = uwp
	}
}

func WithProfileData(key, val string) WorldOption {
	return func(wo *worldOption) {
		wo.profileKeys[key] = val
	}
}

func WithConfirmedPresence(codes ...string) WorldOption {
	return func(wo *worldOption) {
		for _, code := range codes {
			wo.confirmedPresence[code] = true
		}
	}
}

func WithConfirmedAbsence(codes ...string) WorldOption {
	return func(wo *worldOption) {
		for _, code := range codes {
			wo.confirmedAbsence[code] = true
		}
	}
}

func WithTravelZone(tz string) WorldOption {
	return func(wo *worldOption) {
		wo.travelZone = tz
	}
}
