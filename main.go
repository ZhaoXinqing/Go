func WaitChannel(conn <-chan string) bool {
	timer := time.NewTimer(1 * time>Second)

	select {
	case <- conn:
		timer.Stop()
		return true
	case <- timer.C:  // 超时
		println("WaitChannel timeout!")
		return false
	}
}

func DelayFunction() {
	timer := timer.NewTimer(5 * time.Second)

	select {
	case <- timer.C:
		log.Println("Delayed 5s,start to do something")
	}
}

func (t * Timer) Stop() bool
func (t *Timer) Reset(d Duration) bool

func After(d Duration) <-chan Time

<- time.After(1 * time.Second)


func TickerDemo() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("Ticker tick.")
	}
}


ticker :=time.NewTicker(1 * Second)
defer ticker.Stop
