package main

import (
	"fmt"
	"math/rand"
	"time"
)


type Shop struct {
	Verbose        bool
	Cakes          int           // number of cakes to bake
	BakeTime       time.Duration // time to bake one cake
	BakeStdDev     time.Duration // standard deviation of baking time
	BakeBuf        int           // buffer slots between baking and icing
	NumIcers       int           // number of cooks doing icing
	IceTime        time.Duration // time to ice one cake
	IceStdDev      time.Duration // standard deviation of icing time
	IceBuf         int           // buffer slots between icing and inscribing
	InscribeTime   time.Duration // time to inscribe one cake
	InscribeStdDev time.Duration // standard deviation of inscribing time
}

type cake int

func (s *Shop) baker(baked chan<- cake) {
	for i := 0; i < s.Cakes; i++ {
		c := cake(i)
		if s.Verbose {
			fmt.Println("backing", c)
		}
		work(s.BakeTime, s.BakeStdDev)
		baked <- c
	}
	close(baked)
}

func (s *Shop) icer(iced chan<- cake, baked <-chan cake) {
	for c := range baked {
		if s.Verbose {
			fmt.Println("icing", c)
		}
		work(s.IceTime, s.IceStdDev)
		iced <- c
	}
}

func (s *Shop) inscriber(iced <-chan cake) {
	for i := 0; i < s.Cakes; i++ {
		c := <-iced
		if s.Verbose {
			fmt.Println("insribing", c)
		}

		work(s.InscribeTime, s.InscribeStdDev)
		if s.Verbose {
			fmt.Println("finished", c)
		}
	}
}

func (s *Shop) Work(runs int) {
	for run := 0; run < runs; run++ {
		baked := make(chan cake, s.BakeBuf)
		iced := make(chan cake, s.IceBuf)
		go s.baker(baked)
		for i := 0; i < s.NumIcers; i++ {
			go s.icer(iced, baked)
		}
		s.inscriber(iced)
	}
}

func work(d, stddev time.Duration) {
	delay := d + time.Duration(rand.NormFloat64()*float64(stddev))
	time.Sleep(delay)
}

func main() {
	var defaults = Shop{
		Verbose:      true,
		Cakes:        20,
		BakeTime:     10 * time.Millisecond,
		NumIcers:     1,
		IceTime:      10 * time.Millisecond,
		InscribeTime: 10 * time.Millisecond,
	}
	defaults.Work(5)
}