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
