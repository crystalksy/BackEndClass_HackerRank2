func priceCheck (products [] string, productPrices [] float32, productSold[] string, soldPrice [] float32) int32 {
	wrongCount := int32 (0)
	for i:=0; i < len (productSold); i++ {
		for j := 0; j < len (products); j++ {
			if productSold[i] == products[j] {
				if productPrices[j] !== soldPrice[i] {
					wrongCount += 1
				}
			}
		}
	}
	return wrongCount
}
