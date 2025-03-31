package parser

type Task struct {
	Name         string
	Requirements map[string]int
	Products     map[string]int
	NbCycles     int
}

type Stock_exchange struct {
	Stock     map[string]int
	Tasks []Task
	To_Optimize  []string
}

func (b *Stock_exchange) SchedualTask(task Task) {
	for req, qua := range task.Requirements {
		b.Stock[req] -= qua
	}
}
func (b *Stock_exchange) CheckStock(task Task) bool {
	for req, qua := range task.Requirements {
		if b.Stock[req]-qua <0 {
			return false
		}
	}
	return true
}
func (b *Stock_exchange) EndTask(task Task) {
	for prd, qua := range task.Products {
		b.Stock[prd] += qua
	}
}
func (b *Stock_exchange) FindProductResource(product string) (Task,bool ){
	Task := Task{}
	current := 0
	for _, prcs := range b.Tasks {
		if amount, exists := prcs.Products[product]; exists && amount > current {
			Task = prcs
			current =amount
		}
	}
	if current==0 {
		return Task,false
	}
	return Task,true
}
