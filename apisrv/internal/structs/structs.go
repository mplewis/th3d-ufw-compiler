package structs

// CompileRequest represents the project configuration used to build custom firmware.
type CompileRequest struct {
	ConfigHeader string `json:"config_header"`
	PioConfig    string `json:"pio_config"`
}
