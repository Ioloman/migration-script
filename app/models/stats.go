package models

import (
	"fmt"
	"time"
)

type Timings struct {
	Select int64
	Insert int64
	Delete int64
	Count  uint64
}

func (t *Timings) Total() int64 {
	return t.Select + t.Insert + t.Delete
}

func (t *Timings) Add(newT *Timings) {
	t.Select += newT.Select
	t.Insert += newT.Insert
	t.Delete += newT.Delete
	t.Count += newT.Count
}

func (t *Timings) String() string {
	total := float64(t.Total()) / 1e3
	s := float64(t.Select) / 1e3
	i := float64(t.Insert) / 1e3
	d := float64(t.Delete) / 1e3
	c := float64(t.Count)
	if t.Count > 1 {
		return fmt.Sprintf(
			"Total: %v s, %v s/op, %v op/s; Select: %v s, %v s/op, %v op/s; Insert: %v s, %v s/op, %v op/s; Delete: %v s, %v s/op, %v op/s",
			total, total/c, c/total,
			s, s/c, c/s,
			i, i/c, c/i,
			d, d/c, c/d,
		)
	} else {
		return fmt.Sprintf(
			"Total: %v s; Select: %v s; Insert: %v s; Delete: %v s",
			total, s, i, d,
		)
	}
}

func (t *Timings) SetSelect(ct time.Time) time.Time {
	now := time.Now()
	t.Select = now.UnixMilli() - ct.UnixMilli()
	return now
}

func (t *Timings) AddSelect(ct time.Time) time.Time {
	now := time.Now()
	t.Select += now.UnixMilli() - ct.UnixMilli()
	return now
}

func (t *Timings) SetInsert(ct time.Time) time.Time {
	now := time.Now()
	t.Insert = now.UnixMilli() - ct.UnixMilli()
	return now
}

func (t *Timings) SetDelete(ct time.Time) time.Time {
	now := time.Now()
	t.Delete = now.UnixMilli() - ct.UnixMilli()
	return now
}
