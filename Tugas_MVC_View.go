package View

import (
	"Project2024kelasV/Node"
	"fmt"
)

func DisplayMembers(members *Node.TableMember) {
	if members == nil {
		fmt.Println("TIDAK ADA DATA MEMBER.")
		return
	}

	temp := members
	for temp != nil {
		fmt.Printf("ID: %d, Nama: %s, Alamat: %s, Point: %.2f\n", temp.Data.Id, temp.Data.Nama, temp.Data.Alamat, temp.Data.Point)
		temp = temp.Next
	}
}

func DisplaySuccess(message string) {
	fmt.Println(message)
}

func DisplayError(message string) {
	fmt.Println(message)
}
