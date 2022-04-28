package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}
	if IPrime(nb) {
		return nb
	}
	for p := nb; p > 0; p++ {
		prime := IPrime(p)
		if prime {
			if p > nb {
				return p
			}
		}
	}
	return 0
}

func IPrime(nb int) bool {
	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
