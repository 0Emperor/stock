package optimizer

import (
	"fmt"
	"slices"
	"stock/parser"
	tfm "stock/stock"
)

type managment struct {
	startCycle int
	task       parser.Task
}

func Optimze(data *parser.BuildData) {
	currentCicle := 0
	mainPrcs, yes := data.FindProductResource(data.Optimize[0])
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

	fmt.Println("no more thing to do at cicle ", currentCicle)
	tfm.Println(data)
}
func notYet(task parser.Task, data *parser.BuildData, currentCicle int, whatever []managment) []managment {
	for need, amount := range task.Requirements {
		i := data.Stock[need]
		prcs, yes := data.FindProductResource(need)
		if yes {
			for i < amount {
				if !data.CheckStock(prcs) {
					return notYet(prcs, data, currentCicle, whatever)
				} else {
					fmt.Printf("%v:%v\n", currentCicle, prcs.Name)
					data.SchedualTask(prcs)
					whatever = append(whatever, managment{currentCicle, prcs})
					i++
				}
			}
		} else if i < amount {
			return whatever
		}
	}
	return whatever
}
