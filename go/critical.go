package main

import (
	"flag"
	"fmt"
	"sort"
)

/*
x ... kill 0.08 cri 0.06
y ... cri 0.09
z ... kill 0.12
x + y + z = 10
*/

type Data struct {
	critical    int
	human       int
	ex_critical float32
	buf_rate    float32
	damage_rate float32
	crew_x      int
	crew_y      int
	crew_z      int
}
type Datas []Data

func (d Datas) Len() int {
	return len(d)
}

func (d Datas) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d Datas) Less(i, j int) bool {
	return d[i].damage_rate < d[j].damage_rate
}

func main() {
	crews := flag.Int("n", 10, "max crews")
	flag.Parse()

	fmt.Println("crew_x = human killer  8% critical 6% .")
	fmt.Println("crew_y = human killer  0% critical 9% .")
	fmt.Println("crew_z = human killer 12% critical 0% .")
	d := Datas{}
	for x := 0; x <= *crews; x++ {
		for y := 0; y <= *crews-x; y++ {
			for z := 0; z <= *crews-(x+y); z++ {
				cri := 6*x + 9*y + 0*z
				exp := float32(1*(100-cri)+3*cri) / 100.0
				human := 8*x + 0*y + 12*z
				buf := float32(human+100) / 100.0
				delta := buf * exp
				if x+y+z == *crews {
					t := Data{}
					t.critical = cri
					t.human = human
					t.ex_critical = exp
					t.buf_rate = buf
					t.damage_rate = delta
					t.crew_x = x
					t.crew_y = y
					t.crew_z = z
					d = append(d, t)
				}
			}
		}
	}
	sort.Sort(d)
	fmt.Printf("cri kill e_cri buf dmg(normal) dmg(good) dmg(bad) crew_x crew_y crew_z\n")
	for i := range d {
		color_good := d[i].buf_rate * 1.5
		if color_good >= 3.0 {
			color_good = 3.0
		}
		color_normal := d[i].buf_rate
		if color_normal >= 3.0 {
			color_normal = 3.0
		}
		color_bad := d[i].buf_rate * 0.5
		if color_bad >= 3.0 {
			color_bad = 3.0
		}

		fmt.Printf("%2d%% %3d%% %.3f %.3f %.3f %.3f %.3f %6d %6d %6d\n",
			d[i].critical, d[i].human, d[i].ex_critical, d[i].buf_rate,
			d[i].ex_critical*color_normal,
			d[i].ex_critical*color_good,
			d[i].ex_critical*color_bad,
			d[i].crew_x, d[i].crew_y, d[i].crew_z)
	}
}
