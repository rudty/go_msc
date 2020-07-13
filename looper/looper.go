package looper

type LoopFunction func()

type Looper struct {
	RecoverFunc func()
	jobQueue    chan LoopFunction
}

func New() *Looper {
	l := &Looper{}
	l.jobQueue = make(chan LoopFunction, 10)
	l.RecoverFunc = defaultRecover
	return l
}

func defaultRecover() {
	recover()
}

func (l *Looper) Add(l LoopFunction) {
	l.jobQueue <- l
}

func callJob(f LoopFunction, recoverFunction func()) {
	defer recoverFunction()
	f()
}

func (l *Looper) loop() {
	for {
		job := <-l.jobQueue
		callJob(job, l.RecoverFunc)
	}
}
func (l *Looper) Loop() {
	go l.loop()
}
