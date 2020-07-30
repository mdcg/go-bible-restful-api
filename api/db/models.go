package db

type Testament struct {
	Id   int    `gorm:"PRIMARY_KEY" json:"id"`
	Name string `gorm:"size:45" json:"name"`
}

func (Testament) TableName() string {
	return "testament"
}

type Books struct {
	Id        int    `gorm:"PRIMARY_KEY" json:"id"`
	Name      string `gorm:"size:45" json:"name"`
	Abbrev    string `gorm:"size:5" json:"abbrev"`
	Testament string `gorm:"size:5" json:"testament"`
}

type Verses struct {
	Id        int    `gorm:"PRIMARY_KEY" json:"id"`
	Version   string `gorm:"size:10" json:"version"`
	Testament int    `json:"testament"`
	Book      int    `json:"book"`
	Chapter   int    `json:"chapter"`
	Verse     int    `json:"verse"`
	Text      string `gorm:"type:text" json:"text"`
}
