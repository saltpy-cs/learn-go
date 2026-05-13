package main

func main() {
	js := person{
		firstName: "Jeff",
		lastName:  "Smith",
		age:       30,
		height:    180,
		contactInfo: contactInfo{
			email:    "jeff.smith@example.com",
			postCode: "EX1 2AB",
		},
	}
	js.updateName("Jeffery")
	js.print()
}
