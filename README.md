# small

small is a ... small tool to translate text to alphabetical unicode variants.

By default, the translation is to ꜱᴍᴀʟʟ ᴄᴀᴘꜱ, hence the name.

```console
$ go install github.com/stefanvanburen/small/cmd/small@latest

$ small "INFORMATION"
ɪɴꜰᴏʀᴍᴀɪᴛᴏɴ
```

But other translations are supported, and more are on the way!

```console
$ small -l
Supported transformations:
  smallcaps
  serif-bold
  serif-italic
  serif-bold-italic
  sans
  sans-bold
  sans-italic
  sans-bold-italic
  script
  fraktur
  fraktur-bold
  doublestruck
  monospace


$ small -t fraktur "american gothic"
𝔞𝔪𝔢𝔯𝔦𝔠𝔞𝔫 𝔤𝔬𝔱𝔥𝔦𝔠
```


It also behaves well as a pipe:

```console
$ echo "here's some text" | small -t sans
𝗁𝖾𝗋𝖾'𝗌 𝗌𝗈𝗆𝖾 𝗍𝖾𝗑𝗍
```
