package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var notepad []string
var counter = 0

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	setupNotepad(scanner)
	for {
		command, note := getInput(scanner)
		switch command {
		case "create":
			createNote(note)
		case "update":
			updateNote(note)
		case "delete":
			deleteNote(note)
		case "list":
			listNotepad()
		case "clear":
			clearNotepad()
		case "exit":
			exitNotepad()
		default:
			fmt.Println("[Error] Unknown command")
		}
	}
}

func setupNotepad(scanner *bufio.Scanner) {
	fmt.Print("Enter the maximum number of notes: ")
	scanner.Scan()
	inputText := scanner.Text()
	notepadSize, _ := strconv.Atoi(inputText)
	notepad = make([]string, notepadSize)
}

func getInput(scanner *bufio.Scanner) (command, note string) {
	fmt.Print("Enter a command and data: ")
	scanner.Scan()
	inputText := scanner.Text()
	command = strings.Split(inputText, " ")[0]
	note = strings.TrimSpace(strings.TrimPrefix(inputText, command))
	return
}

func createNote(note string) {
	if counter >= cap(notepad) {
		fmt.Println("[Error] Notepad is full")
		return
	}
	if isEmpty(note) {
		fmt.Println("[Error] Missing note argument")
		return
	}
	notepad[counter] = note
	counter++
	fmt.Println("[OK] The note was successfully created")
}

func updateNote(line string) {
	position, note := getArguments(line)
	if position == 0 {
		return
	}
	if isEmpty(note) {
		fmt.Print("[Error] Missing note argument\n")
		return
	}
	if position < 1 || position > cap(notepad) {
		fmt.Printf("[Error] Position %d is out of the boundary [1, %d]\n", position, cap(notepad))
		return
	}
	if notepad[position-1] == "" {
		fmt.Println("[Error] There is nothing to update")
		return
	}
	notepad[position-1] = note
	fmt.Printf("[OK] The note at position %d was successfully updated\n", position)
}

func getArguments(line string) (position int, note string) {
	if isEmpty(line) {
		fmt.Printf("[Error] Missing position argument\n")
		return
	}
	pos, err := strconv.ParseInt(strings.Split(line, " ")[0], 10, 0)
	if err != nil {
		fmt.Printf("[Error] Invalid position: %s\n", strings.Split(line, " ")[0])
		return
	}
	position = int(pos)
	note = strings.TrimPrefix(strings.TrimPrefix(line, strconv.Itoa(position)), " ")
	return
}

func deleteNote(note string) {
	position, _ := getArguments(note)
	if position == 0 {
		return
	}
	if position < 1 || position > cap(notepad) {
		fmt.Printf("[Error] Position %d is out of the boundary [1, %d]\n", position, cap(notepad))
		return
	}
	if notepad[position-1] == "" {
		fmt.Println("[Error] There is nothing to delete")
		return
	}
	notepad = append(notepad[:position-1], notepad[position:]...)
	notepad = append(notepad, "")
	counter--
	fmt.Printf("[OK] The note at position %d was successfully deleted\n", position)
}

func listNotepad() {
	if counter == 0 {
		fmt.Println("[Info] Notepad is empty")
		return
	}
	for index, note := range notepad {
		if !isEmpty(note) {
			fmt.Printf("[Info] %v: %v\n", index+1, note)
		}
	}
}

func clearNotepad() {
	notepad = make([]string, cap(notepad))
	counter = 0
	fmt.Println("[OK] All notes were successfully deleted")
}

func exitNotepad() {
	fmt.Println("[Info] Bye!")
	os.Exit(0)
}

func isEmpty(text string) bool {
	if text == "" {
		return true
	}
	return false
}
