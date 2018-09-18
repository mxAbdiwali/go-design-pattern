package singleton

//Singleton add one
type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

//GetInstance return a counter
func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
