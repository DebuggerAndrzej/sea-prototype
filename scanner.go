package main

func getTokensFromText(text string) []Token {
	var tokens []Token
	for _, character := range text {
		tokens = append(tokens, getToken(character))
	}

	return tokens
}

func getToken(character rune) Token {
	switch character {
	case '(':
		return Token{LEFT_PAREN, "(", 1}
	case ')':
		return Token{RIGHT_PAREN, ")", 1}
	case '{':
		return Token{LEFT_BRACE, "{", 1}
	case '}':
		return Token{RIGHT_BRACE, "}", 1}
	case ',':
		return Token{COMMA, ",", 1}
	case '.':
		return Token{DOT, ".", 1}
	case '-':
		return Token{MINUS, "-", 1}
	case '+':
		return Token{PLUS, "+", 1}
	case ';':
		return Token{SEMICOLON, ";", 1}
	case '*':
		return Token{STAR, "*", 1}
	default:
		reportError(1, "Unexpected character")
	}
	return Token{}
}
