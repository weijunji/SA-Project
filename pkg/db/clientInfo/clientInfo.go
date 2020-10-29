package clientInfo

import (
	"github.com/weijunji/SA-Project/pkg/db"
	"gorm.io/gorm/clause"
)

type ClientInfo struct {
	UUID   string `gorm:"primaryKey"`
	Status bool   // true: online false: offline
	Locked bool
}

func init() {
	db.GetDB().AutoMigrate(&ClientInfo{})
}

func NewClientInfo(uuid string, online bool, locked bool) *ClientInfo {
	c := &ClientInfo{uuid, online, locked}
	db.GetDB().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "uuid"}},
		DoUpdates: clause.AssignmentColumns([]string{"status", "locked"}),
	}).Create(c)
	return c
}

func (c *ClientInfo) Offline() {
	if c.Status {
		c.Status = false
		db.GetDB().Model(c).Select("status").Updates(map[string]interface{}{"status": false})
	}
}

func (c *ClientInfo) Online() {
	if !c.Status {
		c.Status = true
		db.GetDB().Model(c).Select("status").Updates(map[string]interface{}{"status": true})
	}
}

func (c *ClientInfo) Unlock() {
	if c.Locked {
		c.Locked = false
		db.GetDB().Model(c).Select("locked").Updates(map[string]interface{}{"locked": false})
	}
}

func (c *ClientInfo) Lock() {
	if !c.Locked {
		c.Locked = true
		db.GetDB().Model(c).Select("locked").Updates(map[string]interface{}{"locked": true})
	}
}
