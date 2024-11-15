package utils

import (
	"go-carbon-tracker/constants"
	"go-carbon-tracker/entities"	
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func GetFilter(c echo.Context) (entities.Filter, error) {
	var startDate, endDate time.Time

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	search := c.QueryParam("search")

	sortBy := c.QueryParam("sort_by")
	if sortBy == "" {
		sortBy = "created_at"
	}

	sort := strings.ToUpper(c.QueryParam("sort"))
	if sort == "" {
		sort = "ASC"
	}

	startDateStr := c.QueryParam("start_date")

	endDateStr := c.QueryParam("end_date")

	if startDateStr != "" {
		startDate, err = time.Parse("2006-01-02", startDateStr)

		if err != nil {
			return entities.Filter{}, constants.Err_INVALID_START_DATE_FORMAT
		}
	}

	if endDateStr != "" {
		endDate, err = time.Parse("2006-01-02", endDateStr)

		if err != nil {
			return entities.Filter{}, constants.Err_INVALID_END_DATE_FORMAT
		}
	}

	params := entities.Filter{
		Page:      page,
		Limit:     limit,
		Search:    search,
		Sort:      sort,
		SortBy:    sortBy,
		StartDate: startDate,
		EndDate:   endDate,
	}

	return params, nil
}
