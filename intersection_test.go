package main

import (
	"reflect"
	"testing"
)

// 	y4	[]  []  []   []    []   []
// 	y3	[]  []  []   [2]   [2]  [2]
// 	y2	[]  []  [1]  [1 2] [2]  [2]
// 	y1	[]  []  [1]  [1 2] [2]  [2]
// 	y0	[]  []  []   []    []   []
//		x0  x1  x2    x3   x4   x5
func TestIntersection(t *testing.T) {
	type fields struct {
		rect []Rect
	}
	tests := []struct {
		name           string
		fields         fields
		expectedResult map[string][]Rect
	}{
		{
			name: "intersection on 2 rectangles",
			fields: fields{
				rect: []Rect{{
					X: 2,
					Y: 2,
					H: 2,
					W: 2,
				}, {
					X: 3,
					Y: 3,
					H: 3,
					W: 3,
				}},
			},
			expectedResult: map[string][]Rect{
				"1,2": {
					{X: 3, Y: 2, H: 2, W: 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mat := printRectangles(tt.fields.rect, 0)
			r := computeIntersection(mat, 5, 5, 0)
			if !reflect.DeepEqual(r, tt.expectedResult) {
				t.Fatalf("expected intersection to be %+v got %+v", tt.expectedResult, r)
			}
		})
	}
}
