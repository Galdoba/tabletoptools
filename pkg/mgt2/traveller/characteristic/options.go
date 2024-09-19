package characteristic

import (
	. "github.com/Galdoba/tabletoptools/pkg/mgt2/key"
)

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
	case CHAR_NAME_STR:
		return "STR", CHAR_CODE_C1, physical
	case CHAR_NAME_DEX:
		return "DEX", CHAR_CODE_C2, physical
	case CHAR_NAME_END:
		return "END", CHAR_CODE_C3, physical
	case CHAR_NAME_INT:
		return "INT", CHAR_CODE_C4, mental
	case CHAR_NAME_EDU:
		return "EDU", CHAR_CODE_C5, mental
	case CHAR_NAME_SOC:
		return "SOC", CHAR_CODE_C6, mental
	case CHAR_NAME_WLT:
		return "WLT", CHAR_CODE_CW, obscure
	case CHAR_NAME_LCK:
		return "LCK", CHAR_CODE_CL, obscure
	case CHAR_NAME_MRL:
		return "MRL", CHAR_CODE_CM, obscure
	case CHAR_NAME_STY:
		return "STY", CHAR_CODE_CS, obscure
	case CHAR_NAME_PSY:
		return "PSY", CHAR_CODE_CP, obscure
	case CHAR_NAME_INS:
		return "INS", CHAR_CODE_C4, mental
	case CHAR_NAME_PCK:
		return "PCK", CHAR_CODE_C6, mental
	case CHAR_NAME_TER:
		return "TER", CHAR_CODE_C6, mental
	case CHAR_NAME_CHA:
		return "CHA", CHAR_CODE_C6, mental
	}
	return "invalid", "invalid", "invalid"
}
