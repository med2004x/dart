package price

import "testing"

func TestAddTax(t *testing.T) {
    tests := []struct {
        name     string
        subtotal int
        rate     float64
        want     int
    }{
        {name: "zero tax", subtotal: 1000, rate: 0, want: 1000},
        {name: "nineteen percent", subtotal: 1000, rate: 0.19, want: 1190},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := AddTax(tt.subtotal, tt.rate)
            if got != tt.want {
                t.Fatalf("got %d, want %d", got, tt.want)
            }
        })
    }
}
