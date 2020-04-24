package useselect

import (
	"testing"
)

func TestCalcFiboni(t *testing.T) {
	CalcFiboni()
}

func TestNotBlock(t *testing.T) {
	// 没有阻塞直接返回了
	NotBlock()
}

func TestRand(t *testing.T) {
	RandSelect()
}
