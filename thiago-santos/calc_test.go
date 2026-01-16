
// calc_test.go – calculator unit tests
package calculator

import "testing"

func TestAdd(t *testing.T) {
    tests := []struct{ a, b, want float64 }{
        {2, 3, 5}, {-1, 4, 3}, {0, 0, 0},
    }
    for _, tt := range tests {
        if got := Add(tt.a, tt.b); got != tt.want {
            t.Errorf("Add(%v,%v) = %v, want %v", tt.a, tt.b, got, tt.want)
        }
    }
}

func TestSub(t *testing.T) {
    tests := []struct{ a, b, want float64 }{
        {5, 3, 2}, {-1, -4, 3}, {0, 0, 0},
    }
    for _, tt := range tests {
        if got := Sub(tt.a, tt.b); got != tt.want {
            t.Errorf("Sub(%v,%v) = %v, want %v", tt.a, tt.b, got, tt.want)
        }
    }
}

func TestMul(t *testing.T) {
    tests := []struct{ a, b, want float64 }{
        {2, 3, 6}, {-1, 4, -4}, {0, 5, 0},
    }
    for _, tt := range tests {
        if got := Mul(tt.a, tt.b); got != tt.want {
            t.Errorf("Mul(%v,%v) = %v, want %v", tt.a, tt.b, got, tt.want)
        }
    }
}

func TestDiv(t *testing.T) {
    // casos válidos
    if res, err := Div(6, 2); err != nil || res != 3 {
        t.Errorf("Div(6,2) = %v,%v; want 3,nil", res, err)
    }

    // divisão por zero
    if _, err := Div(1, 0); err == nil {
        t.Error("Div(1,0) expected error, got nil")
    }
}
