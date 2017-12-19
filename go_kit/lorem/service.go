package lorem

import (
	golorem "github.com/drhodes/golorem"
)

type Service interface {
	word() string
	Sentence() string
	Paragraph() string
}

type loremService struct {


}

func (loremService) Word(min,max int) (s string) {
	return golorem.Word(min,max)


}
func (loremService) Sentence(min,max int) (s string) {
	return golorem.Sentence(min,max)
}

func (loremService) Paragraph(min,max int) (s string) {
	return golorem.Paragraph(min,max)
}

