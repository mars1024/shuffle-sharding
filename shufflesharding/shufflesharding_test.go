package shufflesharding

import "testing"

func TestValidateParameters(t *testing.T) {
	tests := []struct {
		name      string
		handSize  uint32
		queueSize uint32
		success   bool
	}{
		{
			"1",
			8,
			128,
			true,
		},
		{
			"2",
			200,
			512,
			false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if ValidateParameters(test.queueSize, test.handSize) != test.success{
				t.Errorf("fails %s", test.name)
				return
			}

		})
	}
}
