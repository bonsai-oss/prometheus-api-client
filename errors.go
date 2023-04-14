package prometheus

type ErrNoEndpoint struct{}

func (e ErrNoEndpoint) Error() string {
	return "No endpoint specified"
}
