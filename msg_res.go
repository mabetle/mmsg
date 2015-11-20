package mmsg

type MsgRes struct {
	Key  string
	Lang string
	Msg  string
}

type LangMsg struct {
	Lang string
	Msg  string
}

type KeyMsg struct {
	Key      string
	LangMsgs []LangMsg
}

func NewKeyMsg(key string) *KeyMsg {
	return &KeyMsg{Key: key}
}

func (km *KeyMsg) PutLangMsg(lang, msg string) *KeyMsg {
	km.LangMsgs = append(km.LangMsgs, LangMsg{Lang: lang, Msg: msg})
	return km
}

func (km *KeyMsg) Zh(msg string) *KeyMsg {
	return km.PutLangMsg("zh", msg)
}

func (km *KeyMsg) En(msg string) *KeyMsg {
	return km.PutLangMsg("en", msg)
}

func (km *KeyMsg) Put() {
	if km.Key == "" {
		return
	}
	for _, lm := range km.LangMsgs {
		PutMsg(lm.Lang, km.Key, lm.Msg)
	}
}

func PutKeyMsgs(kms []*KeyMsg) {
	for _, km := range kms {
		km.Put()
	}
}

func PutMsgRes(args ...MsgRes) {
	for _, res := range args {
		if res.Key == "" || res.Lang == "" {
			continue
		}
		PutMsg(res.Lang, res.Key, res.Msg)
	}
}

func PutMsgEnZh(key, en, zh string) {
	PutMsg("en", key, en)
	PutMsg("zh", key, zh)
}
