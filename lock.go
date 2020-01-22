package main

type lock struct {
	c chan struct{}
}

func (l *lock) Init() {
	l.c = make(chan struct{}, 1)
}

func (l lock) TryAcquire() (success bool) {
	select {
	case l.c <- struct{}{}:
		return true
	default:
		return false
	}
}

func (l lock) Release() (success bool) {
	select {
	case <-l.c:
		return true
	default:
		return false
	}

}
