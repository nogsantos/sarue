package platform

func (gl *GitLab) GitlabTemplate() []byte {
	return []byte(`#===================================================================
# {{ .Tool }}
#
# {{ .Copyright }}
#
{{ if .Legal.Text }}{{ .Legal.Text }}{{ end }}
#===================================================================
stages:
  - test
  - build
  - deploy

`)
}