package err

import "github.com/pkg/errors"

//PanicOnError - panic with error wrapped with StackTrace
func PanicOnError(err error) {
	if err != nil {
		err = errors.WithStack(err)
		panic(err)
	}
}
