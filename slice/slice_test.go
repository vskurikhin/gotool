package slice

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

type planet struct {
	Name   string
	Radius int
}

func TestAll(t *testing.T) {
	isEven := func(number int) bool {
		return number%2 == 0
	}
	assertEqual(t, All([]int{1, 2, 4, 6, 8}, isEven), false)
	assertEqual(t, All([]int{2, 4, 6, 8}, isEven), true)
}

func TestAny(t *testing.T) {
	isEven := func(number int) bool {
		return number%2 == 0
	}
	assertEqual(t, Any([]int{1, 3, 5, 7, 9}, isEven), false)
	assertEqual(t, Any([]int{1, 3, 2, 7, 9}, isEven), true)
}

func TestAt(t *testing.T) {
	colours := []string{"Cyan", "Magenta", "Yellow"}
	assertEqual(t, At(colours, 1, "Black"), "Magenta")
	assertEqual(t, At(colours, 10, "Black"), "Black")
}

func TestConcat(t *testing.T) {
	colours := []string{"Cyan", "Magenta", "Yellow", "Black"}
	assertEqual(t, Concat([]string{"Cyan", "Magenta"}, []string{"Yellow", "Black"}), colours)
}

func TestCount(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	assertEqual(t, Count(numbers), 9)
}

func TestCountBy(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	isEven := func(number int) bool {
		return number%2 == 0
	}
	assertEqual(t, CountBy(numbers, isEven), 4)
}

func TestEach(t *testing.T) {
	countdown := []string{"3", "2", "1", "Go!"}
	Each(countdown, func(tick string) { fmt.Println(tick) })
}

func TestFilter(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	got := Filter(numbers, func(number int) bool {
		return number%2 == 0
	})
	assertEqual(t, got, []int{2, 4, 6, 8})
}

func TestFlatMap(t *testing.T) {
	numbers := []int{1, 2, 3}
	assertEqual(t, FlatMap(numbers, func(number int) []int {
		return []int{number, number}
	}), []int{1, 1, 2, 2, 3, 3})
}

func TestFrequencies(t *testing.T) {
	frequencies := Frequencies([]string{"aa", "aa", "bb", "cc"})
	expected := map[string]int{"aa": 2, "bb": 1, "cc": 1}
	assertEqual(t, frequencies, expected)
}

func TestFrequenciesBy(t *testing.T) {
	frequencies := FrequenciesBy([]string{"aa", "aA", "bb", "cc"}, func(element string) string {
		return strings.ToLower(element)
	})
	expected := map[string]int{"aa": 2, "bb": 1, "cc": 1}
	assertEqual(t, frequencies, expected)
}

func TestGroupBy(t *testing.T) {
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	expected := map[int][]string{7: {"Mercury", "Jupiter", "Neptune"}, 6: {"Saturn", "Uranus"}, 5: {"Venus", "Earth"}, 4: {"Mars"}}
	got := GroupBy(planets, func(planet string) int {
		return len(planet)
	})
	assertEqual(t, got, expected)
}

func TestIsMember(t *testing.T) {
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	assertEqual(t, IsMember(planets, "Earth"), true)
	assertEqual(t, IsMember(planets, "Pluto"), false)
}

func TestIsMemberBy(t *testing.T) {
	neptune := planet{Name: "Neptune", Radius: 24_622_000}
	neptuneWithDifferentRadius := planet{Name: "Neptune", Radius: 42_000}
	mars := planet{Name: "Mars", Radius: 3_389_500}
	jupiter := planet{Name: "Jupiter", Radius: 69_911_000}
	planets := []planet{neptune, mars, jupiter}

	assertEqual(t, IsMemberBy(planets, neptuneWithDifferentRadius, func(planet planet) string {
		return planet.Name
	}), true)
}

func TestMap(t *testing.T) {
	trafficLights := []string{"red", "amber", "green"}
	got := Map(trafficLights, func(light string) string {
		return light + "!"
	})
	expected := []string{"red!", "amber!", "green!"}
	assertEqual(t, got, expected)

	numbers := []int{1, 2, 3}
	numbersAsStrings := Map(numbers, func(number int) string {
		return strconv.Itoa(number)
	})
	assertEqual(t, numbersAsStrings, []string{"1", "2", "3"})
}

func TestMax(t *testing.T) {
	numbers := []int{6, 4, 8, 2, 1, 9, 4, 7, 5}
	assertEqual(t, Max(numbers), 9)
}

func TestMaxBy(t *testing.T) {
	neptune := planet{Name: "Neptune", Radius: 24_622_000}
	mars := planet{Name: "Mars", Radius: 3_389_500}
	jupiter := planet{Name: "Jupiter", Radius: 69_911_000}

	planets := []planet{neptune, mars, jupiter}
	maxPlanet := MaxBy(planets, func(planet planet) int {
		return planet.Radius
	})
	assertEqual(t, maxPlanet.Name, jupiter.Name)
}

func TestMin(t *testing.T) {
	numbers := []int{6, 4, 8, 2, 1, 9, 4, 7, 5}
	assertEqual(t, Min(numbers), 1)
}

func TestMinBy(t *testing.T) {
	neptune := planet{Name: "Neptune", Radius: 24_622_000}
	mars := planet{Name: "Mars", Radius: 3_389_500}
	jupiter := planet{Name: "Jupiter", Radius: 69_911_000}

	planets := []planet{neptune, mars, jupiter}
	minPlanet := MinBy(planets, func(planet planet) int {
		return planet.Radius
	})
	assertEqual(t, minPlanet.Name, mars.Name)
}

func TestMinMax(t *testing.T) {
	numbers := []int{6, 4, 8, 2, 1, 9, 4, 7, 5}
	minimum, maximum := MinMax(numbers)
	assertEqual(t, minimum, 1)
	assertEqual(t, maximum, 9)
}

func TestMinMaxBy(t *testing.T) {
	neptune := planet{Name: "Neptune", Radius: 24_622_000}
	mars := planet{Name: "Mars", Radius: 3_389_500}
	jupiter := planet{Name: "Jupiter", Radius: 69_911_000}

	planets := []planet{neptune, mars, jupiter}
	minPlanet, maxPlanet := MinMaxBy(planets, func(planet planet) int {
		return planet.Radius
	})
	assertEqual(t, minPlanet.Name, mars.Name)
	assertEqual(t, maxPlanet.Name, jupiter.Name)
}

func TestProduct(t *testing.T) {
	assertEqual(t, Product([]int{2, 3, 4}), 24)
	assertEqual(t, Product([]int{2.0, 3.0, 4.0}), 24.0)
	assertEqual(t, Product([]int{42}), 42)
}

func TestProductBy(t *testing.T) {
	neptune := planet{Name: "Neptune", Radius: 2}
	mars := planet{Name: "Mars", Radius: 3}
	jupiter := planet{Name: "Jupiter", Radius: 4}
	assertEqual(t, ProductBy([]planet{neptune, mars, jupiter}, func(planet planet) int {
		return planet.Radius
	}), 24)
}

func TestRandom(t *testing.T) {
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	assertEqual(t, Random(planets, 42), "Venus")
}

func TestReduce(t *testing.T) {
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	got := Reduce(planets, func(planet string, acc string) string {
		return acc + planet
	}, "")
	expected := "MercuryVenusEarthMarsJupiterSaturnUranusNeptune"
	assertEqual(t, got, expected)
}

func TestReduceWhile(t *testing.T) {
	numbers := []int{40, 2, 8}
	got := ReduceWhile(numbers, func(number int, total int) (Reduction, int) {
		if total >= 42 {
			return Halt, total
		}
		return Cont, number + total
	}, 0)
	assertEqual(t, got, 42)
}

func TestReject(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	got := Reject(numbers, func(number int) bool {
		return number%2 == 0
	})
	assertEqual(t, got, []int{1, 3, 5, 7, 9})
}

func TestReverse(t *testing.T) {
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	expected := []string{"Neptune", "Uranus", "Saturn", "Jupiter", "Mars", "Earth", "Venus", "Mercury"}
	assertEqual(t, Reverse(planets), expected)
}

func TestShuffle(t *testing.T) {
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	assertEqual(t, Shuffle(planets, 42), []string{"Saturn", "Neptune", "Jupiter", "Uranus", "Venus", "Mars", "Mercury", "Earth"})
}

func TestSort(t *testing.T) {
	numbers := []int{5, 6, 1, 3, 7, 8, 2, 4, 9}
	assertEqual(t, Sort(numbers, Asc), []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	assertEqual(t, Sort(numbers, Desc), []int{9, 8, 7, 6, 5, 4, 3, 2, 1})
}

func TestSortBy(t *testing.T) {
	neptune := planet{Name: "Neptune", Radius: 24_622_000}
	mars := planet{Name: "Mars", Radius: 3_389_500}
	jupiter := planet{Name: "Jupiter", Radius: 69_911_000}

	planets := []planet{neptune, mars, jupiter}

	assertEqual(t, SortBy(planets, func(planet planet) int {
		return planet.Radius
	}, Asc), []planet{mars, neptune, jupiter})

	assertEqual(t, SortBy(planets, func(planet planet) int {
		return planet.Radius
	}, Desc), []planet{jupiter, neptune, mars})
}

func TestSplitWhile(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	left, right := SplitWhile(numbers, func(number int) bool {
		return number <= 5
	})
	assertEqual(t, left, []int{1, 2, 3, 4, 5})
	assertEqual(t, right, []int{6, 7, 8, 9})
}

func TestSplitWith(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	evenNumbers, oddNumbers := SplitWith(numbers, func(number int) bool {
		return number%2 == 0
	})
	assertEqual(t, evenNumbers, []int{2, 4, 6, 8})
	assertEqual(t, oddNumbers, []int{1, 3, 5, 7, 9})
}

func TestSum(t *testing.T) {
	numbers := []int{6, 4, 8, 2, 1, 9, 4, 7, 5}
	assertEqual(t, Sum(numbers), 46)
}

func TestSumBy(t *testing.T) {
	neptune := planet{Name: "Neptune", Radius: 24_622_000}
	mars := planet{Name: "Mars", Radius: 3_389_500}
	jupiter := planet{Name: "Jupiter", Radius: 69_911_000}

	planets := []planet{neptune, mars, jupiter}
	sum := SumBy(planets, func(planet planet) int {
		return planet.Radius
	})
	assertEqual(t, sum, mars.Radius+jupiter.Radius+neptune.Radius)
}

func TestTake(t *testing.T) {
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	assertEqual(t, Take(planets, 2), []string{"Mercury", "Venus"})
	assertEqual(t, Take(planets, 10), planets)
	assertEqual(t, Take(planets, 0), []string{})
}

func TestTakeWhile(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	assertEqual(t, TakeWhile(numbers, func(number int) bool {
		return number <= 5
	}), []int{1, 2, 3, 4, 5})
}

func TestUniq(t *testing.T) {
	moves := []string{"Up", "Down", "Up", "Up", "Down", "Left", "Right", "Right", "Right", "Left"}
	assertEqual(t, Uniq(moves), []string{"Up", "Down", "Left", "Right"})
}

func TestUniqBy(t *testing.T) {
	neptune := planet{Name: "Neptune", Radius: 24_622_000}
	mars := planet{Name: "Mars", Radius: 3_389_500}
	sameRadiusAsMars := planet{Name: "Same Radius as Mars", Radius: 3_389_500}

	planets := []planet{mars, neptune, sameRadiusAsMars}
	assertEqual(t, UniqBy(planets, func(planet planet) int {
		return planet.Radius
	}), []planet{mars, neptune})
}

func assertEqual[T any](t *testing.T, got T, expected T) {
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("\n     got: %v\nexpected: %v\n", got, expected)
	}
}

//lint:ignore U1000 Ignore unused function temporarily
//goland:noinspection GoUnusedFunction
func assertNotEqual[T any](t *testing.T, got T, expected T) {
	if reflect.DeepEqual(got, expected) {
		t.Errorf("\n     got: %v\nexpected: %v\n", got, expected)
	}
}
