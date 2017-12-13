package echotool

import (
    "io"
    "os"
    "sync"
    "github.com/labstack/echo"
    "html/template"
)

type H map[string]interface{}
var (
    // Global instance
    defaultE     *echo.Echo
    defaultELock = sync.Mutex{}
)

func DebugMode() bool {
    // TODO: fix env
    return os.Getenv("GIN_MODE") != "release"
}

type Template struct {
     *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.ExecuteTemplate(w, name, data)
}

// Using global instance to manager router packages
func Default() *echo.Echo {
    defaultELock.Lock()
    defer defaultELock.Unlock()
    if defaultE == nil {
        defaultE = echo.New()
        defaultE.Debug = DebugMode()
        defaultE.HideBanner = true
    }
    return defaultE
}
