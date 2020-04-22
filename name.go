package cha


func FirstName() string {
	return PickString(seed.FirstName)
}
func LastName() string {
	return PickString(seed.LastName)
}
func Name() string {
	return FirstName() + " " + LastName()
}
func FullName() string {
	return FirstName() + " " + FirstName() + " " + LastName()
}

func CFirstName() string {
	return PickString(seed.ChineseFirstName)
}
func CLastName() string {
	return PickString(seed.ChineseLastName)
}
func CName() string {
	return CFirstName() + CLastName()
}

