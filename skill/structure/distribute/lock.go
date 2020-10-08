package distribute

type Lock interface {
	Lock() error
	Unlock() error
}
