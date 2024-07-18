package main

import (
	"strconv"
	"syscall/js"
)

type player struct {
	Name string
	Mark string
}

var board [3][3]string
var currentPlayer player
var xUser, oUser player

func main() {
	c := make(chan int, 0)
	js.Global().Set("makeMove", js.FuncOf(makeMove))
	js.Global().Set("restart", js.FuncOf(restart))
	resetBoard()
	getUsers()
	<-c
}

func restart(this js.Value, p []js.Value) interface{} {
	resetBoard()
	return nil
}

func makeMove(this js.Value, p []js.Value) interface{} {
	row := p[0].Int()
	col := p[1].Int()
	if board[row][col] == "" {
		board[row][col] = currentPlayer.Mark
		if checkWin(currentPlayer) {
			setMessage(currentPlayer.Name + " wins!")
		} else if checkDraw() {
			setMessage("It's a draw!")
		} else {
			currentPlayer = switchPlayer(currentPlayer)
		}
		updateBoard()
	}
	return nil
}

func resetBoard() {
	for i := range board {
		for j := range board[i] {
			board[i][j] = ""
		}
	}
	updateBoard()

	activePlayer("playerX")
	deActivePlayer("playerO")

	hideMessage()
}

func updateBoard() {
	for i := range board {
		for j := range board[i] {
			updateCell(i, j, board[i][j])
		}
	}
}

func switchPlayer(player player) player {
	deActivePlayer("player" + player.Mark)

	if player.Mark == "X" {
		player = oUser
	} else {
		player = xUser
	}

	activePlayer("player" + player.Mark)
	return player
}

func checkWin(player player) bool {
	mark := player.Mark
	for i := 0; i < 3; i++ {
		if board[i][0] == mark && board[i][1] == mark && board[i][2] == mark {
			return true
		}
		if board[0][i] == mark && board[1][i] == mark && board[2][i] == mark {
			return true
		}
	}
	if board[0][0] == mark && board[1][1] == mark && board[2][2] == mark {
		return true
	}
	if board[0][2] == mark && board[1][1] == mark && board[2][0] == mark {
		return true
	}
	return false
}

func checkDraw() bool {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == "" {
				return false
			}
		}
	}
	return true
}

func getUsers() {
	//xUserName := js.Global().Call("prompt", "X users' name:").String()
	//oUserName := js.Global().Call("prompt", "O users' name:").String()

	xUserName := "Player X"
	oUserName := "Player O"

	xUser = player{
		Name: xUserName,
		Mark: "X",
	}
	oUser = player{
		Name: oUserName,
		Mark: "O",
	}

	currentPlayer = xUser

	updatePlayerName("playerX", xUserName)
	updatePlayerName("playerO", oUserName)
}

func updateCell(row, col int, value string) {
	js.Global().Get("document").Call("getElementById", "cell-"+strconv.Itoa(row)+"-"+strconv.Itoa(col)).Set("innerText", value)
}

func updatePlayerName(player, name string) {
	js.Global().Get("document").Call("getElementById", player).Set("innerText", name)
}

func activePlayer(player string) {
	element := js.Global().Get("document").Call("getElementById", player)
	element.Get("classList").Call("add", "active")
}

func deActivePlayer(player string) {
	element := js.Global().Get("document").Call("getElementById", player)
	element.Get("classList").Call("remove", "active")
}

func hideMessage() {
	containerElem := js.Global().Get("document").Call("getElementById", "message-container")
	containerElem.Get("classList").Call("add", "hidden")
	containerElem.Get("classList").Call("remove", "flex")
}

func setMessage(message string) {
	messageElem := js.Global().Get("document").Call("getElementById", "message")
	messageElem.Set("innerText", message)

	containerElem := js.Global().Get("document").Call("getElementById", "message-container")
	containerElem.Get("classList").Call("remove", "hidden")
	containerElem.Get("classList").Call("add", "flex")
}
