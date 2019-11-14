package retry

import (
	"time"

	"github.com/jpillora/backoff"
	"golang.org/x/xerrors"
)

var (
	defaultMaxRetries = 3
	defaultTimeout    = -1 * time.Second
	defaultBackoff    = &backoff.Backoff{
		Min:    500 * time.Millisecond,
		Max:    1 * time.Second,
		Factor: 1,
		Jitter: true,
	}
	errorExceededTimeout = xerrors.New("exceeded timeout")
)

type option struct {
	maxRetries int
	timeout    time.Duration
	backoff    *backoff.Backoff
}

type RetryOption func(*option)

func WithMaxRetries(n int) RetryOption {
	return RetryOption(func(opt *option) {
		opt.maxRetries = n
	})
}

func WithTimeout(timeout time.Duration) RetryOption {
	return RetryOption(func(opt *option) {
		opt.timeout = timeout
	})
}

func WithBackoff(backoff *backoff.Backoff) RetryOption {
	return RetryOption(func(opt *option) {
		opt.backoff = backoff
	})
}

type RetryConditionFn func() error

func Retry(fn RetryConditionFn, opts ...RetryOption) error {
	opt := option{
		maxRetries: defaultMaxRetries,
		timeout:    defaultTimeout,
		backoff:    defaultBackoff,
	}

	for _, o := range opts {
		o(&opt)
	}

	for attempt := 0; attempt < opt.maxRetries; attempt++ {
		if err := fn(); err == nil {
			return nil
		}

		select {
		case <-time.After(opt.timeout):
			return errorExceededTimeout
		case <-time.After(opt.backoff.Duration()):
		}
	}

	return nil
}
