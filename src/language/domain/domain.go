package domain


type LanguageType int

const (
	PYTHON LanguageType = iota
	NODEJS
)

func (p LanguageType) String() string {
	return [...]string{"Python", "Nodejs"}[p]
}

var Platfoms = make(map[string]LanguageType)

func Types() []string {
	return []string{PYTHON.String(), NODEJS.String()}
}

type PackageManager struct {
	Packages []string
}

type Framework struct {
	Frameworks []string
}

type Language struct {
	Name string
	PackageManager *PackageManager
	Framework *Framework
}