// i18n package.
// some code copy form revel.
// dependency github.com/robfig/config
// usage:
// Load messages from dir:
// load message first.
// mmsg.Load("xxx dir")
// mmsg.SetLocale(locale)
// mmsg.LocaleMessage(locale, key)
// 
// 
// when we donot want to create messages file,
// we can put messages by code.
// it's usefull when the application is a single one command in go Language.
// Provide one executable command is enough, no extra language resources files. 
// Put messages:
// mmsg.PutMsg(locale, key, value)
package mmsg
