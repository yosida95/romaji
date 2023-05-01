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
			in:  "メーワ",
			out: "MEIWA",
		},
		{
			in:  "メイワ",
			out: "MEIWA",
		},
		{
			in:  "イーヤマ",
			out: "IIYAMA",
		},
		{
			in:  "イイヤマ",
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
			in: "HOTTA",
			out: []string{
				// ホウ|TTA
				"ホウッター",
				"ホウッタ",
				// ホオ|TTA
				"ホオッター",
				"ホオッタ",
				// ホー|TTA
				"ホーッター",
				"ホーッタ",
				// ホ|TTA
				"ホッター",
				"ホッタ",
			},
		},
		{
			in: "KAN-ONJI",
			out: []string{
				// カーンオウ|NJI
				// カーンオウン|JI
				"カーンオウンヂ",
				"カーンオウンジ",

				// カーンオオ|NJI
				// カーンオオン|JI
				"カーンオオンヂ",
				"カーンオオンジ",

				// カーンオー|NJI
				// カーンオーン|JI
				"カーンオーンヂ",
				"カーンオーンジ",

				// カーンオ|NJI
				// カーンオン|JI
				"カーンオンヂ",
				"カーンオンジ",

				// カンオウ|NJI
				// カンオウン|JI
				"カンオウンヂ",
				"カンオウンジ",

				// カンオオ|NJI
				// カンオオン|JI
				"カンオオンヂ",
				"カンオオンジ",

				// カンオー|NJI
				// カンオーン|JI
				"カンオーンヂ",
				"カンオーンジ",

				// カンオ|NJI
				// カンオン|JI
				"カンオンヂ",
				"カンオンジ",
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
				// SKIP: "ホオオ",
				"ホオー",
				"ホオ",
			},
		},
		{
			in: "MEIWA",
			out: []string{
				// メー|WA
				"メーワー",
				"メーワ",

				// メイ|WA
				"メイワー",
				"メイワ",
			},
		},
		{
			in: "IIYAMA",
			out: []string{
				// イー|YAMA
				// イーヤー|MA
				"イーヤーマー",
				"イーヤーマ",
				// イーヤ|MA
				"イーヤマー",
				"イーヤマ",

				// イイ|YAMA
				// イイヤー|MA
				"イイヤーマー",
				"イイヤーマ",
				// イイヤ|MA
				"イイヤマー",
				"イイヤマ",
			},
		},
	} {
		t.Log(c.in)
		out := ToKanaHepburn(c.in)
		if !cmp.Equal(c.out, out) {
			t.Error(cmp.Diff(c.out, out))
		}
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
