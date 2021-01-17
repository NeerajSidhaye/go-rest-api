package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"github.com/bethecodewithyou/restcrud/internal/handlers"
)

func main() {

	// creating logger to write on the standard output. in this case, our command terminal.
	// string "Shoe Product API" is prfix which will print to every log statement.
	logger := log.New(os.Stdout, "Shoe Product API: " , log.LstdFlags)

	//creating product handler
	productHandler := handlers.NewProduct(logger)

	// creating serve Mux and registering the produc thandler
	servMux := http.NewServeMux()
	
	// registers product handler to serve request on url "/products".
	// this is simple http webservice hence you see the url has 'get' in it "/products"
	// this registration makes a call to ServeHTTP method of Product handler and from there response is returned.
	servMux.Handle("/v1/products", productHandler) 


	// create a new server
	server := http.Server {

		Addr: ":7070",  // listen from any ip address on the port 7070
		Handler: servMux,  // setting the default handler
		ErrorLog : logger,  // setting the logger for the server
		ReadTimeout: 10 *time.Second, // max time to read the request from the client.
		WriteTimeout: 5 *time.Second, // max time to write response back to the client.
		IdleTimeout: 100 *time.Second, // max time for connections keeping TCP keep-alive
	}

	// starting the server. This is a go routine.
	go func() {

		logger.Println("starting the server on port 7070")

		e := server.ListenAndServe()

		if e!=nil {
			logger.Fatal(e)
			os.Exit(1)  // non-zero value shows that program is terminating due to error.
		}

		logger.Println("server started on 7070")

	}()

	//this piece of code is to block the go routine. Without this code, 
	//without this code, our go routine will not have enough time to execute and hence our http server will never start.
	c:= make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// blocking routine until a signal is received 
	// as soon as we prcess ctrl+c, the go routine will stop and then server graceful shutdown code will start to execute
	sig := <- c
	logger.Println("got signal to stop the server:", sig)


	// gracefully shutting down the server. 

	// this context tells, when all the handlers finished thier work and after that, 
	// wait for 30 seconds before trigerring server shutdonw.
	// if there is another request comes within that 30 seconds, then this time will be reset again.

	logger.Println("server is stopping ")
	context, cancel := context.WithTimeout(context.Background(), 30 *time.Second)
	
	defer cancel()

	server.Shutdown(context)
	

}