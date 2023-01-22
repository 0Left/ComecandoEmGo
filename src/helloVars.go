// Sempre o arquivo principal precisa de um desses para indicar ser a pagina central
package main

//Importa o pacote para usar o PrintLn
import (
	"fmt"
	"reflect"
)

//toda pagina main precisa de sua função correspondente
func main(){

	//declaração longa + atribuição
	var idadeInt int = 23;
	var nota float32 = 1.5;

	//declaração longa (var + nome + tipo)
	var justDeclared string;

	//declaração com auto atribuiçãol de tipos (usar com tipos faceis)
	var nome = "Pedro";

	//declaração dinamica (nome := valor)
	idadeDin := 10

	//printa tudo pra ver, com os tipos
	fmt.Println("Ola justDeclared!",justDeclared,reflect.TypeOf(justDeclared));
	fmt.Println("Ola nome!",nome,reflect.TypeOf(nome));
	fmt.Println("Ola idadeInt!",idadeInt,reflect.TypeOf(idadeInt));
	fmt.Println("Ola nota!",nota,reflect.TypeOf(nota));
	fmt.Println("Ola idadeDin!",idadeDin,reflect.TypeOf(idadeDin));
}

//; não é obrigatório mas ja é quase um reflexo meu, malz.