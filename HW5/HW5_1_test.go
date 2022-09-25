package main

import (
	"sync"
	"testing"
)

func Benchmark10W90R(b *testing.B) { //10% write 90% read
	var set = NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(10)
		for i := 0; i < 1; i++ {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					set.Add(1)
				}
			})
		}
		for i := 0; i < 9; i++ {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					set.Has(1)
				}
			})
		}
	})
}

func Benchmark50W50R(b *testing.B) { //10% write 90% read
	var set = NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(10)
		for i := 0; i < 5; i++ {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					set.Add(1)
				}
			})
		}
		for i := 0; i < 5; i++ {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					set.Has(1)
				}
			})
		}
	})
}

func Benchmark90W10R(b *testing.B) { //10% write 90% read
	var set = NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(10)
		for i := 0; i < 9; i++ {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					set.Add(1)
				}
			})
		}
		for i := 0; i < 1; i++ {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					set.Has(1)
				}
			})
		}
	})
}

type Set struct {
	sync.RWMutex
	mm map[int]struct{}
}

func NewSet() *Set {
	return &Set{
		mm: map[int]struct{}{},
	}
}
func (s *Set) Add(i int) {
	s.Lock()
	s.mm[i] = struct{}{}
	s.Unlock()
}

func (s *Set) Has(i int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.mm[i]
	return ok
}
