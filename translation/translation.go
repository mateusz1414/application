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
	translated["MenuRegister"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "MenuRegister",
			Other: "Register",
		},
	})
	translated["MenuPermission"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "MenuPermission",
			Other: "Permission",
		},
	})
	translated["MenuGrades"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "MenuGrades",
			Other: "Grades",
		},
	})
	translated["MenuTeachers"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "MenuTeachers",
			Other: "Teachers",
		},
	})
	translated["MenuDepartaments"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "MenuDepartaments",
			Other: "Departaments",
		},
	})
	translated["Display#"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Display#",
			Other: "#",
		},
	})
	translated["DisplayName"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "DisplayName",
			Other: "Name",
		},
	})
	translated["DisplayLastName"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "DisplayLastName",
			Other: "Last name",
		},
	})
	translated["DisplayDateOfBirth"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "DisplayDateOfBirth",
			Other: "Date of birth",
		},
	})
	translated["DisplayDepartment"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "DisplayDepartment",
			Other: "Department",
		},
	})
	translated["DisplayGender"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "DisplayGender",
			Other: "Gender",
		},
	})
	translated["DisplaySelectMen"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "DisplaySelectMen",
			Other: "Men",
		},
	})
	translated["DisplaySelectWomen"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "DisplaySelectWomen",
			Other: "Women",
		},
	})
	translated["DisplayAdd"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "DisplayAdd",
			Other: "Add",
		},
	})
	translated["DisplayDelete"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "DisplayDelete",
			Other: "Delete",
		},
	})
	translated["DisplayEdit"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "DisplayEdit",
			Other: "Edit",
		},
	})
	translated["DisplayLogin"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "DisplayLogin",
			Other: "Login",
		},
	})
	translated["DisplayPassword"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "DisplayPassword",
			Other: "Password",
		},
	})
	translated["DisplayConfirmPassword"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "DisplayPassword",
			Other: "Password",
		},
	})
	translated["DisplayLoginIn"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "DisplayLoginIn",
			Other: "Login in",
		},
	})
	translated["DisplayRegister"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "DisplayRegister",
			Other: "Register",
		},
	})
	translated["DisplayLogged"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "DisplayLogged",
			Other: "Logged",
		},
	})
	translated["DisplayFooter"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "DisplayFooter",
			Other: "Footer",
		},
	})
	translated["MustLoginInError"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "MustLoginInError",
			Other: "Must login in",
		},
	})
	translated["LoginFailed"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "LoginFailed",
			Other: "Login feiled",
		},
	})
	translated["IncorrectLoginOrPassword"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "IncorrectLoginOrPassword",
			Other: "Incorrect login or password",
		},
	})
	translated["RegistrationFailed"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "RegistrationFailed",
			Other: "Registration failed",
		},
	})
	translated["SuccessfullyRegistered"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "SuccessfullyRegistered",
			Other: "Successfully registered",
		},
	})
	translated["LoginIsTooShort"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "LoginIsTooShort",
			Other: "Login is too short",
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
	translated["ThisUserAlreadyExists"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "ThisUserAlreadyExists",
			Other: "This user already exists",
		},
	})
	translated["DisplayChange"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "DisplayChange",
			Other: "Change",
		},
	})
	translated["AddError"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "AddError",
			Other: "Student cannot be added",
		},
	})
	translated["EditError"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "EditError",
			Other: "Editing has failed",
		},
	})
	translated["LogoutButton"] = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "LogoutButton",
			Other: "Logout",
		},
	})
	return translated
}
