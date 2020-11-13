package clientInfo

import (
	"github.com/weijunji/SA-Project/pkg/db"
	"gorm.io/gorm/clause"
)

type ClientInfo struct {
	UUID   string `gorm:"primaryKey"`
	Name   string
	Online bool
	Locked bool
}

func init() {
	db.GetDB().AutoMigrate(&ClientInfo{})
}

func GetClientInfos() []ClientInfo {
	var result []ClientInfo
	db.GetDB().Find(&result)
	return result
}

func GetClient(uuid string) *ClientInfo {
	c := new(ClientInfo)
	db.GetDB().Where("uuid = ?", uuid).First(c)
	return c
}

func NewClientInfo(uuid string, online bool, locked bool) *ClientInfo {
	c := &ClientInfo{uuid, uuid, online, locked}
	db.GetDB().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "uuid"}},
		DoUpdates: clause.AssignmentColumns([]string{"online", "locked"}),
	}).Create(c)
	return c
}

func (c *ClientInfo) OfflineOp() {
	if c.Online {
		c.Online = false
		db.GetDB().Model(c).Select("online").Updates(map[string]interface{}{"online": false})
	}
}

func (c *ClientInfo) OnlineOp() {
	if !c.Online {
		c.Online = true
		db.GetDB().Model(c).Select("online").Updates(map[string]interface{}{"online": true})
	}
}

func (c *ClientInfo) UnlockOp() {
	if c.Locked {
		c.Locked = false
		db.GetDB().Model(c).Select("locked").Updates(map[string]interface{}{"locked": false})
	}
}

func (c *ClientInfo) LockOp() {
	if !c.Locked {
		c.Locked = true
		db.GetDB().Model(c).Select("locked").Updates(map[string]interface{}{"locked": true})
	}
}

func (c *ClientInfo) SetName(name string) {
	db.GetDB().Model(c).Select("name").Updates(map[string]interface{}{"name": name})
}
