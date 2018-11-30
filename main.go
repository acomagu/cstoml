package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"

	"golang.org/x/text/width"
)

type optionalInt int

func (v *optionalInt) UnmarshalJSON(bs []byte) error {
	var str string
	if err := json.Unmarshal(bs, &str); err != nil {
		return err
	}

	if len(str) == 0 {
		*v = 0
		return nil
	}

	n, err := strconv.Atoi(str)
	if err != nil {
		return err
	}

	*v = optionalInt(n)
	return nil
}

type CS struct {
	IsDispStatus      optionalInt   `json:"is_disp_status"`
	VNA1              optionalInt   `json:"V_NA1"`
	NA1               optionalInt   `json:"NA1"`
	VNA2              optionalInt   `json:"V_NA2"`
	NA2               optionalInt   `json:"NA2"`
	VNA3              optionalInt   `json:"V_NA3"`
	NA3               optionalInt   `json:"NA3"`
	VNA4              optionalInt   `json:"V_NA4"`
	NA4               optionalInt   `json:"NA4"`
	VNA5              optionalInt   `json:"V_NA5"`
	NA5               optionalInt   `json:"NA5"`
	VNA6              optionalInt   `json:"V_NA6"`
	NA6               optionalInt   `json:"NA6"`
	VNA7              optionalInt   `json:"V_NA7"`
	NA7               optionalInt   `json:"NA7"`
	VNA8              optionalInt   `json:"V_NA8"`
	NA8               optionalInt   `json:"NA8"`
	NA9               optionalInt   `json:"NA9"`
	NA10              optionalInt   `json:"NA10"`
	NA11              optionalInt   `json:"NA11"`
	NA12              optionalInt   `json:"NA12"`
	NA13              optionalInt   `json:"NA13"`
	NA14              optionalInt   `json:"NA14"`
	NS1               optionalInt   `json:"NS1"`
	NS2               optionalInt   `json:"NS2"`
	NS3               optionalInt   `json:"NS3"`
	NS4               optionalInt   `json:"NS4"`
	NS5               optionalInt   `json:"NS5"`
	NS6               optionalInt   `json:"NS6"`
	NS7               optionalInt   `json:"NS7"`
	NS8               optionalInt   `json:"NS8"`
	NS9               optionalInt   `json:"NS9"`
	NS10              optionalInt   `json:"NS10"`
	NS11              optionalInt   `json:"NS11"`
	NS12              optionalInt   `json:"NS12"`
	NS13              optionalInt   `json:"NS13"`
	NS14              optionalInt   `json:"NS14"`
	NM1               optionalInt   `json:"NM1"`
	NM2               optionalInt   `json:"NM2"`
	NM3               optionalInt   `json:"NM3"`
	NM4               optionalInt   `json:"NM4"`
	NM5               optionalInt   `json:"NM5"`
	NM6               optionalInt   `json:"NM6"`
	NM7               optionalInt   `json:"NM7"`
	NM8               optionalInt   `json:"NM8"`
	NM9               optionalInt   `json:"NM9"`
	NM10              optionalInt   `json:"NM10"`
	NM11              optionalInt   `json:"NM11"`
	NM12              optionalInt   `json:"NM12"`
	NM13              optionalInt   `json:"NM13"`
	NM14              optionalInt   `json:"NM14"`
	NP1               optionalInt   `json:"NP1"`
	NP2               optionalInt   `json:"NP2"`
	NP3               optionalInt   `json:"NP3"`
	NP4               optionalInt   `json:"NP4"`
	NP5               optionalInt   `json:"NP5"`
	NP6               optionalInt   `json:"NP6"`
	NP7               optionalInt   `json:"NP7"`
	NP8               optionalInt   `json:"NP8"`
	NP9               optionalInt   `json:"NP9"`
	NP10              optionalInt   `json:"NP10"`
	NP11              optionalInt   `json:"NP11"`
	NP12              optionalInt   `json:"NP12"`
	NP13              optionalInt   `json:"NP13"`
	NP14              optionalInt   `json:"NP14"`
	IsDispSan         optionalInt   `json:"is_disp_san"`
	SANLeft           optionalInt   `json:"SAN_Left"`
	SANMax            optionalInt   `json:"SAN_Max"`
	SANDanger         optionalInt   `json:"SAN_Danger"`
	IsDispPower       optionalInt   `json:"is_disp_power"`
	TSTotal           optionalInt   `json:"TS_Total"`
	TSMaximum         optionalInt   `json:"TS_Maximum"`
	TSAdd             optionalInt   `json:"TS_Add"`
	TKTotal           optionalInt   `json:"TK_Total"`
	TKMaximum         optionalInt   `json:"TK_Maximum"`
	TKAdd             optionalInt   `json:"TK_Add"`
	IsDispBattlearts  optionalInt   `json:"is_disp_battlearts"`
	TBAU              []optionalInt `json:"TBAU"`
	TBAD              []optionalInt `json:"TBAD"`
	TBAS              []optionalInt `json:"TBAS"`
	TBAK              []optionalInt `json:"TBAK"`
	TBAA              []optionalInt `json:"TBAA"`
	TBAO              []optionalInt `json:"TBAO"`
	TBAP              []optionalInt `json:"TBAP"`
	TBASTotal         optionalInt   `json:"TBAS_Total"`
	TBAKTotal         optionalInt   `json:"TBAK_Total"`
	IsDispFindarts    optionalInt   `json:"is_disp_findarts"`
	TFAU              []optionalInt `json:"TFAU"`
	TFAD              []optionalInt `json:"TFAD"`
	TFAS              []optionalInt `json:"TFAS"`
	TFAK              []optionalInt `json:"TFAK"`
	TFAA              []optionalInt `json:"TFAA"`
	TFAO              []optionalInt `json:"TFAO"`
	TFAP              []optionalInt `json:"TFAP"`
	TFASTotal         optionalInt   `json:"TFAS_Total"`
	TFAKTotal         optionalInt   `json:"TFAK_Total"`
	IsDispActarts     optionalInt   `json:"is_disp_actarts"`
	TAAU              []optionalInt `json:"TAAU"`
	UntenBunya        optionalInt   `json:"unten_bunya"`
	TAAD              []optionalInt `json:"TAAD"`
	TAAS              []optionalInt `json:"TAAS"`
	TAAK              []optionalInt `json:"TAAK"`
	TAAA              []optionalInt `json:"TAAA"`
	TAAO              []optionalInt `json:"TAAO"`
	TAAP              []optionalInt `json:"TAAP"`
	SeisakuBunya      string        `json:"seisaku_bunya"`
	MainSoujuNorimono string        `json:"main_souju_norimono"`
	TAASTotal         optionalInt   `json:"TAAS_Total"`
	TAAKTotal         optionalInt   `json:"TAAK_Total"`
	IsDispCommuarts   optionalInt   `json:"is_disp_commuarts"`
	TCAU              []optionalInt `json:"TCAU"`
	TCAD              []optionalInt `json:"TCAD"`
	TCAS              []optionalInt `json:"TCAS"`
	TCAK              []optionalInt `json:"TCAK"`
	TCAA              []optionalInt `json:"TCAA"`
	TCAO              []optionalInt `json:"TCAO"`
	TCAP              []optionalInt `json:"TCAP"`
	MylangName        optionalInt   `json:"mylang_name"`
	TCASTotal         optionalInt   `json:"TCAS_Total"`
	TCAKTotal         optionalInt   `json:"TCAK_Total"`
	IsDispKnowarts    optionalInt   `json:"is_disp_knowarts"`
	TKAU              []optionalInt `json:"TKAU"`
	TKAD              []optionalInt `json:"TKAD"`
	TKAS              []optionalInt `json:"TKAS"`
	TKAK              []optionalInt `json:"TKAK"`
	TKAA              []optionalInt `json:"TKAA"`
	TKAO              []optionalInt `json:"TKAO"`
	TKAP              []optionalInt `json:"TKAP"`
	GeijutuBunya      string        `json:"geijutu_bunya"`
	TKASTotal         optionalInt   `json:"TKAS_Total"`
	TKAKTotal         optionalInt   `json:"TKAK_Total"`
	IsDispBattle      optionalInt   `json:"is_disp_battle"`
	DmgBonus          optionalInt   `json:"dmg_bonus"`
	ArmsName          []optionalInt `json:"arms_name"`
	ArmsHit           []optionalInt `json:"arms_hit"`
	ArmsDamage        []optionalInt `json:"arms_damage"`
	ArmsRange         []optionalInt `json:"arms_range"`
	ArmsAttackCount   []optionalInt `json:"arms_attack_count"`
	ArmsLastShot      []optionalInt `json:"arms_last_shot"`
	ArmsVitality      []optionalInt `json:"arms_vitality"`
	ArmsSonota        []optionalInt `json:"arms_sonota"`
	IsDispItem        optionalInt   `json:"is_disp_item"`
	ItemName          []optionalInt `json:"item_name"`
	ItemTanka         []optionalInt `json:"item_tanka"`
	ItemNum           []optionalInt `json:"item_num"`
	ItemPrice         []optionalInt `json:"item_price"`
	ItemMemo          []optionalInt `json:"item_memo"`
	PriceItemSum      optionalInt   `json:"price_item_sum"`
	Money             optionalInt   `json:"money"`
	Debt              optionalInt   `json:"debt"`
	IsDispPersonal    optionalInt   `json:"is_disp_personal"`
	PCName            string        `json:"pc_name"`
	PCTags            string        `json:"pc_tags"`
	Shuzoku           string        `json:"shuzoku"`
	Age               optionalInt   `json:"age"`
	Sex               string        `json:"sex"`
	PCHeight          optionalInt   `json:"pc_height"`
	PCWeight          optionalInt   `json:"pc_weight"`
	PCKigen           string        `json:"pc_kigen"`
	ColorHair         string        `json:"color_hair"`
	ColorEye          string        `json:"color_eye"`
	ColorSkin         string        `json:"color_skin"`
	PCMakingMemo      string        `json:"pc_making_memo"`
	PCMakingMemoRows  optionalInt   `json:"pc_making_memo_rows"`
	Message           string        `json:"message"`
	Game              string        `json:"game"`
	DataID            int           `json:"data_id"`
	Phrase            string        `json:"phrase"`
	DataTitle         string        `json:"data_title"`
	DodontofSendTo    string        `json:"dodontof_send_to"`
	DodontofURL       string        `json:"dodontof_url"`
	DodontofRoom      string        `json:"dodontof_room"`
	KomaName          string        `json:"koma_name"`
	DodontofImage     string        `json:"dodontof_image"`
	DodontofScColors  string        `json:"dodontof_sc_colors"`
	Elysion           string        `json:"elysion"`
	Nechro            string        `json:"nechro"`
}

var canBeBared = regexp.MustCompile(`^[A-Za-z0-9_\-]+$`).MatchString

type tomlhash []struct {
	k string
	v interface{}
}

func (h tomlhash) WriteTo(w io.Writer) {
	maxw := 0
	for _, r := range h {
		w := h.width(r.k)
		if w > maxw {
			maxw = w
		}
	}

	for _, r := range h {
		var kexp string
		if canBeBared(r.k) {
			kexp = fmt.Sprintf("%s%s", r.k, h.genpad(maxw-h.width(r.k)))
		} else {
			kexp = fmt.Sprintf(`"%s"%s`, r.k, h.genpad(maxw-h.width(r.k)))
		}

		var vexp string
		switch r.v.(type) {
		case string:
			vexp = fmt.Sprintf("%q", r.v)
		default:
			vexp = fmt.Sprint(r.v)
		}

		fmt.Fprintf(w, "%s = %s\n", kexp, vexp)
	}
}

func (tomlhash) width(s string) int {
	w := 0
	for _, c := range []rune(s) {
		switch width.LookupRune(c).Kind() {
		case width.EastAsianWide, width.EastAsianFullwidth:
			w += 2
		default:
			w += 1
		}
	}
	return w
}

func (tomlhash) genpad(n int) string {
	res := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		res = append(res, ' ')
	}
	return string(res)
}

func main() {
	if err := run(os.Stdout, os.Stdin); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(w io.Writer, r io.Reader) error {
	cs := new(CS)
	if err := json.NewDecoder(r).Decode(cs); err != nil {
		return err
	}

	profile := tomlhash{
		{"キャラクター名", cs.PCName},
		{"職業", cs.Shuzoku},
		{"年齢", int(cs.Age)},
		{"性別", cs.Sex},
		{"髪の色", cs.ColorHair},
		{"身長", int(cs.PCHeight)},
		{"体重", int(cs.PCWeight)},
	}
	fmt.Fprintln(w, `["プロフィール"]`)
	profile.WriteTo(w)

	characteristics := tomlhash{
		{"STR", int(cs.NA1)},
		{"CON", int(cs.NA2)},
		{"POW", int(cs.NA3)},
		{"DEX", int(cs.NA4)},
		{"APP", int(cs.NA5)},
		{"SIZ", int(cs.NA6)},
		{"INT", int(cs.NA7)},
		{"EDU", int(cs.NA8)},
	}
	fmt.Fprintln(w)
	fmt.Fprintln(w, `["技能"]`)
	characteristics.WriteTo(w)

	skills := tomlhash{
		{"言いくるめ", int(cs.TCAS[0]) + int(cs.TCAK[0])},
		{"医学", int(cs.TKAS[0]) + int(cs.TKAK[0])},
		{"運転", int(cs.TAAS[0]) + int(cs.TAAK[0])},
		{"応急手当", int(cs.TFAS[0]) + int(cs.TFAK[0])},
		{"オカルト", int(cs.TKAS[1]) + int(cs.TKAK[1])},
		{"回避", int(cs.TBAS[0]) + int(cs.TBAK[0])},
		{"化学", int(cs.TKAS[2]) + int(cs.TKAK[2])},
		{"鍵開け", int(cs.TFAS[1]) + int(cs.TFAK[1])},
		{"隠す", int(cs.TFAS[2]) + int(cs.TFAK[2])},
		{"隠れる", int(cs.TFAS[3]) + int(cs.TFAK[3])},
		{"機械修理", int(cs.TAAS[1]) + int(cs.TAAK[1])},
		{"聞き耳", int(cs.TFAS[4]) + int(cs.TFAK[4])},
		{"キック", int(cs.TBAS[1]) + int(cs.TBAK[1])},
		{"クトゥルフ神話", int(cs.TKAS[3]) + int(cs.TKAK[3])},
		{"組みつき", int(cs.TBAS[2]) + int(cs.TBAK[2])},
		{"芸術", int(cs.TKAS[4]) + int(cs.TKAK[4])},
		{"経理", int(cs.TKAS[5]) + int(cs.TKAK[5])},
		{"拳銃", int(cs.TBAS[7]) + int(cs.TBAK[7])},
		{"考古学", int(cs.TKAS[6]) + int(cs.TKAK[6])},
		{"こぶし/パンチ", int(cs.TBAS[3]) + int(cs.TBAK[3])},
		{"コンピューター", int(cs.TKAS[7]) + int(cs.TKAK[7])},
		{"サブマシンガン", int(cs.TBAS[8]) + int(cs.TBAK[8])},
		{"忍び歩き", int(cs.TFAS[5]) + int(cs.TFAK[5])},
		{"写真術", int(cs.TFAS[6]) + int(cs.TFAK[6])},
		{"重機械操作", int(cs.TAAS[2]) + int(cs.TAAK[2])},
		{"乗馬", int(cs.TAAS[3]) + int(cs.TAAK[3])},
		{"ショットガン", int(cs.TBAS[9]) + int(cs.TBAK[9])},
		{"信用", int(cs.TCAS[1]) + int(cs.TCAK[1])},
		{"心理学", int(cs.TKAS[8]) + int(cs.TKAK[8])},
		{"人類学", int(cs.TKAS[9]) + int(cs.TKAK[9])},
		{"水泳", int(cs.TAAS[4]) + int(cs.TAAK[4])},
		{"制作", int(cs.TAAS[5]) + int(cs.TAAK[5])},
		{"精神分析", int(cs.TFAS[7]) + int(cs.TFAK[7])},
		{"生物学", int(cs.TKAS[10]) + int(cs.TKAK[10])},
		{"説得", int(cs.TCAS[2]) + int(cs.TCAK[2])},
		{"操縦", int(cs.TAAS[6]) + int(cs.TAAK[6])},
		{"地質学", int(cs.TKAS[11]) + int(cs.TKAK[11])},
		{"跳躍", int(cs.TAAS[7]) + int(cs.TAAK[7])},
		{"追跡", int(cs.TFAS[8]) + int(cs.TFAK[8])},
		{"頭突き", int(cs.TBAS[4]) + int(cs.TBAK[4])},
		{"電気修理", int(cs.TAAS[8]) + int(cs.TAAK[8])},
		{"電子工学", int(cs.TKAS[12]) + int(cs.TKAK[12])},
		{"天文学", int(cs.TKAS[13]) + int(cs.TKAK[13])},
		{"投擲", int(cs.TBAS[5]) + int(cs.TBAK[5])},
		{"登攀", int(cs.TFAS[9]) + int(cs.TFAK[9])},
		{"図書館", int(cs.TFAS[10]) + int(cs.TFAK[10])},
		{"ナビゲート", int(cs.TAAS[9]) + int(cs.TAAK[9])},
		{"値切り", int(cs.TCAS[3]) + int(cs.TCAK[3])},
		{"博物学", int(cs.TKAS[14]) + int(cs.TKAK[14])},
		{"物理学", int(cs.TKAS[15]) + int(cs.TKAK[15])},
		{"変装", int(cs.TAAS[10]) + int(cs.TAAK[10])},
		{"法律", int(cs.TKAS[16]) + int(cs.TKAK[16])},
		{"ほかの言語", 0},
		{"母国語", int(cs.TCAS[4]) + int(cs.TCAK[4])},
		{"マーシャルアーツ", int(cs.TBAS[6]) + int(cs.TBAK[6])},
		{"マシンガン", int(cs.TBAS[10]) + int(cs.TBAK[10])},
		{"目星", int(cs.TFAS[11]) + int(cs.TFAK[11])},
		{"薬学", int(cs.TKAS[17]) + int(cs.TKAK[17])},
		{"ライフル", int(cs.TBAS[11]) + int(cs.TBAK[11])},
		{"歴史", int(cs.TKAS[18]) + int(cs.TKAK[18])},
	}
	fmt.Fprintln(w)
	fmt.Fprintln(w, `["能力値"]`)
	filtered := make(tomlhash, 0, len(skills))
	for _, s := range skills {
		if s.v.(int) != 0 {
			filtered = append(filtered, s)
		}
	}
	filtered.WriteTo(w)

	return nil
}
