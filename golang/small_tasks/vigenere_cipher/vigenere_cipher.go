package vigenerecipher

type VigenèreCipher struct {
	Key                     string       // Key of arbitrary length. Name is defined by Codewars kata.
	Alpha                   string       // Alphabet the encoder / decoder works with. Name is defined by Codewars kata.
	keyRuneArray            []rune       // Rune array representation to easily index non-utf8 characters
	alphabetRuneArray       []rune       // Rune array representation to easily index non-utf8 characters
	runeAlphabetPositionMap map[rune]int // For internals to get rune index in the alphabet quickly

}

func (c *VigenèreCipher) initialize() {
	if c.keyRuneArray == nil {
		c.keyRuneArray = []rune(c.Key)
	}
	if c.alphabetRuneArray == nil {
		c.alphabetRuneArray = []rune(c.Alpha)
	}
	if c.runeAlphabetPositionMap == nil {
		c.runeAlphabetPositionMap = make(map[rune]int)
		for index, alphabetRune := range c.alphabetRuneArray {
			c.runeAlphabetPositionMap[alphabetRune] = index
		}
	}
}

func (c *VigenèreCipher) Encode(str string) string {
	c.initialize()
	inputRuneArray := []rune(str)
	result := make([]rune, len(inputRuneArray))

	for index, inputRune := range inputRuneArray {
		if position, found := c.runeAlphabetPositionMap[inputRune]; found {
			currentKeyRune := c.keyRuneArray[index%len(c.keyRuneArray)]
			encodedRune := c.alphabetRuneArray[(position+c.runeAlphabetPositionMap[currentKeyRune])%len(c.alphabetRuneArray)]
			result[index] = encodedRune
		} else {
			result[index] = inputRune
		}
	}
	return string(result)
}

func (c *VigenèreCipher) Decode(str string) string {
	c.initialize()
	inputRuneArray := []rune(str)
	result := make([]rune, len(inputRuneArray))

	for index, inputRune := range inputRuneArray {
		if position, found := c.runeAlphabetPositionMap[inputRune]; found {
			currentKeyRune := c.keyRuneArray[index%len(c.keyRuneArray)]
			encodedRune := c.alphabetRuneArray[(position-c.runeAlphabetPositionMap[currentKeyRune]+len(c.alphabetRuneArray))%len(c.alphabetRuneArray)]
			result[index] = encodedRune
		} else {
			result[index] = inputRune
		}
	}
	return string(result)
}
