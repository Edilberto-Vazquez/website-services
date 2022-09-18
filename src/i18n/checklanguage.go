package i18n

import (
	"strings"

	"github.com/Edilberto-Vazquez/website-services/src/constants"
)

func CheckLanguage(lang string) (value string, exists bool) {
	langToLower := strings.ToLower(lang)
	value, exists = constants.ACCEPTED_LANGUAGES[langToLower]
	return value, exists
}
