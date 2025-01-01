package controller

import (
	"fmt"

	"net/http"

	"github.com/labstack/echo/v4"
)

// ------------- Relations -------------
// ShowAllRelations
func (mc ModelController) StartGeneration(c echo.Context) error {
	if err := mc.repo.StartGeneration(); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("Generation failed %v", err))
	}

	return mc.ShowAllEntities(c)

}
