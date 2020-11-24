package clientInfo

import (
	"time"

	"github.com/weijunji/SA-Project/pkg/db"
	"gorm.io/gorm/clause"
)

type ClientInfo struct {
	UUID   string `gorm:"primaryKey"`
	Name   string
	Online bool
	Locked bool
}

type ClientLog struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time
	UUID      string
	Operation uint // 0:lock 1:unlock 2:Online 3:Offline
}

const (
	Lock = iota
	Unlock
	Online
	Offline
)

func init() {
	db.GetDB().AutoMigrate(&ClientInfo{})
	db.GetDB().AutoMigrate(&ClientLog{})
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
	db.GetDB().Create(&ClientLog{UUID: c.UUID, Operation: Online})
	return c
}

func (c *ClientInfo) OfflineOp() {
	if c.Online {
		c.Online = false
		db.GetDB().Model(c).Select("online").Updates(map[string]interface{}{"online": false})
		db.GetDB().Create(&ClientLog{UUID: c.UUID, Operation: Offline})
		if !c.Locked {
			db.GetDB().Create(&ClientLog{UUID: c.UUID, Operation: Lock})
		}
	}
}

func (c *ClientInfo) OnlineOp() {
	if !c.Online {
		c.Online = true
		db.GetDB().Model(c).Select("online").Updates(map[string]interface{}{"online": true})
		db.GetDB().Create(&ClientLog{UUID: c.UUID, Operation: Online})
	}
}

func (c *ClientInfo) UnlockOp() {
	if c.Locked {
		c.Locked = false
		db.GetDB().Model(c).Select("locked").Updates(map[string]interface{}{"locked": false})
		db.GetDB().Create(&ClientLog{UUID: c.UUID, Operation: Unlock})
	}
}

func (c *ClientInfo) LockOp() {
	if !c.Locked {
		c.Locked = true
		db.GetDB().Model(c).Select("locked").Updates(map[string]interface{}{"locked": true})
		db.GetDB().Create(&ClientLog{UUID: c.UUID, Operation: Lock})
	}
}

func (c *ClientInfo) SetName(name string) {
	db.GetDB().Model(c).Select("name").Updates(map[string]interface{}{"name": name})
}

func GetClientLogs(UUID string) (logs []ClientLog) {
	db.GetDB().Where("UUID = ?", UUID).Find(&logs)
	return
}
