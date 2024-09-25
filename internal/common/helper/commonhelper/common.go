package commonhelper

type Environment string

const (
	ENV__DEV Environment = "dev"
	ENV__PRD Environment = "prd"
)

type LogLevel string

const (
	LOG_LEVEL__WARN  LogLevel = "warn"
	LOG_LEVEL__DEBUG LogLevel = "debug"
	LOG_LEVEL__INFO  LogLevel = "info"
)

type APIResponseStatus string

const (
	API_RESP_STATUS__PROCESSING APIResponseStatus = "PROCESSING"
	API_RESP_STATUS__ACCEPT     APIResponseStatus = "ACCEPT"
	API_RESP_STATUS__REJECT     APIResponseStatus = "REJECT"
	API_RESP_STATUS__FAILED     APIResponseStatus = "FAILED"
	API_RESP_STATUS__DONE       APIResponseStatus = "DONE"
)
