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
		{30001, 1,     1,     UnknownTriangle}, 
        {1,     20001, 1,     UnknownTriangle}, 
        {1,     1,     10001, UnknownTriangle}, 
        {0,     1,     1,     UnknownTriangle}, 
        {1,     0,     1,     UnknownTriangle}, 
        {1,     1,     0,     UnknownTriangle}, 
        {-1,    1,     1,     UnknownTriangle},
        {1,    -1,     1,     UnknownTriangle}, 
        {1,     1,    -1,     UnknownTriangle}, 

        // ==== InvalidTriangle (triangle inequality fails) ====
        {1, 1, 2, InvalidTriangle}, 
        {2, 1, 1, InvalidTriangle},  
        {1, 2, 1, InvalidTriangle}, 

        // ==== RightTriangle (only when a*a == b*b + c*c) ====
        {5, 3, 4, RightTriangle},    
        {13, 5, 12, RightTriangle},
        {17, 8, 15, RightTriangle},  

        // ==== AcuteTriangle (only when a*a < b*b + c*c) ====
        {2, 3, 4, AcuteTriangle},    
        {3, 4, 2, AcuteTriangle},   
        {5, 5, 5, AcuteTriangle},   
        // ==== ObtuseTriangle (only when a*a > b*b + c*c) ====
        {7, 3, 5, ObtuseTriangle}, 
        {9, 4, 7, ObtuseTriangle},  
        {29999, 20000, 10000, ObtuseTriangle}, // large obtuse

	}

	for _, test := range tests {
		actual := getTriangleType(test.a, test.b, test.c)
		if actual != test.expected {
			t.Errorf("getTriangleType(%d, %d, %d)=%v; want %v", test.a, test.b, test.c, actual, test.expected)
		}
	}
}
