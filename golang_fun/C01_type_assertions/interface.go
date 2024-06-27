package main

import (
	"fmt"
)

type Guitarist interface {
	PlayGuitar()
}

type BassGuitarist struct {
	name string
}

type AcousticGuitarist struct {
	name string
}

func (b BassGuitarist) PlayGuitar() {
	fmt.Printf("%s plays bejs guitar.\n Bam bam bam\n", b.name)
}

func (a AcousticGuitarist) PlayGuitar() {
	fmt.Printf("%s plays acoustic guitar.\n Dring dring drin.\n", a.name)
}

func main() {

	var player01 BassGuitarist
	player01.name = "Maks"
	var player02 AcousticGuitarist
	player02.name = "Mirko"

	player01.PlayGuitar()
	player02.PlayGuitar()

	var guitarists []Guitarist
	guitarists = append(guitarists, player01)
	guitarists = append(guitarists, player02)

	fmt.Printf("The lenght of guitarists is %d\n", len(guitarists))
	fmt.Printf("The cap of guitarists is %d\n", cap(guitarists))

	for _, guitarist := range guitarists {
		fmt.Printf("%s\n", guitarist)
	}
}
