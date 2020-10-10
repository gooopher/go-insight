package distribute

type Lock interface {
	Lock() error
	Unlock() error
	Proccess(dealFunc func() error) error
}
