package domain

type PlatfomType int

const (
	GITHUB PlatfomType = iota
	GITLAB
)

func (p PlatfomType) String() string {
	return [...]string{"Github", "Gitlab"}[p]
}

var Platfoms = make(map[string]Platfom)

func Types() []string {
	return []string{GITHUB.String(), GITLAB.String()}
}

type Platfom struct {
	Name PlatfomType
	ConfigPath string
	ConfigFile string
	Template string
}