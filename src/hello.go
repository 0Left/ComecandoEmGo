//Sempre o arquivo principal precisa de um desses para indicar ser a pagina central
package main

//Importa o pacote para usar o PrintLn
import "fmt"

//toda pagina main precisa de sua função correspondente
func main(){
	fmt.Println("Ola Mundo!");
}

//go build hello.go Compila o script
//go run hello.go Compila e executa o script