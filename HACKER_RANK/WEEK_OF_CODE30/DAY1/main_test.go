package main

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestExecuteGuestStealingCandy(t *testing.T) {
	Convey("Testing Candy Stealing", t, func() {
		current_candy_bowl := candybowl{8, 8, 0}

		Convey("Sanity check; cannot steal more than you have!", func() {
			_, myerr := current_candy_bowl.ExecuteGuestStealingCandy(30)
			So(myerr.Error(), ShouldEqual, "Cannot steal more than available!!")
		})

		Convey("Steal 2 candies, check correct reduction", func() {
			// Steal everything
			current_candy_bowl.ExecuteGuestStealingCandy(2)
			// Steal just enough
			// Steal more; but now the time is the last minute
			// DEBUG:
			// fmt.Println(current_candy_bowl)
			So(current_candy_bowl.number_of_candies, ShouldEqual, 6)
		})

		Convey("Steal until number of candies equals 5 (edge case)", func() {
			should_refill, _ := current_candy_bowl.ExecuteGuestStealingCandy(3)
			So(should_refill, ShouldEqual, false)
		})

		Convey("Steal until number of candies less than 5", func() {
			should_refill, _ := current_candy_bowl.ExecuteGuestStealingCandy(4)
			So(should_refill, ShouldEqual, true)
		})
	})
}

func TestExecuteRobotInspection(t *testing.T) {
	Convey("", t, func() {
		party_today := party{4, 0}
		current_candy_bowl := candybowl{8, 8, 0}

		// Steal everything
		Convey("Start of party phase, steal ALL candy", func() {
			// Steal everything
			current_candy_bowl.number_of_candies = 0
			should_refill_bowl := true

			how_many_refilled, _ := party_today.ExecuteRobotInspection(should_refill_bowl, &current_candy_bowl)
			So(how_many_refilled, ShouldEqual, 8)

		})
		// Steal everything
		Convey("Start of party phase, steal ALL candy, check structure states", func() {
			current_candy_bowl.number_of_candies = 0
			should_refill_bowl := true

			party_today.ExecuteRobotInspection(should_refill_bowl, &current_candy_bowl)
			So(current_candy_bowl, ShouldResemble, candybowl{8, 8, 8})
			So(party_today, ShouldResemble, party{4, 1})
		})
		// Steal just enough
		Convey("Mid of party phase, steal enough candy", func() {
			current_candy_bowl.number_of_candies = 4
			party_today.current_party_phase = 2
			should_refill_bowl := true

			how_many_refilled, _ := party_today.ExecuteRobotInspection(should_refill_bowl, &current_candy_bowl)
			So(how_many_refilled, ShouldEqual, 4)
		})
		// Steal too little; right on the border
		Convey("Mid of party phase, steal too little candy (edge case)", func() {
			current_candy_bowl.number_of_candies = 5
			party_today.current_party_phase = 2
			should_refill_bowl := false
			// Too little in the mid prt of the party; shoudl not adbance ..
			how_many_refilled, _ := party_today.ExecuteRobotInspection(should_refill_bowl, &current_candy_bowl)
			So(how_many_refilled, ShouldEqual, 0)
		})
		// Steal more; but now the time is the last minute
		Convey("End of party phase, steal too much candy", func() {
			// Steal everything
			current_candy_bowl.number_of_candies = 4
			party_today.current_party_phase = 3
			should_refill_bowl := true
			// Since it is the end of the party; should not advance it ..
			how_many_refilled, _ := party_today.ExecuteRobotInspection(should_refill_bowl, &current_candy_bowl)
			So(how_many_refilled, ShouldEqual, 0)
		})
		// DEBUG:
		// fmt.Println(current_candy_bowl)

	})

}
