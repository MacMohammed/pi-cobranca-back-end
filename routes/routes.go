package routes

import (
	"fatec/controllers"
	"fatec/cookie"

	"github.com/gorilla/mux"
)

func LoadRoutes(rtr *mux.Router) {

	//Teste de Cookie
	rtr.HandleFunc("/deleteCookie", cookie.DeleteCookie).Methods("GET")

	//Teste de User
	rtr.HandleFunc("/users", controllers.AllUsers).Methods("GET")
	rtr.HandleFunc("/GetUser", controllers.GetUser).Methods("POST")
	rtr.HandleFunc("/UpdateUser", controllers.UpdateUser).Methods("POST")
	rtr.HandleFunc("/InsertUser", controllers.InsertUser).Methods("POST")
	rtr.HandleFunc("/DeletetUser", controllers.DeletetUser).Methods("POST")

	//Teste de Bank
	rtr.HandleFunc("/bancos", controllers.AllBank).Methods("GET")
	rtr.HandleFunc("/GetBank", controllers.GetBank).Methods("POST")
	rtr.HandleFunc("/UpdateBank", controllers.UpdateBank).Methods("POST")
	rtr.HandleFunc("/InsertBank", controllers.InsertBank).Methods("POST")
	rtr.HandleFunc("/DeleteBank", controllers.DeleteBank).Methods("POST")

	//Teste de Client
	rtr.HandleFunc("/clientes", controllers.AllClient).Methods("GET")
	rtr.HandleFunc("/GetClient", controllers.GetClient).Methods("POST")
	rtr.HandleFunc("/UpdateClient", controllers.UpdateClient).Methods("POST")
	rtr.HandleFunc("/InsertClient", controllers.InsertClient).Methods("POST")
	rtr.HandleFunc("/DeleteClient", controllers.DeleteClient).Methods("POST")

	//Teste de Office
	rtr.HandleFunc("/AllOffice", controllers.AllOffice).Methods("GET")
	rtr.HandleFunc("/GetOffice", controllers.GetOffice).Methods("POST")
	rtr.HandleFunc("/UpdateOffice", controllers.UpdateOffice).Methods("POST")
	rtr.HandleFunc("/InsertOffice", controllers.InsertOffice).Methods("POST")
	rtr.HandleFunc("/DeleteOffice", controllers.DeleteOffice).Methods("POST")

	//Teste de Transaction
	rtr.HandleFunc("/transacoes", controllers.AllTransaction).Methods("GET")
	rtr.HandleFunc("/GetTransaction", controllers.GetTransaction).Methods("POST")
	rtr.HandleFunc("/UpdateTransaction", controllers.UpdateTransaction).Methods("POST")
	rtr.HandleFunc("/transacao", controllers.InsertTransaction).Methods("POST")
	rtr.HandleFunc("/DeleteTransaction", controllers.DeleteTransaction).Methods("POST")
	rtr.HandleFunc("/trasancao/{id_transacao}", controllers.TransacaoBaixar).Methods("PUT")
	rtr.HandleFunc("/trasancao/cancelar/{id_transacao}", controllers.TransacaoCancelar).Methods("PUT")

	//Teste de File
	rtr.HandleFunc("/AllFile", controllers.AllFile).Methods("GET")
	rtr.HandleFunc("/GetFile", controllers.GetFile).Methods("POST")
	rtr.HandleFunc("/UpdateFile", controllers.UpdateFile).Methods("POST")
	rtr.HandleFunc("/InsertFile", controllers.InsertFile).Methods("POST")
	rtr.HandleFunc("/DeleteFile", controllers.DeleteFile).Methods("POST")
}
