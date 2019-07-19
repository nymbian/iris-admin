package libs

import (
	htmpl "html/template"
	"time"

	config "github.com/spf13/viper"
)

func TimeToDate(s time.Time) string {
	return s.Format(config.GetString("site.TimeFormat"))
}

func StrToHtml(raw string) htmpl.HTML {
	return htmpl.HTML(raw)
}
