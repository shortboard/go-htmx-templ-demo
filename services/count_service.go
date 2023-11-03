package services

type CountService struct {
	count int
}

func NewCountService() *CountService {
	return &CountService{
		count: 0,
	}
}

func (c *CountService) Increment() {
	c.count++
}

func (c *CountService) GetCount() int {
	c.Increment()
	return c.count
}
