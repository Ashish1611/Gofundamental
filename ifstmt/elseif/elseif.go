package elseif


func Elseifex(no int) string{

	//num := 9

	if no <= 50 {

		return "No is Less than or equal to 50 "
	} else if no >= 51 && no <= 100 {
		return "No is between 51 and 100"
	} else {
		return "No is greater than 100"
	}

}
