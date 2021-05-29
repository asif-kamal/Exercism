package space



// type Planet string

//Age is...
func Age(Planet string, seconds int) float64 {
	
	
	if Planet == "Mercury" {
		return 0.2408467 * float64(seconds)
	}
	if Planet == "Venus" {
		return 0.61519726 * float64(seconds)
	}
	if Planet == "Earth" {
		return 31557600 * float64(seconds)
	}
	if Planet == "Mars" {
		return 0.61519726 * float64(seconds)
	}
	if Planet == "Jupiter" {
		return 11.862615 * float64(seconds)
	}
	if Planet == "Saturn" {
		return 29.447498 * float64(seconds)
	}
	if Planet == "Uranus" {
		return 84.016846 * float64(seconds)
	}
	if Planet == "Neptune" {
		return 164.79132 * float64(seconds)
	}
	return 0
}
