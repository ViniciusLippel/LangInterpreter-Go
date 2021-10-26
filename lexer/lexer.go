package lexer

import "Interpreter/token"

type Lexer struct {
	input        string
	position     int  //Current position in input (points to current char)
	readPosition int  //Current reading position in input (after current char)
	ch           byte //Current char under examination
}

//Initialize Lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

//Give us the next char and advance our position in input
func (l *Lexer) readChar() {
	//Check if reached the end of input
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		//Sets ch to next char (index in readPosition)
		l.ch = l.input[l.readPosition]
	}
	//Update position and readPosition
	l.position = l.readPosition
	l.readPosition += 1
}

//Look at the current char and return a corresponding token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) { //Identifier or keyword
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) { //Number
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else { //Illegal statement
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	//Advance our pointers into the input
	l.readChar()
	return tok
}

//Create new token
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

//Read identifier and advance lexer positions until non-letter-char
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

//Read a number
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

//Skip/eat/consume whitespaces
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

//Return True if char is valid for identifiers
//Need to be updated so identifiers with numbers are recognized
func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' ||
		ch == '_')
}

//Return true if char is a digit between 0 and 9
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
