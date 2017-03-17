package main

import (
	"fmt"
	"errors"
)

func main() {

	/*

	Input Format
	=============
	The first line contains two space-separated integers describing the respective values of n and t.
	 n = bowl_size, t = end_party_phase (inclusive)
	The second line contains  space-separated integers describing the respective values of c0, c1, ct-1.

	 */
	// INput
	bowl_size := 0
	end_party_phase := 0
	fmt.Scanf("%d %d", &bowl_size, &end_party_phase)
	// DEBUG:
	// fmt.Println("BOWL_SIZE: ", bowl_size, " PARTY ENDS MIN: ", end_party_phase)
	// Init structs
	party_today := party{end_party_phase, 0}
	current_candy_bowl := candybowl{bowl_size, bowl_size, 0}
	// DEBUG:
	// fmt.Println("PARTY: ", party_today)
	// fmt.Println("BOWL:", current_candy_bowl)
	// Input candies taken; size=t, c0,c1,ct-1
	timeline_candies_taken := make([]int, end_party_phase)
	for iter := 0; iter < end_party_phase; iter++ {
		fmt.Scanf("%d", &timeline_candies_taken[iter])
		// Execute here no need twice!!
		// DEBUG:
		// fmt.Println("Number of Candies STOLEN: ", timeline_candies_taken[iter])
		ProcessPhase(timeline_candies_taken[iter], &party_today, &current_candy_bowl)
		// Init CandyBowl; attach it to Party
		// p.GuestTakeFromBowl(number_of_candies)
		//	p.c.TakeFromCandyBowl(number_of_candies)
		// p.RobotInspectsBowl(current_phase_of_party)
		//	if NOT p.IsEndOfParty(current_phase_of_party) &&
		// 	if p.IsRefillNeededForCurrentCandyBowl() [[ p.c.IsRefillNeeded() ]]
		//		p.NoteCandiesAddedToCurrentCandyBowl(p.c.Refill())
		// p.TotalCandiesAddedDuringParty
	}
	fmt.Println(current_candy_bowl.number_of_refills)
	// DEBUG:
	// fmt.Println("TOTAL_REFILL: ", current_candy_bowl.number_of_refills)
}

type Party struct {
	total_duration_party_minutes int
	total_candies_added          int
	current_candy_bowl           CandyBowl
}

type CandyBowl struct {
	bowl_size         int
	number_of_candies int
}

type party struct {
	end_party_phase     int
	current_party_phase int
}

type candybowl struct {
	bowl_size         int
	number_of_candies int
	number_of_refills int
}

func ProcessPhase(number_of_candies_stolen int, party_today *party, current_candy_bowl *candybowl) {
	// In this phase; number of stolen candies ...
	// number_of_candies_stolen = 2
	should_refill, myerr := current_candy_bowl.ExecuteGuestStealingCandy(number_of_candies_stolen)
	if myerr != nil {
		// handle error
		fmt.Println("FATAL: ", myerr.Error())
	} else {
		// DEBUG:
		// fmt.Println("REFILL? ", should_refill)
		_, myerr := party_today.ExecuteRobotInspection(should_refill, current_candy_bowl)
		if myerr != nil {
			// handle error
			fmt.Println("FATAL: ", myerr.Error())
		} else {
			// All OK!!?
			// fmt.Println("Refill: ", how_many_to_refill)
		}
	}
}

func (c *candybowl) ExecuteGuestStealingCandy(number_of_candies_stolen int) (should_refill bool, err error) {
	if number_of_candies_stolen > c.number_of_candies {
		return false, errors.New("Cannot steal more than available!!")
	}
	// Reduce the candies in the bowl
	c.number_of_candies = c.number_of_candies - number_of_candies_stolen
	// Check the limit; whic is a hard code
	if c.number_of_candies < 5 {
		return true, nil
	}
	return false, nil
}

func (p *party) ExecuteRobotInspection(should_refill bool, c *candybowl) (how_many_refilled int, err error) {
	// Update party phase info
	p.current_party_phase++
	if p.current_party_phase < p.end_party_phase {
		// DO something ONLY if number_of_candies is less than 5!
		if (c.number_of_candies < 5) {
			how_many_refilled = c.bowl_size - c.number_of_candies
			// Update state; tightly coupled?
			c.number_of_candies = c.bowl_size
			c.number_of_refills += how_many_refilled
			return how_many_refilled, nil
		}
	}
	return 0, nil
}
