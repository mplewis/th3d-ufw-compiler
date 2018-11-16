package structs

// CompileRequest represents the project configuration used to build custom firmware.
type CompileRequest struct {
	ID           string
	ConfigHeader string `json:"config_header"`
	PioConfig    string `json:"pio_config"`
}

// CompileResult represents the results of a single compilation.
type CompileResult struct {
	Request  CompileRequest
	IntelHex string
	Error    error
}

// NewCompileSuccess creates a CompileResult for a successful compilation
func NewCompileSuccess(req CompileRequest, hex string) CompileResult {
	return CompileResult{Request: req, IntelHex: hex}
}

// NewCompileFailure creates a CompileResult for a failed compilation
func NewCompileFailure(req CompileRequest, e error) CompileResult {
	return CompileResult{Request: req, Error: e}
}
