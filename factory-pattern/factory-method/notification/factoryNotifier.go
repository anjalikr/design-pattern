package notification

type NotifierFactory interface {
	Create() Notifier
}
