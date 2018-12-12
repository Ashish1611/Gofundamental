package days



func Days(ch int) string{

	switch ch {

	case 1:
		return "Sunday"

	case 2:
		return "Monday"

	case 3:
        return "Tue"
	case 4:
		return "Wed"

	case 5:
        return "Thu"
	case 6:
		return "Fri"

	case 7:
		return "Sat"

	}

	return "Not a day of week"
}
