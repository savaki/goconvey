/*

Reference: http://butunclebob.com/ArticleS.UncleBob.TheBowlingGameKata

See the very first link (which happens to be the very first word of
the first paragraph) on the page for a tutorial.

*/

package examples

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestScoring(t *testing.T) {
	Convey("Subject: Bowling Game Scoring", t, func(c *Context, so Assert) {
		var game *Game // Whatever you do, don't do this: game := NewGame()
		// Otherwise nested closures won't reference the correct instance

		Convey("Given a fresh score card", c, func() {
			game = NewGame()

			Convey("When all gutter balls are thrown", c, func() {
				game.rollMany(20, 0)

				Convey("The score should be zero", c, func() {
					so(game.Score(), ShouldEqual, 0)
				})
			})

			Convey("When all throws knock down only one pin", c, func() {
				game.rollMany(20, 1)

				Convey("The score should be 20", c, func() {
					so(game.Score(), ShouldEqual, 20)
				})
			})

			Convey("When a spare is thrown", c, func() {
				game.rollSpare()
				game.Roll(3)
				game.rollMany(17, 0)

				Convey("The score should include a spare bonus.", c, func() {
					so(game.Score(), ShouldEqual, 16)
				})
			})

			Convey("When a strike is thrown", c, func() {
				game.rollStrike()
				game.Roll(3)
				game.Roll(4)
				game.rollMany(16, 0)

				Convey("The score should include a strike bonus.", c, func() {
					so(game.Score(), ShouldEqual, 24)
				})
			})

			Convey("When all strikes are thrown", c, func() {
				game.rollMany(21, 10)

				Convey("The score should be 300.", c, func() {
					so(game.Score(), ShouldEqual, 300)
				})
			})
		})
	})
}

func (self *Game) rollMany(times, pins int) {
	for x := 0; x < times; x++ {
		self.Roll(pins)
	}
}
func (self *Game) rollSpare() {
	self.Roll(5)
	self.Roll(5)
}
func (self *Game) rollStrike() {
	self.Roll(10)
}
