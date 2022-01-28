package configs

import "freecharge/light/api/helpers/logger"

func NewLogConfig(o *logger.Options) (*logger.Options, error) {
	o.JSONEncoding = true
	o.IncludeCallerSourceLocation = true
	o.LogGrpc = true
	return o, nil
}
