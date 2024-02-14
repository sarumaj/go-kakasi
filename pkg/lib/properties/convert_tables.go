package properties

// convertTables is a set of conversion tables.
// It is used to generate symbol tables.
// It is also used to generate alphabet tables.
// It is based on Original KAKASI's EUC_JP - alphabet converter table.
var ConvertTables = convertTables{}

type convertTables struct{}

func (convertTables) generateSymbolTable(lo, hi int) []rune {
	var out []rune
	for i := lo; i <= hi; i++ {
		out = append(out, rune(i))
	}

	return out
}

// U3000 - 301F
// 　、。〃〄々〆〇〈〉《》「」『』【】〒〓〔〕〖〗〘〙〚〛〜〝〞〟
// Ideographic space and punctuation
func (c convertTables) SymbolTable1() []rune {
	return c.generateSymbolTable(0x3000, 0x301F)
}

// U3030 - 3040
// 〰〱〲〳〴〵〶〷〸〹〺〻〼〽〾〿぀
// Wavy dash and punctuation
func (c convertTables) SymbolTable2() []rune {
	return c.generateSymbolTable(0x3030, 0x3040)
}

// U0391 - 03A9
// ΑΒΓΔΕΖΗΘΙΚΛΜΝΞΟΠΡΣΤΥΦΧΨΩ
// Greek capital letters
func (c convertTables) SymbolTable3() []rune {
	return c.generateSymbolTable(0x0391, 0x03A9)
}

// U03B1 - 03C9
// αβγδεζηθικλμνξοπρστυφχψω
// Greek small letters
func (c convertTables) SymbolTable4() []rune {
	return c.generateSymbolTable(0x03B1, 0x03C9)
}

// UFF01-FF0F
// ！＂＃＄％＆＇（）＊＋，－．／
func (c convertTables) SymbolTable5() []rune {
	return c.generateSymbolTable(0xFF01, 0xFF0F)
}

// Cyrillic
// АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ
// абвгдеёжзийклмнопрстуфхцчшщъыьэюя
func (c convertTables) CyrillicTable() []rune {
	return c.generateSymbolTable(0x0410, 0x044F)
}

// Alpha1
// 　！＂＃＄％＆＇（）＊＋，－．／０１２３４５６７８９：；＜＝＞？＠
// Ideographic space and punctuation
func (c convertTables) AlphaTable1() []rune {
	return append([]rune{0x3000}, c.generateSymbolTable(0xFF01, 0xFF20)...)
}

// Alpha2
// ［＼］＾＿｀
func (c convertTables) AlphaTable2() []rune {
	return c.generateSymbolTable(0xFF3B, 0xFF40)
}

// Alpha3
// ｛｜｝～
func (c convertTables) AlphaTable3() []rune {
	return c.generateSymbolTable(0xFF5B, 0xFF5E)
}

// Latin1
// ¡¢£¤¥¦§¨©ª«¬­®¯
// °±²³´µ¶·¸¹º»¼½¾¿
// ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏ
// ÐÑÒÓÔÕÖ×ØÙÚÛÜÝÞß
// àáâãäåæçèéêëìíîï
// ðñòóôõö÷øùúûüýþÿ
func (c convertTables) Latin1Table() []rune {
	return c.generateSymbolTable(0x00A1, 0x00FF)
}
