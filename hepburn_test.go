// Copyright (C) 2023 Kohei YOSHIDA. All rights reserved.
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of The BSD 3-Clause License
// that can be found in the LICENSE file.

package romaji

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func init() {
	debug = true
}

func TestFromKanaHepburn(t *testing.T) {
	for _, c := range []struct {
		in  string
		out string
	}{
		{
			in:  "クッチャン",
			out: "KUTCHAN",
		},
		{
			in:  "ホッタ",
			out: "HOTTA",
		},
		{
			in:  "カンオンジ",
			out: "KAN-ONJI",
		},
		{
			in:  "ゼンツウジ",
			out: "ZENTSUJI",
		},
		{
			in:  "ホウオウ",
			out: "HOO",
		},
		{
			in:  "クーベツ",
			out: "KUBETSU",
		},
		{
			in:  "エーデル",
			out: "EIDERU",
		},
		{
			in:  "イーヤマ",
			out: "IIYAMA",
		},
		{
			in:  "第3シン TOKYO シ",
			out: "第3SHIN TOKYO SHI",
		},
		{
			in:  "ャン",
			out: "ャN",
		},
	} {
		t.Log(c.in)
		out := FromKanaHepburn(c.in)
		if out != c.out {
			t.Errorf("expect %q, got %q", c.out, out)
		}
	}
}

func TestToKanaHepburn(t *testing.T) {
	for _, c := range []struct {
		in  string
		out []string
	}{
		{
			in: "KUTCHAN",
			out: []string{
				// クウ|TCHAN
				// クウッチャー|N
				"クウッチャーンー",
				"クウッチャーン",
				// クウッチャ|N
				"クウッチャンー",
				"クウッチャン",

				// クー|TCHAN
				// クーッチャー|N
				"クーッチャーンー",
				"クーッチャーン",
				// クーッチャ|N
				"クーッチャンー",
				"クーッチャン",

				// ク|TCHAN
				// クッチャー|N
				"クッチャーンー",
				"クッチャーン",
				// クッチャ|N
				"クッチャンー",
				"クッチャン",
			},
		},
		{
			in: "KAN-ONJI",
			out: []string{
				// カーンオウ|NJI
				// カーンオウン|JI
				"カーンオウンヂー",
				"カーンオウンヂ",
				"カーンオウンジー",
				"カーンオウンジ",

				// カーンオオ|NJI
				// カーンオオン|JI
				"カーンオオンヂー",
				"カーンオオンヂ",
				"カーンオオンジー",
				"カーンオオンジ",

				// カーンオー|NJI
				// カーンオーン|JI
				"カーンオーンヂー",
				"カーンオーンヂ",
				"カーンオーンジー",
				"カーンオーンジ",

				// カーンオ|NJI
				// カーンオン|JI
				"カーンオンヂー",
				"カーンオンヂ",
				"カーンオンジー",
				"カーンオンジ",

				// カンオウ|NJI
				// カンオウン|JI
				"カンオウンヂー",
				"カンオウンヂ",
				"カンオウンジー",
				"カンオウンジ",

				// カンオオ|NJI
				// カンオオン|JI
				"カンオオンヂー",
				"カンオオンヂ",
				"カンオオンジー",
				"カンオオンジ",

				// カンオー|NJI
				// カンオーン|JI
				"カンオーンヂー",
				"カンオーンヂ",
				"カンオーンジー",
				"カンオーンジ",

				// カンオ|NJI
				// カンオン|JI
				"カンオンヂー",
				"カンオンヂ",
				"カンオンジー",
				"カンオンジ",
			},
		},
		{
			in: "ZENTSUJI",
			out: []string{
				// ゼー|NTSUJI
				// ゼーン|TSUJI
				// ゼーンツウ|JI
				"ゼーンツウヂー",
				"ゼーンツウヂ",
				"ゼーンツウジー",
				"ゼーンツウジ",
				// ゼーンツー|JI
				"ゼーンツーヂー",
				"ゼーンツーヂ",
				"ゼーンツージー",
				"ゼーンツージ",
				// ゼーンツ|JI
				"ゼーンツヂー",
				"ゼーンツヂ",
				"ゼーンツジー",
				"ゼーンツジ",

				// ぜ|NTSUJI
				// ゼン|TSUJI
				// ゼンツウ|JI
				"ゼンツウヂー",
				"ゼンツウヂ",
				"ゼンツウジー",
				"ゼンツウジ",
				// ゼンツー|JI
				"ゼンツーヂー",
				"ゼンツーヂ",
				"ゼンツージー",
				"ゼンツージ",
				// ゼンツ|JI
				"ゼンツヂー",
				"ゼンツヂ",
				"ゼンツジー",
				"ゼンツジ",
			},
		},
		{
			in: "HOO",
			out: []string{
				// ホウ|O
				"ホウヲウ",
				"ホウヲオ",
				"ホウヲー",
				"ホウヲ",
				"ホウオウ",
				"ホウオオ",
				"ホウオー",
				"ホウオ",

				// ホオ|O
				"ホオヲウ",
				"ホオヲオ",
				"ホオヲー",
				"ホオヲ",
				"ホオオウ",
				"ホオオオ",
				"ホオオー",
				"ホオオ",

				// ホー|O
				"ホーヲウ",
				"ホーヲオ",
				"ホーヲー",
				"ホーヲ",
				"ホーオウ",
				"ホーオオ",
				"ホーオー",
				"ホーオ",

				// ホ|O
				"ホヲウ",
				"ホヲオ",
				"ホヲー",
				"ホヲ",
				"ホオウ",
				// "ホオオ",
				"ホオー",
				"ホオ",
			},
		},
		{
			in: "KUBETSU",
			out: []string{
				// クウ|BETSU
				// クウベー|TSU
				"クウベーツウ",
				"クウベーツー",
				"クウベーツ",
				// クウベ|TSU
				"クウベツウ",
				"クウベツー",
				"クウベツ",

				// クー|BETSU
				// クーベー|TSU
				"クーベーツウ",
				"クーベーツー",
				"クーベーツ",
				// クーベ|TSU
				"クーベツウ",
				"クーベツー",
				"クーベツ",

				// ク|BETSU
				// クベー|TSU
				"クベーツウ",
				"クベーツー",
				"クベーツ",
				// クベ|TSU
				"クベツウ",
				"クベツー",
				"クベツ",
			},
		},
	} {
		t.Log(c.in)
		out := ToKanaHepburn(c.in)
		if !cmp.Equal(c.out, out) {
			t.Error(cmp.Diff(c.out, out))
		}
		t.Log(len(out))
	}
}

func TestHepburn(t *testing.T) {
	const expect = "TOKYO"

	kana := ToKanaHepburn(expect)
	for _, kana := range kana {
		t.Log(kana)
		roman := FromKanaHepburn(kana)
		if roman != expect {
			t.Errorf("expect %q, got %q", expect, roman)
		}
	}
}

func BenchmarkHepburnToKana(b *testing.B) {
	debug = false
	for i := 0; i < b.N; i++ {
		ToKanaHepburn("ゼンツウジ")
	}
}
