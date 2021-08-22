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

      - name: Install dependencies
        run: |
          python -m pip install --upgrade pipenv wheel

      - name: Execute Linter
        run: |
          python -m pip install flake8
          flake8

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

      - name: Install dependencies
        run: |
          python -m pip install --upgrade pipenv wheel

      - name: Execute Tests
        run: |
          python -m pip install flake8
          flake8

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
  pull_request:
    paths:{{ range .Language.Extension }}
      - '**.{{ . }}'{{ end }}

jobs:
  format:
    runs-on: ubuntu-latest
      - name: Check out repository code
        uses: actions/checkout@v2

      - name: Set up {{ .Language.Name }}
        uses: {{ .Language.GitHubCiBuilder }}
        with:
          {{ .Language.Name | ToLower }}-version: {{ .Language.Version }}

      - name: Install dependencies
        run: |
          python -m pip install --upgrade pipenv wheel

      - name: Execute Formatter
        run: |
          python -m pip install flake8
          flake8

`)
}
