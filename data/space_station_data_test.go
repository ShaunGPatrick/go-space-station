package data

import "testing"

func TestGetLat(t *testing.T) {

	tests := []struct {
		name string
		got  float64
		want float64
	}{
		{
			name: "Test lat",
			got:  GetLat(),
			want: 1.123,
		},
		{
			name: "Test long",
			got:  GetLong(),
			want: 14.123,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			if test.got != test.want {
				t.Errorf("got %f, want %f", test.got, test.want)
			}
		})
	}

}
