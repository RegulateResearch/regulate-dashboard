package xmlutils

import (
	"encoding/xml"
	"fmt"
	"frascati/exception"

	"github.com/go-playground/validator/v10"
)

// to ensure that the value is properly parsed,
// please annotate each of the desired fields with:
// `validate:"required"`
func ParseGeneric[T any](data []byte, initFn func() T) (T, exception.Exception) {
	res := initFn()
	err := xml.Unmarshal(data, &res)
	if err != nil {
		newErr := fmt.Errorf("xml unmarshal error: %w", err)
		return res, exception.NewBaseException(exception.CAUSE_INTERNAL, "xml_parser", exception.INTERNAL, newErr)
	}

	checker := validator.New()
	if validateErr := checker.Struct(res); validateErr != nil {
		newErr := fmt.Errorf("xml validate error: %w", validateErr)
		return res, exception.NewBaseException(exception.CAUSE_INTERNAL, "xml_parser", exception.INTERNAL, newErr)
	}

	return res, nil
}
