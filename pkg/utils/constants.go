package utils

import "time"

const (
	ErrorInputRequired                   = "error missing required input"
	ErrorInputFail                       = "error input failed"
	ErrorEmailFail                       = "error email failed"
	ErrorPasswordFail                    = "error password failed"
	ErrorCheckMaxLengthUnder50Characters = "error check max length under 50 characters"
	ErrorInputCharacterLimit             = "error input character limit"
	ErrorInputByteLimit                  = "error input byte limit"
	ErrorInvalidDomain                   = "error invalid domain"

	DefaultPageNo   = 1
	DefaultPageSize = 30

	ExecutableFilePermission = 0o750

	FormatYearISO     = "2006"
	FormatDateISO     = time.DateOnly
	FormatTimeHHMM    = "15:04"
	FormatTimeHHMMSS  = time.TimeOnly
	FormatDateTimeISO = time.RFC3339
	FormatDateTimeSQL = time.DateTime
	FormatDateCompact = "20060102150405"

	ProdEnv    = "prod"
	DevEnv     = "dev"
	TestingEnv = "testing"

	ErrorInternalServer    = "Internal server error"
	ErrorLogRequestBody    = "failed to log request body"
	ErrorCloseResponseBody = "failed to close response body"
	ErrorCloseRows         = "failed to close Rows"
	ErrorCloseReader       = "failed to close reader"
	ErrorCloseWriter       = "failed to close writer"
	ErrorCloseFile         = "failed to close file"
	ErrorCloseSftp         = "failed to close sftp"
	ErrorParseUrl          = "failed to parse url"
	ErrorReadBody          = "failed to read body"
	ErrorDecodeBody        = "failed to decode body"
	ErrorMapToStruct       = "failed to map to struct"
	ErrorGetTx             = "failed to get Tx"
	ErrorGetSpec           = "failed to get specification"
	ErrorCreateReq         = "failed to create new HTTP request"
	ErrorMarshalFailed     = "failed to marshal object to JSON"
)
