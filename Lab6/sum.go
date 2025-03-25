func calculaSoma(files []string) map[string]string{
	sumChan := make(chan struct{path , sum string})

	var wg sync.WaitGroup

	for _, file := range files{
		wg.Add(1)
		go func(f string){
			defer wg.Done()
			sum:= calculaSum(f)
			sumChan <- struct{path, sum string}{f, sum}
		}(file)
	}

	go func(){
		wg.wait()
		close(sumChan)
	}()

	sums := make(map[string]string)
	for result := range sumChan{
		sums[result.path] = result.sum
	}

	return sums
}