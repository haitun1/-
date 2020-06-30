package PDKlogic

/*
   玩家操作： 出牌，过牌，等操作
*/

// Player 玩家信息 后期可考虑将Player设为私有
type Player struct {
	Poker       []int32   // 玩家手牌 切片，严格按照append和[:i]之类增删 ，故可省略玩家手牌数量
	PokerCnt    int       // 手牌数量，考虑加上手牌数量，可以少多次len()函数调用
	Score       ScoreInfo // 玩家分数信息
	CurPoker    []int32   // 当前出牌 PS：可以锁定上个出牌
	CurPType    int       // 当前出牌牌型
	WarnOnlyOne bool      // 单张报警， 报警true，默认false

}

// ScoreInfo 玩家分数信息
type ScoreInfo struct {
	MaxWinScore    int // 单局最高得分
	BombTotalCount int // 总炸弹数
	WinCount       int // 赢次数
	LostCount      int // 数次数
	CurScore       int // 当前分/总分
	GameScore      int // 单局结算分数
	BombScore      int // 单局炸弹分数
	BombCount      int // 单局炸弹数
}

// Init 全局就初始化一次
func (p *Player) Init() {
	p.Poker = p.Poker[:0]
	p.CurPoker = p.CurPoker[:0]
	p.CurPType = 0
	p.WarnOnlyOne = false
	p.Score.BombCount = 0
	p.Score.BombScore = 0
	p.Score.BombTotalCount = 0
	p.Score.CurScore = 0
	p.Score.GameScore = 0
	p.Score.LostCount = 0
	p.Score.MaxWinScore = 0
	p.Score.WinCount = 0
}

// Release 每局游戏结束清空上局信息
func (p *Player) Release() {
	p.Poker = p.Poker[:0]
	p.CurPoker = p.CurPoker[:0]
	p.CurPType = 0
	p.WarnOnlyOne = false
	p.Score.GameScore = 0
	p.Score.BombScore = 0
	p.Score.BombCount = 0
}
