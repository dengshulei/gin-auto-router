package gin_auto_router

import (
	"reflect"
	"strings"
	"unicode"
)

// ToKebabCase converts a string to kebab-case format
// Examples:
//
//	ListGet   → list-get
//	infoPush  → info-push
//	User_Info → user-info
//	UserAPI   → user-api
//	HTTPModel → http-model
func ToKebabCase(s string) string {
	if s == "" {
		return ""
	}

	var buf strings.Builder
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n; i++ {
		c := runes[i]

		// 1. Skip leading non-alphabetic characters (underscore, dash, space)
		if i == 0 {
			if c == '_' || c == '-' || c == ' ' {
				continue
			}
			buf.WriteRune(unicode.ToLower(c))
			continue
		}

		// 2. Skip useless symbols
		if c == '_' || c == '-' || c == ' ' {
			// Avoid consecutive --
			if buf.Len() > 0 && buf.String()[buf.Len()-1] != '-' {
				buf.WriteRune('-')
			}
			continue
		}

		// 3. Handle uppercase letters (core: avoid consecutive dashes)
		if unicode.IsUpper(c) {
			last := runes[i-1]
			if !unicode.IsUpper(last) && last != '-' && last != '_' {
				buf.WriteRune('-')
			}
			buf.WriteRune(unicode.ToLower(c))
			continue
		}

		// 4. Normal characters
		buf.WriteRune(c)
	}

	return strings.Trim(buf.String(), "-")
}

// KebabCaseToSnakeCase converts kebab-case to snake_case
// Examples: list-get => list_get; info-push => info_push; code => code; setting => setting
func KebabCaseToSnakeCase(s string) string {
	return strings.ReplaceAll(s, "-", "_")
}

// KebabCaseToCamelCase converts kebab-case to camelCase
// Examples: list-get => listGet; info-push => infoPush; code => code; setting => setting
func KebabCaseToCamelCase(s string) string {
	return KebabCaseToOther(s, false)
}

// KebabCaseToPascalCase converts kebab-case to PascalCase
// Examples: list-get => ListGet; info-push => InfoPush; code => Code; setting => Setting
func KebabCaseToPascalCase(s string) string {
	return KebabCaseToOther(s, true)
}

// KebabCaseToOther converts kebab-case to camelCase or PascalCase
// s: string to convert
// firstUpper: true for PascalCase, false for camelCase
func KebabCaseToOther(s string, firstUpper bool) string {
	var output []rune

	for _, r := range s {
		// When encountering dash, mark next character to be uppercase, skip current
		if r == '-' {
			firstUpper = true
			continue
		}

		// If needs uppercase, convert and reset flag
		if firstUpper {
			output = append(output, unicode.ToUpper(r))
			firstUpper = false
		} else {
			output = append(output, r)
		}
	}

	return string(output)
}

// NormalizeNamingConvention gets the naming convention, supports: kebab-case (default), snake_case, camelCase, PascalCase
func NormalizeNamingConvention(args ...string) (nom string) {
	nom = GetOptionalArg(0, args...)
	switch nom {
	case "kebab-case":
	case "snake_case":
	case "camelCase":
	case "PascalCase":
	default:
		nom = "kebab-case"
	}
	return nom
}

// GetOptionalArg gets the specified arg from a batch of args
// i: the index of the arg to get, starting from 0
func GetOptionalArg(i int, args ...string) (arg string) {
	if len(args) > i {
		arg = args[i]
	}
	return
}

// GetControllerName gets the controller struct name
// Example:
//
//	*controller.Article => Article
func GetControllerName(controller interface{}) (ControllerName string) {
	// Get the controller type (reflect.Type), result similar to "*controller.Article"
	ControllerName = reflect.TypeOf(controller).String()
	if strings.Contains(ControllerName, ".") {
		// At this point the result is similar to "Article"
		ControllerName = ControllerName[strings.Index(ControllerName, ".")+1:]
	}
	return
}
