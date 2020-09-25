package main

var msgFormat =
`{{- if .JiraNumber}}[{{.JiraTag}}-{{ .JiraNumber }}] {{end}}{{ .Message }}
{{- if .Pair}}
{{println}}{{- range .Pair }}
{{- if .Email}}
Co-Authored-by: {{ .FullName }} <{{ .Email }}>{{- end }}
{{- end }}
{{- end }}
`
