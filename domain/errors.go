package domain

import "errors"

var (
	ErrParseArguments         = errors.New("couldn't parse arguments")
	ErrParseConfig            = errors.New("couldn't parse config")
	ErrConvertConfigToJson    = errors.New("couldn't conver config to JSON")
	ErrInitLogger             = errors.New("couldn't init logger")
	ErrParseSchemeConfig      = errors.New("couldn't parse scheme config")
	ErrUnknownNodeId          = errors.New("unknown node id")
	ErrUnknownNodeType        = errors.New("unknown node type")
	ErrUnknownActionDirection = errors.New("unknown action direction")
	ErrUnknownModel           = errors.New("unknown model")
	ErrUnknownModelAction     = errors.New("unknown model's action")
)
