package main

import (
	"reflect"
	"testing"
)

func TestParseJunctionBox(t *testing.T) {
	line := "14445,17266,2972"
	jb := ParseJunctionBox(line)
	expected := JunctionBox{14445, 17266, 2972}
	if !reflect.DeepEqual(jb, expected) {
		t.Errorf("ParseJunctionBox(\"%s\") expected %v but got %v\n", line, expected, jb)
	}
}
