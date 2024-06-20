package Controller

import (
	"Project2024kelasV/Model"
	"Project2024kelasV/Node"
)

func AddMember(id int, nama string, alamat string, point float32) {
	Model.InsertMember(id, nama, alamat, point)
}

func GetAllMembers() *Node.TableMember {
	return Model.ReadAllMember()
}

func RemoveMember(id int) bool {
	return Model.DeleteMember(id)
}

func ModifyMember(id int, nama string, alamat string, point float32) bool {
	return Model.UpdateMember(id, nama, alamat, point)
}
