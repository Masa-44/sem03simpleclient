package main

import (
	"log"
	"net"
	"os"

	"github.com/Masa-44/mycrypt" // Anta at dette er importen for den fiktive mycrypt-pakken
)

func main() {
	conn, err := net.Dial("tcp", "172.17.0.3:8300")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("os.Args[1] = ", os.Args[1])

	// Krypterer input fra kommandolinjen med mycrypt-pakken
	kryptertMelding := mycrypt.Krypter([]rune(os.Args[1]), mycrypt.ALF_SEM03, 4)
	log.Println("Kryptert melding: ", string(kryptertMelding))

	// Sender de krypterte dataene til serveren
	_, err = conn.Write([]byte(string(kryptertMelding)))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	response := string(buf[:n])
	log.Printf("Svar fra server: %s", response)
}
