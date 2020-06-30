package PDKlogic

const (
	PyShow           int64 = 0x0000000001 //显示牌
	PyNoShow         int64 = 0x0000000002 //不显示牌
	PyMustSpadeThree int64 = 0x0000000004 //首局黑桃3先出
	PyMust           int64 = 0x0000000008 //必须管
	PyMaybe          int64 = 0x0000000010 //可不管
	PyZhaNiao        int64 = 0x0000000020 //红桃10扎鸟
	PyFifteen        int64 = 0x0000000040 //15张
	PySixteen        int64 = 0x0000000080 //16张
	PyOneDeckPoker   int64 = 0x0000000100 //1副牌
	PyTwoDeckPoker   int64 = 0x0000000200 //2副牌
	PyLaiZi          int64 = 0x0000000400 //赖子
	PyOverlordBank   int64 = 0x0000000800 //霸王庄
	Py2AChunTian     int64 = 0x0000001000 //2AAA春天
	Py0BombBet       int64 = 0x0000002000 //炸弹不算分
	Py5BombBet       int64 = 0x0000004000 //5分炸弹
	Py10BombBet      int64 = 0x0000008000 //10分炸弹
	PySeparateBomb   int64 = 0x0000010000 //炸弹可拆
	Py4TakeTwo       int64 = 0x0000020000 //可以四带二
	PyOlnyOneMax     int64 = 0x0000040000 //报单出最大
	Py3TakePair      int64 = 0x0000080000 //可以三带一对
	Py3Takesingle    int64 = 0x0000100000 //可以三带一
	Py4TakeOne       int64 = 0x0000200000 //可以四带一
	PyOnlySpadeThree int64 = 0x0000400000 //首句单出黑桃3
	PyAllSpadeThree  int64 = 0x0000800000 //每局黑桃3先出
	PyOnlyThreeOne   int64 = 0x0001000000 //仅三带一

)

const (
	CtError              int32 = 0  //错误
	CtSingle             int32 = 1  //单牌
	CtSingleLine         int32 = 2  //单连 5张以上
	CtDouble             int32 = 3  //单对子 55
	CtDoubleLine         int32 = 4  //连对 5566
	CtThree              int32 = 5  //三条
	CtThreeTakeSingle    int32 = 6  //三带一
	CtThreeTakeDouble    int32 = 7  //三带二
	CtThreeLine          int32 = 8  //飞机
	CtBomb               int32 = 9  //炸弹
	CtKingBomb           int32 = 10 //王炸
	CtStraightFlush      int32 = 11 //同花顺
	CtThreeLineTakeLess  int32 = 12 //飞机少带牌
	CtFourTakePair       int32 = 13 //四带一对
	Ct3ABomb             int32 = 14 //3A炸弹
	CtThreeTakePair      int32 = 15 //三带一对
	CtFourTakeDoublePair int32 = 16 //四带二对
	CtFourTakeSingle     int32 = 17 //四带一
	CtFourTakeDouble     int32 = 18 //四带两张单牌
)

//内部维护游戏状态
type GameStatusType int

const (
	StatusFree GameStatusType = 1 //空闲场景
	StatusPlay GameStatusType = 2 //游戏场景
)

type UserOutPokerCode int32

const (
	UserOutPokerSuccess                 UserOutPokerCode = 0
	UserOutPokerError                   UserOutPokerCode = 1
	UserOutPokerMustSpadeThree          UserOutPokerCode = 2
	UserOutPokerMustMax                 UserOutPokerCode = 3
	UserOutPokerFirstSpade              UserOutPokerCode = 4
	UserOutPokerMustMode                UserOutPokerCode = 5
	UserOutPokerCanntThree              UserOutPokerCode = 6
	UserOutPokerThreeTakeSingle         UserOutPokerCode = 7
	UserOutPokerThreeLineTakeLess       UserOutPokerCode = 8
	UserOutPokerMustHeartThree          UserOutPokerCode = 9
	UserOutPokerStatus                  UserOutPokerCode = 10
	UserOutPokerCurrent                 UserOutPokerCode = 11
	UserOutPokerNotExist                UserOutPokerCode = 12
	UserOutPokerMustA                   UserOutPokerCode = 13
	UserOutPokerMustAA                  UserOutPokerCode = 14
	UserOutPokerMustTwo                 UserOutPokerCode = 15
	UserOutPokerMustBomb                UserOutPokerCode = 16
	UserOutPokerThreeLineTakeDouble     UserOutPokerCode = 17 // 不能三带二
	UserOutPokerMustBySingle            UserOutPokerCode = 18
	UserOutPokerMustSkyBomb             UserOutPokerCode = 19
	UserOutPokerCanntThreeTakePair      UserOutPokerCode = 20 // 三带一对
	UserOutPokerCanntFourTakeDoublePair UserOutPokerCode = 21 // 四带二对
	UserOutPokerSeparateBomb            UserOutPokerCode = 22 // 炸弹不可拆
	UserOutPokerCanntThreeLinePair      UserOutPokerCode = 23 // 不能飞机带对
	UserOutPokerCanntFourTakeSingle     UserOutPokerCode = 24 // 不能四带一
	UserOutPokerCanntFourTakeTwo        UserOutPokerCode = 25 // 不能四带二
)

const (
	GAME_TIMER_PLAY = 15
)

const (
	MAIN_GAME_ID = 301 //游戏内部的主命令id
)

//S->C
const (
	SUB_S_NotifyGameStart        = 100
	SUB_S_NotifyOutCard          = 101
	SUB_S_BroadcastOutCard       = 102
	SUB_S_BroadcastGameEnd       = 103
	SUB_S_BroadcastGameOver      = 104
	SUB_S_BroadcastSceneGameFree = 105
	SUB_S_BroadcastSceneGamePlay = 106
)

//C->S
const (
	SUB_C_OutPoker = 1
)
