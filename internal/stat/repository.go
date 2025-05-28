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
