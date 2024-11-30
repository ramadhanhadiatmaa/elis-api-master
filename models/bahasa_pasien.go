package models

type BahasaPasien struct {
	Id         int    `gorm:"type:int(11);primaryKey" json:"id"`
	NamaBahasa string `gorm:"type:varchar(30);not null" json:"nama_bahasa"`
}

func (BahasaPasien) TableName() string {
	return "bahasa_pasien"
}
