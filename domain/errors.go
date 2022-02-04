package domain

import "errors"

var (
	ERROR_PARSE_ARGUMENTS        = errors.New("couldn't parse arguments")
	ERROR_PARSE_CONFIG           = errors.New("couldn't parse config")
	ERROR_CONVERT_CONFIG_TO_JSON = errors.New("couldn't conver config to JSON")
	ERROR_INIT_LOGGER            = errors.New("couldn't init logger")
	ERROR_PARSE_SCHEME_CONFIG    = errors.New("couldn't parse scheme config")
	UNKNOWN_NODE_TYPE            = errors.New("unknown node type")
)
