package characteristic

type CharacteristicOption func(*charOpt)

type charOpt struct {
	effectiveScore int
	maxScore       int
	creationMod    int
	creationDice   int
}

func defaultOptions() charOpt {
	return charOpt{
		effectiveScore: 0,
		maxScore:       0,
		creationMod:    0,
		creationDice:   2,
	}
}

func EffectiveScore(es int) CharacteristicOption {
	return func(co *charOpt) {
		co.effectiveScore = es
	}
}

func MaxScore(ms int) CharacteristicOption {
	return func(co *charOpt) {
		co.maxScore = ms
	}
}

func CreationMod(cm int) CharacteristicOption {
	return func(co *charOpt) {
		co.creationMod = cm
	}
}

func CreationDice(dc int) CharacteristicOption {
	return func(co *charOpt) {
		co.creationDice = dc
	}
}

//////////////////////////////////

func abbCodeType(chName string) (string, string, string) {
	switch chName {
	case STR:
		return "STR", C1, physical
	case DEX:
		return "DEX", C2, physical
	case END:
		return "END", C3, physical
	case INT:
		return "INT", C4, mental
	case EDU:
		return "EDU", C5, mental
	case SOC:
		return "SOC", C6, mental
	case WLT:
		return "WLT", CW, obscure
	case LCK:
		return "LCK", CL, obscure
	case MRL:
		return "MRL", CM, obscure
	case STY:
		return "STY", CS, obscure
	case PSY:
		return "PSY", CP, obscure
	case INS:
		return "INS", C4, mental
	case PCK:
		return "PCK", C6, mental
	case TER:
		return "TER", C6, mental
	case CHA:
		return "CHA", C6, mental
	}
	return "invalid", "invalid", "invalid"
}
