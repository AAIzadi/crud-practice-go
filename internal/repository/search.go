package repository

type PagingAndSorting struct {
	Limit     int
	Offset    int
	SortBy    string
	SortOrder string
}
