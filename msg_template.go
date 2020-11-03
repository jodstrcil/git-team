package main

var msgFormat =
`{{- if .JiraNumber}}[{{.TicketTag}}-{{ .JiraNumber }}] {{end}}{{ .Message }}
{{- if .Collaborators}}
{{println}}{{- range .Collaborators }}
{{- if .Email}}
Co-Authored-by: {{ .FullName }} <{{ .Email }}>{{- end }}
{{- end }}
{{- end }}
`
