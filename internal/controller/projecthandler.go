package controller

import (
	model "crudgengui/model"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
)

type projectdao struct {
	Name       string `form:"field_name" `
	Currency   string `form:"field_currency" `
	Decimal    string `form:"field_decimal" `
	Thousend   string `form:"field_thousend" `
	TimeFormat string `form:"field_timeformat" `
	DateFormat string `form:"field_dateformat" `
}

// ShowProject shows detail page to model or edit screen for new model
func (mc ModelController) ShowProject(c echo.Context) error {
	m := mc.repo.GetModel()

	text := map[string]string{
		"title": "Project",
		"menu":  "menu_project",
	}

	rdao := projectdao{Name: m.Name, Currency: m.CurrencySymbol, Decimal: m.DecimalSeparator, Thousend: m.ThousendSeparator, TimeFormat: m.TimeFormat, DateFormat: m.DateFormat}

	rd := newRequestData(text, map[string]interface{}{
		"model": rdao,
	})

	return c.Render(http.StatusOK, "project.html", rd)

}

// SaveProject
func (mc ModelController) SaveProject(c echo.Context) error {
	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()

	rdao := projectdao{}

	if err := c.Bind(&rdao); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	rsettings := model.Settings{CurrencySymbol: rdao.Currency, DecimalSeparator: rdao.Decimal, ThousendSeparator: rdao.Thousend, TimeFormat: rdao.TimeFormat, DateFormat: rdao.DateFormat}
	if err := mc.repo.SaveModel(rdao.Name, rsettings); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return mc.ShowAllEntities(c)
}
