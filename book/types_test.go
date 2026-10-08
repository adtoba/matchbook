package book

import "testing"

func TestSideBetter(t *testing.T) {
	tests := []struct {
		name string
		side Side
		a    int64
		b    int64
		want bool
	}{
		{side: Buy, a: 101, b: 100, name: "buy: higher price is better", want: true},
		{side: Buy, a: 100, b: 101, name: "buy: lower price is not better", want: false},
		{side: Sell, a: 99, b: 100, name: "sell: lower price is better", want: true},
		{side: Sell, a: 100, b: 99, name: "sell: higher price is not better", want: false},
		{side: Buy, a: 100, b: 100, name: "buy: equal prices not better", want: false},
		{side: Sell, a: 100, b: 100, name: "sell: equal prices not better", want: false},
		{side: SideUnknown, a: 101, b: 100, name: "unknown side: never better", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.side.better(tt.a, tt.b)

			if res != tt.want {
				t.Errorf("%v.better(%d, %d) = %v, want %v", tt.side, tt.a, tt.b, res, tt.want)
			}
		})
	}
}
