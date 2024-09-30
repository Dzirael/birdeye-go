package birdeye

func toString(arr []string) (str string) {
	for _, v := range arr {
		str += v + ","
	}
	return str[:len(str)-1]
}
