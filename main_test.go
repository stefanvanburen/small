package main

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestTransform(t *testing.T) {
	tests := map[string]struct {
		trans    transform
		given    string
		expected string
	}{
		"boldSansStr": {
			sansBold,
			"pack my box with five dozen liquor jugs",
			"𝗽𝗮𝗰𝗸 𝗺𝘆 𝗯𝗼𝘅 𝘄𝗶𝘁𝗵 𝗳𝗶𝘃𝗲 𝗱𝗼𝘇𝗲𝗻 𝗹𝗶𝗾𝘂𝗼𝗿 𝗷𝘂𝗴𝘀",
		},
		"italicSerif": {
			italicSerif,
			"pack my box with five dozen liquor jugs",
			"𝑝𝑎𝑐𝑘 𝑚𝑦 𝑏𝑜𝑥 𝑤𝑖𝑡ℎ 𝑓𝑖𝑣𝑒 𝑑𝑜𝑧𝑒𝑛 𝑙𝑖𝑞𝑢𝑜𝑟 𝑗𝑢𝑔𝑠",
		},
		"italicSans": {
			sansItalic,
			"pack my box with five dozen liquor jugs",
			"𝘱𝘢𝘤𝘬 𝘮𝘺 𝘣𝘰𝘹 𝘸𝘪𝘵𝘩 𝘧𝘪𝘷𝘦 𝘥𝘰𝘻𝘦𝘯 𝘭𝘪𝘲𝘶𝘰𝘳 𝘫𝘶𝘨𝘴",
		},
		"boldItalicSerif": {
			boldItalicSerif,
			"pack my box with five dozen liquor jugs",
			"𝒑𝒂𝒄𝒌 𝒎𝒚 𝒃𝒐𝒙 𝒘𝒊𝒕𝒉 𝒇𝒊𝒗𝒆 𝒅𝒐𝒛𝒆𝒏 𝒍𝒊𝒒𝒖𝒐𝒓 𝒋𝒖𝒈𝒔",
		},
		"boldItalicSans": {
			sansBoldItalic,
			"pack my box with five dozen liquor jugs",
			"𝙥𝙖𝙘𝙠 𝙢𝙮 𝙗𝙤𝙭 𝙬𝙞𝙩𝙝 𝙛𝙞𝙫𝙚 𝙙𝙤𝙯𝙚𝙣 𝙡𝙞𝙦𝙪𝙤𝙧 𝙟𝙪𝙜𝙨",
		},
		"boldSansCaps": {
			sansBold,
			"pack my box with five dozen liquor jugs",
			"𝗣𝗔𝗖𝗞 𝗠𝗬 𝗕𝗢𝗫 𝗪𝗜𝗧𝗛 𝗙𝗜𝗩𝗘 𝗗𝗢𝗭𝗘𝗡 𝗟𝗜𝗤𝗨𝗢𝗥 𝗝𝗨𝗚𝗦",
		},
		"italicSerifCaps": {
			italicSerif,
			"pack my box with five dozen liquor jugs",
			"𝑃𝐴𝐶𝐾 𝑀𝑌 𝐵𝑂𝑋 𝑊𝐼𝑇𝐻 𝐹𝐼𝑉𝐸 𝐷𝑂𝑍𝐸𝑁 𝐿𝐼𝑄𝑈𝑂𝑅 𝐽𝑈𝐺𝑆",
		},
		"italicSansCaps": {
			sansItalic,
			"pack my box with five dozen liquor jugs",
			"𝘗𝘈𝘊𝘒 𝘔𝘠 𝘉𝘖𝘟 𝘞𝘐𝘛𝘏 𝘍𝘐𝘝𝘌 𝘋𝘖𝘡𝘌𝘕 𝘓𝘐𝘘𝘜𝘖𝘙 𝘑𝘜𝘎𝘚",
		},
		"boldItalicSerifCaps": {
			boldItalicSerif,
			"pack my box with five dozen liquor jugs",
			"𝑷𝑨𝑪𝑲 𝑴𝒀 𝑩𝑶𝑿 𝑾𝑰𝑻𝑯 𝑭𝑰𝑽𝑬 𝑫𝑶𝒁𝑬𝑵 𝑳𝑰𝑸𝑼𝑶𝑹 𝑱𝑼𝑮𝑺",
		},
		"boldItalicSansCaps": {
			sansBoldItalic,
			"pack my box with five dozen liquor jugs",
			"𝙋𝘼𝘾𝙆 𝙈𝙔 𝘽𝙊𝙓 𝙒𝙄𝙏𝙃 𝙁𝙄𝙑𝙀 𝘿𝙊𝙕𝙀𝙉 𝙇𝙄𝙌𝙐𝙊𝙍 𝙅𝙐𝙂𝙎",
		},
		"boldSerif": {
			boldSerif,
			"pack my box with five dozen liquor jugs",
			"𝐩𝐚𝐜𝐤 𝐦𝐲 𝐛𝐨𝐱 𝐰𝐢𝐭𝐡 𝐟𝐢𝐯𝐞 𝐝𝐨𝐳𝐞𝐧 𝐥𝐢𝐪𝐮𝐨𝐫 𝐣𝐮𝐠𝐬",
		},
		"boldSerifCaps": {
			boldSerif,
			"pack my box with five dozen liquor jugs",
			"𝐏𝐀𝐂𝐊 𝐌𝐘 𝐁𝐎𝐗 𝐖𝐈𝐓𝐇 𝐅𝐈𝐕𝐄 𝐃𝐎𝐙𝐄𝐍 𝐋𝐈𝐐𝐔𝐎𝐑 𝐉𝐔𝐆𝐒",
		},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			out := performTransform(tc.trans, tc.given)
			c := qt.New(t)
			c.Assert(tc.expected, qt.DeepEquals, out)
		})
	}

}
