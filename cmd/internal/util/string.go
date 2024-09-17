package util

func ToString(value []string) string {
	stringFormat := ""
	for i := range len(value) {
		stringFormat += value[i] + " "
	}

	return stringFormat
}
