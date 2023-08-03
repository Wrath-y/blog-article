package event

type PublisherI interface {
	Publish() error
}
