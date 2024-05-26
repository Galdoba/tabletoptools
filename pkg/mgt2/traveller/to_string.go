package traveller

import (
	"fmt"
	"strings"

	"github.com/Galdoba/tabletoptools/pkg/characteristic"
	"github.com/charmbracelet/lipgloss"
)

var long = "Jack-of-All-Trades 1"
var personalBlockWidth = len(long)*3 + 2

const (
	RED   = lipgloss.Color("1")
	BLACK = lipgloss.Color("0")
	WHITE = lipgloss.Color("15")
	GRAY  = lipgloss.Color("#808080")

	// FG_BLACK = lipgloss.Color("0")
	// FG_BLACK = lipgloss.Color("0")
	// FG_BLACK = lipgloss.Color("0")
)

func (tr *Traveller) Render() string {
	s := ""
	chrs := coreCharsBlock(tr.Characteristics)
	pers := personalBlock(tr.Personal)
	s = lipgloss.JoinVertical(lipgloss.Left, chrs, pers)
	return s
}

///////////CORE CHARS BLOCK

func coreCharsBlock(cs *characteristic.Set) string {
	st := titleStyle(lipgloss.NewStyle())
	cd := newCharacteristicsData(cs)
	title := "  CHARACTERISTICS "
	for len(title) < personalBlockWidth {
		title += " "
	}
	title = st.Render(title)

	out := title
	st = normalStyle(st)
	out += "\n" + st.Render(cd.Header)
	out += "\n" + st.Render(cd.Vals)

	return st.Render(out)
}

// /////////PERSONAL BLOCK
func personalBlock(pd *PersonalData) string {
	st := lipgloss.NewStyle()
	//wd := len(long)*3 + 2
	output := headerPersonal(st)
	output += "\n" + namePersonal(st, pd)
	output += "\n" + speciesPersonal(st, pd)
	output += "\n" + traitsPersonal(st, pd)

	return st.Render(output)
}

func headerPersonal(st lipgloss.Style) string {
	text := "  PERSONAL DATA FILE"
	text = setWidth(text, personalBlockWidth)
	st = titleStyle(st)
	return st.Render(text)
}

func namePersonal(st lipgloss.Style, pd *PersonalData) string {
	text := setWidth(fmt.Sprintf(" NAME: %v", pd.Name), personalBlockWidth-14)
	text = setWidth(fmt.Sprintf("%v AGE: %v", text, pd.Age), personalBlockWidth)
	st = normalStyle(st)
	return st.Render(text)
}

func speciesPersonal(st lipgloss.Style, pd *PersonalData) string {
	text := " SPECIES: " + pd.Species
	text = setWidth(text, personalBlockWidth/2)
	text = setWidth(text+" HOMEWORLD: "+pd.Homeworld, personalBlockWidth)
	st = normalStyle(st)
	return st.Render(text)
}

func traitsPersonal(st lipgloss.Style, pd *PersonalData) string {
	text := " TRAITS: "
	if len(pd.Traits) == 0 {
		text += "NONE"
	}
	text = setWidth(text, personalBlockWidth)
	st = normalStyle(st)
	st.Render(text)
	for _, trt := range pd.Traits {
		st = normalStyle(st)
		text += "\n" + st.Render(setWidth(fmt.Sprintf("     %v", trt.Name), personalBlockWidth))
		st = secondaryStyle(st)
		text += "\n" + st.Render(setWidth(fmt.Sprintf("      -%v", trt.Description), personalBlockWidth))
	}
	return text
}

// STYLES
func normalStyle(st lipgloss.Style) lipgloss.Style {
	return st.Background(BLACK).Foreground(WHITE)
}

func secondaryStyle(st lipgloss.Style) lipgloss.Style {
	return st.Background(BLACK).Foreground(GRAY)
}

func titleStyle(st lipgloss.Style) lipgloss.Style {
	return st.Background(RED).Foreground(BLACK)
}

//HELPERS

func setWidth(text string, w int) string {
	letters := split(text)
	out := ""
	for i, l := range letters {
		if i >= w {
			return out
		}
		out += l
	}
	for len(out) < w {
		out += " "
	}
	return out
}

func split(text string) []string {
	return strings.Split(text, "")
}
