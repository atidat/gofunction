package math

import (
    "testing"
)

func TestMaxA(t *testing.T) {
    if IntMax(3, 2) != 3 {
        t.Errorf("function error")
    }
}

func TestMaxB(t *testing.T) {
    if IntMax(0, 2) != 2 {
        t.Errorf("function error")
    }
}

func TestMaxC(t *testing.T) {
    if IntMax(-1, 2) != 2 {
        t.Errorf("function error")
    }
}

