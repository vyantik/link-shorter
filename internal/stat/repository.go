package stat

import (
	"app/test/pkg/db"
	"time"

	"gorm.io/datatypes"
)

type StatRepository struct {
	database *db.Db
}

func NewStatRepository(db *db.Db) *StatRepository {
	return &StatRepository{
		database: db,
	}
}

func (r *StatRepository) AddClick(linkId uint) {
	var stat Stat
	currentDate := datatypes.Date(time.Now())

	r.database.DB.Find(&stat, "link_id = ? and date = ?", linkId, currentDate)
	if stat.ID == 0 {
		r.database.DB.Create(&Stat{
			LinkID: linkId,
			Date:   currentDate,
			Clicks: 1,
		})
	} else {
		stat.Clicks++
		r.database.DB.Save(&stat)
	}
}

func (r *StatRepository) GetStats(by string, from time.Time, to time.Time) []GetStatResponse {
	var stats []GetStatResponse
	var selectQuery string

	switch by {
	case GroupByDay:
		selectQuery = "TO_CHAR(date, 'YYYY-MM-DD') AS period, SUM(clicks) as total"
	case GroupByMonth:
		selectQuery = "TO_CHAR(date, 'YYYY-MM') AS period, SUM(clicks) as total"
	}

	r.database.Table("stats").
		Select(selectQuery).
		Where("date BETWEEN ? AND ?", from, to).
		Group("period").
		Order("period").
		Scan(&stats)

	return stats
}
