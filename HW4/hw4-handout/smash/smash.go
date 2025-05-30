package smash

import (
	"io"
	"bufio"
	"sync"
)

type word string

func Smash(r io.Reader, smasher func(word) uint32) map[uint32]uint {
	m := make(map[uint32]uint)

	scanner := bufio.NewScanner(r) //initialize the scanner
	scanner.Split(bufio.ScanWords) // Use bufio.Scanner to read words from the input

	var wg sync.WaitGroup
	var mu sync.Mutex

	// For each word, launch a goroutine 
	for scanner.Scan() {
		w := word(scanner.Text()) // Convert the scanned text to a word type
		wg.Add(1)
		go func(w word) {
			defer wg.Done()
			h := smasher(w)
			mu.Lock()
			m[h]++
			mu.Unlock()
		}(w)
	}

	// Wait for all  goroutines to finish
	wg.Wait()
	return m
}
 