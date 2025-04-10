package triangle

import "testing"

func TestGetTriangleType(t *testing.T) {
	type Test struct {
		a, b, c  int
		expected triangleType
	}

	var tests = []Test{
		{30001, 6, 2, UnknownTriangle},
		// TODO add more tests for 100% test coverage
		{3, 20001, 2, UnknownTriangle},
		{3, 6, 10001, UnknownTriangle},
		{0, 5, 5, UnknownTriangle},
		{6, -5, 5, UnknownTriangle},
		{6, 5, -5, UnknownTriangle},
		{1,1,20, InvalidTriangle},
		{20,1,1, InvalidTriangle},
		{1,20,1, InvalidTriangle},
		{5, 3, 4, RightTriangle},   
		{5, 5, 5, AcuteTriangle}, 
		{7, 4, 5, ObtuseTriangle}, 


	}

	for _, test := range tests {
		actual := getTriangleType(test.a, test.b, test.c)
		if actual != test.expected {
			t.Errorf("getTriangleType(%d, %d, %d)=%v; want %v", test.a, test.b, test.c, actual, test.expected)
		}
	}
}
