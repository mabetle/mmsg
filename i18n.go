package mmsg

// modify from revel
import (
	"fmt"
	"github.com/mabetle/mcore"
	"github.com/mabetle/mcore/mconf/wrobfig"
	"github.com/robfig/config"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// check if key exists.
func Contains(locale string, key string) bool {
	language, region := parseLocale(locale)
	messageConfig, knownLanguage := messages[language]
	if !knownLanguage {
		return false
	}
	_, err := messageConfig.String(region, key)
	if err != nil {
		return false
	}
	return true
}

// return value and if contain
func MessageContains(locale string, key string, args ...interface{}) (string, bool) {
	language, region := parseLocale(locale)
	messageConfig, knownLanguage := messages[language]

	keyLabel := mcore.ToLabel(key)

	if !knownLanguage {
		return keyLabel, false
	}
	// check message.
	if v, err := messageConfig.String(region, key); err != nil {
		return keyLabel, false
	} else {
		return v, true
	}
}

func UpdateMsg(locale, key, value string) {
	_, region := parseLocale(locale)
	c := config.NewDefault()
	b := c.AddOption(region, key, value)
	if !b {
		logger.Infof("Put Msg faild. Key: %s Value: %s", key, value)
	}
	//
	PutConfig(locale, c)
}

// PutMsg
func PutMsg(locale, key, value string) {
	if Contains(locale, key) {
		logger.Tracef("Key Exists, skip put. Key: %s, Locale: %s", key, locale)
		return
	}
	if key == "" || locale == "" {
		return
	}
	UpdateMsg(locale, key, value)
}

func PutConfig(locale string, c *config.Config) {
	lang, _ := parseLocale(locale)
	if _, exists := messages[lang]; exists {
		messages[lang].Merge(c)
	} else {
		messages[lang] = c
	}
}

func PutMsgText(locale, text string) {
	c, err := wrobfig.ReadConfigFromString(text)
	if err != nil {
		logger.Warn("put msg text error: ", err)
	}
	PutConfig(locale, c)
}

func SetDefaultLanguage(v string) {
	defaultLanguage = v
}

// Return all currently loaded message languages.
func MessageLanguages() []string {
	languages := make([]string, len(messages))
	i := 0
	for language, _ := range messages {
		languages[i] = language
		i++
	}
	return languages
}

func LocaleMessage(key string, args ...interface{}) string {
	locale := GetLocale()
	return Message(locale, key, args...)
}

func GetLocale() string {
	return DefaultLocale
}

// Perform a message look-up for the given locale and message using the given arguments.
// When either an unknown locale or message is detected, a specially formatted string is returned.
func Message(locale, message string, args ...interface{}) string {
	language, region := parseLocale(locale)

	messageConfig, knownLanguage := messages[language]
	if !knownLanguage {
		logger.Warnf("Unsupported language for locale '%s' and message '%s', trying default language", locale, message)
		messageConfig, knownLanguage = messages[defaultLanguage]
		if !knownLanguage {
			logger.Warnf("Unsupported default language for locale '%s' and message '%s'", defaultLanguage, message)
			return fmt.Sprintf(unknownValueFormat, message)
		}
	}

	// This works because unlike the mconfig documentation suggests it will actually
	// try to resolve message in DEFAULT if it did not find it in the given section.
	if value, err := messageConfig.String(region, message); err == nil {
		return doMsgArgs(value, args...)
	}
	// try default region
	if value, err := messageConfig.String("", message); err == nil {
		return doMsgArgs(value, args...)
	}
	// cannot find
	logger.Debugf("Cannot found message. Locale:%s, Key:%s", locale, message)
	return mcore.ToLabel(message)
}

func doMsgArgs(value string, args ...interface{}) string {
	if len(args) > 0 {
		logger.Infof("Arguments detected, formatting '%s' with %v", value, args)
		value = fmt.Sprintf(value, args...)
	}
	return value
}

func parseLocale(locale string) (language, region string) {
	if strings.Contains(locale, "-") {
		languageAndRegion := strings.Split(locale, "-")
		return languageAndRegion[0], languageAndRegion[1]
	}
	return locale, ""
}

// LoadDefaultMessages
// Default work dir messages
func LoadDefaultMessages() {
	LoadMessages("messages")
}

// Recursively read and cache all available messages from all message files on the given path.
func LoadMessages(path string) {
	if error := filepath.Walk(path, LoadMessageFile); error != nil && !os.IsNotExist(error) {
		logger.Warn("Reading messages files error:", error)
	}
}

// LoadPrefixMessages know path and prefix, not all dir files.
func LoadPrefixMessages(path string, prefix string) {
	// not found dir
	var fs []os.FileInfo
	var err error
	fs, err = ioutil.ReadDir(path)
	// path not found
	if logger.CheckError(err) {
		return
	}
	for _, f := range fs {
		// skip sub dir
		if f.IsDir() {
			continue
		}
		// load messages
		if strings.HasPrefix(f.Name(), prefix) {
			fpath := fmt.Sprintf("%s/%s", path, f.Name())
			if err := LoadMessageFile(fpath, f, nil); err != nil {
				logger.Warnf("Load messages file error, path: %s, file name: %s", fpath, f.Name())
			}
		}
	}
}

func FprintMessages(locale string, w io.Writer) {
	if locale == "" {
		logger.Warn("Locale should not be blank")
		return
	}
	language, _ := parseLocale(locale)
	if c, ok := messages[language]; ok {
		for _, s := range c.Sections() {
			fmt.Fprintf(w, "[%s]\n", s)
			opts, _ := c.Options(s)
			for _, opt := range opts {
				optValue, _ := c.RawString(s, opt)
				fmt.Fprintf(w, "%s: %s\n", opt, optValue)
			}
		}
	}

	logger.Infof("Not found messages for locale: %s", locale)
}

// SaveMessageFile
func SaveMessageFile(path string, prefix string) {
	for locale, c := range messages {
		if !strings.HasSuffix(path, "/") {
			path = path + "/"
		}
		fn := fmt.Sprintf("%s%s.%s", path, prefix, locale)
		err := c.WriteFile(fn, 0666, "")
		logger.CheckError(err)
	}
}

// Load a single message file
func LoadMessageFile(path string, info os.FileInfo, osError error) error {
	if osError != nil {
		return osError
	}
	if info.IsDir() {
		return nil
	}
	//if matched, _ := regexp.MatchString(messageFilePattern, info.Name()); matched {
	if config, err := parseMessagesFile(path); err != nil {
		logger.Error("Error parse Message file:", path, "Error:", err)
		return err
	} else {
		locale := parseLocaleFromFileName(info.Name())
		PutConfig(locale, config)
		logger.Trace("Successfully loaded messages from file: ", path)
	}
	//} else {
	//logger.Warnf("Ignoring file %s because it did not have a valid extension", path)
	//}
	return nil
}

// LoadMatchMessageFile
func LoadMatchMessageFile(path string, info os.FileInfo, osError error) error {
	if matched, _ := regexp.MatchString(messageFilePattern, info.Name()); matched {
		return LoadMessageFile(path, info, osError)
	} else {
		logger.Warnf("Ignoring file %s because it did not have a valid extension", path)
	}
	return nil
}

func parseMessagesFile(path string) (messageConfig *config.Config, error error) {
	messageConfig, error = config.ReadDefault(path)
	return
}

func parseLocaleFromFileName(file string) string {
	extension := filepath.Ext(file)[1:]
	return strings.ToLower(extension)
}
