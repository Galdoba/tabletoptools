package profile

func exportRule(prType string) []string {
	switch prType {
	default:
		return nil
	case UWP:
		return []string{KEY_Starport, KEY_Size, KEY_Atmo, KEY_Hydr,
			KEY_Pops, KEY_Govr, KEY_Laws, SEPARATOR1, KEY_TL}
	}
}
