package key

import (
	"fmt"
	"strings"
)

const (
	CHARACTERISTIC = "characteristic"
	ITEM           = "item"
	SKILL          = "skill"
	WEAPON         = ITEM + "|weapon"
)

func New(properties ...string) string {
	key := ""
	for _, p := range properties {
		key += p + "|"
	}
	return strings.TrimSuffix(key, "|")
}

func ExtractData(key string) ([]string, error) {
	data := strings.Split(key, "|")
	basicPointsFound := false
	for _, point := range data {
		if isBasicPoint(point) {
			basicPointsFound = true
		}
	}
	if !basicPointsFound {
		return nil, fmt.Errorf("key '%v' contains no basic points")
	}
	return data, nil
}

func isBasicPoint(p string) bool {
	switch p {
	default:
		return false
	case
		CHARACTERISTIC,
		ITEM,
		SKILL:
		return true
	}
}

type KeyMaker interface {
}
