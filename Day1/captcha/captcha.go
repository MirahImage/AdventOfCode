package captcha

//Captcha contains digits
type Captcha struct {
	Digits []int
}

func (c *Captcha) MatchPrev(index int) bool {
	if index == 0 {
		return c.Digits[index] == c.Digits[len(c.Digits)-1]
	}
	return c.Digits[index] == c.Digits[index-1]
}

func (c *Captcha) MatchHalfway(index int) bool {
	return c.Digits[index] == c.Digits[(index+len(c.Digits)/2)%len(c.Digits)]
}

//Sum computes sum of repeated digits
func (c *Captcha) Sum() int {
	sum := 0
	for i, digit := range c.Digits {
		if c.MatchPrev(i) {
			sum += digit
		}
	}
	return sum
}

func (c *Captcha) SumHalfway() int {
	sum := 0
	for i, digit := range c.Digits {
		if c.MatchHalfway(i) {
			sum += digit
		}
	}
	return sum
}
