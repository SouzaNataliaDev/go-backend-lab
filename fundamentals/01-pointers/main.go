package main

import "fmt"

type Account struct {
	Owner   string
	Balance int64
}

// Sem ponteiro, a struct é copiada. A mudança fica só na cópia local.
func depositCopy(account Account, cents int64) { account.Balance += cents }

// *Account é um ponteiro: permite acessar o objeto original.
// Go continua passando por valor: aqui o valor copiado é o endereço.
func deposit(account *Account, cents int64) {
	// Nunca desreferencie um ponteiro nil.
	if account == nil {
		return
	}
	// account.Balance equivale a (*account).Balance para este acesso.
	account.Balance += cents
}

func main() {
	account := Account{Owner: "Natália", Balance: 10000}
	depositCopy(account, 5000)
	fmt.Println("Após cópia:", account.Balance)    // 10000
	deposit(&account, 5000)                        // & obtém o endereço da variável.
	fmt.Println("Após ponteiro:", account.Balance) // 15000
	pointer := &account
	fmt.Println("Dono:", (*pointer).Owner) // * acessa o valor apontado.
	// USO: métodos que atualizam estado. Ponteiros não tornam acesso concorrente seguro.
}
