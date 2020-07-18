package bazi

//  {* 纳音五行，与相邻一对六十干支对应}
// 甲子乙丑海中金丙寅丁卯炉中火戊辰己巳大林木
// 庚午辛未路旁土壬申癸酉剑锋金甲戌乙亥山头火
// 丙子丁丑涧下水戊寅己卯城头土庚辰辛巳白蜡金
// 壬午癸未杨柳木甲申乙酉井泉水丙戌丁亥屋上土
// 戊子己丑霹雳火庚寅辛卯松柏木壬辰癸巳长流水
// 甲午乙未砂中金丙申丁酉山下火戊戌己亥平地木
// 庚子辛丑壁上土壬寅癸卯金箔金甲辰乙巳覆灯火
// 丙午丁未天河水戊申己酉大驿土庚戌辛亥钗钏金
// 壬子癸丑桑柘木甲寅乙卯大溪水丙辰丁巳砂中土
// 戊午己未天上火庚申辛酉石榴木壬戌癸亥大海水

// GetNaYinFromNumber 从数字获得纳音名, 0-29
func GetNaYinFromNumber(nValue int) string {
	switch nValue {
	case 0:
		return "海中金"
	case 1:
		return "炉中火"
	case 2:
		return "大林木"

	case 3:
		return "路旁土"
	case 4:
		return "剑锋金"
	case 5:
		return "山头火"

	case 6:
		return "涧下水"
	case 7:
		return "城墙土"
	case 8:
		return "白蜡金"

	case 9:
		return "杨柳木"
	case 10:
		return "泉中水"
	case 11:
		return "屋上土"

	case 12:
		return "霹雷火"
	case 13:
		return "松柏木"
	case 14:
		return "长流水"

	case 15:
		return "沙中金"
	case 16:
		return "山下火"
	case 17:
		return "平地木"

	case 18:
		return "壁上土"
	case 19:
		return "金箔金"
	case 20:
		return "佛灯火"

	case 21:
		return "天河水"
	case 22:
		return "大驿土"
	case 23:
		return "钗钏金"

	case 24:
		return "桑柘木"
	case 25:
		return "大溪水"
	case 26:
		return "沙中土"

	case 27:
		return "天上火"
	case 28:
		return "石榴木"
	case 29:
		return "大海水"
	}

	return ""
}

// NewNaYin 纳音
func NewNaYin(nValue int) *TNaYin {
	nValue %= 30

	NaYin := TNaYin(nValue)
	return &NaYin
}

// TNaYin 纳音
type TNaYin int

// Value 转换成int
func (self *TNaYin) Value() int {
	return (int)(*self)
}

// String 转换成可阅读的字符串
func (self *TNaYin) String() string {
	return GetNaYinFromNumber(self.Value())
}
