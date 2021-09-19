package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/bethecodewithyou/gorest/gorilla/internal/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var totalRequest = prometheus.NewCounterVec(
	prometheus.CounterOpts{
			Name: "http_request_total",
			Help: "Number of get request",
	},
	[]string{"path"},
)

func prometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		
		next.ServeHTTP(w, r)

		totalRequest.With(prometheus.Labels{"path":"/products"}).Inc()

	})
}

func init(){
	prometheus.Register(totalRequest)
}

func main() {

	// string "Product API" is prefix which will print to every log statement.
	logger := log.New(os.Stdout, "Product API: " , log.LstdFlags)

	productHandler := handlers.NewProduct(logger)

	servMux := mux.NewRouter()

	servMux.Use(prometheusMiddleware)

	servMux.Path("/metrics").Handler(promhttp.Handler())

	// registers product handler methods to serve request on api end points with specific http methods.
	getHandler := servMux.Methods(http.MethodGet).Subrouter()
	getHandler.HandleFunc("/products", productHandler.GetProducts)

	postHandler := servMux.Methods(http.MethodPost).Subrouter()
	postHandler.HandleFunc("/products", productHandler.AddProduct)

	putHandler := servMux.Methods(http.MethodPut).Subrouter()
	putHandler.HandleFunc("/products/{id:[0-9]+}", productHandler.UpdateProduct)

	patchHandler := servMux.Methods(http.MethodPatch).Subrouter()
	patchHandler.HandleFunc("/products/{id:[0-9]+}", productHandler.UpdateProductAttribute)

	deleteHandler := servMux.Methods(http.MethodDelete).Subrouter()
	deleteHandler.HandleFunc("/products/{id:[0-9]+}", productHandler.DeleteProduct)

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

	// this context tells, when all the handlers finished their work and after that, 
	// wait for 30 seconds before trigerring server shutdown.
	// if there is another request comes within that 30 seconds, then this time will be reset again.

	logger.Println("server is stopping ")
	context, cancel := context.WithTimeout(context.Background(), 30 *time.Second)
	
	defer cancel()

	server.Shutdown(context)
	

}