package main

import (
    "testing"
)

func TestParseUpdateAmountMsg(t *testing.T) {
    input := "/income 42"
    expected := UpdateData{42}
    actual, _ := ParseUpdateAmountMsg(input)
    if actual != expected {
        t.Errorf("Parse `%s`, expected %v; get %v", input, expected, actual)
    }
}

func TestValidateAmount(t *testing.T) {
    var tests = []struct {
        val int64
        expected bool
    }{
        {42, true},
        {-1, false},
        {0, false},
    }

    for _, rec := range tests {
        ans := validateAmount(rec.val)
        if rec.expected != ans {
            t.Errorf("validateAmount: expected %v get %v", rec.expected, ans)
        }
    }
}
