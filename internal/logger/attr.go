package logger

// ErrMsg returns error content string.
//
// Mainly used to prevent memory panic caused by nil error.
func ErrMsg(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
