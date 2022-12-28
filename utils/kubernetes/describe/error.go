package describe

import "github.com/fenixvlabs/meshkit/pkg/errors"

var (
	ErrGetDescriberFuncCode = "not set"
)

func ErrGetDescriberFunc() error {
	return errors.NewErrorDescription(
		ErrGetDescriberFuncCode,
		errors.Fatal,
		[]string{"Failed to get describer for the resource"},
		[]string{"invalid kubernetes object type or object type not supported in meshkit", "Describer not found for the defined Resource"},
		nil, nil,
	)
}
