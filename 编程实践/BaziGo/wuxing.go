package bazi

// 五行

// GetWuXingFromNumber  从数字获得五行名, 0-4
func GetWuXingFromNumber(nValue int) string {
	// {* 五行字符串，以通常的金木水火土为顺序 }
	// 这里没用五行相生或者相克来排列
	switch nValue {
	case 0:
		return "金"
	case 1:
		return "木"
	case 2:
		return "水"
	case 3:
		return "火"
	case 4:
		return "土"
	}
	return ""
}

// GetWuXingFromGan 获得某干的五行，0-4 对应 金木水火土
// 甲乙为木，丙丁为火，戊己为土，庚辛为金，壬癸为水，
func GetWuXingFromGan(pGan *TGan) *TWuXing {
	return pGan.ToWuXing()
}

// NewWuXing 创建五行
func NewWuXing(nValue int) *TWuXing {
	nValue %= 5
	wuxing := TWuXing(nValue)
	return &wuxing
}

// TWuXing 五行
type TWuXing int

// ToString 转换成可阅读的字符串
func (self *TWuXing) ToString() string {
	return self.String()
}

// ToInt 转换成int
func (self *TWuXing) ToInt() int {
	return self.Value()
}

// Value 转换成int
func (self *TWuXing) Value() int {
	return (int)(*self)
}

// String 转换成可阅读的字符串
func (self *TWuXing) String() string {
	return GetWuXingFromNumber(self.Value())
}
