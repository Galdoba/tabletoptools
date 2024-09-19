package skill

type SkillOption func(*skillOpt)

type skillOpt struct {
	effectiveScore int
	maxScore       int
}

func defaultOptions() skillOpt {
	return skillOpt{
		effectiveScore: 0,
		maxScore:       15,
	}
}

func EffectiveScore(es int) SkillOption {
	return func(so *skillOpt) {
		so.effectiveScore = es
	}
}

func MaxScore(ms int) SkillOption {
	return func(so *skillOpt) {
		so.maxScore = ms
	}
}

//////////////////////////////////
