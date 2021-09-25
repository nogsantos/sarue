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

func (gb *Github) GithubLintTemplate() []byte {
	return []byte(`#===================================================================
# {{ .Tool }}
#
# {{ .Copyright }}
#
{{ if .Legal.Text }}{{ .Legal.Text }}{{ end }}
#===================================================================
name: Lint {{ .Language.Name }}

on:
  push:
    branches:
      - '*'
      - '*/*'
      - '**'
      - '!master'

  pull_request:
    paths:{{ range .Language.Extension }}
      - '**.{{ . }}'{{ end }}

jobs:
  lint:
    runs-on: ubuntu-latest
    steps:
      - name: Check out repository code
        uses: actions/checkout@v2

      - name: Set up {{ .Language.Name }}
        uses: {{ .Language.GitHubCiBuilder }}
        with:
          {{ .Language.Name | ToLower }}-version: {{ .Language.Version }}

      - name: Lint Execution
        run: |{{range .LintCommands}}{{ . }}
          {{ end }}
`)
}

func (gb *Github) GithubTestTemplate() []byte {
	return []byte(`#===================================================================
# {{ .Tool }}
#
# {{ .Copyright }}
#
{{ if .Legal.Text }}{{ .Legal.Text }}{{ end }}
#===================================================================
name: Test {{ .Language.Name }}

on:
  push:
    branches:
      - '*'
      - '*/*'
      - '**'
      - '!master'

  pull_request:
    paths:{{ range .Language.Extension }}
      - '**.{{ . }}'{{ end }}

jobs:
  tests:
    runs-on: ubuntu-latest
    steps:
      - name: Check out repository code
        uses: actions/checkout@v2

      - name: Set up {{ .Language.Name }}
        uses: {{ .Language.GitHubCiBuilder }}
        with:
          {{ .Language.Name | ToLower }}-version: {{ .Language.Version }}

      - name: Test Execution
        run: |{{range .TestCommands}}{{ . }}
          {{ end }}
`)
}

func (gb *Github) GithubFormatTemplate() []byte {
	return []byte(`#===================================================================
# {{ .Tool }}
#
# {{ .Copyright }}
#
{{ if .Legal.Text }}{{ .Legal.Text }}{{ end }}
#===================================================================
name: Format {{ .Language.Name }}

on:
  push:
    branches:
      - '*'
      - '*/*'
      - '**'
      - '!master'

  pull_request:
    paths:{{ range .Language.Extension }}
      - '**.{{ . }}'{{ end }}

jobs:
  format:
    runs-on: ubuntu-latest
    steps:
      - name: Check out repository code
        uses: actions/checkout@v2

      - name: Set up {{ .Language.Name }}
        uses: {{ .Language.GitHubCiBuilder }}
        with:
          {{ .Language.Name | ToLower }}-version: {{ .Language.Version }}

      - name: Format Execution
        run: |{{range .FormatCommands}}{{ . }}
          {{ end }}
`)
}
