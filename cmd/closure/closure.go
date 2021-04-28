package main

/*  closure

Bevor eine Funktion (hier mkmycounter) eine neu konstruierte Funktion an den Aufrufer zurückreicht,
darf sie vorab Variablen definieren, deren Daten die generierte Funktion umschlingt
und die anschliesend nur für sie global erscheinen. So gehören die umschlossenen Variablen zur Funktion,
ähnlich wie Instanzvariablen in der objektorientierten programmierung zu einem Objekt gehören.

Thanks to mschilli@perlmeister.com
*/

import "fmt"

func main() {
	mycounter := mkmycounter()
	mycounter()
	mycounter()
	mycounter()
}

func mkmycounter() func() {
	count := 1
	return func() {
		fmt.Printf("%d\n", count)
		count++
	}
}
