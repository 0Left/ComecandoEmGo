// Sempre o arquivo principal precisa de um desses para indicar ser a pagina central
package main

//Importa o pacote para usar o PrintLn
import (
	"fmt"
	"reflect"
)

//toda pagina main precisa de sua função correspondente
func main(){
	//declaração padrão(var + nome + tipo)
	var justDeclared string;
	//pode adicionar o valor para atribuir (= valor)
	var nome string = "Pedro";
	var idadeInt int = 23;
	var nota float32 = 1.0;
	//declaração dinamica (nome := valor)
	idadeDin := 10
	fmt.Println("Ola justDeclared!",justDeclared);
	fmt.Println("Ola nome!",nome);
	fmt.Println("Ola idadeInt!",idadeInt);
	fmt.Println("Ola nota!",nota);
	fmt.Println("Ola idadeDin!",idadeDin);

	fmt.Println("My type is",reflect.TypeOf(justDeclared));
	fmt.Println("My type is",reflect.TypeOf(idadeInt));
	fmt.Println("My type is",reflect.TypeOf(nota));
}

//; não é obrigatório mas ja é quase um reflexo meu, malz.