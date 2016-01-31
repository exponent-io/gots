package gots

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type StructA struct {
	A float64 `json:"a"`
	B string  `json:"b,omitempty"`
	C string  `json:"-"`
	D string  `json:",omitempty"`
}

type StructB struct {
	A string    `json:"a"`
	B []int     `json:"b"`
	C []StructB `json:"c"`
	D StructA   `json:"d"`
}

type StructC struct {
	A string
	B []int
	c []StructB
	D StructA
}

func TestStructA(t *testing.T) {
	exp := `
export interface StructA {
  a: number;
  b: string;
}
`
	ts, err := ConvertToString(StructA{})
	assert.NoError(t, err)
	assert.Equal(t, exp, ts)
}

func TestStructB(t *testing.T) {
	exp := `
export interface StructA {
  a: number;
  b: string;
}

export interface StructB {
  a: string;
  b: number[];
  c: StructB[];
  d: StructA;
}
`
	ts, err := ConvertToString(StructB{})
	assert.NoError(t, err)
	assert.Equal(t, exp, ts)
}

func TestStructC(t *testing.T) {
	exp := `
export interface StructA {
  a: number;
  b: string;
}

export interface StructB {
  a: string;
  b: number[];
  c: StructB[];
  d: StructA;
}
`
	ts, err := ConvertToString(StructB{}, StructA{})
	assert.NoError(t, err)
	assert.Equal(t, exp, ts)
}

func TestStructD(t *testing.T) {
	exp := `
export interface StructA {
  a: number;
  b: string;
}

export interface StructC {
  A: string;
  B: number[];
  D: StructA;
}
`
	ts, err := ConvertToString(StructC{})
	assert.NoError(t, err)
	assert.Equal(t, exp, ts)
}
