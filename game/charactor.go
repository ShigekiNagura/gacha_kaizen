package game

import (
	"math/rand"
	"time"
)

type Charactor struct {
	Name      string
	Rarity    int
	Image     string
	SoundFile string
}

func (c *Charactor) RaritytoStar() string {
	var star string
	for i := 0; i < c.Rarity; i++ {
		star += "★"
	}
	return star
}

var dog = Charactor{
	Name:   "イヌ",
	Rarity: 1,
	Image: `
__    __
o-''))_____\\
"--__/ * * * )
c_c__/-c____/
`,
	SoundFile: "ken",
}

var cat = Charactor{
	Name:   "ネコ",
	Rarity: 10,
	Image: `
/\___/\
( o   o )
(  =^=  )   ニャー
(        )
(         )
___v__^__v___
/             \
|               |
|               |
\             /
""""""""""""""
`,
	SoundFile: "door_lock",
}

var shiro = Charactor{
	Name:   "しろ",
	Rarity: 3,
	Image: `
	　　　　　　_ -‐──-ﾆ二ー-､
	　 _／ ﾌ‐´　　　　　´｀ヽﾍ｀ー'
	　/　,//´｀　　　　　　 ・　 ',
	　ゝ' ,' 　・　　　　　　　　　 l
	　　　!　　　　　 ━　 ＿, -'
	　　　ヽ＿＿,-＝＝i´
	　　　　　　 ,'　　　　.',
	　　　／,'., '　　　　　 ',
	　　　',_.ﾆl 　 .l l　.l l. ﾉ
	　　　　　 ｀￣U─U´
`,
	SoundFile: "level_up",
}

var Charactors = map[int]Charactor{
	0: dog,
	1: cat,
	2: shiro,
}

func generateCharactor() Charactor {

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(3)

	return Charactors[n]
}
