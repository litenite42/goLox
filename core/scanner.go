package core

import (
	"strconv"
	// "fmt"
	"golox/loxerr"
	"golox/tern"
)

type Scanner struct {
	src    string
	tokens []*Token

	start, current, line int
}

func NewScanner(src string) *Scanner {
	return &Scanner{src, nil, 0, 0, 1}
}

func (s *Scanner) ScanTokens() []*Token {
	for !s.atEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.addToken(EOF, "", nil, s.line)
	return s.tokens
}

func (s *Scanner) addToken(ttype TokenType, lexeme string, literal interface{}, line int) {
	s.tokens = append(s.tokens, newToken(ttype, lexeme, literal, line))
}

func (s *Scanner) addLiteral(ttype TokenType, literal interface{}) {
	text := s.src[s.start:s.current]
	s.addToken(ttype, text, literal, s.line)
}

func (s *Scanner) addSingle(ttype TokenType) {
	s.addLiteral(ttype, nil)
}

func (s *Scanner) addSingleT(m byte, a TokenType, b TokenType) {
	s.addSingle(tern.Q(s.match(m), a, b))
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case '{':
		s.addSingle(LPAREN)
	case '}':
		s.addSingle(RPAREN)
	case '(':
		s.addSingle(LBRACE)
	case ')':
		s.addSingle(RBRACE)
	case ',':
		s.addSingle(COMMA)
	case '.':
		s.addSingle(DOT)
	case '-':
		s.addSingle(MINUS)
	case '+':
		s.addSingle(PLUS)
	case ';':
		s.addSingle(SEMICOLON)
	case '*':
		s.addSingle(STAR)
	case '!':
		s.addSingleT('=', BANG_EQ, BANG)
	case '=':
		s.addSingleT('=', EQ_EQ, EQ)
	case '<':
		s.addSingleT('=', LS_EQ, LS)
	case '>':
		s.addSingleT('=', GTR_EQ, GTR)
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.atEnd() {
				s.advance()
			}
		} else {
			s.addSingle(SLASH)
		}
	case ' ':
	case '\r':
	case '\t': // ignore whitespace
		break
	case '\n':
		s.line++
	case '\'':
		s.string()
	default:
		if s.isDigit(c) {
			s.number()
		} else {
			loxerr.Error(s.line, "Unexpected character!")
		}
	}
}

func (s *Scanner) string() {
	for s.peek() != '"' && !s.atEnd() {
		if s.peek() == '\n' {
			s.line++
		}

		s.advance()
	}

	if s.atEnd() {
		loxerr.Error(s.line, "Unterminated String")
		return
	}

	s.advance()

	val := s.src[s.start+1 : s.current-1]
	s.addLiteral(STRING, val)
}

func (s Scanner) isDigit(b byte) bool {
	return b >= '0' && b <= '9'	
}

func (s *Scanner) number() {
	for s.isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		s.advance()

		for s.isDigit(s.peek()) {
			s.advance()
		}
	}

	valF, err := strconv.ParseFloat(s.src[s.start:s.current], 64)

	if err != nil {
		loxerr.Error(s.line, "Unable to parse expected float")
		return
	}

	s.addLiteral(NUMBER, valF) 
}

func (s *Scanner) advance() byte {
	result := s.src[s.current]
	s.current++

	return result
}

func (s *Scanner) match(expected byte) bool {
	if s.atEnd() {
		return false
	}

	if s.src[s.current] != expected {
		return false
	}

	s.current++

	return true
}

func (s *Scanner) peek() byte {
	if s.atEnd() {
		return 0 
	}

	return s.src[s.current]
}

func (s *Scanner) peekNext() byte {
	if s.current + 1 >= len(s.src) { return 0 }
	return s.src[s.current + 1]
}

func (s Scanner) atEnd() bool {
	return s.current >= len(s.src)
}
