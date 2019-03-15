package restfool

import (
	log "github.com/sirupsen/logrus"
)

var usedLogger = log.StandardLogger()

// SetLogger replaces the standard logger with custom configuration
func SetLogger(newLogger *log.Logger) {
	usedLogger = newLogger
}

// Log is the generic logging function used by Debug, Error and Info
func Log(msg string, params map[string]interface{}, loglevel string) {
	switch loglevel {
	case INFO:
		usedLogger.WithFields(params).Info(msg)
	case DEBUG:
		usedLogger.WithFields(params).Debug(msg)
	case ERROR:
		usedLogger.WithFields(params).Error(msg)
	}
}

// Error logs out errors using fields
// e.g. Error("error msg", fmt.Errorf("my error"))
// or using log.Fields / map[string]interface type
// Error("error msg", map[string]interface{}{"val1": "foo", "val2": "bar"}
func Error(msg string, params interface{}) {
	switch params.(type) {
	case map[string]interface{}:
		fields := params.(map[string]interface{})
		if fields["err"] != nil {
			Log(msg, fields, ERROR)
		}
	case error:
		if params != nil {
			fields := map[string]interface{}{"msg": params}
			Log(msg, fields, ERROR)
		}
	}
}

// Debug uses the same syntax as Error function but does
// not support error type and does not check for errors
// e.g. Debug("debug", fmt.Errorf("my debug msg"))
// or using log.Fields / map[string]interface type
// Debug("debug msg", map[string]interface{}{"val1": "foo", "val2": "bar"}
func Debug(msg string, params interface{}) {
	switch params.(type) {
	case map[string]interface{}:
		fields := params.(map[string]interface{})
		Log(msg, fields, DEBUG)
	case string, error:
		fields := map[string]interface{}{"msg": params}
		Log(msg, fields, DEBUG)
	}
}

// Info uses the same syntax as Error function but does
// not support error type and does not check for errors
// e.g. Info("info", fmt.Errorf("my info msg"))
// or using log.Fields / map[string]interface type
// Info("info msg", map[string]interface{}{
//	"err": "foo",
//	"someField": "bar"
// }
func Info(msg string, params interface{}) {
	switch params.(type) {
	case map[string]interface{}:
		fields := params.(map[string]interface{})
		Log(msg, fields, INFO)
	case string, error:
		fields := map[string]interface{}{"msg": params}
		Log(msg, fields, INFO)
	}
}

// ErrorMsg used for simple error logging without fields
func ErrorMsg(msg string) {
	usedLogger.Error(msg)
}

// DebugMsg used for simple debug logging without fields
func DebugMsg(msg string) {
	usedLogger.Debug(msg)
}

// InfoMsg used for simple info logging without fields
func InfoMsg(msg string) {
	usedLogger.Info(msg)
}
