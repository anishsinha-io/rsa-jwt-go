package errors

type AppError struct {
	typ string
	msg error
}

func IOError(err error) AppError {
	return AppError{typ: "io", msg: err}
}

func ParseError(err error) AppError {
	return AppError{typ: "parse", msg: err}
}

func InvalidTokenError(err error) AppError {
	return AppError{typ: "invalid_token", msg: err}
}

func NonceError(err error) AppError {
	return AppError{typ: "nonce", msg: err}
}

func EnvironmentVariableError(err error) AppError {
	return AppError{typ: "environment_variable", msg: err}
}

func SignTokenError(err error) AppError {
	return AppError{typ: "sign_token", msg: err}
}

func ValidateTokenError(err error) AppError {
	return AppError{typ: "invalid_token", msg: err}
}
