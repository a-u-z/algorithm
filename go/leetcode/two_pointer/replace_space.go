package main

import "log"

func Re(s string) string {
	log.Printf("here is s:%+v", s)
	spaceCounter := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' { // 因為 range s 會變成 byte
			spaceCounter++
		}
	}

	newS := make([]byte, len(s)+spaceCounter*2) // ' '=>%20 所以多出兩個 spaceCounter
	newSPointer := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			newS[newSPointer] = '%'
			newS[newSPointer+1] = '2'
			newS[newSPointer+2] = '0'
			newSPointer += 3
		} else {
			newS[newSPointer] = s[i]
			newSPointer++
		}
	}
	log.Printf("here is string(newS):%+v", string(newS))
	return string(newS)
}
