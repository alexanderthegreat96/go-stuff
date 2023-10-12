package variables

var username string = "Alexander"
var password string = "thisismypassword123"

const weight int = 85
const height float64 = 6.3

func protected_method() string {
	return "Protected Method"
}

func PublicFunction() string {
	return "Calling a protected method here: " + protected_method()
}

// multiple returns
// each with it's own type

func ReturnVariables() (string, string, int, float64) {
	return username, password, weight, height
}
