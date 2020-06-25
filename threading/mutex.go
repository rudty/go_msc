package threading

// Mutex unix pthread_mutex, windows CriticalSection impl
type Mutex struct {
	data [65]byte
}
