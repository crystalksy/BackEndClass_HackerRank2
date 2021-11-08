func funWithAnagrams (text [] string) [] string {
	intText := []int{}
	for_,val := range text {
		runeSum := 0
		for_,char := range val {
			runeSum += int (char)
		}
		intText = append(intText,runeSum)
	}
	for i, val := range intText {
		for j := i+1; j < len(intText); j++ {
			if val == intText[j]{
				intText = append(intText[:j], intText[j+1:]...)
				j = i
			}
		}
	}
	result := [] string{}
	for i := 0; i < len (intText); i++{
		runeSum := 0
		for_,char := range text [i] {
			runeSum += int (char)
		}
		if runeSum == intText [i] {
			result = append (result,text[i])
		}
	}
}


