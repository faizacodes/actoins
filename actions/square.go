package actions

type Side struct{
	SideLen int
}


type Actions interface{
	Area() int
	Perimeter() int

}

func(s Side) Area() int{
	return s.SideLen * s.SideLen
}

func(s Side) Perimeter() int{
	return  2*(s.SideLen + s.SideLen)
}


