package piscine

func IterativePower(nb int, power int) int {
	nbPower := nb
	//if nb == 0 {
	//	return 0
	//}
	if power == 0 {
		return 1
	}
	if power <= -1 {
		return 0
	}
	for i := 1; i < power; i++ {
		nb = nb * nbPower
	}
	return nb
}
