package q1

import "fmt"

func CalculateDiscount(currentPurchase float64, purchaseHistory []float64) (float64, error) {

	if currentPurchase <= 0 {
		return 0, fmt.Errorf("valor da compra inválido")
	}

	sum := 0.0
	for i := 0; i < len(purchaseHistory); i++ {
		sum += purchaseHistory[i]
	}

	discount := 0.0

	if sum/float64(len(purchaseHistory)) > 1000 {
		discount = 0.2 * currentPurchase
	} else if sum > 1000 {
		discount = 0.1 * currentPurchase
	} else if len(purchaseHistory) == 0 {
		discount = 0.1 * currentPurchase
	} else if sum <= 1000 && sum > 500 {
		discount = 0.05 * currentPurchase
	} else if sum <= 500 {
		discount = 0.02 * currentPurchase
	}

	return discount, nil
}
