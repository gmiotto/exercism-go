package letter

// ConcurrentFrequency 1
func ConcurrentFrequency(texts []string) FreqMap {
	outputFreq := make(FreqMap)
	joiner := make(chan FreqMap)
	go producer(texts, joiner)
	for m := range joiner {
		for k, v := range m {
			outputFreq[k] += v
		}
	}
	return outputFreq
}

func producer(texts []string, chnl chan FreqMap) {
	for i := 0; i < len(texts); i++ {
		chnl <- Frequency(texts[i])
	}
	close(chnl)
}
