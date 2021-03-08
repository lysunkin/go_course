package main

import (
	"fmt"
	"sync"
	"time"
)

const maxNumberOfMeals = 3

// How many philosophers and chopsticks
const countP = 5

// how many simultaneous eaters
const maxNumberOfEaters = 2

var dining sync.WaitGroup

var hostWG sync.WaitGroup

type ChopS struct{ sync.Mutex }

type Philo struct {
	id              int
	leftCS, rightCS *ChopS
	mealsEaten      int
}

func (p *Philo) eat() {
	if p.mealsEaten == maxNumberOfMeals {
		// already fed and had all meals
		dining.Done()
		return
	}

	p.leftCS.Lock()
	p.rightCS.Lock()

	fmt.Println("starting to eat", p.id)
	time.Sleep(time.Second)

	p.rightCS.Unlock()
	p.leftCS.Unlock()

	fmt.Println("finishing eating", p.id)
	p.mealsEaten++
	time.Sleep(time.Second)
	dining.Done()
}

func host(philos []*Philo) {

	i := 1
	for c := 0; c < countP*maxNumberOfMeals; c++ {
		for n := 0; n < maxNumberOfEaters; n++ {
			dining.Add(1)
			go philos[i].eat()
			i = (i + 2) % countP
		}
		dining.Wait() // wait for simultaneous eaters to finish
	}

	hostWG.Done()
}

func main() {

	CSticks := make([]*ChopS, countP)
	for i := 0; i < countP; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, countP)
	for i := 0; i < countP; i++ {
		philos[i] = &Philo{i + 1, CSticks[i], CSticks[(i+1)%countP], 0}
	}

	hostWG.Add(1)

	go host(philos)

	hostWG.Wait()
}
