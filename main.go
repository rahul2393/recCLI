package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	maxInputRectangle = 10
)

var (
	inputFileName *string
	intersections map[string][]Rect
)

type Rect struct {
	X, Y, W, H int64
}

type input struct {
	Rectangles []Rect `json:"rects"`
}

func init() {
	inputFileName = flag.String("input", "input.json", "input file to read rectangle cooridnates")
	flag.Parse()
}

type mat []string

func main() {
	content, err := ioutil.ReadFile(*inputFileName)
	if err != nil {
		log.Fatalf("error reading input file: %v", err)
	}
	var input input
	if err := json.Unmarshal(content, &input); err != nil {
		log.Fatalf("error unmarshalling input file: %v", err)
	}
	var minY, maxX, maxY int64
	var rectangles []Rect
	for i, r := range input.Rectangles {
		if i == maxInputRectangle {
			break
		}
		if minY > r.Y-r.H {
			minY = r.Y - r.H
		}
		if maxX < r.X+r.W {
			maxX = r.X + r.W
		}
		if maxY < r.Y {
			maxY = r.Y
		}
		rectangles = append(rectangles, r)
	}
	shift := int64(0)
	if minY < 0 {
		shift = -1 * minY
	}
	mat := printRectangles(rectangles, shift)
	//for i := 25; i >= 0; i-- {
	//	for j := 0; j <= 25; j++ {
	//		if mat[j][i] == nil {
	//			mat[j][i] = []string{}
	//		}
	//		fmt.Printf(" %v ", mat[j][i])
	//	}
	//	fmt.Println()
	//}
	intersections = computeIntersection(mat, maxX, maxY, shift)

	printIntersections()
}

func computeIntersection(mat [1000][1000]mat, maxX, maxY, shift int64) map[string][]Rect {
	intersections = map[string][]Rect{}
	for i := int64(0); i <= maxY+shift; i++ {
		for j := int64(0); j <= maxX; j++ {
			key := mat[j][i]
			if len(key) > 1 {
				w := int64(1)
				for {
					if strings.Join(mat[j+w][i], ",") != strings.Join(key, ",") {
						break
					}
					w++
				}
				h := int64(1)
				found := false
				for {
					for l := j; l < j+w; l++ {
						if strings.Join(mat[l][i+h], ",") != strings.Join(key, ",") {
							found = true
							break
						}
					}
					if found {
						break
					}
					h++
				}
				for m := int64(0); m < w; m++ {
					for n := int64(0); n < h; n++ {
						mat[j+m][i+n] = []string{}
					}
				}
				intersections[strings.Join(key, ",")] = append(intersections[strings.Join(key, ",")],
					Rect{X: j, Y: i + h - 1 - shift, W: w, H: h})
			}
		}
	}
	return intersections
}

// printRectangles prints the rectangles
func printRectangles(rectangles []Rect, shift int64) [1000][1000]mat {
	mat := [1000][1000]mat{}
	fmt.Println("Input:")
	for i, r := range rectangles {
		fmt.Printf("\t %v: Rectangle at (%v,%v), w=%v, h=%v.\n", i+1, r.X, r.Y, r.W, r.H)
		for j := r.X; j < (r.X + r.W); j++ {
			for k := r.Y - r.H + 1 + shift; k <= r.Y+shift; k++ {
				if mat[j][k] == nil {
					mat[j][k] = []string{}
				}
				mat[j][k] = append(mat[j][k], strconv.Itoa(i+1))
			}
		}
	}
	return mat
}

// printIntersections prints the intersections
func printIntersections() {
	fmt.Println("Intersections:")
	count := 1
	for k, v := range intersections {
		keys := strings.Split(k, ",")
		for _, r := range v {
			fmt.Printf("\t %v: Between rectangle %v", count, keys[0])
			count++
			for _, rk := range keys[1 : len(keys)-1] {
				fmt.Printf(", %v", rk)
			}
			fmt.Printf(" and %v", keys[len(keys)-1])
			fmt.Printf(" at (%v,%v), w=%v, h=%v.\n", r.X, r.Y, r.W, r.H)
		}
	}
}
