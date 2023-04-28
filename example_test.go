// Copyright (C) 2023 Kohei YOSHIDA. All rights reserved.
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of The BSD 3-Clause License
// that can be found in the LICENSE file.

package romaji_test

import (
	"fmt"

	"github.com/yosida95/romaji"
)

func Example() {
	kana := "トウキョウ"
	roman := romaji.FromKanaHepburn(kana)
	fmt.Printf("%q -> %q\n", kana, roman)

	kanaCandidates := romaji.ToKanaHepburn(roman)
	fmt.Printf("%q in kanaCandidates: %v\n", kana, StringContains(kanaCandidates, kana))

	// Output:
	// "トウキョウ" -> "TOKYO"
	// "トウキョウ" in kanaCandidates: true
}

func StringContains(in []string, expect string) bool {
	for _, have := range in {
		if have == expect {
			return true
		}
	}
	return false
}
