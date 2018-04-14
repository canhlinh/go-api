package utils

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/nicksnyder/go-i18n/i18n"
)

var T i18n.TranslateFunc
var locales map[string]string = make(map[string]string)

func Init(i18nPath string) {
	InitTranslationsWithDir(i18nPath)
	T = TfuncWithFallback("en")
}

func InitTranslationsWithDir(i18nDirectory string) {
	files, _ := ioutil.ReadDir(i18nDirectory)
	for _, f := range files {
		if filepath.Ext(f.Name()) == ".json" {
			filename := f.Name()
			locales[strings.Split(filename, ".")[0]] = i18nDirectory + "/" + filename
			i18n.MustLoadTranslationFile(i18nDirectory + "/" + filename)
		}
	}
}

func TfuncWithFallback(pref string) i18n.TranslateFunc {
	t, _ := i18n.Tfunc(pref)
	return func(translationID string, args ...interface{}) string {
		if translated := t(translationID, args...); translated != translationID {
			return translated
		}

		t, _ := i18n.Tfunc("en")
		return t(translationID, args...)
	}
}
