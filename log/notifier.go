package log

type Notifier interface {
	Send(msg string) (err error)
}
