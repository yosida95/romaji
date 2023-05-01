// Copyright (C) 2023 Kohei YOSHIDA. All rights reserved.
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of The BSD 3-Clause License
// that can be found in the LICENSE file.

package romaji

import (
	"log"
	"strings"
	"unicode/utf8"
)

type debugT bool

var debug = debugT(false)

func (debug debugT) Printf(format string, in ...any) {
	if debug {
		log.Printf(format, in...)
	}
}

func fromKana(table map[string]string, in string) string {
	type point struct{ start, stop int }
	var (
		stop int
		last byte
		out  strings.Builder

		queue = []point{}
	)
	for {
		if stop >= len(in) {
			if len(queue) == 0 {
				return out.String()
			}
		} else {
			rune, size := utf8.DecodeRuneInString(in[stop:])
			switch {
			case size == 1:
				if len(queue) == 0 {
					out.WriteByte(in[stop])
					stop++
					last = 0
					continue
				}
			case rune == 'ッ' || rune == 'ン':
				queue = append(queue, point{stop, stop + size})
				stop += size
				if len(queue) == 1 {
					continue
				}
			case strings.ContainsRune("ィェャュョ", rune):
				if len(queue) == 0 {
					out.WriteRune(rune)
					stop += size
					last = 0
					continue
				}
				queue = append(queue, point{stop, stop + size})
				stop += size
			case rune == 'ー':
				if len(queue) == 0 {
					if last == 'I' || last == 'E' {
						out.WriteByte('I')
					}
					stop += size
					last = 0
					continue
				}
			default:
				queue = append(queue, point{stop, stop + size})
				stop += size
				if len(queue) < 3 {
					continue
				}
			}
		}
		var sokuon bool
		if e := queue[0]; in[e.start:e.stop] == "ッ" {
			sokuon = true
			queue = queue[1:]
		}
		var roman string
		for n := len(queue); n > 0; n-- {
			kana := in[queue[0].start:queue[n-1].stop]
			roman = table[kana]
			if roman != "" {
				queue = queue[n:]
				break
			}
		}
		if roman == "" {
			e := queue[0]
			queue = queue[1:]
			out.WriteString(in[e.start:e.stop])
			last = 0
			continue
		}
		switch first := roman[0]; {
		case last == 'O' && (first == 'O' || first == 'U'):
			out.WriteString(roman[1:])
			fallthrough
		case last == 'U' && first == 'U':
			out.WriteString(roman[1:])
			last = 0
		case sokuon:
			if strings.HasPrefix(roman, "CH") {
				out.WriteByte('T')
			} else {
				out.WriteByte(first)
			}
			fallthrough
		default:
			out.WriteString(roman)
			last = roman[len(roman)-1]
		}
	}
}

func toKana(table map[string][]string, in string) []string {
	type candidate struct {
		out  strings.Builder
		next string
		i    int
	}
	type cache struct {
		out string
		i   int
	}
	var (
		stack = []*candidate{{}}
		seen  = map[cache]struct{}{}
	)
	var ret []string
	for len(stack) > 0 {
		n := len(stack) - 1
		if debug {
			log.Printf("====================")
			for i, p := range stack {
				if n > i {
					log.Printf("%3d: %s|%s", i, p.out.String(), in[p.i:])
				} else {
					log.Printf("***: %s|%s", p.out.String(), in[p.i:])
				}
			}
		}
		p := stack[n]
		stack = stack[:n]

		key := cache{i: p.i, out: p.out.String()}
		if _, ok := seen[key]; ok {
			debug.Printf("Skip as duplicate")
			continue
		}
		seen[key] = struct{}{}

		roman := in[p.i:]
		if n := len(roman); n > 0 {
			if vowel := strings.IndexAny(roman, "AIUEO"); 0 <= vowel && vowel < n /* BCE */ {
				roman = roman[0 : vowel+1]
			}
		}
		if roman == "" {
			if p.next == "" {
				ret = append(ret, p.out.String())
			}
			continue
		}
		if p.next != "" {
			rune, size := utf8.DecodeRuneInString(roman)
			if !strings.ContainsRune(p.next, rune) {
				continue
			}
			p.i += size
			p.next = ""
			stack = append(stack, p)
			continue
		}

		if len(roman) >= 2 && roman[0] == 'N' && roman[1] != '-' {
			alt := &candidate{i: p.i + 1}
			alt.out.WriteString(p.out.String())
			alt.out.WriteString("ン")
			stack = append(stack, alt)
		}
		if strings.HasPrefix(roman, "TCH") || len(roman) >= 3 && roman[0] == roman[1] {
			alt := &candidate{i: p.i + 1}
			alt.out.WriteString(p.out.String())
			alt.out.WriteString("ッ")
			stack = append(stack, alt)
		}
		p.i += len(roman)
		for _, kana := range table[roman] {
			next := &candidate{i: p.i}
			next.out.WriteString(p.out.String())
			next.out.WriteString(kana)
			stack = append(stack, next)

			alt := &candidate{i: p.i}
			if last := roman[len(roman)-1]; last == 'I' || last == 'E' {
				alt.next = "I"
			}
			alt.out.WriteString(next.out.String())
			alt.out.WriteString("ー")
			stack = append(stack, alt)

			switch roman[len(roman)-1] {
			case 'O':
				alt := &candidate{i: p.i}
				alt.out.WriteString(next.out.String())
				alt.out.WriteString("オ")
				stack = append(stack, alt)
				fallthrough
			case 'U':
				alt := &candidate{i: p.i}
				alt.out.WriteString(next.out.String())
				alt.out.WriteString("ウ")
				stack = append(stack, alt)
			}
		}
	}
	return ret
}
