package config

import (
	"net/http"
	"database/sql"
)

type HandlerWithDB func(http.ResponseWriter, *http.Request, *sql.DB)

type TemplateVaribles struct {
	Title string
}