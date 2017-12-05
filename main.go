package main
/*
#include <stdio.h>
int datspecialnumber() {
    return 2015;
}
*/
import "C"
import "fmt"
import "sysmeter/meter"

func main() {
	fmt.Println(C.datspecialnumber())
	fmt.Println(meter.Mem())
}
