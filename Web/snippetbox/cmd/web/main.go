package main

import (
  "database/sql"
  "log"
  "flag"
  "net/http"
  "os"

  _ "github.com/go-sql-driver/mysql"
  "github.com/Arctic-Husky/Aprendendo-GO/pkg/models/mysql"
)

type application struct{
  errorLog *log.Logger
  infoLog *log.Logger
  snippets *mysql.SnippetModel
}

// curl -i -X POST http://localhost:4000/snippet/create

//router www.site.com/coisa1/coisa2
//handlers tratar de requisições
//servidor

func main() {
              // nome da flag, valor e descricao
  addr := flag.String("addr", ":4000", "Porta da Rede")

  dsn := flag.String("dsn",
                     "1rJx8vlHTM:6IuGT3vqrE@tcp(remotemysql.com)/1rJx8vlHTM?parseTime=true", 
                     "MySql DSN")
  
  flag.Parse()

  infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
  errorLog := log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)

  db, err := openDB(*dsn)
  if err != nil{
    errorLog.Fatal(err)
  }
  defer db.Close()
  
  // new application
  app := &application{
    errorLog:errorLog,
    infoLog:infoLog,
    snippets: &mysql.SnippetModel{DB:db},
  }

  //mux := app.routes()

  srv := &http.Server{
    Addr: *addr,
    ErrorLog: errorLog,
    Handler: app.routes(),
  }
  
  app.infoLog.Printf("Inicializando o servidor na porta: %s\n", *addr)
  err = srv.ListenAndServe()
  app.errorLog.Fatal(err)
  
//   //multiplexador em concorrência, o escutador
//   mux := http.NewServeMux()
//   //uma rota, chamar "home" a cada /
//   mux.HandleFunc("/", home) // '/' funciona como um wildcard
//   mux.HandleFunc("/snippet", showSnippet)
//   mux.HandleFunc("/snippet/create", createSnippet)

//   log.Printf("\n\tInicializando servidor na porta: \"4000\"\n")

//   err := http.ListenAndServe(":4000", mux)

//   log.Fatal(err)
// }

// func home(rw http.ResponseWriter, r *http.Request) {
//   if r.URL.Path != "/"{
//     http.NotFound(rw,r)
//     return
//   }
//   rw.Write([]byte("Bem vindo ao SnippetBox"))
// }

// func showSnippet(rw http.ResponseWriter, r *http.Request) {
//   rw.Write([]byte("Mostrando um snippet espicífico"))
// }

// func createSnippet(rw http.ResponseWriter, r *http.Request) {

//   //para testar: curl -i -X POST http://localhost:4000
//   if r.Method != "POST" {
    
//     rw.Header().Set("Allow","POST")
//     //rw.WriteHeader(405)
//     http.Error(rw, "Metodo não permitido", http.StatusMethodNotAllowed)
//     //rw.Write([]byte("Metodo não permitido"))
//     return
//   }
//   rw.Write([]byte("Criando um novo snippet"))
}

func openDB(dsn string) (*sql.DB, error){
  db, err := sql.Open("mysql", dsn)
  if err!= nil{
    return nil, err
  }
  if err = db.Ping(); err != nil{
    return nil, err
  }
  return db, nil
}