## Route Manager

Um básico roteador de rotas, criádo como intuito de melhorar minhas  hábilidades em Golang. Fique avontade para 
dar sugestões de melhorias, no momento este mini projeto é uma MVP, pretendo ir adicionando novas funcionalidades 
conforme minhas hábilidades em Golang forem sendo aperfeiçoadas.

Exemplo:

````
package main

import (
	"fmt"
	"net/http"
	"os"
	_ "routemanager/utils"
)

func main() {

	rm := RouteManagerInit()

	rm.Get("/user/all", func(w http.ResponseWriter, r *http.Request, params map[string]interface{}) {
		fmt.Println("All users")
	})

	rm.Get("/user/select/field:id", func(w http.ResponseWriter, r *http.Request, params map[string]interface{}) {
		fmt.Println("select user by id:", params["id"])
	})

	rm.Get("/user/delete/field:id", func(w http.ResponseWriter, r *http.Request, params map[string]interface{}) {
		fmt.Println("delete user by id:", params["id"])
	})

	rm.Get("/user/update/field:id", func(w http.ResponseWriter, r *http.Request, params map[string]interface{}) {
		fmt.Println("update user by id:", params["id"])
	})

	addr := fmt.Sprintf("%s%s", os.Getenv("HOST_NAME"), os.Getenv("HOST_PORT"))
	http.HandleFunc("/", rm.HandleFunc)
	http.ListenAndServe(addr, nil)
}

````


### Road Maps
- Melhorar legibilidade do código
- Testes unitários