package domains

type SlowQuery struct {
	ID        int64   `gorm:"column:queryid" json:"id"`
	Query     string  `gorm:"column:query" json:"query"`
	MaxTime   float64 `gorm:"column:max_time" json:"max_time"`
	MinTime   float64 `gorm:"column:min_time" json:"min_time"`
	MeanTime  float64 `gorm:"column:mean_time" json:"mean_time"`
	TotalTime float64 `gorm:"column:total_time" json:"total_time"`
}

func (q SlowQuery) TableName() string {
	return "pg_stat_statements"
}
