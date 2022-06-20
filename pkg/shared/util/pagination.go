package util

type PaginationHelper struct {
	Total       int
	PerPage     int
	CurrentPage int
	LastPage    int
	FormIndex   int
	ToIndex     int
	HasNext     string
	HasPrev     string
}

func MakePagination(curPage, limit, countAll int) *PaginationHelper {
	if countAll == 0 {
		return &PaginationHelper{
			HasNext: "N",
			HasPrev: "N",
		}
	}
	if curPage == 0 {
		curPage = 1
	}

	if limit == 0 {
		limit = countAll
	}
	var lastPage, remaining int

	remaining = countAll % limit

	if remaining != 0 {
		lastPage = countAll/limit + 1
	} else {
		lastPage = countAll / limit
	}

	fromIndex := (curPage-1)*limit + 1

	toIndex := fromIndex - 1 + limit
	if curPage == lastPage {
		toIndex = fromIndex - 1 + remaining
	}

	hasPrev := "N"
	if curPage > 1 {
		hasPrev = "Y"
	}
	hasNext := "N"
	if curPage < lastPage && curPage >= 1 {
		hasNext = "Y"
	}

	return &PaginationHelper{
		ToIndex:     toIndex,
		FormIndex:   fromIndex,
		Total:       countAll,
		PerPage:     limit,
		LastPage:    lastPage,
		CurrentPage: curPage,
		HasPrev:     hasPrev,
		HasNext:     hasNext,
	}
}
