package board_game

import (
	"testing"
)

func TestThrowingDice(t *testing.T) {
	var result uint64

	result = boardGame(0)
	if result != 1 {
		t.Errorf("expect boardGame(0) to be 1, but got %d", result)
	}

	result = boardGame(1)
	if result != 1 {
		t.Errorf("expect boardGame(1) to be 1, but got %d", result)
	}

	result = boardGame(2)
	if result != 2 {
		t.Errorf("expect boardGame(2) to be 2, but got %d", result)
	}

	result = boardGame(3)
	if result != 4 {
		t.Errorf("expect boardGame(3) to be 4, but got %d", result)
	}

	result = boardGame(4)
	if result != 8 {
		t.Errorf("expect boardGame(4) to be 8, but got %d", result)
	}

	result = boardGame(5)
	if result != 16 {
		t.Errorf("expect boardGame(5) to be 16, but got %d", result)
	}

	result = boardGame(6)
	if result != 32 {
		t.Errorf("expect boardGame(6) to be 32, but got %d", result)
	}

	result = boardGame(7)
	if result != 63 {
		t.Errorf("expect boardGame(7) to be 63, but got %d", result)
	}

	result = boardGame(15)
	if result != 15109 {
		t.Errorf("expect boardGame(15) to be 15109, but got %d", result)
	}

	result = boardGame(18)
	if result != 117920 {
		t.Errorf("expect boardGame(18) to be 117920, but got %d", result)
	}

	result = boardGame(19)
	if result != 233904 {
		t.Errorf("expect boardGame(19) to be 233904, but got %d", result)
	}

	result = boardGame(20)
	if result != 463968 {
		t.Errorf("expect boardGame(20) to be 463968, but got %d", result)
	}

	result = boardGame(50)
	if result != 389043663364337 {
		t.Errorf("expect boardGame(50) to be 389043663364337, but got %d", result)
	}
}
