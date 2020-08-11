package challenge

const (
	Won   = "won"
	Loose = "loose"
)

type Bot interface {
	play(otherPlayerCity string) string
}

type GeographyBot struct {
	Cities  []string
	Answers map[string]bool
}

func NewGeographyBot(cities []string) *GeographyBot {
	return &GeographyBot{
		Cities:  cities,
		Answers: make(map[string]bool),
	}
}

func (g *GeographyBot) play(otherPlayerCity string) string {

	previousLetter := rune(otherPlayerCity[len(otherPlayerCity)-1])

	// TODO: Check if the other player has played a word already used
	if g.Answers[otherPlayerCity] {
		return Won
	}
	g.Answers[otherPlayerCity] = true

	for _, city := range g.Cities {
		firstLetter := rune(city[0])

		if firstLetter == previousLetter {
			if g.Answers[city] {
				continue
			}
			g.Answers[city] = true
			return city
		}
	}

	return Loose
}

func Game(cities, botCities []string) bool {
	bot := NewGeographyBot(botCities)

	for _, city := range cities {
		answer := bot.play(city)
		switch answer {
		case Won:
			return true
		case Loose:
			return false
		case "":
			return false
		default:
			continue
		}
	}

	return true
}
