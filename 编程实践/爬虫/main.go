var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

go func() {
	os.Stdin.Read(make([]byte,1))
	close(done)
}

for {
	select {
	case <- done:
		for range fileSizes {
			// do nothing
		}
		return
	case size,ok := <-fileSizes:
		// ...
	}
}