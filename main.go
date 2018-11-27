package main

import "fmt"

// Struct to hold name and votes of a single candidate
type candidate struct {
    name string
    votes []int
}

/*
 * Main function, entry point of program
 */
func main() {
	// Set up integer slice of votes
	votes := []int{0, 0, 0, 1, 1, 1, -1, -1, -1, 1}

	// Set up candidate slice of candidates
	candidates := []candidate{
		candidate{"Adams", 	 []int{ 1,  1,  1,  1,  1,  1,  1,  1,  1,  1}},
		candidate{"Grant", 	 []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1}},
		candidate{"Polk", 	 []int{ 1, -1,  1, -1,  1, -1,  1, -1,  1, -1}},
		candidate{"Jackson", []int{ 1,  0,  1,  0,  1,  0,  1,  0,  1,  0}},
		candidate{"Taft", 	 []int{ 0, -1,  0, -1,  0, -1,  0, -1,  0, -1}},
		candidate{"Ford", 	 []int{ 1,  1,  1,  1,  0,  0,  0,  0,  0,  0}},
		candidate{"Madison", []int{ 0,  0,  0,  1, -1,  0,  0, -1,  1,  1}},
	}

	// Create list of agreement scores for each candidate
	var agreement_scores []int = agreement_scores(votes, candidates)

	// Calculate the highest agreement score among all candiates
	var highest_agreement_score int = max_in_list(agreement_scores);

	// Setup empty string slice to hold names of candidates with highest
	// agreement score
	var highest_agreement_scores []string

	// Populate highest_agreement_scores slice with names of candidates with
	// the highest agreement score
	for i := 0; i < len(candidates); i++ {
		if (agreement_scores[i] == highest_agreement_score) {
			highest_agreement_scores = append(highest_agreement_scores,
											  candidates[i].name)
		}
	}

	// Print final list of candidates
	fmt.Println(highest_agreement_scores)
}

/*
 * Returns the largest integer in the list slice
 */
func max_in_list(list []int) int {
	// Set largest value to initially be the first item
	var largest int = list[0]

	// Loop through every item in the list and if it is larger than largest,
	// update largest
	for i := 0; i < len(list); i++ {
		if (list[i] > largest) {
			largest = list[i]
		}
	}

	// Return the largest item
	return largest
}

/*
 * Puts together a list of the agreement scores for each candidate in the
 * candidate slice compared to the votes slice
 */
func agreement_scores(votes []int, candidates []candidate) []int {
	// Empty slice that we will populate and return
	agreement_scores := []int{}

	// Populate the agreement_scores slice with the agreement score of each
	// candidates votes and the votes slice
	for i := 0; i < len(candidates); i++ {
		agreement_scores = append(agreement_scores,
								  agreement_score(votes, candidates[i].votes))
	}

	// Return populated list of agreement scores
	return agreement_scores
}

/*
 * Calculates the agreement score between two arrays of votes. Agreement score
 * is calculated as the difference between the number of same nonzero votes and
 * the number of different nonzero votes
 */
func agreement_score(votes1 []int, votes2 []int) int {
	// Local variables used for final calculation
	var same_answers, different_answers int

	// Loop over each vote in both arrays, determine if either same_answer or
	// different_answer should be incremented
	for i := 0; i < len(votes1); i++ {
		if (votes1[i] != 0 && votes2[i] != 0) {
			if (votes1[i] == votes2[i]) {
				same_answers++
			} else {
				different_answers++
			}
		}
	}

	// Return final agreement score calculated as different between same and
	// different answers
	return same_answers - different_answers
}