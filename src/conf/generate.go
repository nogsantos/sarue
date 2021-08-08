package conf

import (
	language "github.com/nogsantos/sarue/src/language/domain"
	platform "github.com/nogsantos/sarue/src/platform/domain"
)

type Generate struct {
	Platfom platform.Platfom
	Language language.Language
	Steps []string
	Template string
	InitAt string
	FinishAt string
}