//SPDX-FileCopyrightText: © 2023 3nets, Inc. <it@3nets.io>
package utils

import (
	"regexp"

	validator "github.com/go-playground/validator/v10"
)

func matchRegex(regex *regexp.Regexp, fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return regex.MatchString(value)
}

// Functions to generate regular expressions. It could be placed per entity, in different folders.

func ResourceId(fl validator.FieldLevel) bool {
	regex := regexp.MustCompile(`^urn:[0-9a-fA-F]{24}:[0-9a-fA-F]{24}$`)
	return matchRegex(regex, fl)
}

func Name(fl validator.FieldLevel) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9-]{3,64}$`)
	return matchRegex(regex, fl)
}

func Description(fl validator.FieldLevel) bool {
	regex := regexp.MustCompile(`^.{5,512}$`)
	return matchRegex(regex, fl)
}
