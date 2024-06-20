package Model

import (
	"Project2024kelasV/Database"
	"Project2024kelasV/Node"
)

func InsertMember(id int, nama string, alamat string, point float32) {
	newDataMember := Node.TableMember{
		Data: Node.Member{Id: id, Nama: nama, Alamat: alamat, Point: point},
		Next: nil,
	}
	if Database.DBmember.Next == nil {
		Database.DBmember.Next = &newDataMember
	} else {
		temp := &Database.DBmember
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &newDataMember
	}
}

func ReadAllMember() *Node.TableMember {
	if Database.DBmember.Next == nil {
		return nil
	}
	return Database.DBmember.Next
}

func DeleteMember(id int) bool {
	if Database.DBmember.Next == nil {
		return false
	}
	temp := &Database.DBmember
	for temp.Next != nil {
		if temp.Next.Data.Id == id {
			temp.Next = temp.Next.Next
			return true
		}
		temp = temp.Next
	}
	return false
}

func UpdateMember(id int, nama string, alamat string, point float32) bool {
	temp := &Database.DBmember
	for temp.Next != nil {
		if temp.Next.Data.Id == id {
			temp.Next.Data.Nama = nama
			temp.Next.Data.Alamat = alamat
			temp.Next.Data.Point = point
			return true
		}
		temp = temp.Next
	}
	return false
}
