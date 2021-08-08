package gitlab

func MainTemplate() []byte {
	return []byte(`/*
{{ .Copyright }}
{{ if .Legal.Header }}{{ .Legal.Header }}{{ end }}
*/
stages:
	- test
    - build
    - deploy

`)
}