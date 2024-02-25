package iteration

func Repeat(text string, repeatTime int) string {
	var repeated string
	for i := 0; i < repeatTime; i++ {
		repeated += text
	}
	return repeated
}
