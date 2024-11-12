package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("Ошибка ", err)
		os.Exit(1)
	}
	defer logger.Sync()
	sugar := logger.Sugar()

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
