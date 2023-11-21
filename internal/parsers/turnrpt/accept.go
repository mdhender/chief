// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package turnrpt

func (r *Report) accept(lex func([]byte) ([]byte, []byte)) string {
	token, rest := lex(r.input)
	if token == nil {
		return ""
	}
	r.input = rest
	return string(token)
}

func (r *Report) acceptDelimiter(ch byte) bool {
	token, rest := lexDelimiter(r.input, ch)
	if token == nil {
		return false
	}
	r.input = rest
	return true
}

func (r *Report) acceptDirection() string {
	token, rest := lexDirection(r.input)
	if token == nil {
		return ""
	}
	r.input = rest
	return string(token)
}

func (r *Report) acceptGoods() string {
	var goods string
	if token, rest := lexGoods(r.input); token != nil {
		goods, r.input = string(token), rest
	}
	return goods
}

func (r *Report) acceptLiteral(literal string) bool {
	token, rest := lexLiteral(r.input, []byte(literal))
	if token == nil {
		return false
	}
	r.input = rest
	return true
}

func (r *Report) peekLiteral(literal string) bool {
	saved := r.input
	if !r.acceptLiteral(literal) {
		return false
	}
	r.input = saved
	return true
}
