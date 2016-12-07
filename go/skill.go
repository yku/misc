package main

import (
	"flag"
	"fmt"
)

func main() {
	skill_time := flag.Float64("t", 10, "skill charge time")
	flag.Parse()

	for reload := 1.0; reload <= 2.0; reload += 0.03 {
		apply_time := float32(*skill_time) / float32(reload)
		decrease := apply_time / float32(*skill_time) * 100.0
		fmt.Printf("%3d%%\tcharge_time=%.3f[sec]\t%.3f%%\n", int32(reload*100.0)-100, apply_time, decrease)
	}
}
