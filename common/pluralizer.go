package common

func Plural(count int, singular string, plural string) string {
	if count == 1 {
		return singular
	}

	return plural
}
