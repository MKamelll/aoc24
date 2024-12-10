package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type TokenType string

const (
	tt_mul         TokenType = "tt_mul"
	tt_left_paren  TokenType = "tt_left_paren"
	tt_right_paren TokenType = "tt_right_paren"
	tt_int         TokenType = "tt_int"
	tt_comma       TokenType = "tt_comma"
	tt_do          TokenType = "tt_do"
	tt_do_not      TokenType = "tt_don_not"
	tt_unsuported  TokenType = "tt_unsuported"
)

type Token struct {
	kind   TokenType
	lexeme string
}

func NewToken(kind TokenType, lexeme string) Token {
	return Token{kind: kind, lexeme: lexeme}
}

func (token *Token) String() string {
	return "Token(type: " + string(token.kind) + ", lexeme: '" + token.lexeme + "')"
}

func tokenize(line string) []Token {
	var result []Token
	var i int
	for i < len(line) {
		switch line[i] {
		case '(':
			result = append(result, NewToken(tt_left_paren, "("))
			i++
			break
		case ')':
			result = append(result, NewToken(tt_right_paren, ")"))
			i++
			break
		case ',':
			result = append(result, NewToken(tt_comma, ","))
			i++
			break
		case 'm':
			{
				if i < len(line) && line[i+1] == 'u' {
					i++
					if i < len(line) && line[i+1] == 'l' {
						i++
						result = append(result, NewToken(tt_mul, "mul"))
						i++
						break
					}
				}

				result = append(result, NewToken(tt_unsuported, string(line[i])))
				i++
				break
			}

		case 'd':
			{
				var old_i = i
				id := ""
				for i < len(line) && unicode.IsLetter(rune(line[i])) || line[i] == '\'' {
					id += string(line[i])
					i++
				}

				if id == "do" {
					result = append(result, NewToken(tt_do, "do"))
					break
				}

				if id == "don't" {
					result = append(result, NewToken(tt_do_not, "don't"))
					break
				}

				i = old_i
				result = append(result, NewToken(tt_unsuported, string(line[i])))
				i++
				break

			}
		default:
			{
				if unicode.IsDigit(rune(line[i])) {
					number := ""
					for i < len(line) && unicode.IsDigit(rune(line[i])) {
						number += string(line[i])
						i++
					}
					result = append(result, NewToken(tt_int, number))
					break
				}

				result = append(result, NewToken(tt_unsuported, string(line[i])))
				i++
			}
		}
	}

	return result
}

type mul_expr struct {
	first_digit  int
	second_digit int
}

func NewMulExpr(first_digit, second_digit int) mul_expr {
	return mul_expr{first_digit: first_digit, second_digit: second_digit}
}

func (expr *mul_expr) String() string {
	return "mul(" + strconv.Itoa(expr.first_digit) + ", " + strconv.Itoa(expr.second_digit) + ")"
}

func parse_mul_exprs(line_tokens []Token) []mul_expr {
	var valid_mul_exprs []mul_expr
	var i int

	for i < len(line_tokens) {
		if line_tokens[i].kind == tt_mul {
			i++
			if i < len(line_tokens) && line_tokens[i].kind == tt_left_paren {
				i++
				if i < len(line_tokens) && line_tokens[i].kind == tt_int {
					first_digit, _ := strconv.Atoi(line_tokens[i].lexeme)
					i++
					if i < len(line_tokens) && line_tokens[i].kind == tt_comma {
						i++
					}
					if i < len(line_tokens) && line_tokens[i].kind == tt_int {
						second_digit, _ := strconv.Atoi(line_tokens[i].lexeme)
						i++
						if i < len(line_tokens) && line_tokens[i].kind == tt_right_paren {
							valid_mul_exprs = append(valid_mul_exprs, NewMulExpr(first_digit, second_digit))
						}
					}
				}
			}
		}
		i++
	}

	return valid_mul_exprs
}

var do_we_collect = true

func parse_mul_exprs2(line_tokens []Token) []mul_expr {
	var valid_mul_exprs []mul_expr
	var i int

	for i < len(line_tokens) {
		if line_tokens[i].kind == tt_do {
			i++
			if line_tokens[i].kind == tt_left_paren {
				i++
				if line_tokens[i].kind == tt_right_paren {
					i++
					do_we_collect = true
				}
			}
		}

		if line_tokens[i].kind == tt_do_not {
			i++
			if line_tokens[i].kind == tt_left_paren {
				i++
				if line_tokens[i].kind == tt_right_paren {
					i++
					do_we_collect = false
				}
			}
		}

		if line_tokens[i].kind == tt_mul {
			i++
			if i < len(line_tokens) && line_tokens[i].kind == tt_left_paren {
				i++
				if i < len(line_tokens) && line_tokens[i].kind == tt_int {
					first_digit, _ := strconv.Atoi(line_tokens[i].lexeme)
					i++
					if i < len(line_tokens) && line_tokens[i].kind == tt_comma {
						i++
					}
					if i < len(line_tokens) && line_tokens[i].kind == tt_int {
						second_digit, _ := strconv.Atoi(line_tokens[i].lexeme)
						i++
						if i < len(line_tokens) && line_tokens[i].kind == tt_right_paren {
							if do_we_collect {
								valid_mul_exprs = append(valid_mul_exprs, NewMulExpr(first_digit, second_digit))
							}
						}
					}
				}
			}
		}

		i++
	}

	return valid_mul_exprs
}

func part1(content string) {
	lines := strings.Split(content, "\n")
	var result int
	for _, line := range lines {
		line_tokens := tokenize(line)
		exprs := parse_mul_exprs(line_tokens)

		for _, expr := range exprs {
			result += expr.first_digit * expr.second_digit
		}

	}

	fmt.Println(result)

}

func part2(content string) {
	lines := strings.Split(content, "\n")
	var result int
	for _, line := range lines {
		line_tokens := tokenize(line)
		exprs := parse_mul_exprs2(line_tokens)

		for _, expr := range exprs {
			result += expr.first_digit * expr.second_digit
		}

	}

	fmt.Println(result)

}

func main() {

	//file, _ := os.ReadFile("test.txt")
	file, _ := os.ReadFile("input.txt")
	content := string(file)
	part2(content)
}
