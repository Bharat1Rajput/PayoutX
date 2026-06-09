package port

import "context"

type DLQPublisher interface {
	Publish(
		ctx context.Context,
		key []byte,
		value []byte,
	) error
}
