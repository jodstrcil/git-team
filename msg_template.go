package main

var msgFormat =
`{{- if .TicketNumber}}[{{.TicketTag}}-{{ .TicketNumber }}] {{end}}{{ .Message }}
{{- if .Collaborators}}
{{println}}{{- range .Collaborators }}
{{- if .Email}}
Co-Authored-by: {{ .FullName }} <{{ .Email }}>{{- end }}
{{- end }}
{{- end }}
`
