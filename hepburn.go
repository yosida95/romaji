// Copyright (C) 2023 Kohei YOSHIDA. All rights reserved.
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of The BSD 3-Clause License
// that can be found in the LICENSE file.

package romaji

import (
	"sort"
)

var (
	hepburnToKana   = map[string][]string{}
	hepburnFromKana = map[string]string{
		"ア": "A", "イ": "I", "ウ": "U", "エ": "E", "オ": "O",
		"カ": "KA", "キ": "KI", "ク": "KU", "ケ": "KE", "コ": "KO",
		"サ": "SA", "シ": "SHI", "ス": "SU", "セ": "SE", "ソ": "SO",
		"タ": "TA", "チ": "CHI", "ツ": "TSU", "テ": "TE", "ト": "TO",
		"ナ": "NA", "ニ": "NI", "ヌ": "NU", "ネ": "NE", "ノ": "NO",
		"ハ": "HA", "ヒ": "HI", "フ": "FU", "ヘ": "HE", "ホ": "HO",
		"マ": "MA", "ミ": "MI", "ム": "MU", "メ": "ME", "モ": "MO",
		"ヤ": "YA", "ユ": "YU", "ヨ": "YO",
		"ラ": "RA", "リ": "RI", "ル": "RU", "レ": "RE", "ロ": "RO",
		"ワ": "WA", "ヲ": "O", "ン": "N",

		"ガ": "GA", "ギ": "GI", "グ": "GU", "ゲ": "GE", "ゴ": "GO",
		"ザ": "ZA", "ジ": "JI", "ズ": "ZU", "ゼ": "ZE", "ゾ": "ZO",
		"ダ": "DA", "ヂ": "JI", "ヅ": "ZU", "デ": "DE", "ド": "DO",
		"バ": "BA", "ビ": "BI", "ブ": "BU", "ベ": "BE", "ボ": "BO",
		"パ": "PA", "ピ": "PI", "プ": "PU", "ペ": "PE", "ポ": "PO",

		"キャ": "KYA", "キュ": "KYU", "キョ": "KYO",
		"シャ": "SHA", "シュ": "SHU", "ショ": "SHO",
		"チャ": "CHA", "チュ": "CHU", "チョ": "CHO",
		"ニャ": "NYA", "ニュ": "NYU", "ニョ": "NYO",
		"ヒャ": "HYA", "ヒュ": "HYU", "ヒョ": "HYO",
		"ミャ": "MYA", "ミュ": "MYU", "ミョ": "MYO",
		"リャ": "RYA", "リュ": "RYU", "リョ": "RYO",

		"ギャ": "GYA", "ギュ": "GYU", "ギョ": "GYO",
		"ジャ": "JA", "ジュ": "JU", "ジョ": "JO",
		"ヂャ": "JA", "ヂュ": "JU", "ヂョ": "JO",
		"ビャ": "BYA", "ビュ": "BYU", "ビョ": "BYO",
		"ピャ": "PYA", "ピュ": "PYU", "ピョ": "PYO",

		"キェ": "KYE", "シェ": "SHE", "チェ": "CHE", "ニェ": "NYE",
		"ヒェ": "HYE", "ミェ": "MYE", "イェ": "YE", "リェ": "RYE",
		"ギェ": "GYE", "ジェ": "JE", "ビェ": "BYE", "ピェ": "PYE",

		"ティ": "TI", "トゥ": "TU",
		"ディ": "DI", "ドゥ": "DU",
		"デュ": "DYU",
		"ツァ": "TSA", "ツィ": "TSI", "ツェ": "TSE", "ツォ": "TSO",
		"ファ": "FA", "フィ": "FI", "フェ": "FE", "フォ": "FO",
		"フャ": "FYA", "フュ": "FYU", "フョ": "FYO",
		"ヴァ": "VA", "ヴィ": "VI", "ヴ": "VU", "ヴェ": "VE", "ヴォ": "VO",
		"ウィ": "WI", "ウェ": "WE", "ウォ": "WO",

		"ンア": "N-A", "ンイ": "N-I", "ンウ": "N-U", "ンエ": "N-E", "ンオ": "N-O",
		"ンヤ": "N-YA", "ンユ": "N-YU", "ンヨ": "N-YO",
		"ンイェ": "N-YE",

		"ンバ": "MBA", "ンビ": "MBI", "ンブ": "MBU", "ンベ": "MBE", "ンボ": "MBO",
		"ンビャ": "MBYA", "ンビュ": "MBYU", "ンビョ": "MBYO",
		"ンビェ": "MBYE",

		"ンマ": "MMA", "ンミ": "MMI", "ンム": "MMU", "ンメ": "MME", "ンモ": "MMO",
		"ンミャ": "MMYA", "ンミュ": "MMYU", "ンミョ": "MMYO",
		"ンミェ": "MMYE",

		"ンパ": "MPA", "ンピ": "MPI", "ンプ": "MPU", "ンペ": "MPE", "ンポ": "MPO",
		"ンピャ": "MPYA", "ンピュ": "MPYU", "ンピョ": "MPYO",
		"ンピェ": "MPYE",
	}
)

func init() {
	for kana, roman := range hepburnFromKana {
		hepburnToKana[roman] = append(hepburnToKana[roman], kana)
	}
	for roman, kana := range hepburnToKana {
		sort.Strings(kana)
		hepburnToKana[roman] = kana
	}
}

// FromKanaHepburn translates Zenkaku Katakana to Hepburn Romaji. If the
// translation is successful, it returns the resulting string. If it encounters
// a character is not able to translate to Romaji, it will emit the character
// as-is.
func FromKanaHepburn(kana string) string {
	return fromKana(hepburnFromKana, kana)
}

// ToKanaHepburn translates Hepburn romaji to Zenkaku Katakana and returns
// all possible reprensentations. The return values are not sorted in any order
//and its order is subject to change.
func ToKanaHepburn(roman string) []string {
	return toKana(hepburnToKana, roman)
}
