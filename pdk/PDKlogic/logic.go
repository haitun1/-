package PDKlogic

import (
	"math/rand"
	"mjserver/gameprivate/gdefine/pdk"
	"mjserver/utils"
	"time"
)

/*
  游戏逻辑： 比牌， 提示， 牌型判断
*/

// Card48 ... 去掉1个A,3个2,2个王
var Card48 = []int32{
	0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, // 方块
	0x11, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, // 梅花
	0x21, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2A, 0x2B, 0x2C, 0x2D, // 红桃
	0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3A, 0x3B, 0x3C, 0x3D, // 黑桃
}

// SortInt32 排序
type SortInt32 []int32

func (a SortInt32) Len() int      { return len(a) }
func (a SortInt32) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortInt32) Less(i, j int) bool {

	b := a.LogicValue(a[i])
	c := a.LogicValue(a[j])

	if b == c {
		return b > c
	}
	return b > c

}

// LogicValue 排序规则
func (a SortInt32) LogicValue(value int32) int32 {
	cardValue := value & 0x0f

	if cardValue >= 0xe {
		return cardValue + 2
	}

	if cardValue <= 2 {
		return cardValue + 13
	}
	return cardValue
}

// RandomShuffle 洗牌
func RandomShuffle(src []int32) []int32 {
	dest := make([]int32, len(src))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	perm := r.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}

	return dest
}

// Logic  玩法规则
type Logic struct {
	log       *utils.Logger // 获取全局唯一的日志句柄
	rule      int64         // 游戏配置
	threeFour bool          // 三带二何时可以压三带对
}

// Init .
func (g *Logic) Init() {
	g.threeFour = false
}

// SetThreeFour 设置threeFour
func (g *Logic) SetThreeFour(th bool) {
	g.threeFour = th
}

// AnalyseResult .
type AnalyseResult struct {
	EightCount       int          // 八张数目
	SevenCount       int          // 七张数目
	SixCount         int          // 六张数目
	FiveCount        int          // 五张数目
	FourCount        int          // 四张数目
	ThreeCount       int          // 三张数目
	DoubleCount      int          // 两张数目
	SignedCount      int          // 单张数目
	EightLogicVolue  [3]int32     // 八张列表
	SevenLogicVolue  [7]int32     // 七张列表
	SixLogicVolue    [4]int32     // 六张列表
	FiveLogicVolue   [5]int32     // 五张列表
	FourLogicVolue   [7]int32     // 四张列表
	ThreeLogicVolue  [9]int32     // 三张列表
	DoubleLogicVolue [14]int32    // 两张列表
	SignedLogicVolue [27]int32    // 单张列表
	EightCardData    [27]int32    // 八张列表
	SevenCardData    [27]int32    // 七张列表
	SixCardData      [27]int32    // 六张列表
	FiveCardData     [27]int32    // 五张列表
	FourCardData     [27]int32    // 四张列表
	ThreeCardData    [27]int32    // 三张列表
	DoubleCardData   [27]int32    // 两张列表
	SignedCardData   [27]int32    // 单张数目
	PokerData        [8][27]int32 // 扑克数据
	BlockCount       [8]int
}

// AnalysebCardData 牌型组合
func (g *Logic) AnalysebCardData(cardData []int32, cardCount int, analyseResult *AnalyseResult) {

	for i := 0; i < cardCount; i++ {
		// 变量定义
		sameCount := 1
		sameCardData := []int32{cardData[i], 0, 0, 0, 0, 0, 0, 0}
		logicValue := g.GetPokerLogicValue(cardData[i])

		// 获取同牌
		for j := i + 1; j < cardCount; j++ {
			// 逻辑对比
			if g.GetPokerLogicValue(cardData[j]) != logicValue {
				break
			}

			// 设置扑克
			sameCardData[sameCount] = cardData[j]
			sameCount++
		}

		// 设置结果
		//analyseResult.BlockCount[sameCount-1]++
		index := analyseResult.BlockCount[sameCount-1]
		analyseResult.BlockCount[sameCount-1]++
		for k := 0; k < sameCount; k++ {
			analyseResult.PokerData[sameCount-1][index*sameCount+k] = cardData[i+k]
		}

		// 保存结果
		switch sameCount {
		case 1: // 单张
			analyseResult.SignedLogicVolue[analyseResult.SignedCount] = logicValue
			copy(analyseResult.SignedCardData[(analyseResult.SignedCount)*sameCount:], sameCardData)
			analyseResult.SignedCount++
			break

		case 2: // 两张
			analyseResult.DoubleLogicVolue[analyseResult.DoubleCount] = logicValue
			copy(analyseResult.DoubleCardData[(analyseResult.DoubleCount)*sameCount:], sameCardData)
			analyseResult.DoubleCount++
			break

		case 3: // 三张
			analyseResult.ThreeLogicVolue[analyseResult.ThreeCount] = logicValue
			copy(analyseResult.ThreeCardData[(analyseResult.ThreeCount)*sameCount:], sameCardData)
			analyseResult.ThreeCount++
			break

		case 4: // 四张
			analyseResult.ThreeLogicVolue[analyseResult.ThreeCount] = logicValue
			copy(analyseResult.ThreeCardData[(analyseResult.ThreeCount)*(sameCount-1):], sameCardData[:3]) // 隐藏bug，提示三张牌型不对，以下更改类似
			analyseResult.ThreeCount++

			analyseResult.FourLogicVolue[analyseResult.FourCount] = logicValue
			copy(analyseResult.FourCardData[(analyseResult.FourCount)*sameCount:], sameCardData)
			analyseResult.FourCount++
			break
		case 5: // 五张
			analyseResult.ThreeLogicVolue[analyseResult.ThreeCount] = logicValue
			copy(analyseResult.ThreeCardData[(analyseResult.ThreeCount)*(sameCount-2):], sameCardData[:3])
			analyseResult.ThreeCount++

			analyseResult.FourLogicVolue[analyseResult.FourCount] = logicValue
			copy(analyseResult.FourCardData[(analyseResult.FourCount)*(sameCount-1):], sameCardData[:4])
			analyseResult.FourCount++

			analyseResult.FiveLogicVolue[analyseResult.FourCount] = logicValue
			copy(analyseResult.FiveCardData[(analyseResult.FiveCount)*sameCount:], sameCardData)
			analyseResult.FiveCount++
			break
		case 6: // 六张
			analyseResult.ThreeLogicVolue[analyseResult.ThreeCount] = logicValue
			copy(analyseResult.ThreeCardData[(analyseResult.ThreeCount)*(sameCount-3):], sameCardData[:3])
			analyseResult.ThreeCount++

			analyseResult.FourLogicVolue[analyseResult.FourCount] = logicValue
			copy(analyseResult.FourCardData[(analyseResult.FourCount)*(sameCount-2):], sameCardData[:4])
			analyseResult.FourCount++

			analyseResult.FiveLogicVolue[analyseResult.FourCount] = logicValue
			copy(analyseResult.FiveCardData[(analyseResult.FiveCount)*(sameCount-1):], sameCardData[:5])
			analyseResult.FiveCount++

			analyseResult.SixLogicVolue[analyseResult.SixCount] = logicValue
			copy(analyseResult.SixCardData[(analyseResult.SixCount)*sameCount:], sameCardData)
			analyseResult.SixCount++
			break
		case 7: // 七张
			analyseResult.ThreeLogicVolue[analyseResult.ThreeCount] = logicValue
			copy(analyseResult.ThreeCardData[(analyseResult.ThreeCount)*(sameCount-4):], sameCardData[:3])
			analyseResult.ThreeCount++

			analyseResult.FourLogicVolue[analyseResult.FourCount] = logicValue
			copy(analyseResult.FourCardData[(analyseResult.FourCount)*(sameCount-3):], sameCardData[:4])
			analyseResult.FourCount++

			analyseResult.FiveLogicVolue[analyseResult.FourCount] = logicValue
			copy(analyseResult.FiveCardData[(analyseResult.FiveCount)*(sameCount-2):], sameCardData[:5])
			analyseResult.FiveCount++

			analyseResult.SixLogicVolue[analyseResult.SixCount] = logicValue
			copy(analyseResult.SixCardData[(analyseResult.SixCount)*(sameCount-3):], sameCardData[:6])
			analyseResult.SixCount++

			analyseResult.SevenLogicVolue[analyseResult.SevenCount] = logicValue
			copy(analyseResult.SevenCardData[(analyseResult.SevenCount)*sameCount:], sameCardData)
			analyseResult.SevenCount++
			break
		case 8: // 八张
			analyseResult.ThreeLogicVolue[analyseResult.ThreeCount] = logicValue
			copy(analyseResult.ThreeCardData[(analyseResult.ThreeCount)*(sameCount-5):], sameCardData[:3])
			analyseResult.ThreeCount++

			analyseResult.FourLogicVolue[analyseResult.FourCount] = logicValue
			copy(analyseResult.FourCardData[(analyseResult.FourCount)*(sameCount-4):], sameCardData[:4])
			analyseResult.FourCount++

			analyseResult.FiveLogicVolue[analyseResult.FourCount] = logicValue
			copy(analyseResult.FiveCardData[(analyseResult.FiveCount)*(sameCount-3):], sameCardData[:5])
			analyseResult.FiveCount++

			analyseResult.SixLogicVolue[analyseResult.SixCount] = logicValue
			copy(analyseResult.SixCardData[(analyseResult.SixCount)*(sameCount-2):], sameCardData[:6])
			analyseResult.SixCount++

			analyseResult.SevenLogicVolue[analyseResult.SevenCount] = logicValue
			copy(analyseResult.SevenCardData[(analyseResult.SevenCount)*(sameCount-1):], sameCardData[:7])
			analyseResult.SevenCount++

			analyseResult.EightLogicVolue[analyseResult.EightCount] = logicValue
			copy(analyseResult.EightCardData[(analyseResult.EightCount)*sameCount:], sameCardData)
			analyseResult.EightCount++
			break
		}

		// 设置递增
		i += (sameCount - 1)
	}
}

// SetGameRule 获取游戏规则
func (g *Logic) SetGameRule(rule int64) {
	g.rule = rule
}

// GetPokerValue 牌值
func (g *Logic) GetPokerValue(poker int32) int32 {
	return poker & 0x0f
}

// GetPokerColor 花色
func (g *Logic) GetPokerColor(poker int32) int32 {
	return poker & 0xf0
}

// GetPokerLogicValue 真实牌值
func (g *Logic) GetPokerLogicValue(cardData int32) int32 {
	if cardData == 0 {
		return 0
	}

	cardValue := cardData & 0x0f

	if cardValue >= 0xe {
		return cardValue + 2
	}

	if cardValue <= 2 {
		return cardValue + 13
	}
	return cardValue
}

// GetPokerType 牌型
func (g *Logic) GetPokerType(poker []int32) (pokerType int32, leftCount int) {
	switch len(poker) {
	case 1: // 单牌
		return pdk.CtSingle, 0
	case 2: // 对牌
		if g.GetPokerLogicValue(poker[0]) == g.GetPokerLogicValue(poker[1]) {
			return pdk.CtDouble, 0
		}

	}

	tempPoker := make([]int32, len(poker))
	copy(tempPoker, poker)

	return g.AnalysebPokerType(tempPoker)

}

// AnalysebPokerType . PS 可以考虑将牌型放在这里判断 后期可考虑拆分 先默认支持所有
func (g *Logic) AnalysebPokerType(poker []int32) (pokerType int32, leftCount int) {
	pokerCount := len(poker)
	var analyseResult AnalyseResult
	g.AnalysebCardData(poker, pokerCount, &analyseResult)
	if analyseResult.FourCount == 1 && pokerCount == 4 {
		return pdk.CtBomb, 0
	} else if analyseResult.FourCount == 1 && analyseResult.SignedCount == 1 && pokerCount == 5 { // 四带一
		return pdk.CtFourTakeSingle, 0
	} else if analyseResult.FiveCount == 1 && analyseResult.DoubleCount == 1 && pokerCount == 6 { // 四带对
		return pdk.CtFourTakePair, 0
	} else if analyseResult.FiveCount == 1 && analyseResult.SignedCount == 2 && pokerCount == 6 { // 四带两张单牌
		return pdk.CtFourTakeDouble, 0
	} else if analyseResult.FiveCount == 1 && analyseResult.DoubleCount == 2 && pokerCount == 8 { // 四带二对
		return pdk.CtFourTakeDoublePair, 0
	}
	// 对牌连子判断 后期可考虑拆分
	if analyseResult.DoubleCount > 1 && analyseResult.DoubleCount*2 == pokerCount {
		seriesFlag := false
		i := 1
		logicValue := analyseResult.DoubleLogicVolue[0]
		if logicValue < 15 {
			for ; i < analyseResult.DoubleCount; i++ {
				if analyseResult.DoubleLogicVolue[i] != (logicValue-1) && logicValue != 15 {
					g.log.Infof("连牌判断logicValue:%d,analyseResult.DoubleLogicVolue[i]:%d", logicValue, analyseResult.DoubleLogicVolue[i])
					break
				}
				logicValue = analyseResult.DoubleLogicVolue[i]
			}
		}
		if i == analyseResult.DoubleCount {
			seriesFlag = true
		}
		// 连对判断
		if seriesFlag == true && analyseResult.DoubleCount*2 == pokerCount {
			g.log.Infoln("连牌判断成功")
			return pdk.CtDoubleLine, 0
		}
	}
	// 单连判断
	if analyseResult.SignedCount > 4 && analyseResult.SignedCount == pokerCount {
		g.log.Infoln("开始单连判断")
		// 变量定义
		seriesFlag := false
		logicValue := g.GetPokerLogicValue(poker[0])
		// 连牌判断
		if logicValue < 15 {
			i := 1
			for ; i < analyseResult.SignedCount; i++ {
				if g.GetPokerLogicValue(poker[i]) != logicValue-1 {
					break
				}
				logicValue = g.GetPokerLogicValue(poker[i])
			}

			if i == analyseResult.SignedCount && analyseResult.SignedCount == pokerCount {
				seriesFlag = true
			}
		}
		// 单连判断
		if seriesFlag == true {
			g.log.Infoln("单连判断成功")
			return pdk.CtSingleLine, 0
		}
	}
	if analyseResult.ThreeCount == 1 && pokerCount < 6 {
		if analyseResult.ThreeCount*3 == pokerCount { // PS:可以考虑再这里限制三条
			return pdk.CtThree, 0
		} else if analyseResult.ThreeCount*3+1 == pokerCount {
			return pdk.CtThreeTakeSingle, 1
		} else if analyseResult.ThreeCount*3+2 == pokerCount && analyseResult.DoubleCount == 1 {
			return pdk.CtThreeTakePair, 2
		} else if analyseResult.ThreeCount*3+2 == pokerCount && analyseResult.SignedCount == 2 {
			return pdk.CtThreeTakePair, 2
		}
		return pdk.CtError, 0
	}
	if analyseResult.ThreeCount > 1 && analyseResult.ThreeCount*5 >= pokerCount { // 飞机问题 PS:在不清楚飞机最后一手能不能少带或者不带的情况下先默认少牌不让出
		for i := 0; i != analyseResult.ThreeCount-1; i++ {
			if analyseResult.ThreeLogicVolue[i] != analyseResult.ThreeLogicVolue[i+1]-1 {
				g.log.Infoln("飞机判断错误.", analyseResult.ThreeLogicVolue)
				return pdk.CtError, 0 // 飞机判断错误
			}
		}
		if analyseResult.ThreeCount*4 == pokerCount {
			return pdk.CtThreeLine, analyseResult.ThreeCount
		} else if analyseResult.DoubleCount == analyseResult.ThreeCount {
			return pdk.CtThreeLine, analyseResult.DoubleCount // 可以用pokerCount%5 == 0 || leftcount == 奇数来判断。
		} else if analyseResult.ThreeCount*5 == pokerCount {
			return pdk.CtThreeLine, analyseResult.ThreeCount * 2
		}
	}
	return pdk.CtError, 0
}

// RemovePoker 删除扑克 -->删除后手牌len(Poker)减少
func (g *Logic) RemovePoker(removePoker []int32, removeCount int, pokerData *[]int32, pokerCount int) bool {

	deleteCount := 0
	tempPokerData := make([]int32, pokerCount)
	if pokerCount > len(tempPokerData) {
		return false
	}

	copy(tempPokerData[0:], (*pokerData)[0:pokerCount])

	for i := 0; i < removeCount; i++ {
		for j := 0; j < pokerCount; j++ {
			if removePoker[i] == tempPokerData[j] {
				deleteCount++
				g.log.Infof("找到待删除扑克： %x\n", removePoker[i])
				tempPokerData[j] = 0
				break
			}
		}
	}
	if deleteCount != removeCount {
		return false
	}

	//清理扑克
	pos := 0
	for i := 0; i < pokerCount; i++ {
		if tempPokerData[i] != 0 {
			(*pokerData)[pos] = tempPokerData[i]
			pos++
		}
	}
	(*pokerData) = (*pokerData)[:pos] // 删除扑克 1244 delete 2 --> 144
	return true
}

// CompareCard 对比扑克 PS: copy PDK 因为不知道有的跑得快都有那些规则 在明确后可以考虑修改
func (g *Logic) CompareCard(firstList []int32, nextList []int32, firstCount int, nextCount int) bool {

	// 获取类型
	nextType, _ := g.GetPokerType(nextList)
	firstType, _ := g.GetPokerType(firstList)

	// 类型判断
	if firstType == pdk.CtError {
		return false
	}

	// 3A炸弹
	if firstType == pdk.Ct3ABomb {
		return true
	} else if nextType == pdk.Ct3ABomb {
		return false
	}

	// 王炸
	if firstType == pdk.CtKingBomb {
		return true
	} else if nextType == pdk.CtKingBomb {
		return false
	}

	// 炸弹判断
	if firstType == pdk.CtBomb && nextType != pdk.CtBomb {
		return true
	}
	if firstType != pdk.CtBomb && nextType == pdk.CtBomb {
		return false
	}

	if firstType == pdk.CtSingleLine && nextType == pdk.CtStraightFlush {
		return false
	} else if firstType == pdk.CtStraightFlush && nextType == pdk.CtSingleLine {
		return true
	} else if firstType == pdk.CtStraightFlush && nextType == pdk.CtStraightFlush {
		nextLogicValue := g.GetPokerLogicValue(nextList[0])
		firstLogicValue := g.GetPokerLogicValue(firstList[0])
		return firstLogicValue > nextLogicValue
	}

	// 规则判断
	if firstType != nextType || (firstCount != nextCount && firstType != pdk.CtBomb && nextType != pdk.CtBomb) {
		if firstType == pdk.CtThreeTakePair && nextType == pdk.CtThreeTakeDouble {

		} else if (g.rule&pdk.Py3Takesingle == 0 || g.threeFour) && firstType == pdk.CtThreeTakeDouble && nextType == pdk.CtThreeTakePair {

		} else if firstType == pdk.CtFourTakePair && nextType == pdk.CtFourTakeDouble {

		} else if (g.rule&pdk.Py4TakeOne == 0 || g.threeFour) && firstType == pdk.CtFourTakeDouble && nextType == pdk.CtFourTakePair {

		} else {
			return false
		}
	}

	// 开始对比
	switch nextType {
	case pdk.CtSingle, pdk.CtDouble, pdk.CtSingleLine, pdk.CtDoubleLine:
		nextLogicValue := g.GetPokerLogicValue(nextList[0])
		firstLogicValue := g.GetPokerLogicValue(firstList[0])
		return firstLogicValue > nextLogicValue
	case pdk.CtThree, pdk.CtThreeTakeSingle, pdk.CtThreeTakeDouble, pdk.CtThreeLine, pdk.CtThreeTakePair:
		var nextResult AnalyseResult
		var firstResult AnalyseResult
		g.AnalysebCardData(nextList, nextCount, &nextResult)
		g.AnalysebCardData(firstList, firstCount, &firstResult)
		return firstResult.ThreeLogicVolue[0] > nextResult.ThreeLogicVolue[0]
	case pdk.CtFourTakePair, pdk.CtFourTakeDoublePair, pdk.CtFourTakeSingle, pdk.CtFourTakeDouble:
		// 需要测试
		var nextResult AnalyseResult
		var firstResult AnalyseResult
		g.AnalysebCardData(nextList, nextCount, &nextResult)
		g.AnalysebCardData(firstList, firstCount, &firstResult)
		return firstResult.FourLogicVolue[0] > nextResult.FourLogicVolue[0]
	case pdk.CtBomb:
		if firstCount != nextCount {
			return firstCount > nextCount
		}
		nextLogicValue := g.GetPokerLogicValue(nextList[0])
		firstLogicValue := g.GetPokerLogicValue(firstList[0])
		return firstLogicValue > nextLogicValue
	}

	return false
}

/*
  提示问题， 需求，最好将所有提示都完整提示处理，并要尽量提高复用，方便后期维护和新增功能
  PS： 暂时考虑将飞机摘除，因为飞机可能存在带单又带对情况
*/

// realSingleHint 单牌提示
func (g *Logic) realSingleHint(handPoker []int32, handPokerCount int32) []int32 {
	var singleTemp []int32
	singleTemp = append(singleTemp, handPoker[0])
	for j := int32(0); j < handPokerCount-1; j++ {
		if g.GetPokerValue(handPoker[j]) == g.GetPokerValue(handPoker[j+1]) {
			continue
		}
		singleTemp = append(singleTemp, handPoker[j+1])
	}
	return singleTemp
}

// attachSingle 带单牌, poker 主牌型牌值，singleTemp 单牌数量，pokerLen 牌型长度， ply 一个牌型带几张 考虑所有牌型提示
func (g *Logic) attachSingle(pokerValue []int32, singleTemps []int32, pokerLen int, ply int) (singPokers [][]int32) {
	var sing []int32
	var singleTemp []int32
	singleTemp = append(singleTemp, singleTemps...)
	for i := 0; i != pokerLen; i++ {

		for m := 0; m != len(singleTemp); m++ { // PS: 本来考虑dfs深度递归变量， 但在考虑二次迭代最多不到百次循环时候不必考虑算法优化，飞机多情况暂不考虑
			if ply == 2 { // 两单 暂时考虑
				for n := m + 1; n < len(singleTemp); n++ {
					if g.GetPokerValue(singleTemp[m]) == pokerValue[i] || g.GetPokerValue(singleTemp[n]) == pokerValue[i] {
						continue
					}
					sing = append(sing, singleTemp[m])
					sing = append(sing, singleTemp[n])
					// g.RemovePoker(sing, 2, &singleTemp, len(singleTemp))
					//g.attachSingle(pokerValue[pokerLen+1:], singleTemp[:], pokerLen-1, ply) // PS: 尝试使用递归循环如飞机带多单情况 飞机情况复杂，暂不考虑
					singPokers = append(singPokers, sing)
					sing = sing[:0]
				}
			} else {
				if g.GetPokerValue(singleTemp[m]) == pokerValue[i] {
					continue
				}
				sing = append(sing, singleTemp[m])
				singPokers = append(singPokers, sing)
				sing = sing[:0]
			}
		}
	}
	return
}

// NewLogic new
func NewLogic() *Logic {
	gameLogic := Logic{
		log: utils.NewLogger(),
	}

	gameLogic.Init()

	return &gameLogic
}
