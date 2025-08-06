package pagination

import "gorm.io/gorm"

type Meta struct {
	CurrentPage int   `json:"current_page"`
	From        int   `json:"from"`
	To          int   `json:"to"`
	LastPage    int   `json:"last_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
}
type PaginateResult struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta"`
}

func Paginate[T any, R any](
	db *gorm.DB,
	page, limit int,
	rawFunc func(*gorm.DB) *gorm.DB,
	output *[]T,
	mapFunc func([]T) []R,
) (PaginateResult, error) {
	offset := (page - 1) * limit

	query := db
	if rawFunc != nil {
		query = rawFunc(query)
	}

	var total int64
	query.Model(output).Count(&total)

	err := query.Offset(offset).Limit(limit).Find(output).Error
	if err != nil {
		return PaginateResult{}, err
	}

	to := offset + limit
	if to > int(total) {
		to = int(total)
	}

	dtoData := mapFunc(*output)

	return PaginateResult{
		Data: dtoData,
		Meta: Meta{
			CurrentPage: page,
			From:        offset + 1,
			To:          to,
			LastPage:    (int(total) + limit - 1) / limit,
			PerPage:     limit,
			Total:       total,
		},
	}, nil
}
