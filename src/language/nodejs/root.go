package nodejs

import "fmt"

type packageManager int
func (p packageManager) String() string {
	return [...]string{"npm", "yarn"}[p]
}

type Framework int
func (f Framework) String() string {
	return [...]string{"React", "Vue"}[f]
}

type NodeJs struct {
	Manager packageManager
	Framework Framework
}

func Init() {
	fmt.Print("NodeJS initialized")
}