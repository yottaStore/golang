package block

import "testing"

func TestSize(t *testing.T) {

	tests := [][3]int{
		{1, 1, 1},
		{4080, 1, 4080},
		{4085, 2, 4085},
		{4088, 2, 0},
		{4090, 2, 2},
		{8168, 2, 4080},
		{8168, 2, 4080},
		{8176, 3, 0},
	}

	for _, test := range tests {

		s, r := GetSize(test[0])
		if s != test[1] {
			t.Error("Expected count not matched: ", test[0], s, test[1])
		}
		if r != test[2] {
			t.Error("Expected remainder not matched: ", test[0], r, test[2])
		}

	}

}