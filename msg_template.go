package main

var msgFormat =
`{{- if .JiraNumber}}[{{.JiraTag}}-{{ .JiraNumber }}] {{end}}{{ .Message }}
{{- if .Collaborators}}
{{println}}{{- range .Collaborators }}
{{- if .Email}}
Co-Authored-by: {{ .FullName }} <{{ .Email }}>{{- end }}
{{- end }}
{{- end }}
`
