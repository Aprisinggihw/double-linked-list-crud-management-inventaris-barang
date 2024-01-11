package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	id, nama string
	stok     int
	next     *Node
	prev     *Node
}

var head, tail *Node

func addNode( id string, nama string, stok int) *Node{
	newNode := &Node{
		id:   id,
		nama: nama,
		stok: stok,
		next: nil,
		prev: tail,
	}

	if head == nil {
		head = newNode
		tail = head
	}else{
		tail.next = newNode
		tail = newNode
	}
	return head
}

func readFileToNode(filePath string) *Node {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("Error open file", err)
	}
	defer file.Close()

	var node *Node
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) == 3 {
			id := fields[0]
			nama := fields[1]
			stok, err := strconv.Atoi(fields[2])
			if err != nil {
				fmt.Println("error parsing stok", err)
				continue
			}
			node = addNode(id, nama, stok)
		}
		if len(fields) > 3 {
			id := fields[0]
			nama := strings.Join(fields[1:len(fields)-1], " ") // Gabungkan kata-kata untuk mendapatkan nama barang
			stok, err := strconv.Atoi(fields[len(fields)-1])
			if err != nil {
				fmt.Println("error parsing stok", err)
				continue
			}
			node = addNode(id, nama, stok)

		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Gagal baca file", err)
		}
	}

	return node
}

func displayDataFromDatabase(node *Node) {
	current := node
	for current != nil {
		fmt.Println("-----------------------------")
		fmt.Printf("Id: %s\n", current.id)
		fmt.Printf("Nama: %s\n", current.nama)
		fmt.Printf("Stok: %d\n", current.stok)
		fmt.Println("-----------------------------")
		current = current.next
	}
}

func addToFile(filePath string) {
	current := head
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("Gak Bisa Buka File", err)
	}
	for current != nil {
		line := strconv.Itoa(current.stok)
		file.WriteString(current.id + " " + current.nama + " " + line + "\n")
		current = current.next
	}
}

func deleteDataBarangBYId(node *Node, id string) {
	for node != nil {
		if node.id == id {
			if node.prev != nil {
				node.prev.next = node.next
			} else {
				head = node.next
			}
			if node.next != nil {
				node.next.prev = node.prev
			} else {
				tail = node.prev
				node = nil
			}
			break
		}
		node = node.next
	}
	fmt.Println("---------------------------------------------------")
	fmt.Printf("Barang Dengan Id : %s Telah Berhasil Dihapus\n", id)
	fmt.Println("---------------------------------------------------")

}

func updateStok(id string, stok int) {
	current := head
	for current != nil {
		if current.id == id {
			current.stok = stok
		}
		current = current.next
	}

}

func main() {
	var newId string
	var newStok int
	filePath := "data\\dataBarang.txt"
	node := readFileToNode(filePath)
	var pilih int
	yes := true
	for yes {
		fmt.Println("\n\n--------------------------------")
		fmt.Println("\tMENU MANAGEMENT BARANG")
		fmt.Println("------------------------------------")
		fmt.Println("1. TAMPILKAN BARANG")
		fmt.Println("2. TAMBAH BARANG")
		fmt.Println("3. HAPUS BARANG")
		fmt.Println("4. UBAH STOK BARANG")
		fmt.Println("5. KELUAR")
		fmt.Println("------------------------------------")
		fmt.Print("Masukan Pilihan: ")
		fmt.Scanln(&pilih)
		fmt.Print("-----------------------------------\n\n\n")
		switch pilih {
		case 1:
			displayDataFromDatabase(node)
		case 2:
			fmt.Println("\n\n========================================")
			fmt.Println("\tTAMBAH DATA BARANG")
			fmt.Print("id: ")
			fmt.Scanln(&newId)
			fmt.Print("nama: ")
			input := bufio.NewReader(os.Stdin)
			newNama, err := input.ReadString('\n') //dibaca sampai newline
			if err != nil {
				fmt.Println("error reading newNama", err)
			}
			newNama = strings.TrimSpace(newNama)
			fmt.Print("stok: ")
			fmt.Scanln(&newStok)
			fmt.Print("========================================\n\n\n")
			addNode(newId, newNama, newStok)
			addToFile(filePath)
		case 3:
			var deleteById string
			fmt.Println("\n\n---------------------------------")
			fmt.Print("Masukan Id Barang Yang Akan Dihapus: ")
			fmt.Scanln(&deleteById)
			fmt.Print("-------------------------------------\n\n\n")
			deleteDataBarangBYId(node, deleteById)
			addToFile(filePath)
		case 4:
			fmt.Println("\n\n---------------------------------")
			fmt.Print("Masukan Id Barang Yang Akan Diupdate Stoknya: ")
			fmt.Scanln(&newId)
			fmt.Print("Stok Baru: ")
			fmt.Scanln(&newStok)
			updateStok(newId, newStok)
			addToFile(filePath)
			fmt.Print("-------------------------------------\n\n\n")
		case 5:
			fmt.Println("================================")
			fmt.Println("\t\tBYE BYE")
			fmt.Println("================================")
			yes = false
		}
	}
}
