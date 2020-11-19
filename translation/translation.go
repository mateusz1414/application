package translation

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func LoadBundles() *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	//bundle.MustLoadMessageFile()
	return bundle
}

func LoadTranslation(localizer *i18n.Localizer) map[string]string {
	translated := make(map[string]string)
	translated["H1Banner"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "H1Banner",
			Other: "My owner application",
		},
	})
	translated["MenuDisplay"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "MenuDisplay",
			Other: "Display",
		},
	})
	translated["MenuAdd"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "MenuAdd",
			Other: "Addition",
		},
	})
	return translated
}
