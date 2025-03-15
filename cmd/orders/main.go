package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"go.uber.org/zap"
)

// ~/go/bin/goose postgres "postgres://postgres:postgres@0.0.0.0:5432/postgres" status
// psql -h 0.0.0.0 -p 5432 -U postgres
func main() {

	ctx := context.Background()

	// Open a PostgreSQL database.
	dsn := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	// Create a Bun db on top of it.
	db := bun.NewDB(pgdb, pgdialect.New())

	// Print all queries to stdout.
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	type User struct {
		ID  uuid.UUID
		Inn string
	}

	users := make([]User, 0)
	// insert new clients
	err := db.NewRaw(
		"SELECT id, inn FROM ? LIMIT ?",
		bun.Ident("clients"), 100,
	).Scan(ctx, &users)
	for _, u := range users {
		fmt.Println(u.ID)
		fmt.Println(u.Inn)
	}
	/* var rnd float64

	Select a random number.
	if err := db.NewSelect().ColumnExpr("random()").Scan(ctx, &rnd); err != nil {
		panic(err)
	}

	fmt.Println(rnd)
	*/
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("Ошибка ", err)
		os.Exit(1)
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	// id := uuid.New()
	// fmt.Println(id)
	str := "42"
	fmt.Printf("Строка %s\n", str)
	i, _ := strconv.Atoi(str)
	fmt.Printf("int %d\n", i)

	if err = run(sugar); err != nil {
		sugar.Errorw("Стартап", "err", err)
		logger.Sync()
		os.Exit(1)
	}

}
func sayHello(w http.ResponseWriter, r *http.Request) {
	//Body, err := io.ReadAll(r.Body)
	/*if err!= nil{
		return
	}*/
	//defer r.Body.Close()

	//token := r.Header.Get("Токен авторизации")
	//fmt.Fprintln(w, "Вызван метод ТИП", r.Method, r.URL.Path, "Тело", r.Body, "Токен", token)
}

type Cfg struct {
	Env string
}

func run(log *zap.SugaredLogger) error {

	var c Cfg
	err := envconfig.Process("APP", &c)
	if err != nil {
		log.Errorw("Запуск", "Инициализация конфига", err.Error())
	}
	log.Infow("Test", "Env", c.Env)

	h := http.HandlerFunc(sayHello)
	http.ListenAndServe("localhost:7077", h)

	/*shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	sig := <-shutdown
	log.Infow("Завершение", "Статус", "Старт завершения", " Сигнал", sig)*/
	return nil
}
