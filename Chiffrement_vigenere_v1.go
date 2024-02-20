package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
)

var (
	alphabet = [27]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", " "}
)
var inputReader *bufio.Reader
var clee string
var entree string
var err error

func chiff_vigenere(mot string, cle string) string {
	var res string
	decc := 0
	dec := 0
	compt := 0
	var slice_pos []int = make([]int, len(cle))

	for k:=0; k<len(cle); k++ {
		for l:= 0; l<len(alphabet); l++ {
			if string(cle[k]) == alphabet[l] {
				slice_pos[k] = l
				// fmt.Printf("%d\n", slice_pos)
			}
		}
	}
	for i:=0; i<len(mot); i++ {
		for j:=0; j<len(alphabet); j++ {
			if string(mot[i]) == alphabet[j] {
				if compt<len(slice_pos) {
					decc = slice_pos[compt]
					compt++
				} else if compt>=len(slice_pos){
					compt = 0
					decc = slice_pos[compt]
					compt++
				}
				
				if j+decc > len(alphabet)-1 {
					dec = (j+decc)-(len(alphabet)-1)
					res += alphabet[dec-1]
					
				} else {
					res += alphabet[j+decc]
				}			
			}
		}			
	}
	return res
}

func dechiff_vigenere(chiffre string, cle string) string {
	resu := ""
	decaa := 0
	decb := 0
	compt := 0
	var slice_pos []int = make([]int, len(cle))

	for k:=0; k<len(cle); k++ {
		for l:= 0; l<len(alphabet); l++ {
			if string(cle[k]) == alphabet[l] {
				slice_pos[k] = l
			}
		}
	}

	for i:=0; i<len(chiffre); i++ {
		for j:=0; j<len(alphabet); j++ {
			if string(chiffre[i]) == alphabet[j] {
				if compt<len(slice_pos) {
					decb = slice_pos[compt]
					compt++
				} else if compt>=len(slice_pos){
					compt = 0
					decb = slice_pos[compt]
					compt++
				}

				if j-decb < 0 {
					decaa = len(alphabet) - (decb-j)
					resu += alphabet[decaa]
				} else {
					resu += alphabet[j-decb]
				}
			}
		}
	}
	return resu
}

func main() {
	fmt.Println("Chiffrement de Vigenere : ")

	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Input your encryption key (a word without number or space) : ")
	clee, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("Input : %s\n", clee)
	}
	
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("The sentence you want to encrypt : ")
	entree, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("Input : %s\n", entree)
	}
	mot_a_chiffrer := strings.ToLower(entree)
	fmt.Printf("%s\n", mot_a_chiffrer)

	fmt.Println("Encryption : ")
	chiffrement := chiff_vigenere(mot_a_chiffrer, clee)
	fmt.Printf("%s", chiffrement)
	fmt.Printf("\n\n")

	fmt.Println("Decryption : ")
	dechiffrement := dechiff_vigenere(chiffrement, clee)
	fmt.Printf("%s\n", dechiffrement)
	fmt.Println()
}
