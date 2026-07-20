package fallback

// Execute wraps primary function blocks, running the fallback function on primary errors.
func Execute(fn func() (interface{}, error), fallback func(err error) (interface{}, error)) (interface{}, error) {
	res, err := fn()
	if err != nil {
		return fallback(err)
	}
	return res, nil
}
