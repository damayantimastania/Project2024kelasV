package main

import (
	"Project2024kelasV/Controller"
	"Project2024kelasV/View"
)

func main() {
	Controller.AddMember(1, "Damayanti", "Surabaya", 78.5)
	Controller.AddMember(2, "Joe Alwyn", "London", 150.75)

	members := Controller.GetAllMembers()
	View.DisplayMembers(members)

	success := Controller.ModifyMember(1, "Damayanti", "Surabaya", 100)
	if success {
		View.DisplaySuccess("Update member berhasil.")
	} else {
		View.DisplayError("Update member gagal.")
	}

	members = Controller.GetAllMembers()
	View.DisplayMembers(members)

	success = Controller.RemoveMember(2)
	if success {
		View.DisplaySuccess("Delete member berhasil")
	} else {
		View.DisplayError("Delete member gagal.")
	}

	members = Controller.GetAllMembers()
	View.DisplayMembers(members)
}
