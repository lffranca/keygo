package oauth2

import "errors"

type Options struct {
	ID     *string
	Secret *string
	Domain *string
}

func (pkg *Options) validate() error {
	if pkg.ID == nil {
		return errors.New("id param is required")
	}

	if pkg.Secret == nil {
		return errors.New("secret param is required")
	}

	if pkg.Domain == nil {
		return errors.New("domain param is required")
	}

	return nil
}
