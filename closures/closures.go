package closures

func closer() func()int  {
	i :=0
	return func ()int  {
		i++
		return i
	}
}