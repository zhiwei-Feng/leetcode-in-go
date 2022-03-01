package n1117

import (
	"context"
	"fmt"
	"sync"

	"github.com/marusama/cyclicbarrier"
	"golang.org/x/sync/semaphore"
)

type H2O struct {
	semaH semaphore.Weighted
	semaO semaphore.Weighted
	barr  cyclicbarrier.CyclicBarrier
}

func (this *H2O) hydrogen() {
	this.semaH.Acquire(context.Background(), 1)
	fmt.Print("H")
	this.barr.Await(context.Background())
	this.semaH.Release(1)
}
func (this *H2O) oxygen() {
	this.semaO.Acquire(context.Background(), 1)
	fmt.Print("O")
	this.barr.Await(context.Background())
	this.semaO.Release(1)
}

func Print(water string) {
	h2o := H2O{
		semaH: *semaphore.NewWeighted(2),
		semaO: *semaphore.NewWeighted(1),
		barr:  cyclicbarrier.New(3),
	}
	var wg sync.WaitGroup
	wg.Add(len(water))
	for i := range water {
		if water[i] == 'H' {
			go func() {
				h2o.hydrogen()
				wg.Done()
			}()
		} else {
			go func() {
				h2o.oxygen()
				wg.Done()
			}()
		}
	}
	wg.Wait()
}
