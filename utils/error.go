package utils

import (
	"reflect"
	"strconv"

	"github.com/fenixvlabs/meshkit/pkg/errors"
)

var (
	ErrUnmarshalCode                 = "11043"
	ErrUnmarshalInvalidCode          = "11044"
	ErrUnmarshalSyntaxCode           = "11045"
	ErrUnmarshalTypeCode             = "11046"
	ErrUnmarshalUnsupportedTypeCode  = "11047"
	ErrUnmarshalUnsupportedValueCode = "11048"
	ErrMarshalCode                   = "11049"
	ErrGetBoolCode                   = "11050"
	ErrInvalidProtocolCode           = "11051"
	ErrRemoteFileNotFoundCode        = "11052"
	ErrReadingRemoteFileCode         = "11053"
	ErrReadingLocalFileCode          = "11054"
	ErrGettingLatestReleaseTagCode   = "11055"
	ErrInvalidProtocol               = errors.NewErrorDescription(ErrInvalidProtocolCode, errors.Alert, []string{"invalid protocol: only http, https and file are valid protocols"}, []string{}, []string{"Network protocol is incorrect"}, []string{"Make sure to specify the right network protocol"})
	ErrMissingFieldCode              = "11076"
	ErrExpectedTypeMismatchCode      = "11079"
	ErrJsonToCueCode                 = "11085"
	ErrYamlToCueCode                 = "11086"
	ErrJsonSchemaToCueCode           = "11087"
	ErrCueLookupCode                 = "11089"
)

func ErrCueLookup(err error) error {
	return errors.NewErrorDescription(ErrCueLookupCode, errors.Alert, []string{"Could not lookup the given path in the CUE value"}, []string{err.Error()}, []string{""}, []string{"make sure that the path is a valid cue expression and is correct", "make sure that there exists a field with the given path", "make sure that the given root value is correct"})
}

func ErrJsonSchemaToCue(err error) error {
	return errors.NewErrorDescription(ErrJsonSchemaToCueCode, errors.Alert, []string{"Could not convert given JsonSchema into a CUE Value"}, []string{err.Error()}, []string{"Invalid jsonschema"}, []string{"Make sure that the given value is a valid JSONSCHEMA"})
}

func ErrYamlToCue(err error) error {
	return errors.NewErrorDescription(ErrYamlToCueCode, errors.Alert, []string{"Could not convert given yaml object into a CUE Value"}, []string{err.Error()}, []string{"Invalid yaml"}, []string{"Make sure that the given value is a valid YAML"})
}

func ErrJsonToCue(err error) error {
	return errors.NewErrorDescription(ErrJsonToCueCode, errors.Alert, []string{"Could not convert given json object into a CUE Value"}, []string{err.Error()}, []string{"Invalid json object"}, []string{"Make sure that the given value is a valid JSON"})
}

func ErrExpectedTypeMismatch(err error, expectedType string) error {
	return errors.NewErrorDescription(ErrExpectedTypeMismatchCode, errors.Alert, []string{"Expected the type to be: ", expectedType}, []string{err.Error()}, []string{"Invalid manifest"}, []string{"Make sure that the value provided in the manifest has the needed type."})
}

func ErrMissingField(err error, missingFieldName string) error {
	return errors.NewErrorDescription(ErrMissingFieldCode, errors.Alert, []string{"Missing field or property with name: ", missingFieldName}, []string{err.Error()}, []string{"Invalid manifest"}, []string{"Make sure that the concerned data type has all the required fields/values."})
}

func ErrUnmarshal(err error) error {
	return errors.NewErrorDescription(ErrUnmarshalCode, errors.Alert, []string{"Unmarshal unknown error: "}, []string{err.Error()}, []string{"Invalid object format"}, []string{"Make sure to input a valid JSON object"})
}

func ErrUnmarshalInvalid(err error, typ reflect.Type) error {
	return errors.NewErrorDescription(ErrUnmarshalInvalidCode, errors.Alert, []string{"Unmarshal invalid error for type: ", typ.String()}, []string{err.Error()}, []string{"Invalid object format"}, []string{"Make sure to input a valid JSON object"})
}

func ErrUnmarshalSyntax(err error, offset int64) error {
	return errors.NewErrorDescription(ErrUnmarshalSyntaxCode, errors.Alert, []string{"Unmarshal syntax error at offest: ", strconv.Itoa(int(offset))}, []string{err.Error()}, []string{"Invalid object format"}, []string{"Make sure to input a valid JSON object"})
}

func ErrUnmarshalType(err error, value string) error {
	return errors.NewErrorDescription(ErrUnmarshalTypeCode, errors.Alert, []string{"Unmarshal type error at key: %s. Error: %s", value}, []string{err.Error()}, []string{"Invalid object format"}, []string{"Make sure to input a valid JSON object"})
}

func ErrUnmarshalUnsupportedType(err error, typ reflect.Type) error {
	return errors.NewErrorDescription(ErrUnmarshalUnsupportedTypeCode, errors.Alert, []string{"Unmarshal unsupported type error at key: ", typ.String()}, []string{err.Error()}, []string{"Invalid object format"}, []string{"Make sure to input a valid JSON object"})
}

func ErrUnmarshalUnsupportedValue(err error, value reflect.Value) error {
	return errors.NewErrorDescription(ErrUnmarshalUnsupportedValueCode, errors.Alert, []string{"Unmarshal unsupported value error at key: ", value.String()}, []string{err.Error()}, []string{"Invalid object format"}, []string{"Make sure to input a valid JSON object"})
}

func ErrMarshal(err error) error {
	return errors.NewErrorDescription(ErrMarshalCode, errors.Alert, []string{"Marshal error, Description: %s"}, []string{err.Error()}, []string{"Invalid object format"}, []string{"Make sure to input a valid JSON object"})
}

func ErrGetBool(key string, err error) error {
	return errors.NewErrorDescription(ErrGetBoolCode, errors.Alert, []string{"Error while getting Boolean value for key: %s, error: %s", key}, []string{err.Error()}, []string{"Not a valid boolean"}, []string{"Make sure it is a boolean"})
}

func ErrRemoteFileNotFound(url string) error {
	return errors.NewErrorDescription(ErrRemoteFileNotFoundCode, errors.Alert, []string{"remote file not found at", url}, []string{}, []string{"File doesnt exist in the location", "File name is incorrect"}, []string{"Make sure to input the right file name and location"})
}

func ErrReadingRemoteFile(err error) error {
	return errors.NewErrorDescription(ErrReadingRemoteFileCode, errors.Alert, []string{"error reading remote file"}, []string{err.Error()}, []string{"File doesnt exist in the location", "File name is incorrect"}, []string{"Make sure to input the right file name and location"})
}

func ErrReadingLocalFile(err error) error {
	return errors.NewErrorDescription(ErrReadingLocalFileCode, errors.Alert, []string{"error reading local file"}, []string{err.Error()}, []string{"File doesnt exist in the location", "File name is incorrect"}, []string{"Make sure to input the right file name and location"})
}

func ErrGettingLatestReleaseTag(err error) error {
	return errors.NewErrorDescription(
		ErrGettingLatestReleaseTagCode,
		errors.Alert,
		[]string{"Could not fetch latest stable release from github"},
		[]string{err.Error()},
		[]string{"Failed to make GET request to github", "Invalid response received on github.com/<org>/<repo>/releases/stable"},
		[]string{"Make sure Github is reachable", "Make sure a valid response is available on github.com/<org>/<repo>/releases/stable"},
	)
}
