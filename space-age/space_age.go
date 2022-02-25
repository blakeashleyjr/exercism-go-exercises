package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	var mapDB = map[Planet]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	earthYearInSeconds := 31557600

	ageInEarthYears := seconds / float64(earthYearInSeconds)
	return (1 / mapDB[planet]) * ageInEarthYears
}
