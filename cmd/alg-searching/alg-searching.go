package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/leptitnoir/go-boyer-alg/packages/searching"
)

func main() {
	text := "coucou uco dzauiouco hedzuyikg fukyegzzv yhjfgev ukgvf ezuktygv fdukzegbv dukgzv egykjtgcv zy bceky gcekuyezioudze gcuikyzeg ilu hcuyzeg duloygze cuyzeg dhgv dx kjho_e yrier gged icu hez"
	key := "ou"
	wrapper := &searching.SearchingContext{}
	timer := time.Now().Nanosecond()
	wrapper.SetAlgorithm(&searching.BoyerMooreHorspoolSearch{})

	fmt.Println(wrapper.FindAllString(text, key))
	fmt.Printf("Boyer Moore Horspool algorithm : %d ms\n", (time.Now().Nanosecond()-timer)/1000)

	timer = time.Now().Nanosecond()
	wrapper.SetAlgorithm(&searching.NaiveSearch{})
	fmt.Println(wrapper.FindAllString(text, key))
	fmt.Printf("Naive search algorithm : %d ms\n", (time.Now().Nanosecond()-timer)/1000)

	timer = time.Now().Nanosecond()
	strings.Contains(text, key)
	fmt.Printf("Native search algorithm : %d ms\n", (time.Now().Nanosecond()-timer)/1000)

}
