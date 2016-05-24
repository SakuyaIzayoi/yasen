package main

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getSuffixString(suffixIndex int) string {
	var suffix string

	switch suffixIndex {
	case 1:
		suffix = " Kai"
	case 5:
		suffix = " Kou Kai"
	case 4:
		suffix = " Kou"
	case 6:
		suffix = " Kou Kai2"
	case 9:
		suffix = " Kai2 A"
	case 2:
		suffix = " Kai2"
	case 8:
		suffix = " drei"
	case 3:
		suffix = " Kou"
	case 10:
		suffix = " Kai2 B"
	case 7:
		suffix = " zwei"
	default:
		suffix = ""
	}
	return suffix
}
