package main

const entityTmpl = `// Code generated by mkentity, DO NOT EDIT.
package {{ $.PackageName }}

import (
	"context"

	"github.com/gotd/td/tg"
)

var (
	_ = tg.Invoker(nil)
	_ = context.Context(nil)
)

{{- /*gotype: github.com/gotd/td/telegram/message/internal/mkrun.Config*/ -}}
{{- range $typ := $.Data }}
{{ $helperName := trimPrefix ( trimPrefix $typ.Name "Input" ) "MessageEntity" -}}
// {{ $helperName }} creates Formatter of {{ $helperName }} message entity.
//
// See https://core.telegram.org/constructor/{{ $typ.SchemaType.Name }}.
func {{ $helperName }}({{- range $f := $typ.Fields }}{{ lowerFirst $f.Name }} {{ $f.Type }}{{- end }}) Formatter {
	return func(offset, length int) tg.MessageEntityClass {
		return &tg.{{ $typ.Name }}{
			Offset: offset,
			Length: length,
			{{- range $f := $typ.Fields }}
			{{ $f.Name }}: {{ lowerFirst $f.Name }},
			{{- end }}
		}
	}
}

// {{ $helperName }} adds and formats message as {{ $helperName }} message entity.
//
// See https://core.telegram.org/constructor/{{ $typ.SchemaType.Name }}.
func (b *Builder) {{ $helperName }}(s string,
{{- range $f := $typ.Fields }}{{ lowerFirst $f.Name }} {{ $f.Type }}{{- end }}) *Builder {
	return b.Format(s, {{ $helperName }}({{- range $f := $typ.Fields }}{{ lowerFirst $f.Name }},{{- end }}))
}
{{- end }}
`
