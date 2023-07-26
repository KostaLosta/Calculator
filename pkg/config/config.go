package config

import (
	"io"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

const formatJSON = "json"

// ServiceConfig ..
type ServiceConfig struct {
	LogLevel     string `envconfig:"LOG_LEVEL" default:"trace"`
	LogFormat    string `envconfig:"LOG_FORMAT" default:"console"`
	ReportCaller bool   `envconfig:"LOG_REPORT_CALLER" default:"false"`

	// ACTS STORAGE
	DBEngine string `envconfig:"DB_ENGINE" required:"true" default:"mssql"`
	DBHost   string `envconfig:"DB_HOST" required:"true"`
	DBUser   string `envconfig:"DB_USER" required:"true"`
	DBPass   string `envconfig:"DB_PASS" required:"true"`
	DBPort   string `envconfig:"DB_PORT" required:"true"`
	DBName   string `envconfig:"DB_NAME" required:"true"`

	DBActsPoolSize int           `envconfig:"DB_ACTS_POOL_SIZE" default:"2"`
	DBConnTimeout  time.Duration `envconfig:"DB_CONN_TIMEOUT" default:"5m"`
}

var service *ServiceConfig

// Service ..
func Service() ServiceConfig {
	if service != nil {
		return *service
	}
	service = &ServiceConfig{}
	if err := envconfig.Process("", service); err != nil {
		panic(err)
	}
	return *service
}

// Logger ..
func (cfg ServiceConfig) Logger() (logger zerolog.Logger) {
	level := zerolog.InfoLevel
	if newLevel, err := zerolog.ParseLevel(cfg.LogLevel); err == nil {
		level = newLevel
	}
	var out io.Writer = os.Stdout
	if cfg.LogFormat != formatJSON {
		out = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.StampMicro}
	}
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	ctxLog := zerolog.New(out).Level(level).With().Timestamp().Stack()
	if cfg.ReportCaller {
		ctxLog = ctxLog.Caller()
	}
	return ctxLog.Logger()
}
