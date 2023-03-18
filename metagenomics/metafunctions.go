package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println("hello")
}

/*
// function BrayCurtis takes two frequency maps as input
// and calculates the Bray-Curtis distance between the two maps as output
func BrayCurtisDistance(map1, map2 map[string]int) float64 {
	allspecies := MergeSpecies(map1, map2)
	mincount := 0
	sumcount := 0

	for _, sp := range allspecies {
		mincount += min(map1[sp], map2[sp])
		sumcount += map1[sp] + map2[sp]
	}

	avgcount := float64(sumcount) / 2
	dis := 1 - float64(mincount)/avgcount

	return dis
}

// function BrayCurtis takes two frequency maps as input
// and calculates the Jaccard distance between the two maps as output
func JaccardDistance(map1, map2 map[string]int) float64 {
	allspecies := MergeSpecies(map1, map2)
	mincount := 0
	maxcount := 0

	for _, sp := range allspecies {
		mincount += min(map1[sp], map2[sp])
		maxcount += max(map1[sp], map2[sp])
	}

	dis := 1 - float64(mincount)/float64(maxcount)

	return dis
}
*/

// funtion MergeSpecies takes two maps as input
// and returns a slice with all the keys in the two maps
func MergeSpecies(map1, map2 map[string]int) []string {
	species := make([]string, 0)

	for keys := range map1 {
		species = append(species, keys)
	}

	for keys := range map2 {
		if !ElementInSlice(keys, species) {
			species = append(species, keys)
		}

	}

	return species
}

// function ElementInSlice takes an element and a slice as input
// and tells whether the element is in the slice or not
func ElementInSlice(element string, slice []string) bool {
	elementin := false

	for _, sl := range slice {
		if element == sl {
			elementin = true
			break
		}
	}

	return elementin
}

/*
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
*/

//range over the keys in one map, check if map2[key1] is 0 or not
// _, exists := sample2[key]
//if exists{
//sum += min(sample1[key], sample2[key])

//}

// for pattern := range sample1{
//_,exists := sample2[pattern]
//if exists{
//sum  += Max2(sample1[patterm],sample2[pattern])

//}else{
// sum += sample1[pattern]
//}
//}
//then range over sample2 if a key in sample2 does not exist in sample1, sum += sample2[key]

func MultipleRichness(allmaps map[string](map[string]int)) map[string]int {
	R := make(map[string]int)

	for mapname, submap := range allmaps {
		R[mapname] = Richness(submap)
	}

	for sampleID := range allmaps {
		currentSample := allmaps[sampleID]
		currentRichness := Richness(currentSample)
		R[sampleID] = currentRichness
	}
	return R
}

func Richness(sample map[string]int) int {
	Richness := 0
	for keys := range sample {
		if keys != "" {
			Richness += 1
		}
	}

	return Richness
}

/*
func MultipleBD(allmaps map[string](map[string]int), disMetric string) ([]string, [][]float64) {
	// first sort the keys(sampleIDs) of the map

	sampleIDs := make([]string, 0)
	for sampleID := range allmaps {
		sampleIDs = append(sampleIDs, sampleID)
	}
	sort.Strings(sampleIDs)
	numSamples := len(allmaps)
	distanceMatrix := InitializeSquareMatrix(numSamples)

	for row := range distanceMatrix {
		for col := range distanceMatrix {
			freqMap1 := allmaps[samples[row]]
			freqMap2 := allmaps[samples[col]]

			if distMetric == "Jaccard" {
				distanceMatrix[row][col] = Jaccard(freqMap1, freqMap2)
			}
		}
	}

	return samples, distanceMatrix
}

func InitializeSquareMatrix(n int) [][]float64 {
	mtx := make([]int, n)
	for i := range mtx {
		mtx[n] = make([]int, n)
	}
	return mtx
}
*/

// BetaDiversityMatrix
// Input: A map of frequency maps allMaps and a string distanceMetric representing "Bray-Curtis" or "Jaccard".
// Output: A sorted list samples of the sample names, as well as a 2-D array D such that D[i][j] is the distance from samples[i] to samples[j] using the distance metric indicated by distanceMetric.
func BetaDiversityMatrix(allMaps map[string](map[string]int), distMetric string) ([]string, [][]float64) {
	//first, let's sort the keys (sample IDs) of allMaps

	samples := make([]string, 0)

	//range over the maps and append each sample ID
	for sampleID := range allMaps {
		samples = append(samples, sampleID)
	}

	//let's sort these!
	sort.Strings(samples)

	//now build the distance matrix
	numSamples := len(allMaps)

	//make a distance matrix with zero values
	distanceMatrix := InitializeSquareMatrix(numSamples)

	//range over distance matrix and set all values
	for row := range distanceMatrix {
		for col := row; col < numSamples; col++ {

			freqMap1 := allMaps[samples[row]]
			freqMap2 := allMaps[samples[col]]
			if row == col {
				distanceMatrix[row][col] = 0
			} else if distMetric == "Jaccard" {
				distanceMatrix[row][col] = JaccardDistance(freqMap1, freqMap2)
			} else if distMetric == "Bray-Curtis" {
				distanceMatrix[row][col] = BrayCurtisDistance(freqMap1, freqMap2)
			} else {
				panic("no")
			}
		}
	}

	return samples, distanceMatrix

}

func InitializeSquareMatrix(n int) [][]float64 {
	mtx := make([][]float64, n)
	for i := range mtx {
		mtx[i] = make([]float64, n)
	}
	return mtx
}

// BrayCurtisDistance computes the Bray-Curtis distance between two samples.
// Input: two frequency maps (of strings to ints)
// Output: A float64 representing the Bray-Curtis distance between these maps.
func BrayCurtisDistance(sample1, sample2 map[string]int) float64 {
	total1 := SumOfValues(sample1)
	total2 := SumOfValues(sample2)

	av := Average(float64(total1), float64(total2))

	sum := SumOfMinima(sample1, sample2)

	return 1.0 - float64(sum)/av
}

func Average(x, y float64) float64 {
	return (x + y) / 2.0
}

// JaccardDistance computes the Jaccard distance between two samples.
// Input: two frequency maps (of strings to ints)
// Output: A float64 representing the Jaccard distance between these maps.
func JaccardDistance(sample1, sample2 map[string]int) float64 {
	sumMin := SumOfMinima(sample1, sample2)
	sumMax := SumOfMaxima(sample1, sample2)

	return 1.0 - float64(sumMin)/float64(sumMax)
}

// SumOfMinima
// Input: Two frequency maps of strings to integers
// Output: Sum of minimum values over the two maps for each shared key in the frequency maps
func SumOfMinima(sample1, sample2 map[string]int) int {
	sum := 0

	// range through the keys of one of the maps.
	for key := range sample1 {
		//is this key present in the other map?

		_, exists := sample2[key] // exists = true if sample2[key] exists

		if exists { // yes
			//add the minimum to current sum
			sum += Min2(sample1[key], sample2[key])
		}
		//if no, take no action (or add zero)\
	}
	return sum
}

func Min2(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// SumOfMaxima
// Input: Two frequency maps of strings to integers
// Output: Sum of maximum values over two maps for each key present in either frequency map; if a key is present in one map but not the other, add its value to the sum.
func SumOfMaxima(sample1, sample2 map[string]int) int {
	sum := 0
	//range through all keys of sample 1
	for pattern := range sample1 {
		_, exists := sample2[pattern] //does this key occur in sample 2?
		//if yes, add max of the two values to sum
		if exists {
			sum += Max2(sample1[pattern], sample2[pattern])
		} else {
			//if no, add value of sample1[key] to sum
			sum += sample1[pattern]
		}
	}

	//range through the keys of sample 2.
	for pattern := range sample2 {
		//does this key occur in sample 1?
		_, exists := sample1[pattern]

		//if yes, take no action.
		if !exists {
			//if no, add value of sample2[key] to sum
			sum += sample2[pattern]
		}
	}

	return sum
}

func Max2(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// SumOfValues sums all values in a frequency map.
// Input: a map of strings to integers.
// Output: the sum of all values. Panics if there is a negative value.
func SumOfValues(freq map[string]int) int {
	total := 0

	for _, val := range freq {
		if val < 0 {
			panic("Error: negative value given to SumOfValues.")
		}
		total += val
	}

	return total
}

func SimpsonsIndex(freq map[string]int) float64 {
	simpson := 0.0

	//need to know the sum of all values in the frequency map
	total := SumOfValues(freq)

	//iterate over map, and square the probability of "choosing" the current element twice with replacement
	for _, val := range freq {
		probability := float64(val) / float64(total)
		simpson += probability * probability
	}

	return simpson
}

func ListMersennePrimes(n int) []int {
	var Mer []int
	for i := 1; i <= n; i++ {
		mers := int(math.Pow(2, float64(i))) - 1
		if IsPrime(mers) {
			Mer = append(Mer, mers)
		}
	}
	return Mer
}

func IsPrime(p int) bool {
	if p == 1 {
		return false
	} else {
		for k := 2; float64(k) <= math.Sqrt(float64(p)); k++ {
			if p%k == 0 {
				return false
			}
		}
	}
	return true
}
