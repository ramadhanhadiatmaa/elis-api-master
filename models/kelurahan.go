package models

type Kelurahan struct {
	KdKel int    `gorm:"type:int(11);primaryKey;autoIncrement" json:"kd_kel"`
	NmKel string `gorm:"type:varchar(60)" json:"nm_kel"`
}

func (Kelurahan) TableName() string {
	return "kelurahan"
}
