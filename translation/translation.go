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
			Other: "Students",
		},
	})
	translated["MenuDisplay"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "MenuDisplay",
			Other: "Display",
		},
	})
	translated["MenuDelete"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "MenuDelete",
			Other: "Delete",
		},
	})
	translated["MenuEdit"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "MenuEdit",
			Other: "Edit",
		},
	})
	translated["RegisterNoun"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "RegisterNoun",
			Other: "Register",
		},
	})
	translated["Name"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Name",
			Other: "Name",
		},
	})
	translated["Surname"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Surname",
			Other: "Surname",
		},
	})
	translated["DOB"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "DOB",
			Other: "DOB",
		},
	})
	translated["Department"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Department",
			Other: "Department",
		},
	})
	translated["Sex"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Sex",
			Other: "Sex",
		},
	})
	translated["Male"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Male",
			Other: "Male",
		},
	})
	translated["Female"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Female",
			Other: "Female",
		},
	})
	translated["LoginNoun"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "LoginNoun",
			Other: "Login",
		},
	})
	translated["Password"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Password",
			Other: "Password",
		},
	})
	translated["ConfirmPassword"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "ConfirmPassword",
			Other: "Confirm password",
		},
	})
	translated["IncorrectEmailOrPassword"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "IncorrectEmailOrPassword",
			Other: "Incorrect email or password",
		},
	})
	translated["EmailIsTooShort"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "EmailIsTooShort",
			Other: "Email is too short",
		},
	})
	translated["PasswordIsTooShort"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "PasswordIsTooShort",
			Other: "Password is too short",
		},
	})
	translated["PasswordsDoNotMatch"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "PasswordsDoNotMatch",
			Other: "Passwords do not match",
		},
	})
	translated["BusyEmail"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "BusyEmail",
			Other: "This email is busy",
		},
	})
	translated["Modify"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Modify",
			Other: "Modify",
		},
	})
	translated["LogoutButton"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "LogoutButton",
			Other: "Logout",
		},
	})
	translated["Teachers"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Teachers",
			Other: "Teachers",
		},
	})
	translated["NO"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "NO",
			Other: "No",
		},
	})
	translated["Subject"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Subject",
			Other: "Subject.",
		},
	})
	translated["Email"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Email",
			Other: "Email",
		},
	})
	translated["LoginVerb"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "LoginVerb",
			Other: "Login",
		},
	})
	translated["ComingSoon"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "ComingSoon",
			Other: "Coming soon",
		},
	})
	translated["RegisterVerb"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "RegisterVerb",
			Other: "Register",
		},
	})
	translated["WaitingList"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "WaitingList",
			Other: "Waiting list",
		},
	})
	translated["EditStudent"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "EditStudent",
			Other: "Edit student",
		},
	})
	translated["Select"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Select",
			Other: "Wybierz",
		},
	})
	translated["Cancel"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Cancel",
			Other: "Cancel",
		},
	})
	translated["AddGrades"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "AddGrades",
			Other: "Add grades",
		},
	})
	translated["AddGrade"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "AddGrade",
			Other: "Add grade",
		},
	})
	translated["ViewGrades"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "ViewGrades",
			Other: "View grades",
		},
	})
	translated["Join"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Join",
			Other: "Join",
		},
	})
	translated["JoinToWaiting"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "JoinToWaiting",
			Other: "Join to witing list",
		},
	})
	translated["SendRequest"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "SendRequest",
			Other: "Send request",
		},
	})
	translated["AddOnWaiting"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "AddOnWaiting",
			Other: "You have been added on waiting list.",
		},
	})
	translated["SessionExpired"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "SessionExpired",
			Other: "Your session has expired.",
		},
	})
	translated["Average"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Average",
			Other: "Average:",
		},
	})
	translated["IncorrectGrade"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "IncorrectGrade",
			Other: "This grade is incorrect",
		},
	})
	translated["IncorrectEmail"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "IncorrectEmail",
			Other: "This email is incorrect",
		},
	})
	translated["AllFields"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "AllFields",
			Other: "Fill in all fields",
		},
	})
	translated["OnList"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "OnList",
			Other: "You are on list.",
		},
	})
	translated["ServerError"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "ServerError",
			Other: "Oh no, we have some problem.",
		},
	})

	return translated
}
