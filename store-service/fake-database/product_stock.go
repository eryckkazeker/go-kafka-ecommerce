package fakedatabase

var stockMap = map[string]int{
	"p1": 3,
	"p2": 5,
}

func GetProductStockById(id string) int {
	return stockMap[id]
}
