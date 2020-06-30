package PDKlogic

import (
	"lcm1/pdk/pdkdefine"
)

/*
   洗牌，发牌，比牌校验，出牌玩家轮换，计算结果等功能。负责跟游戏框架连接，尽量不要使用框架消息。
*/

// Table 桌子
type Table struct {
	PlayerInfo    []Player                // 玩家信息
	CurChairID    int                     // 当前玩家座位号 -->PlayerINfo[CurChairId] 当前玩家信息
	TurnChairID   int                     // 上个玩家
	GameRule      int64                   // 游戏规则
	RoomType      int32                   // 房间类型
	PlayerCount   int                     // 玩家数
	currentStatus PDKlogic.GameStatusType // 游戏状态
	GameLogic     Logic                   // 游戏规则
}

// Init 桌子初始化 PS 玩家人数，规则，房间类型需要在初始化之前获取
func (t *Table) Init() {
	t.PlayerInfo = make([]Player, t.PlayerCount)
	for i := 0; i != t.PlayerCount; i++ {
		t.PlayerInfo[i].Init()
	}
	t.CurChairID = -1 // define.INVALID_CHAIR_ID
	t.TurnChairID = -1
	t.currentStatus = PDKlogic.StatusFree
	t.GameLogic.SetGameRule(t.GameRule)
}

// Release 清理桌子
func (t *Table) Release() {
	for i := 0; i != t.PlayerCount; i++ {
		t.PlayerInfo[i].Release()
	}
}

// SetPlayCnt 设置当前人数
func (t *Table) SetPlayCnt(cnt int) {
	t.PlayerCount = cnt
}

// SetGameRule 获取游戏规则
func (t *Table) SetGameRule(rule int64) {
	t.GameRule = rule
}

// SetRoomType 获取房间类型
func (t *Table) SetRoomType(ty int32) {
	t.RoomType = ty
}
