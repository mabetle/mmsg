package mmsg

import (
	"github.com/mabetle/mlog"
	"github.com/robfig/config"
)

const (
	messageFilePattern = `^\w+\.[a-zA-Z]{2}$`
	unknownValueFormat = "[%s]"
)

var (
	logger = mlog.GetLogger("github.com/mabetle/mmsg")

	// default message directory
	MessageFilesDirectory = "messages"

	defaultLanguage = "en"

	// DefaultLocale run time locale
	DefaultLocale = "en-US"

	// All currently loaded message configs.
	messages map[string]*config.Config = make(map[string]*config.Config)
)
