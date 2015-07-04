package lib

const infoNumber = 1
const infoColor = 2

const movePlay = 1
const moveDiscard = 2
const moveHint = 3

var colors = [...]string{"red", "green", "blue", "yellow", "white"}
var numbers = [...]int{0, 3, 2, 2, 2, 1}  // this represents the COUNTS of each number (0 added for simplicity)
const maxHints = 8
const startingHints	= 8
const startingBombs	= 3
var cardsInHand = [...]int{0, 0, 5, 5, 4, 4} // this represents the COUNTS for each # of players