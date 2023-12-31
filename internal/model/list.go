package model

import (
	"fmt"
	"strings"
)

var titleStatus = [...]string{"watching", "completed", "on-hold", "dropped", "plan to watch"}

type ListInfo struct {
	ListID  int64 `json:"list_id"`
	OwnerID int64 `json:"user_id"`
}

type SearchResult struct {
	Docs []Movie `json:"docs"`
}

type Movie struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type ListUnit struct {
	Movie
	Status     string `json:"status"`
	Score      uint8  `json:"score"`
	IsFavorite bool   `json:"is_favorite"`
	ListInfo   `json:"-"`
}

/*
couldn't think of anything else to implement the functionality
of partial update (PATCH HTTP method) of the list_titles table fields..
so, I had to duplicate the code of the structure, but with pointers
*/
type ListUnitPatch struct {
	ListID     *int64  `json:"-"`
	OwnerID    *int64  `json:"-"`
	MovieID    *int64  `json:"-"`
	Status     *string `json:"status"`
	Score      *uint8  `json:"score"`
	IsFavorite *bool   `json:"is_favorite"`
}

func (u *ListUnit) Validate() error {
	if len(u.Name) == 0 {
		return fmt.Errorf("empty movie name")
	}
	if u.Score > 10 {
		return fmt.Errorf("score cannot be greater than 10")
	}
	for _, status := range titleStatus {
		if strings.EqualFold(status, u.Status) {
			u.Status = status
			return nil
		}
	}
	return fmt.Errorf("invalid title status")
}

func (u *ListUnitPatch) Validate() error {
	if u.Score != nil && *u.Score > 10 {
		return fmt.Errorf("score cannot be greater than 10")
	}
	if u.Status == nil {
		return nil
	}
	for _, status := range titleStatus {
		if strings.EqualFold(status, *u.Status) {
			*u.Status = status
			return nil
		}
	}
	return fmt.Errorf("invalid title status")
}
