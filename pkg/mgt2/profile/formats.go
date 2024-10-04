package v2

import "fmt"

const (
	UWP  = "Universal World Profile"
	Size = "Size"
)

func formatKeys(f string) []string {
	switch f {
	default:
		return nil
	case UWP:
		return []string{KEY_Port, KEY_Size, KEY_Atmo, KEY_Hydr, KEY_Pops, KEY_Govr, KEY_Laws, SEP1, KEY_TL}
	case Size:
		return []string{KEY_Size, SEP1, KEY_Size_Dkm, SEP1, KEY_Size_D, SEP1, KEY_Size_G, SEP1, KEY_Size_M}
	}
	panic(fmt.Sprintf("profile format '%v' is not implemented", f))
	return nil
}
