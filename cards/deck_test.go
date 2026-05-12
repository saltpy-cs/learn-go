package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	_ = w.Close()
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	os.Stdout = old
	return buf.String()
}

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)
	deck := newDeck()
	deck.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)
	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(loadedDeck))
	}
	os.Remove(filename)
}

func TestShuffle(t *testing.T) {
	deck := newDeck()
	shuffledDeck := newDeck()
	shuffledDeck.shuffle()

	if len(deck) != len(shuffledDeck) {
		t.Errorf("Expected shuffled deck to have the same length as original deck, but got %v", len(shuffledDeck))
	}

	samePositionCount := 0
	for i := range deck {
		if deck[i] == shuffledDeck[i] {
			samePositionCount++
		}
	}

	if samePositionCount > 5 {
		t.Errorf("Expected shuffled deck to have different card positions than original deck, but found %v cards in the same position", samePositionCount)
	}
}

func TestDeal(t *testing.T) {
	deck := newDeck()
	handSize := 5
	hand, remainingDeck := deal(deck, handSize)

	if len(hand) != handSize {
		t.Errorf("Expected hand size of %v, but got %v", handSize, len(hand))
	}

	if len(remainingDeck) != len(deck)-handSize {
		t.Errorf("Expected remaining deck size of %v, but got %v", len(deck)-handSize, len(remainingDeck))
	}
}

func TestToString(t *testing.T) {
	deck := newDeck()
	deckString := deck.toString()
	expectedString := "Ace of Spades,Two of Spades,Three of Spades,Four of Spades,Five of Spades,Six of Spades,Seven of Spades,Eight of Spades,Nine of Spades,Ten of Spades,Jack of Spades,Queen of Spades,King of Spades,Ace of Diamonds,Two of Diamonds,Three of Diamonds,Four of Diamonds,Five of Diamonds,Six of Diamonds,Seven of Diamonds,Eight of Diamonds,Nine of Diamonds,Ten of Diamonds,Jack of Diamonds,Queen of Diamonds,King of Diamonds,Ace of Hearts,Two of Hearts,Three of Hearts,Four of Hearts,Five of Hearts,Six of Hearts,Seven of Hearts,Eight of Hearts,Nine of Hearts,Ten of Hearts,Jack of Hearts,Queen of Hearts,King of Hearts,Ace of Clubs,Two of Clubs,Three of Clubs,Four of Clubs,Five of Clubs,Six of Clubs,Seven of Clubs,Eight of Clubs,Nine of Clubs,Ten of Clubs,Jack of Clubs,Queen of Clubs,King of Clubs"
	if deckString != expectedString {
		t.Errorf("Expected deck string to be %v, but got %v", expectedString, deckString)
	}
}

func TestPrint(t *testing.T) {
	deck := newDeck()
	output := captureStdout(func() {
		deck.print()
	})

	expectedOutput := "0 Ace of Spades\n1 Two of Spades\n2 Three of Spades\n3 Four of Spades\n4 Five of Spades\n5 Six of Spades\n6 Seven of Spades\n7 Eight of Spades\n8 Nine of Spades\n9 Ten of Spades\n10 Jack of Spades\n11 Queen of Spades\n12 King of Spades\n13 Ace of Diamonds\n14 Two of Diamonds\n15 Three of Diamonds\n16 Four of Diamonds\n17 Five of Diamonds\n18 Six of Diamonds\n19 Seven of Diamonds\n20 Eight of Diamonds\n21 Nine of Diamonds\n22 Ten of Diamonds\n23 Jack of Diamonds\n24 Queen of Diamonds\n25 King of Diamonds\n26 Ace of Hearts\n27 Two of Hearts\n28 Three of Hearts\n29 Four of Hearts\n30 Five of Hearts\n31 Six of Hearts\n32 Seven of Hearts\n33 Eight of Hearts\n34 Nine of Hearts\n35 Ten of Hearts\n36 Jack of Hearts\n37 Queen of Hearts\n38 King of Hearts\n39 Ace of Clubs\n40 Two of Clubs\n41 Three of Clubs\n42 Four of Clubs\n43 Five of Clubs\n44 Six of Clubs\n45 Seven of Clubs\n46 Eight of Clubs\n47 Nine of Clubs\n48 Ten of Clubs\n49 Jack of Clubs\n50 Queen of Clubs\n51 King of Clubs\n"

	if output != expectedOutput {
		t.Errorf("Expected print output to be %v, but got %v", expectedOutput, output)
	}
}

func TestNewDeckFromFileWithNonExistentFile(t *testing.T) {
	filename := "non_existent_file"
	_ = os.Remove(filename)

	output := captureStdout(func() {
		newDeckFromFile(filename)
	})

	expectedOutput := "Error: open non_existent_file: no such file or directory\n"

	if output != expectedOutput {
		t.Errorf("Expected output to be %v, but got %v", expectedOutput, output)
	}
}
