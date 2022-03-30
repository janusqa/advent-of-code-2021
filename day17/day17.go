package day17

import (
	"adventofcode/aocutils"
	"fmt"
	"regexp"
)

var input []string

type target struct {
	xmin int
	xmax int
	ymin int
	ymax int
}

type probe struct {
	x  []int
	y  []int
	xv []int
	yv []int
	mh int
}

func PartI(filename string) {
	input = aocutils.GetInputFromFile(filename)

	re_target_area, _ := regexp.Compile(`x=(-?\d{1,})\.\.(-?\d{1,}), y=(-?\d{1,})\.\.(-?\d{1,})$`)
	target_area := re_target_area.FindStringSubmatch(input[0])

	t := target{
		xmin: aocutils.StringToInt(target_area[1], 10),
		xmax: aocutils.StringToInt(target_area[2], 10),
		ymin: aocutils.StringToInt(target_area[3], 10),
		ymax: aocutils.StringToInt(target_area[4], 10),
	}

	p := probe{
		x:  []int{0},
		y:  []int{0},
		xv: []int{0},
		yv: []int{0},
		mh: 0,
	}

	// xv := 0
	// for (((xv + 1) * xv) / 2) <= t.xmax {
	// 	xv_current := ((xv + 1) * xv) / 2
	// 	if t.xmin <= xv_current && xv_current <= t.xmax {
	// 		fmt.Println(xv, xv_current)
	// 	}
	// 	fmt.Println(xv, xv_current)
	// 	xv += 1
	// }

	// yv := 0
	// for t.ymin <= (((yv+1)*yv)/2)+(((-1+t.ymin)*aocutils.Abs(t.ymin))/2) {
	// 	yv_current := (((yv + 1) * yv) / 2) + (((-1 + t.ymin) * aocutils.Abs(t.ymin)) / 2)
	// 	fmt.Println(yv, yv_current)
	// 	if t.ymin <= yv_current && yv_current <= t.ymax {
	// 		fmt.Println(yv, yv_current)
	// 	}
	// 	yv += 1
	// }

	mh := 0
	successful_bullseyes := 0
	for ixv := 0; ixv <= t.xmax; ixv++ {
		for iyv := t.ymin; iyv <= aocutils.Abs(t.ymin)-1; iyv++ {
			if !(ixv == 0 && iyv == 0) {
				p.xv = nil
				p.yv = nil
				p.x = nil
				p.y = nil
				p.x = append(p.x, 0)
				p.y = append(p.y, 0)
				p.xv = append(p.xv, ixv)
				p.yv = append(p.yv, iyv)
				p.mh = 0
				if bullseye(&p, &t) {
					successful_bullseyes += 1
					if p.mh > mh {
						mh = p.mh
					}
				}
			}
		}
	}

	fmt.Println("Max Height :", mh)
	fmt.Println("# bullseyes:", successful_bullseyes)
}

func bullseye(p *probe, t *target) bool {

	x := (*p).x[len((*p).x)-1]
	y := (*p).y[len((*p).y)-1]
	xv := (*p).xv[len((*p).xv)-1]
	yv := (*p).yv[len((*p).yv)-1]

	if x > (*t).xmax || y < (*t).ymin {
		return false
	}

	if ((*t).xmin <= x) && (x <= (*t).xmax) && ((*t).ymin <= y) && (y <= (*t).ymax) {
		return true
	}

	if !(x == 0 && y == 0) {
		if xv > 0 {
			xv -= 1
		} else if xv < 0 {
			xv += 1
		}
		yv -= 1
	}
	x += xv
	y += yv

	if y > (*p).mh {
		(*p).mh = y
	}
	(*p).x = append((*p).x, x)
	(*p).y = append((*p).y, y)
	(*p).xv = append((*p).xv, xv)
	(*p).yv = append((*p).yv, yv)

	return bullseye(p, t)
}
