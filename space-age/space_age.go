// Package space exports a type Planet and a function Age
package space

// Planet is a alias for string
type Planet string

// Constants year seconds per planet
const (
	Earth   = 31557600           // Earth year in seconds
	Mercury = 0.2408467 * Earth  // Mercury in earth years
	Venus   = 0.61519726 * Earth // Venus in earth years
	Mars    = 1.8808158 * Earth  // Mars in earth years
	Jupiter = 11.862615 * Earth  // Jupiter in earth years
	Saturn  = 29.447498 * Earth  // Saturn in earth years
	Uranus  = 84.016846 * Earth  // Uranus in earth years
	Neptune = 164.79132 * Earth  // Neptune in earth years
)

var planetSeconds = map[Planet]float64{
	"Earth":   Earth,
	"Mercury": Mercury,
	"Venus":   Venus,
	"Mars":    Mars,
	"Jupiter": Jupiter,
	"Saturn":  Saturn,
	"Uranus":  Uranus,
	"Neptune": Neptune,
}

// Age receives secods and the aimed planet
func Age(seconds float64, planet Planet) float64 {
	return seconds / planetSeconds[planet]
}
