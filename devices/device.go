package devices


type Cpu struct{
	Prod string
	Core int
}

type Pc struct{
	Cpu
}

func (pc Pc) On(){
	println("[On] Включение Pc")
}

func (pc Pc) Off(){
	println("[Off] Выключение Pc" )

}

func (pc Pc) initMethod() {
	println("[Init] ...")
}
