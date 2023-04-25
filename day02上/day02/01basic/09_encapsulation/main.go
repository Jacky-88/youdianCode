package encapsulation

type Student struct {
	Name string
	score int
}

//封装一个new方法
func NewStudent(name string, score int) *Student {
	stu := &Student{
		Name:  name,
		score: score,
	}
	return stu
}

//获取分数
func (s *Student) GetScore() int {
	return s.score
}

//设置分数
func (s *Student) SetScore(score int) {
	s.score = score
}