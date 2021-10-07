/*
Copyright 2021 The Sarue Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
  {{range .Stages.DefinedBy}}- {{ . }}
  {{ end -}}
{{ $blankLine := '\n' }}
{{ if Contains .Stages.DefinedBy "lint" -}}
lint:
  stage: lint
  image: {{ .Language.GitLabCiBuilder }}
  script:
    - python -m pip install flake8
    - flake8 . --ignore E203,E501,W503 --count --select=E9,F63,F7,F82 --show-source --exit-zero --max-complexity=10 --max-line-length=127 --statistics
  tags:
    - {{ .Language.Name | ToLower }}
{{ $blankLine := '\n' }}
{{ end -}}

{{ if Contains .Stages.DefinedBy "format" -}}
format:
  stage: format
  image: {{ .Language.GitLabCiBuilder }}
  script:
    - black --check .
  tags:
    - {{ .Language.Name | ToLower }}
{{ $blankLine := '\n' }}
{{ end -}}

{{ if Contains .Stages.DefinedBy "test" -}}
test:
  stage: test
  image: {{ .Language.GitLabCiBuilder }}
  script:
    - python -m pip install --upgrade pip
    - python -m pip install -r requirements.txt
    - python ./manage.py test --noinput --failfast -v 2
  tags:
    - {{ .Language.Name | ToLower }}
{{ end -}}
`)
}