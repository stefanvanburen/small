// Package small performs alphabetical translations from ascii to various
// unicode variants.
package small

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

type Transform map[string]string

func PerformTransform(t Transform, input io.Reader, output io.Writer) error {
	bufReader := bufio.NewReader(input)
	bufWriter := bufio.NewWriter(output)
	defer bufWriter.Flush()

	for {
		c, err := bufReader.ReadByte()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}

			return err
		}

		if val, ok := t[string(c)]; ok {
			if _, err := bufWriter.WriteString(val); err != nil {
				return err
			}
		} else {
			if err := bufWriter.WriteByte(c); err != nil {
				return err
			}
		}
	}
}

func SupportedTransformations() map[string]Transform {
	return map[string]Transform{
		"smallcaps":         smallcaps,
		"serif-bold":        boldSerif,
		"serif-italic":      italicSerif,
		"serif-bold-italic": boldItalicSerif,
		"sans":              sans,
		"sans-bold":         boldSans,
		"sans-italic":       italicSans,
		"sans-bold-italic":  boldItalicSans,
		"script":            script,
		"fraktur":           fraktur,
		"fraktur-bold":      frakturBold,
		"doublestruck":      doublestruck,
		"monospace":         monospace,
	}
}

func GetTransform(name string) (Transform, error) {
	switch name {
	case "", "smallcaps":
		return smallcaps, nil
	case "serif-bold":
		return boldSerif, nil
	case "serif-italic":
		return italicSerif, nil
	case "serif-bold-italic":
		return boldItalicSerif, nil
	case "sans":
		return sans, nil
	case "sans-bold":
		return boldSans, nil
	case "sans-italic":
		return italicSans, nil
	case "sans-bold-italic":
		return boldItalicSans, nil
	case "script":
		return script, nil
	case "fraktur":
		return fraktur, nil
	case "fraktur-bold":
		return frakturBold, nil
	case "doublestruck":
		return doublestruck, nil
	case "monospace":
		return monospace, nil
	default:
		return nil, fmt.Errorf("invalid transform: %s", name)
	}
}

var smallcaps = Transform{
	"A": "ᴀ",
	"B": "ʙ",
	"C": "ᴄ",
	"D": "ᴅ",
	"E": "ᴇ",
	"F": "ꜰ",
	"G": "ɢ",
	"H": "ʜ",
	"I": "ɪ",
	"J": "ᴊ",
	"K": "ᴋ",
	"L": "ʟ",
	"M": "ᴍ",
	"N": "ɴ",
	"O": "ᴏ",
	"P": "ᴘ",
	"Q": "ꞯ",
	"R": "ʀ",
	"S": "ꜱ",
	"T": "ᴛ",
	"U": "ᴜ",
	"V": "ᴠ",
	"W": "ᴡ",
	"X": "X", // not transformed as no Unicode smallcaps
	"Y": "ʏ",
	"Z": "ᴢ",
	"a": "ᴀ",
	"b": "ʙ",
	"c": "ᴄ",
	"d": "ᴅ",
	"e": "ᴇ",
	"f": "ꜰ",
	"g": "ɢ",
	"h": "ʜ",
	"i": "ɪ",
	"j": "ᴊ",
	"k": "ᴋ",
	"l": "ʟ",
	"m": "ᴍ",
	"n": "ɴ",
	"o": "ᴏ",
	"p": "ᴘ",
	"q": "ꞯ",
	"r": "ʀ",
	"s": "ꜱ",
	"t": "ᴛ",
	"u": "ᴜ",
	"v": "ᴠ",
	"w": "ᴡ",
	"x": "x", // not transformed as no Unicode smallcaps
	"y": "ʏ",
	"z": "ᴢ",
}

var boldSerif = Transform{
	"A": "𝐀",
	"B": "𝐁",
	"C": "𝐂",
	"D": "𝐃",
	"E": "𝐄",
	"F": "𝐅",
	"G": "𝐆",
	"H": "𝐇",
	"I": "𝐈",
	"J": "𝐉",
	"K": "𝐊",
	"L": "𝐋",
	"M": "𝐌",
	"N": "𝐍",
	"O": "𝐎",
	"P": "𝐏",
	"Q": "𝐐",
	"R": "𝐑",
	"S": "𝐒",
	"T": "𝐓",
	"U": "𝐔",
	"V": "𝐕",
	"W": "𝐖",
	"X": "𝐗",
	"Y": "𝐘",
	"Z": "𝐙",
	"a": "𝐚",
	"b": "𝐛",
	"c": "𝐜",
	"d": "𝐝",
	"e": "𝐞",
	"f": "𝐟",
	"g": "𝐠",
	"h": "𝐡",
	"i": "𝐢",
	"j": "𝐣",
	"k": "𝐤",
	"l": "𝐥",
	"m": "𝐦",
	"n": "𝐧",
	"o": "𝐨",
	"p": "𝐩",
	"q": "𝐪",
	"r": "𝐫",
	"s": "𝐬",
	"t": "𝐭",
	"u": "𝐮",
	"v": "𝐯",
	"w": "𝐰",
	"x": "𝐱",
	"y": "𝐲",
	"z": "𝐳",
}

var italicSerif = Transform{
	"A": "𝐴",
	"B": "𝐵",
	"C": "𝐶",
	"D": "𝐷",
	"E": "𝐸",
	"F": "𝐹",
	"G": "𝐺",
	"H": "𝐻",
	"I": "𝐼",
	"J": "𝐽",
	"K": "𝐾",
	"L": "𝐿",
	"M": "𝑀",
	"N": "𝑁",
	"O": "𝑂",
	"P": "𝑃",
	"Q": "𝑄",
	"R": "𝑅",
	"S": "𝑆",
	"T": "𝑇",
	"U": "𝑈",
	"V": "𝑉",
	"W": "𝑊",
	"X": "𝑋",
	"Y": "𝑌",
	"Z": "𝑍",
	"a": "𝑎",
	"b": "𝑏",
	"c": "𝑐",
	"d": "𝑑",
	"e": "𝑒",
	"f": "𝑓",
	"g": "𝑔",
	"h": "ℎ",
	"i": "𝑖",
	"j": "𝑗",
	"k": "𝑘",
	"l": "𝑙",
	"m": "𝑚",
	"n": "𝑛",
	"o": "𝑜",
	"p": "𝑝",
	"q": "𝑞",
	"r": "𝑟",
	"s": "𝑠",
	"t": "𝑡",
	"u": "𝑢",
	"v": "𝑣",
	"w": "𝑤",
	"x": "𝑥",
	"y": "𝑦",
	"z": "𝑧",
}

var boldItalicSerif = Transform{
	"A": "𝑨",
	"B": "𝑩",
	"C": "𝑪",
	"D": "𝑫",
	"E": "𝑬",
	"F": "𝑭",
	"G": "𝑮",
	"H": "𝑯",
	"I": "𝑰",
	"J": "𝑱",
	"K": "𝑲",
	"L": "𝑳",
	"M": "𝑴",
	"N": "𝑵",
	"O": "𝑶",
	"P": "𝑷",
	"Q": "𝑸",
	"R": "𝑹",
	"S": "𝑺",
	"T": "𝑻",
	"U": "𝑼",
	"V": "𝑽",
	"W": "𝑾",
	"X": "𝑿",
	"Y": "𝒀",
	"Z": "𝒁",
	"a": "𝒂",
	"b": "𝒃",
	"c": "𝒄",
	"d": "𝒅",
	"e": "𝒆",
	"f": "𝒇",
	"g": "𝒈",
	"h": "𝒉",
	"i": "𝒊",
	"j": "𝒋",
	"k": "𝒌",
	"l": "𝒍",
	"m": "𝒎",
	"n": "𝒏",
	"o": "𝒐",
	"p": "𝒑",
	"q": "𝒒",
	"r": "𝒓",
	"s": "𝒔",
	"t": "𝒕",
	"u": "𝒖",
	"v": "𝒗",
	"w": "𝒘",
	"x": "𝒙",
	"y": "𝒚",
	"z": "𝒛",
}

var script = Transform{
	"A": "𝒜",
	"B": "ℬ",
	"C": "𝒞",
	"D": "𝒟",
	"E": "ℰ",
	"F": "ℱ",
	"G": "𝒢",
	"H": "ℋ",
	"I": "ℐ",
	"J": "𝒥",
	"K": "𝒦",
	"L": "ℒ",
	"M": "ℳ",
	"N": "𝒩",
	"O": "𝒪",
	"P": "𝒫",
	"Q": "𝒬",
	"R": "ℛ",
	"S": "𝒮",
	"T": "𝒯",
	"U": "𝒰",
	"V": "𝒱",
	"W": "𝒲",
	"X": "𝒳",
	"Y": "𝒴",
	"Z": "𝒵",
	"a": "𝒶",
	"b": "𝒷",
	"c": "𝒸",
	"d": "𝒹",
	"e": "ℯ",
	"f": "𝒻",
	"g": "ℊ",
	"h": "𝒽",
	"i": "𝒾",
	"j": "𝒿",
	"k": "𝓀",
	"l": "𝓁",
	"m": "𝓂",
	"n": "𝓃",
	"o": "ℴ",
	"p": "𝓅",
	"q": "𝓆",
	"r": "𝓇",
	"s": "𝓈",
	"t": "𝓉",
	"u": "𝓊",
	"v": "𝓋",
	"w": "𝓌",
	"x": "𝓍",
	"y": "𝓎",
	"z": "𝓏",
}

var fraktur = Transform{
	"A": "𝔄",
	"B": "𝔅",
	"C": "ℭ",
	"D": "𝔇",
	"E": "𝔈",
	"F": "𝔉",
	"G": "𝔊",
	"H": "ℌ",
	"I": "ℑ",
	"J": "𝔍",
	"K": "𝔎",
	"L": "𝔏",
	"M": "𝔐",
	"N": "𝔑",
	"O": "𝔒",
	"P": "𝔓",
	"Q": "𝔔",
	"R": "ℜ",
	"S": "𝔖",
	"T": "𝔗",
	"U": "𝔘",
	"V": "𝔙",
	"W": "𝔚",
	"X": "𝔛",
	"Y": "𝔜",
	"Z": "ℨ",
	"a": "𝔞",
	"b": "𝔟",
	"c": "𝔠",
	"d": "𝔡",
	"e": "𝔢",
	"f": "𝔣",
	"g": "𝔤",
	"h": "𝔥",
	"i": "𝔦",
	"j": "𝔧",
	"k": "𝔨",
	"l": "𝔩",
	"m": "𝔪",
	"n": "𝔫",
	"o": "𝔬",
	"p": "𝔭",
	"q": "𝔮",
	"r": "𝔯",
	"s": "𝔰",
	"t": "𝔱",
	"u": "𝔲",
	"v": "𝔳",
	"w": "𝔴",
	"x": "𝔵",
	"y": "𝔶",
	"z": "𝔷",
}

var doublestruck = Transform{
	"A": "𝔸",
	"B": "𝔹",
	"C": "ℂ",
	"D": "𝔻",
	"E": "𝔼",
	"F": "𝔽",
	"G": "𝔾",
	"H": "ℍ",
	"I": "𝕀",
	"J": "𝕁",
	"K": "𝕂",
	"L": "𝕃",
	"M": "𝕄",
	"N": "ℕ",
	"O": "𝕆",
	"P": "ℙ",
	"Q": "ℚ",
	"R": "ℝ",
	"S": "𝕊",
	"T": "𝕋",
	"U": "𝕌",
	"V": "𝕍",
	"W": "𝕎",
	"X": "𝕏",
	"Y": "𝕐",
	"Z": "ℤ",
	"a": "𝕒",
	"b": "𝕓",
	"c": "𝕔",
	"d": "𝕕",
	"e": "𝕖",
	"f": "𝕗",
	"g": "𝕘",
	"h": "𝕙",
	"i": "𝕚",
	"j": "𝕛",
	"k": "𝕜",
	"l": "𝕝",
	"m": "𝕞",
	"n": "𝕟",
	"o": "𝕠",
	"p": "𝕡",
	"q": "𝕢",
	"r": "𝕣",
	"s": "𝕤",
	"t": "𝕥",
	"u": "𝕦",
	"v": "𝕧",
	"w": "𝕨",
	"x": "𝕩",
	"y": "𝕪",
	"z": "𝕫",
}

var frakturBold = Transform{
	"A": "𝕬",
	"B": "𝕭",
	"C": "𝕮",
	"D": "𝕯",
	"E": "𝕰",
	"F": "𝕱",
	"G": "𝕲",
	"H": "𝕳",
	"I": "𝕴",
	"J": "𝕵",
	"K": "𝕶",
	"L": "𝕷",
	"M": "𝕸",
	"N": "𝕹",
	"O": "𝕺",
	"P": "𝕻",
	"Q": "𝕼",
	"R": "𝕽",
	"S": "𝕾",
	"T": "𝕿",
	"U": "𝖀",
	"V": "𝖁",
	"W": "𝖂",
	"X": "𝖃",
	"Y": "𝖄",
	"Z": "𝖅",
	"a": "𝖆",
	"b": "𝖇",
	"c": "𝖈",
	"d": "𝖉",
	"e": "𝖊",
	"f": "𝖋",
	"g": "𝖌",
	"h": "𝖍",
	"i": "𝖎",
	"j": "𝖏",
	"k": "𝖐",
	"l": "𝖑",
	"m": "𝖒",
	"n": "𝖓",
	"o": "𝖔",
	"p": "𝖕",
	"q": "𝖖",
	"r": "𝖗",
	"s": "𝖘",
	"t": "𝖙",
	"u": "𝖚",
	"v": "𝖛",
	"w": "𝖜",
	"x": "𝖝",
	"y": "𝖞",
	"z": "𝖟",
}

var sans = Transform{
	"A": "𝖠",
	"B": "𝖡",
	"C": "𝖢",
	"D": "𝖣",
	"E": "𝖤",
	"F": "𝖥",
	"G": "𝖦",
	"H": "𝖧",
	"I": "𝖨",
	"J": "𝖩",
	"K": "𝖪",
	"L": "𝖫",
	"M": "𝖬",
	"N": "𝖭",
	"O": "𝖮",
	"P": "𝖯",
	"Q": "𝖰",
	"R": "𝖱",
	"S": "𝖲",
	"T": "𝖳",
	"U": "𝖴",
	"V": "𝖵",
	"W": "𝖶",
	"X": "𝖷",
	"Y": "𝖸",
	"Z": "𝖹",
	"a": "𝖺",
	"b": "𝖻",
	"c": "𝖼",
	"d": "𝖽",
	"e": "𝖾",
	"f": "𝖿",
	"g": "𝗀",
	"h": "𝗁",
	"i": "𝗂",
	"j": "𝗃",
	"k": "𝗄",
	"l": "𝗅",
	"m": "𝗆",
	"n": "𝗇",
	"o": "𝗈",
	"p": "𝗉",
	"q": "𝗊",
	"r": "𝗋",
	"s": "𝗌",
	"t": "𝗍",
	"u": "𝗎",
	"v": "𝗏",
	"w": "𝗐",
	"x": "𝗑",
	"y": "𝗒",
	"z": "𝗓",
}

var boldSans = Transform{
	"A": "𝗔",
	"B": "𝗕",
	"C": "𝗖",
	"D": "𝗗",
	"E": "𝗘",
	"F": "𝗙",
	"G": "𝗚",
	"H": "𝗛",
	"I": "𝗜",
	"J": "𝗝",
	"K": "𝗞",
	"L": "𝗟",
	"M": "𝗠",
	"N": "𝗡",
	"O": "𝗢",
	"P": "𝗣",
	"Q": "𝗤",
	"R": "𝗥",
	"S": "𝗦",
	"T": "𝗧",
	"U": "𝗨",
	"V": "𝗩",
	"W": "𝗪",
	"X": "𝗫",
	"Y": "𝗬",
	"Z": "𝗭",
	"a": "𝗮",
	"b": "𝗯",
	"c": "𝗰",
	"d": "𝗱",
	"e": "𝗲",
	"f": "𝗳",
	"g": "𝗴",
	"h": "𝗵",
	"i": "𝗶",
	"j": "𝗷",
	"k": "𝗸",
	"l": "𝗹",
	"m": "𝗺",
	"n": "𝗻",
	"o": "𝗼",
	"p": "𝗽",
	"q": "𝗾",
	"r": "𝗿",
	"s": "𝘀",
	"t": "𝘁",
	"u": "𝘂",
	"v": "𝘃",
	"w": "𝘄",
	"x": "𝘅",
	"y": "𝘆",
	"z": "𝘇",
}

var italicSans = Transform{
	"A": "𝘈",
	"B": "𝘉",
	"C": "𝘊",
	"D": "𝘋",
	"E": "𝘌",
	"F": "𝘍",
	"G": "𝘎",
	"H": "𝘏",
	"I": "𝘐",
	"J": "𝘑",
	"K": "𝘒",
	"L": "𝘓",
	"M": "𝘔",
	"N": "𝘕",
	"O": "𝘖",
	"P": "𝘗",
	"Q": "𝘘",
	"R": "𝘙",
	"S": "𝘚",
	"T": "𝘛",
	"U": "𝘜",
	"V": "𝘝",
	"W": "𝘞",
	"X": "𝘟",
	"Y": "𝘠",
	"Z": "𝘡",
	"a": "𝘢",
	"b": "𝘣",
	"c": "𝘤",
	"d": "𝘥",
	"e": "𝘦",
	"f": "𝘧",
	"g": "𝘨",
	"h": "𝘩",
	"i": "𝘪",
	"j": "𝘫",
	"k": "𝘬",
	"l": "𝘭",
	"m": "𝘮",
	"n": "𝘯",
	"o": "𝘰",
	"p": "𝘱",
	"q": "𝘲",
	"r": "𝘳",
	"s": "𝘴",
	"t": "𝘵",
	"u": "𝘶",
	"v": "𝘷",
	"w": "𝘸",
	"x": "𝘹",
	"y": "𝘺",
	"z": "𝘻",
}

var boldItalicSans = Transform{
	"A": "𝘼",
	"B": "𝘽",
	"C": "𝘾",
	"D": "𝘿",
	"E": "𝙀",
	"F": "𝙁",
	"G": "𝙂",
	"H": "𝙃",
	"I": "𝙄",
	"J": "𝙅",
	"K": "𝙆",
	"L": "𝙇",
	"M": "𝙈",
	"N": "𝙉",
	"O": "𝙊",
	"P": "𝙋",
	"Q": "𝙌",
	"R": "𝙍",
	"S": "𝙎",
	"T": "𝙏",
	"U": "𝙐",
	"V": "𝙑",
	"W": "𝙒",
	"X": "𝙓",
	"Y": "𝙔",
	"Z": "𝙕",
	"a": "𝙖",
	"b": "𝙗",
	"c": "𝙘",
	"d": "𝙙",
	"e": "𝙚",
	"f": "𝙛",
	"g": "𝙜",
	"h": "𝙝",
	"i": "𝙞",
	"j": "𝙟",
	"k": "𝙠",
	"l": "𝙡",
	"m": "𝙢",
	"n": "𝙣",
	"o": "𝙤",
	"p": "𝙥",
	"q": "𝙦",
	"r": "𝙧",
	"s": "𝙨",
	"t": "𝙩",
	"u": "𝙪",
	"v": "𝙫",
	"w": "𝙬",
	"x": "𝙭",
	"y": "𝙮",
	"z": "𝙯",
}

var monospace = Transform{
	"A": "𝙰",
	"B": "𝙱",
	"C": "𝙲",
	"D": "𝙳",
	"E": "𝙴",
	"F": "𝙵",
	"G": "𝙶",
	"H": "𝙷",
	"I": "𝙸",
	"J": "𝙹",
	"K": "𝙺",
	"L": "𝙻",
	"M": "𝙼",
	"N": "𝙽",
	"O": "𝙾",
	"P": "𝙿",
	"Q": "𝚀",
	"R": "𝚁",
	"S": "𝚂",
	"T": "𝚃",
	"U": "𝚄",
	"V": "𝚅",
	"W": "𝚆",
	"X": "𝚇",
	"Y": "𝚈",
	"Z": "𝚉",
	"a": "𝚊",
	"b": "𝚋",
	"c": "𝚌",
	"d": "𝚍",
	"e": "𝚎",
	"f": "𝚏",
	"g": "𝚐",
	"h": "𝚑",
	"i": "𝚒",
	"j": "𝚓",
	"k": "𝚔",
	"l": "𝚕",
	"m": "𝚖",
	"n": "𝚗",
	"o": "𝚘",
	"p": "𝚙",
	"q": "𝚚",
	"r": "𝚛",
	"s": "𝚜",
	"t": "𝚝",
	"u": "𝚞",
	"v": "𝚟",
	"w": "𝚠",
	"x": "𝚡",
	"y": "𝚢",
	"z": "𝚣",
}