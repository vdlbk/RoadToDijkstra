package strings_challenges

import (
	"fmt"
	"testing"
)

func BenchmarkRemoveDuplicates(b *testing.B) {
	benchmarks := []struct {
		n []byte
	}{
		{[]byte{byte('H'), byte('a'), byte('n'), byte('n'), byte('a'), byte('h')}},
		{[]byte{byte('H'), byte('a'), byte('n'), byte('n'), byte('a'), byte('h'), byte('0'), byte('1'), byte(' ')}},
		{[]byte{byte('H'), byte('a'), byte('n'), byte('n'), byte('a'), byte('h'), byte('0'), byte('1'), byte(' '),
			byte('H'), byte('a'), byte('n'), byte('n'), byte('a'), byte('h'), byte('0'), byte('1'), byte(' '),
			byte('H'), byte('a'), byte('n'), byte('n'), byte('a'), byte('h'), byte('0'), byte('1'), byte(' ')}},
	}

	for i, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("reverseString %v", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				reverseString(benchmark.n)
			}
		})
	}

	for i, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("reverseStringOptimal %v", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				reverseStringOptimal(benchmark.n)
			}
		})
	}
}

func BenchmarkIsAnagram(b *testing.B) {
	benchmarks := []struct {
		s string
		t string
	}{
		{"", ""},
		{"foo", "oof"},
		{"foobar", "barfoo"},
		{"anagram", "nagaram"},
		{"wazaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaazaw",
			"zawaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaawaz"},
	}

	for i, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("isAnagram %v", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				isAnagram(benchmark.s, benchmark.t)
			}
		})
	}

	for i, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("isAnagramOpti %v", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				isAnagramOpti(benchmark.s, benchmark.t)
			}
		})
	}
}

func BenchmarkStrStr(b *testing.B) {
	benchmarks := []struct {
		s string
		t string
	}{
		{"", ""},
		{"foo", "oo"},
		{"foobar", "bar"},
		{"anagram", "gr"},
		{"abcdefghijklmnopqrstuvwxyz", "mno"},
		{"abcdefghijklmnopqrstuvwxyz", ""},
		{"abcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxyz0"},
	}

	for i, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("strStr %v", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				strStr(benchmark.s, benchmark.t)
			}
		})
	}

	for i, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("strStrOpti %v", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				strStrOpti(benchmark.s, benchmark.t)
			}
		})
	}
}
