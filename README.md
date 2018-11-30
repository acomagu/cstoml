# TRPG キャラクターシートTOML変換コマンド

[キャラクター保管所](https://charasheet.vampire-blood.net)さんに保存されているキャラクターシートをTOMLで出力します。

## 例

```
$ curl 'https://charasheet.vampire-blood.net/mbdd0709d9532378710ceffc69301efe5.json' | cstoml
["プロフィール"]
"キャラクター名" = "すー"
"職業"           = "アパレルショップスタッフ"
"年齢"           = 25
"性別"           = "女性"
"髪の色"         = "明るい茶"
"身長"           = 150
"体重"           = 49

["技能"]
STR = 10
CON = 8
POW = 12
DEX = 13
APP = 15
SIZ = 12
INT = 10
EDU = 13

["能力値"]
"応急手当"      = 30
"鍵開け"        = 40
"芸術"          = 10
"経理"          = 10
"こぶし/パンチ" = 30
"心理学"        = 60
"制作"          = 30
"説得"          = 60
"変装"          = 30
"目星"          = 60
```

## インストール

```bash
$ go get github.com/acomagu/cstoml
```
