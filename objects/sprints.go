package objects

import (
	"time"
)

type Sprint struct {
	ID            int        `json:"id"`
	Name          string     `json:"name"`
	Self          string     `json:"self"`
	State         string     `json:"state"`
	Goal          string     `json:"goal"`
	OriginBoardID int        `json:"originBoardId"`
	EndDate       *time.Time `json:"endDate"`
	StartDate     *time.Time `json:"startDate"`
	OldValue      *Sprint    `json:"oldValue"`
}
