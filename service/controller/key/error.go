package key

import "github.com/giantswarm/microerror"

var invalidMemoryConfigurationError = &microerror.Error{
	Kind: "invalidMemoryConfigurationError",
}

func IsInvalidMemoryConfigurationError(err error) bool {
	return microerror.Cause(err) == invalidMemoryConfigurationError
}

var missingAnnotationError = &microerror.Error{
	Kind: "missingAnnotationError",
}

func IsMissingAnnotationError(err error) bool {
	return microerror.Cause(err) == missingAnnotationError
}

var missingNodeInternalIP = &microerror.Error{
	Kind: "missingNodeInternalIP",
}

func IsMissingNodeInternalIP(err error) bool {
	return microerror.Cause(err) == missingNodeInternalIP
}

var missingVersionError = &microerror.Error{
	Kind: "missingVersionError",
}

func IsMissingVersionError(err error) bool {
	return microerror.Cause(err) == missingVersionError
}

var wrongTypeError = &microerror.Error{
	Kind: "wrongTypeError",
}

// IsWrongTypeError asserts wrongTypeError.
func IsWrongTypeError(err error) bool {
	return microerror.Cause(err) == wrongTypeError
}
