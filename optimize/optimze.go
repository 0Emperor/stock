package optimizer

import (
	"fmt"
	"slices"
	"stock/parser"
)

type managment struct {
	startCycle int
	task       parser.Task
}

func Optimize(data *parser.Stock_exchange, done chan int) {
	currentCicle := 0
	mainPrcs, yes := data.FindProductResource(data.To_Optimize[len(data.To_Optimize)-1])
	CurentProcesses := []managment{}

	if yes {
		for {
			for i := 0; i < len(CurentProcesses); i++ {
				prcs := CurentProcesses[i]
				if prcs.startCycle+prcs.task.NbCycles > currentCicle {
					continue
				}
				data.EndTask(prcs.task)
				CurentProcesses = slices.Delete(CurentProcesses, i, i+1)
				i--
			}
			if !data.CheckStock(mainPrcs) {
				CurentProcesses = notYet(mainPrcs, data, currentCicle, CurentProcesses)
			} else {
				fmt.Printf("%v:%v\n", currentCicle, mainPrcs.Name)
				data.SchedualTask(mainPrcs)
				CurentProcesses = append(CurentProcesses, managment{currentCicle, mainPrcs})
			}
			currentCicle++
			if len(CurentProcesses) == 0 {
				break
			}

		}
	}
done<-currentCicle
}

func notYet(task parser.Task, data *parser.Stock_exchange, currentCicle int, whatever []managment) []managment {
	for need, amount := range task.Requirements {
		i := data.Stock[need]
		prcs, yes := data.FindProductResource(need)
		if yes {
			for _, v := range whatever {
				if v.task.Name == prcs.Name {
					i += prcs.Products[need]
				}
			}
			for i < amount {
				if !data.CheckStock(prcs) {
					return notYet(prcs, data, currentCicle, whatever)
				} else {
					fmt.Printf("%v:%v\n", currentCicle, prcs.Name)
					data.SchedualTask(prcs)
					whatever = append(whatever, managment{currentCicle, prcs})
					i += prcs.Products[need]
				}
			}
		} else if i < amount {
			return whatever
		}
	}
	return whatever
}
