package properties

// convertTables is a set of conversion tables.
// It is used to generate symbol tables.
// It is also used to generate alphabet tables.
// It is based on Original KAKASI's EUC_JP - alphabet converter table.
var ConvertTables = convertTables{}

type convertTables struct{}

// Alpha1
// 　！＂＃＄％＆＇（）＊＋，－．／０１２３４５６７８９：；＜＝＞？＠
// Ideographic space and punctuation
func (c convertTables) AlphaTable1() map[rune]string {
	return map[rune]string{
		' ':  "\u3000",
		'!':  "\uFF01",
		'"':  "\uFF02",
		'#':  "\uFF03",
		'$':  "\uFF04",
		'%':  "\uFF05",
		'&':  "\uFF06",
		'\'': "\uFF07",
		'(':  "\uFF08",
		')':  "\uFF09",
		'*':  "\uFF0A",
		'+':  "\uFF0B",
		',':  "\uFF0C",
		'-':  "\uFF0D",
		'.':  "\uFF0E",
		'/':  "\uFF0F",
		'0':  "\uFF10",
		'1':  "\uFF11",
		'2':  "\uFF12",
		'3':  "\uFF13",
		'4':  "\uFF14",
		'5':  "\uFF15",
		'6':  "\uFF16",
		'7':  "\uFF17",
		'8':  "\uFF18",
		'9':  "\uFF19",
		':':  "\uFF1A",
		';':  "\uFF1B",
		'<':  "\uFF1C",
		'=':  "\uFF1D",
		'>':  "\uFF1E",
		'?':  "\uFF1F",
		'@':  "\uFF20",
	}
}

// Alpha2
// ［＼］＾＿｀
func (c convertTables) AlphaTable2() map[rune]string {
	return map[rune]string{
		'[':  "\uFF3B",
		'\\': "\uFF3C",
		']':  "\uFF3D",
		'^':  "\uFF3E",
		'_':  "\uFF3F",
		'`':  "\uFF40",
	}
}

// Alpha3
// ｛｜｝～
func (c convertTables) AlphaTable3() map[rune]string {
	return map[rune]string{
		'{': "\uFF5B",
		'|': "\uFF5C",
		'}': "\uFF5D",
		'~': "\uFF5E",
	}
}

// Cyrillic
// АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ
// абвгдеёжзийклмнопрстуфхцчшщъыьэюяѐё
func (convertTables) CyrillicTable() map[rune]string {
	return map[rune]string{
		0x0400: "E",   // Ѐ
		0x0401: "E",   // Ё
		0x0402: "Dj",  // Ђ
		0x0403: "G",   // Ѓ
		0x0404: "E",   // Є
		0x0405: "Dz",  // Ѕ
		0x0406: "I",   // І
		0x0407: "Yi",  // Ї
		0x0408: "J",   // Ј
		0x0409: "Lj",  // Љ
		0x040A: "Nj",  // Њ
		0x040B: "Tsh", // Ћ
		0x040C: "Kj",  // Ќ
		0x040D: "I",   // Ѝ
		0x040E: "U",   // Ў
		0x040F: "Dz",  // Џ
		0x0410: "A",   // А
		0x0411: "B",   // Б
		0x0412: "V",   // В
		0x0413: "G",   // Г
		0x0414: "D",   // Д
		0x0415: "E",   // Е
		0x0416: "Zh",  // Ж
		0x0417: "Z",   // З
		0x0418: "I",   // И
		0x0419: "Y",   // Й
		0x041A: "K",   // К
		0x041B: "L",   // Л
		0x041C: "M",   // М
		0x041D: "N",   // Н
		0x041E: "O",   // О
		0x041F: "P",   // П
		0x0420: "R",   // Р
		0x0421: "S",   // С
		0x0422: "T",   // Т
		0x0423: "U",   // У
		0x0424: "F",   // Ф
		0x0425: "H",   // Х
		0x0426: "Ts",  // Ц
		0x0427: "Ch",  // Ч
		0x0428: "Sh",  // Ш
		0x0429: "Sch", // Щ
		0x042A: "",    // Ъ
		0x042B: "Y",   // Ы
		0x042C: "",    // Ь
		0x042D: "E",   // Э
		0x042E: "Yu",  // Ю
		0x042F: "Ya",  // Я
		0x0430: "a",   // а
		0x0431: "b",   // б
		0x0432: "v",   // в
		0x0433: "g",   // г
		0x0434: "d",   // д
		0x0435: "e",   // е
		0x0436: "zh",  // ж
		0x0437: "z",   // з
		0x0438: "i",   // и
		0x0439: "y",   // й
		0x043A: "k",   // к
		0x043B: "l",   // л
		0x043C: "m",   // м
		0x043D: "n",   // н
		0x043E: "o",   // о
		0x043F: "p",   // п
		0x0440: "r",   // р
		0x0441: "s",   // с
		0x0442: "t",   // т
		0x0443: "u",   // у
		0x0444: "f",   // ф
		0x0445: "h",   // х
		0x0446: "ts",  // ц
		0x0447: "ch",  // ч
		0x0448: "sh",  // ш
		0x0449: "sch", // щ
		0x044A: "",    // ъ
		0x044B: "y",   // ы
		0x044C: "",    // ь
		0x044D: "e",   // э
		0x044E: "yu",  // ю
		0x044F: "ya",  // я
		0x0450: "e",   // ѐ
		0x0451: "e",   // ё
	}
}

// Latin1
// ¡¢£¤¥¦§¨©ª«¬­®¯
// °±²³´µ¶·¸¹º»¼½¾¿
// ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏ
// ÐÑÒÓÔÕÖ×ØÙÚÛÜÝÞß
// àáâãäåæçèéêëìíîï
// ðñòóôõö÷øùúûüýþÿ
func (c convertTables) Latin1Table() map[rune]string {
	return map[rune]string{
		0x00A1: "!",        // inverted exclamation mark
		0x00A2: "cent",     // cent mark
		0x00A3: "GBP",      // pound mark
		0x00A4: "currency", // currency mark
		0x00A5: "yen",      // yen mark
		0x00A6: "|",        // broken bar
		0x00A7: "ss",       // section sign
		0x00A8: "..",       // diaeresis
		0x00A9: "(c)",      // copyright
		0x00AA: "a",        // feminine ordinal indicator
		0x00AB: "<<",       // left-pointing double angle quotation mark
		0x00AC: "not",      // not sign
		0x00AD: "-",        // soft hyphen
		0x00AE: "(R)",      // registered sign
		0x00AF: "~",        // macron
		0x00B0: ".",        // degree sign
		0x00B1: "+-",       // plus-minus sign
		0x00B2: "^2",       // superscript two
		0x00B3: "^3",       // superscript three
		0x00B4: "`",        // acute accent
		0x00B5: "u",        // micro sign
		0x00B6: "D",        // pilcrow sign
		0x00B7: ".",        // middle dot
		0x00B8: ",",        // cedilla
		0x00B9: "^1",       // superscript one
		0x00BA: "",         // masculine ordinal indicator
		0x00BB: ">>",       // right-pointing double angle quotation mark
		0x00BC: "1/4",      // vulgar fraction one quarter
		0x00BD: "1/2",      // vulgar fraction one half
		0x00BE: "3/4",      // vulgar fraction three quarters
		0x00BF: "?",        // inverted question mark
		0x00C0: "A",        // latin capital letter A with grave
		0x00C1: "A",        // latin capital letter A with acute
		0x00C2: "A",        // latin capital letter A with circumflex
		0x00C3: "A",        // latin capital letter A with tilde
		0x00C4: "A",        // latin capital letter A with diaeresis
		0x00C5: "A",        // latin capital letter A with ring above
		0x00C6: "AE",       // latin capital letter AE
		0x00C7: "C",        // latin capital letter C with cedilla
		0x00C8: "E",        // latin capital letter E with grave
		0x00C9: "E",        // latin capital letter E with acute
		0x00CA: "E",        // latin capital letter E with circumflex
		0x00CB: "E",        // latin capital letter E with diaeresis
		0x00CC: "I",        // latin capital letter I with grave
		0x00CD: "I",        // latin capital letter I with acute
		0x00CE: "I",        // latin capital letter I with circumflex
		0x00CF: "I",        // latin capital letter I with diaeresis
		0x00D0: "Eth",      // latin capital letter ETH
		0x00D1: "N",        // latin capital letter N with tilde
		0x00D2: "O",        // latin capital letter O with grave
		0x00D3: "O",        // latin capital letter O with acute
		0x00D4: "O",        // latin capital letter O with circumflex
		0x00D5: "O",        // latin capital letter O with tilde
		0x00D6: "O",        // latin capital letter O with diaeresis
		0x00D7: "x",        // multiplication sign
		0x00D8: "O",        // latin capital letter O with stroke
		0x00D9: "U",        // latin capital letter U with grave
		0x00DA: "U",        // latin capital letter U with acute
		0x00DB: "U",        // latin capital letter U with circumflex
		0x00DC: "U",        // latin capital letter U with diaeresis
		0x00DD: "Y",        // latin capital letter Y with acute
		0x00DE: "",         // latin capital letter THORN
		0x00DF: "s",        // latin small letter sharp s
		0x00E0: "a",        // latin small letter a with grave
		0x00E1: "a",        // latin small letter a with acute
		0x00E2: "a",        // latin small letter a with circumflex
		0x00E3: "a",        // latin small letter a with tilde
		0x00E4: "a",        // latin small letter a with diaeresis
		0x00E5: "a",        // latin small letter a with ring above
		0x00E6: "ae",       // latin small letter ae
		0x00E7: "c",        // latin small letter c with cedilla
		0x00E8: "e",        // latin small letter e with grave
		0x00E9: "e",        // latin small letter e with acute
		0x00EA: "e",        // latin small letter e with circumflex
		0x00EB: "e",        // latin small letter e with diaeresis
		0x00EC: "i",        // latin small letter i with grave
		0x00ED: "i",        // latin small letter i with acute
		0x00EE: "i",        // latin small letter i with circumflex
		0x00EF: "i",        // latin small letter i with diaeresis
		0x00F0: "eth",      // latin small letter eth
		0x00F1: "n",        // latin small letter n with tilde
		0x00F2: "o",        // latin small letter o with grave
		0x00F3: "o",        // latin small letter o with acute
		0x00F4: "o",        // latin small letter o with circumflex
		0x00F5: "o",        // latin small letter o with tilde
		0x00F6: "o",        // latin small letter o with diaeresis
		0x00F7: "/",        // division sign
		0x00F8: "o",        // latin small letter o with stroke
		0x00F9: "u",        // latin small letter u with grave
		0x00FA: "u",        // latin small letter u with acute
		0x00FB: "u",        // latin small letter u with circumflex
		0x00FC: "u",        // latin small letter u with diaeresis
		0x00FD: "y",        // latin small letter y with acute
		0x00FE: "",         // latin small letter thorn
		0x00FF: "y",        // latin small letter y with diaeresis
	}
}

// U3000 - 301F
// 　、。〃〄々〆〇〈〉《》「」『』【】〒〓〔〕〖〗〘〙〚〛〜〝〞〟
// Ideographic space and punctuation
func (convertTables) SymbolTable1() map[rune]string {
	return map[rune]string{
		0x3000: " ",
		0x3001: ",",
		0x3002: ".",
		0x3003: "\"",
		0x3004: "(kigou)",
		0x3005: "(kurikaesi)",
		0x3006: "(sime)",
		0x3007: "(maru)",
		0x3008: "<",
		0x3009: ">",
		0x300A: "<<",
		0x300B: ">>",
		0x300C: "(",
		0x300D: ")",
		0x300E: "(",
		0x300F: ")",
		0x3010: "(",
		0x3011: ")",
		0x3012: "(kigou)",
		0x3013: "(geta)",
		0x3014: "(",
		0x3015: ")",
		0x3016: "(",
		0x3017: ")",
		0x3018: "(",
		0x3019: ")",
		0x301A: "(",
		0x301B: ")",
		0x301C: "~",
		0x301D: "(kigou)",
		0x301E: "\"",
		0x301F: "(kigou)",
		0x3020: "(kigou)",
	}
}

// U3030 - 3040
// 〰〱〲〳〴〵〶〷〸〹〺〻〼〽〾〿぀
// Wavy dash and punctuation
func (convertTables) SymbolTable2() map[rune]string {
	return map[rune]string{
		0x3030: "-",
		0x3031: "(kurikaesi)",
		0x3032: "(kurikaesi)",
		0x3033: "(kurikaesi)",
		0x3034: "(kurikaesi)",
		0x3035: "(kurikaesi)",
		0x3036: "(kigou)",
		0x3037: "XX",
		0x3038: "",
		0x3039: "",
		0x303A: "",
		0x303B: "",
		0x303C: "(masu)",
		0x303D: "(kurikaesi)",
		0x303E: " ",
		0x303F: " ",
		0x3040: " ",
	}
}

// U0391 - 03A9
// ΑΒΓΔΕΖΗΘΙΚΛΜΝΞΟΠΡΣΤΥΦΧΨΩ
// Greek capital letters
func (convertTables) SymbolTable3() map[rune]string {
	return map[rune]string{
		0x0391: "Alpha",
		0x0392: "Beta",
		0x0393: "Gamma",
		0x0394: "Delta",
		0x0395: "Epsilon",
		0x0396: "Zeta",
		0x0397: "Eta",
		0x0398: "Theta",
		0x0399: "Iota",
		0x039A: "Kappa",
		0x039B: "Lambda",
		0x039C: "Mu",
		0x039D: "Nu",
		0x039E: "Xi",
		0x039F: "Omicron",
		0x03A0: "Pi",
		0x03A1: "Rho",
		0x03A2: "", // not used
		0x03A3: "Sigma",
		0x03A4: "Tau",
		0x03A5: "Upsilon",
		0x03A6: "Phi",
		0x03A7: "Chi",
		0x03A8: "Psi",
		0x03A9: "Omega",
	}
}

// U03B1 - 03C9
// αβγδεζηθικλμνξοπρστυφχψω
// Greek small letters
func (convertTables) SymbolTable4() map[rune]string {
	return map[rune]string{
		0x03B1: "alpha",
		0x03B2: "beta",
		0x03B3: "gamma",
		0x03B4: "delta",
		0x03B5: "epsilon",
		0x03B6: "zeta",
		0x03B7: "eta",
		0x03B8: "theta",
		0x03B9: "iota",
		0x03BA: "kappa",
		0x03BB: "lambda",
		0x03BC: "mu",
		0x03BD: "nu",
		0x03BE: "xi",
		0x03BF: "omicron",
		0x03C0: "pi",
		0x03C1: "rho",
		0x03C2: "final sigma",
		0x03C3: "sigma",
		0x03C4: "tau",
		0x03C5: "upsilon",
		0x03C6: "phi",
		0x03C7: "chi",
		0x03C8: "psi",
		0x03C9: "omega",
	}
}

// UFF01-FF0F
// ！＂＃＄％＆＇（）＊＋，－．／
func (convertTables) SymbolTable5() map[rune]string {
	return map[rune]string{
		0xFF01: "!",
		0xFF02: "\"",
		0xFF03: "#",
		0xFF04: "$",
		0xFF05: "%",
		0xFF06: "&",
		0xFF07: "'",
		0xFF08: "(",
		0xFF09: ")",
		0xFF0A: "*",
		0xFF0B: "+",
		0xFF0C: ",",
		0xFF0D: "-",
		0xFF0E: ".",
		0xFF0F: "/",
	}
}
