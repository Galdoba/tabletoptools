package profile

import (
	"errors"
	"fmt"
	"strings"
)

var ErrNoFeed = errors.New("no feed")

type universalProfile struct {
	profileType   string
	profilePoints map[string]string
	exportRule    []string
}

type Profile interface {
	Profile() string
	GetValue(string) string
	SetValue(string, string) error
	GenerateMissingData(...Roller) error
}

func (up *universalProfile) String() string {
	return up.Profile()
}

func New(profileType string) *universalProfile {
	up := universalProfile{}
	up.profileType = profileType
	up.profilePoints = make(map[string]string)
	up.exportRule = exportRule(up.profileType)
	switch up.profileType {
	case UWP:
		up.profilePoints[KEY_Temp] = "?"
	}
	for _, key := range up.exportRule {
		switch key {
		case SEPARATOR1:
			continue
		}
		up.profilePoints[key] = "?"
	}
	return &up
}

func (up *universalProfile) Profile() string {
	prf := ""
	for _, key := range up.exportRule {
		switch key {
		case SEPARATOR1:
			prf += SEPARATOR1
		default:
			prf += up.GetValue(key)
		}
	}
	return prf
}

func (up *universalProfile) GetValue(key string) string {
	if val, ok := up.profilePoints[key]; ok {
		return val
	}
	return "?"
}

func (up *universalProfile) SetValue(key, value string) error {
	if _, ok := up.profilePoints[key]; !ok {
		return fmt.Errorf("key '%v' does not exeist", key)
	}
	up.profilePoints[key] = value
	return nil
}

func NewUWP(uwp string) (*universalProfile, error) {
	if uwp == "" {
		return nil, ErrNoFeed
	}
	prf := New(UWP)
	valMap, err := parseUWP(uwp)
	if err != nil {
		return nil, fmt.Errorf("failed to parse '%v' feed: %v", uwp, err)
	}
	for key, val := range valMap {
		if err := prf.SetValue(key, val); err != nil {
			return nil, fmt.Errorf("failed to assemple data from '%v' feed: %v", uwp, err)
		}
	}
	return prf, nil
}

func parseUWP(uwp string) (map[string]string, error) {
	valueMap := make(map[string]string)
	if len(uwp) != 9 {
		return nil, fmt.Errorf("uwp profile expect string with 9 glyphs")
	}
	for i, val := range strings.Split(uwp, "") {
		switch i {
		case 0:
			valueMap[KEY_Starport] = val
		case 1:
			valueMap[KEY_Size] = val
		case 2:
			valueMap[KEY_Atmo] = val
		case 3:
			valueMap[KEY_Hydr] = val
		case 4:
			valueMap[KEY_Pops] = val
		case 5:
			valueMap[KEY_Govr] = val
		case 6:
			valueMap[KEY_Laws] = val
		case 8:
			valueMap[KEY_TL] = val
		}
	}
	return valueMap, nil
}
