package main

import "fmt"

/*
x ... kill 0.08 cri 0.06
y ... cri 0.09
z ... kill 0.12
x + y + z = 10
*/

func main() {
	fmt.Println("x := human 8% critical 6% crew.")
	fmt.Println("y := human 0% critical 9% crew.")
	fmt.Println("z := human 12% critical 0% crew.")
	for x := 0; x <= 10; x++ {
		for y := 0; y <= 10-x; y++ {
			for z := 0; z <= 10-(x+y); z++ {
				//cri := 0.06*x + 0.09*y + 0.0*z
				cri := 6*x + 9*y + 0*z
				exp := float32(1*(100-cri)+3*cri) / 100.0
				human := 8*x + 0*y + 12*z
				buf := float32(human+100) / 100.0
				delta := buf * exp
				//fmt.Printf("cri=%d%% exp_cri=%f buf=%f rate=%f rate*0.5=%f rate*1.5=%f (x=%d y=%d z=%d)\n", cri, exp, buf, delta, delta*0.5, delta*1.5, x, y, z)
				if x+y+z == 10 {
					fmt.Printf("%d [cri=%02d%% human=%03d%%] expected_cri=%.2f buf=%.2f rate=%.2f rate*0.5=%.2f rate*1.5=%.2f (x=%d y=%d z=%d)\n", cri, cri, human, exp, buf, delta, delta*0.5, delta*1.5, x, y, z)
				}
			}
		}
	}

}
