package main

import "fmt"

type IMeal interface {
	getCalories() int
	//getDescription() string
}

type PureRice struct{}
type PureMashedPotatoes struct{}

func (p *PureMashedPotatoes) getCalories() int {
	return 130
}

func (r *PureRice) getCalories() int {
	return 100
}

type TomatoGravy struct {
	meal IMeal
}

func (g *TomatoGravy) getCalories() int {
	mealCallories := g.meal.getCalories()
	return mealCallories + 50
}

type RiyalGravy struct {
	meal IMeal
}

func (g *RiyalGravy) getCalories() int {
	mealCallories := g.meal.getCalories()
	return mealCallories + 60
}

type FriedChicken struct {
	meal IMeal
}

func (f *FriedChicken) getCalories() int {
	mealCallories := f.meal.getCalories()
	return mealCallories + 230
}

type BackedChicken struct {
	meal IMeal
}

func (f *BackedChicken) getCalories() int {
	mealCallories := f.meal.getCalories()
	return mealCallories + 150
}

type BoiledChicken struct {
	meal IMeal
}

func (f *BoiledChicken) getCalories() int {
	mealCallories := f.meal.getCalories()
	return mealCallories + 120
}

func main() {
	meal1 := &PureRice{}
	meal2 := &PureMashedPotatoes{}

	riceWithRiyalGravy := &RiyalGravy{meal: meal1}
	riceWithFriedChicken := &FriedChicken{meal: riceWithRiyalGravy}

	mashedPotatoesWithRiyalGravy := &RiyalGravy{meal: meal2}
	mashedPotatoesWithFriedChicken := &FriedChicken{meal: mashedPotatoesWithRiyalGravy}

	fmt.Printf("rice with riyal gravy and fried chicken will have %v callories", riceWithFriedChicken.getCalories())

	fmt.Printf("\n\nmashed potatoes with riyal gravy and fried chicken will have %v callories", mashedPotatoesWithFriedChicken.getCalories())

}
