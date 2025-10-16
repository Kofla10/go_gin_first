package database
import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/jackc/pgx/v4/pgxpool"
)

//Global variable to hold the connection pool
var DB *pgxpool.Pool

func ConnectDB() {
	user:= "postgres"
	password := ""
	host := "localhost"
	port := 5432
	database := "gogin"

	//Data Source Name (DSN)
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", user, password, host, port, database)

	var err error

	DB, err = pgxpool.Connect(context.Background(), dsn)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	//Opcional: probar la conexión
	err = DB.Ping(context.Background())
	if err != nil{
		log.Fatal("Could not ping database:", err)
	}

	fmt.Println("successfully connected to poosgresql")
}

//OJO: Esta función debe ser llamada al incio en el main.go