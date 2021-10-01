package main

import "fmt"

var denuncia string
var Denuncias = []string{}

func main() {

	Menu()


}

func EscreverDenuncias(){
	fmt.Scanf("%s", &denuncia)
}

func MostrarDenuncias() {
	fmt.Println("")
	for indice, valor := range Denuncias {
		fmt.Println("Denuncia ",indice,": ",valor)
	}
	fmt.Println("")
}

func Menu() {
	var opcao int
	fmt.Print("\n***opções:***" +
		"\n1: Adicionar Denuncia" +
		"\n2: Editar denuncia" +
		"\n3: Excluir denuncia" +
		"\n4: Exibir Denuncias"  +
		"\n5: Fechar programa")
	fmt.Print("\nEscolha uma opçao: ")
	fmt.Scanf("%d", &opcao)

	switch opcao {
	case 1:
		AdicionarDenuncias2()
	case 2:
		EditarDenuncias()
	case 3:
		ExcluirDenuncias2()
	case 4:
		MostrarDenuncias()
		fmt.Print("Tecle ENTER para continuar")
		fmt.Scanf("&d")
		Menu()
	case 5:
		fmt.Println("Fim")
		break
	default:
		fmt.Println("\n*********Escolha uma opção valida*********")
		Menu()

	}
}

/*func AdicionarDenuncias(){

	for i:=0;i<5;i++ {

		fmt.Printf("Informe a sua denuncia: ")
		EscreverDenuncias()
		Denuncias[i] = denuncia
		MostrarDenuncias()

	}
	Menu()
}*/

func EditarDenuncias(){
	var num int
	fmt.Print("\nInforme o numero da denunca que deseja editar: ")
	fmt.Scanf("%d", &num)

	fmt.Print("\nReescreva a denuncia: ")
	fmt.Scanf("%s", &Denuncias[num])


	Menu()
}

/*func ExcluirDenuncias()  {
	var num int
	fmt.Print("\nInforme o numero da denunca que deseja editar: ")
	fmt.Scanf("%d", &num)
	Denuncias[num] = ""


	Menu()

}*/

func ExcluirDenuncias2()  {
	var num int
	fmt.Print("\nInforme o numero da denunca que deseja editar: ")
	fmt.Scanf("%d", &num)
	Denuncias = append(Denuncias[:num],Denuncias[num+1:]...)

	Menu()

}

func AdicionarDenuncias2(){
	fmt.Scanf("%s", &denuncia)
	Denuncias = append(Denuncias, denuncia)
	Menu()
}







