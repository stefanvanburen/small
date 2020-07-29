# small

small is a ... small tool to translate text to alphabetical unicode variants.
By default, the translation is to ꜱᴍᴀʟʟ ᴄᴀᴘꜱ, hence the name.
But other translations are supported, and more are on the way!

```commandline
$ go get github.com/svanburen/small/cmd/small

$ small "INFORMATION"
ɪɴꜰᴏʀᴍᴀɪᴛᴏɴ

$ small -transform fraktur "american gothic"
𝔞𝔪𝔢𝔯𝔦𝔠𝔞𝔫 𝔤𝔬𝔱𝔥𝔦𝔠
```
