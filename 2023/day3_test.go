package adoc2023

import "testing"

func TestDay3(t *testing.T) {
	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	expected := 4361
	actual := Day3(input)
	if actual != expected {
		t.Errorf("Day3() = %d, expected %d", actual, expected)
	}

}

func TestDay3LongTest(t *testing.T) {
	input := []string{
		"......@....663*....#..............................................39..................265...................328.......*......*..810.....=186",
	}
	expected := 849
	actual := Day3(input)
	if actual != expected {
		t.Errorf("Day3() = %d, expected %d", actual, expected)
	}
}

//func TestDay3PartB(t *testing.T) {
//	input := []string{
//		"467..114..",
//		"...*......",
//		"..35..633.",
//		"......#...",
//		"617*......",
//		".....+.58.",
//		"..592.....",
//		"......755.",
//		"...$.*....",
//		".664.598..",
//	}
//	expected := int64(46783564)
//	actual := Day3PartB(input)
//	if actual != expected {
//		t.Errorf("Day3() = %d, expected %d", actual, expected)
//	}
//}
