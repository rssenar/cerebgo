package cerebgo

// Parser holds the map to the parsed headers
type Parser struct {
	headerMap map[int]int
}

func (p *Parser) new() *Parser { return nil }
