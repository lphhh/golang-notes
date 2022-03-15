package log

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mattn/go-isatty"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"wasabi_data_acquisition_unit/configs"
)

var (
	green  = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	yellow = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	red    = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue   = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	reset  = string([]byte{27, 91, 48, 109})
)

// LogPushEntry is push response log
type LogPushEntry struct {
	Type     string `json:"type"`
	Platform string `json:"platform"`
	Message  string `json:"message"`
	Error    string `json:"error"`
}

var isTerm bool

func init() {
	isTerm = isatty.IsTerminal(os.Stdout.Fd())
	_ = godotenv.Load(".env")
	if err := InitLog(
		os.Getenv("log_access_level"),
		os.Getenv("log_access_log"),
		os.Getenv("log_error_level"),
		os.Getenv("log_error_log"),
	); err != nil {
		log.Fatalf("can't load log module, error: %v", err)
	}
}

var (
	// LogAccess is log server request log
	LogAccess = logrus.New()
	// LogError is log server error log
	LogError = logrus.New()
)

// InitLog use for initial log module
func InitLog(accessLevel, accessLog, errorLevel, errorLog string) error {
	var err error

	if !isTerm {
		LogAccess.SetFormatter(&logrus.JSONFormatter{})
		LogError.SetFormatter(&logrus.JSONFormatter{})
	} else {
		LogAccess.Formatter = &logrus.TextFormatter{
			TimestampFormat: "2006/01/02 - 15:04:05",
			FullTimestamp:   true,
		}

		LogError.Formatter = &logrus.TextFormatter{
			TimestampFormat: "2006/01/02 - 15:04:05",
			FullTimestamp:   true,
		}
	}

	// set logger
	if err = SetLogLevel(LogAccess, accessLevel); err != nil {
		return errors.New("Set access log level error: " + err.Error())
	}

	if err = SetLogLevel(LogError, errorLevel); err != nil {
		return errors.New("Set error log level error: " + err.Error())
	}

	if err = SetLogOut(LogAccess, accessLog); err != nil {
		return errors.New("Set access log path error: " + err.Error())
	}

	if err = SetLogOut(LogError, errorLog); err != nil {
		return errors.New("Set error log path error: " + err.Error())
	}

	return nil
}

// SetLogOut provide log stdout and stderr output
func SetLogOut(log *logrus.Logger, outString string) error {
	switch outString {
	case "stdout":
		log.Out = os.Stdout
	case "stderr":
		log.Out = os.Stderr
	default:
		f, err := os.OpenFile(outString, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o600)
		if err != nil {
			return err
		}

		log.Out = f
	}

	return nil
}

// SetLogLevel is define log level what you want
// log level: panic, fatal, error, warn, info and debug
func SetLogLevel(log *logrus.Logger, levelString string) error {
	level, err := logrus.ParseLevel(levelString)
	if err != nil {
		return err
	}

	log.Level = level

	return nil
}

func colorForPlatForm(platform int) string {
	switch platform {
	case configs.Binance:
		return yellow
	case configs.Ftx:
		return blue
	default:
		return reset
	}
}

func typeForPlatForm(platform int) string {
	return configs.GetPlatform(platform)
}

// InputLog log request
type InputLog struct {
	Status   string
	Message  string
	Platform int
	Error    error
}

// GetLogPushEntry get push data into log structure
func GetLogPushEntry(input *InputLog) LogPushEntry {
	var errMsg string

	plat := typeForPlatForm(input.Platform)

	if input.Error != nil {
		errMsg = input.Error.Error()
	}

	return LogPushEntry{
		Type:     input.Status,
		Platform: plat,
		Message:  input.Message,
		Error:    errMsg,
	}
}

// LogPush record user push request and server response.
func LogPush(input *InputLog) LogPushEntry {
	var platColor, resetColor, output string

	if isTerm {
		platColor = colorForPlatForm(input.Platform)
		resetColor = reset
	}

	log := GetLogPushEntry(input)

	if os.Getenv("log_format") == "json" {
		logJSON, _ := json.Marshal(log)

		output = string(logJSON)
	} else {
		var typeColor string
		switch input.Status {
		case configs.Succeeded:
			if isTerm {
				typeColor = green
			}

			output = fmt.Sprintf("|%s %s %s| %s%s%s %s",
				typeColor, log.Type, resetColor,
				platColor, log.Platform, resetColor,
				log.Message,
			)
		case configs.Failed:
			if isTerm {
				typeColor = red
			}

			output = fmt.Sprintf("|%s %s %s| %s%s%s | %s | Error Message: %s",
				typeColor, log.Type, resetColor,
				platColor, log.Platform, resetColor,
				log.Message,
				log.Error,
			)
		}
	}

	switch input.Status {
	case configs.Succeeded:
		LogAccess.Info(output)
	case configs.Failed:
		LogError.Error(output)
	}

	return log
}
