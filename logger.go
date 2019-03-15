package restfool

import (
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func (a *RestAPI) initLogger() {
	switch a.Conf.Logging.Type {
	case LOGFORMATJSON:
		usedLogger.SetFormatter(&log.JSONFormatter{})
	case LOGFORMATTEXT:
		formatter := &log.TextFormatter{
			FullTimestamp: true,
		}
		usedLogger.SetFormatter(formatter)
	default:
		usedLogger.WithFields(log.Fields{
			"logformat": a.Conf.Logging.Type,
			"default":   LOGFORMATTEXT,
		}).Error("unknown logformat using default")
	}

	switch a.Conf.Logging.Loglevel {
	case INFO:
		usedLogger.SetLevel(log.InfoLevel)
	case ERROR:
		usedLogger.SetLevel(log.ErrorLevel)
	case DEBUG:
		usedLogger.SetLevel(log.DebugLevel)
	default:
		usedLogger.WithFields(log.Fields{
			"loglevel": a.Conf.Logging.Loglevel,
			"default":  INFO,
		}).Error("unknown loglevel using default")
		usedLogger.SetLevel(log.InfoLevel)
	}

	switch a.Conf.Logging.Output {
	case LOGFILE:
		logfile, err := os.OpenFile(a.Conf.Logging.Logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			usedLogger.WithFields(log.Fields{
				"filepath": a.Conf.Logging.Logfile,
			}).Error("can't open logfile use stdout")
			a.Conf.Logging.Output = LOGSTDOUT
		}
		usedLogger.SetOutput(logfile)
		usedLogger.WithFields(log.Fields{
			"output": LOGFILE,
			"format": a.Conf.Logging.Type,
		}).Debug("initialising logging")
	case LOGSTDOUT:
		usedLogger.WithFields(log.Fields{
			"output": LOGSTDOUT,
			"format": a.Conf.Logging.Type,
		}).Debug("using logging method")
	default:
		usedLogger.WithFields(log.Fields{
			"output":  a.Conf.Logging.Output,
			"default": LOGSTDOUT,
		}).Error("unknown log output using default")
		a.Conf.Logging.Output = LOGSTDOUT
	}
}

// Logger is the default logging handler
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		usedLogger.WithFields(log.Fields{
			"method":      r.Method,
			"request-uri": r.RequestURI,
			"duration":    time.Since(start),
			"name":        name,
			"ip":          r.RemoteAddr,
		}).Info("REQUEST")
	})
}
