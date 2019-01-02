package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"
	goa2sample "github.com/tonouchi510/goa2-sample/controllers"
	securedsvr "github.com/tonouchi510/goa2-sample/gen/http/secured/server"
	swaggersvr "github.com/tonouchi510/goa2-sample/gen/http/swagger/server"
	userssvr "github.com/tonouchi510/goa2-sample/gen/http/users/server"
	vironsvr "github.com/tonouchi510/goa2-sample/gen/http/viron/server"
	secured "github.com/tonouchi510/goa2-sample/gen/secured"
	users "github.com/tonouchi510/goa2-sample/gen/users"
	viron "github.com/tonouchi510/goa2-sample/gen/viron"
	goahttp "goa.design/goa/http"
	"goa.design/goa/http/middleware"
)

func main() {
	// Define command line flags, add any other flag required to configure
	// the service.
	var (
		addr = flag.String("listen", "localhost:8000", "HTTP listen `address`")
		dbg  = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// Setup logger and goa log adapter. Replace logger with your own using
	// your log package of choice.
	var (
		adapter middleware.Logger
		logger  *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[goa2sample] ", log.Ltime)
		adapter = middleware.NewLogger(logger)
	}

	// Initialize service dependencies such as databases.
	var (
		db *sql.DB
	)
	{
		var err error
		db, err = sql.Open("mysql", "test:test@/testdb")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer db.Close()
	}

	// Create the structs that implement the services.
	var (
		usersSvc   users.Service
		securedSvc secured.Service
		vironSvc   viron.Service
	)
	{
		usersSvc = goa2sample.NewUsers(logger, db)
		securedSvc = goa2sample.NewSecured(logger)
		vironSvc = goa2sample.NewViron(logger)
	}

	// Wrap the services in endpoints that can be invoked from other
	// services potentially running in different processes.
	var (
		usersEndpoints   *users.Endpoints
		securedEndpoints *secured.Endpoints
		vironEndpoints   *viron.Endpoints
	)
	{
		usersEndpoints = users.NewEndpoints(usersSvc)
		securedEndpoints = secured.NewEndpoints(securedSvc)
		vironEndpoints = viron.NewEndpoints(vironSvc)
	}

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		usersServer   *userssvr.Server
		swaggerServer *swaggersvr.Server
		securedServer *securedsvr.Server
		vironServer   *vironsvr.Server
	)
	{
		eh := ErrorHandler(logger)
		usersServer = userssvr.New(usersEndpoints, mux, dec, enc, eh)
		swaggerServer = swaggersvr.New(nil, mux, dec, enc, eh)
		securedServer = securedsvr.New(securedEndpoints, mux, dec, enc, eh)
		vironServer = vironsvr.New(vironEndpoints, mux, dec, enc, eh)
	}
	// Configure the mux.
	userssvr.Mount(mux, usersServer)
	swaggersvr.Mount(mux, swaggerServer)
	securedsvr.Mount(mux, securedServer)
	vironsvr.Mount(mux, vironServer)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		if *dbg {
			handler = middleware.Debug(mux, os.Stdout)(handler)
		}
		handler = middleware.Log(adapter)(handler)
		handler = middleware.RequestID()(handler)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)
	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the service to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()
	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: *addr, Handler: handler}
	go func() {
		for _, m := range usersServer.Mounts {
			logger.Printf("method %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
		}
		for _, m := range swaggerServer.Mounts {
			logger.Printf("file %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
		}
		for _, m := range securedServer.Mounts {
			logger.Printf("method %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
		}
		for _, m := range vironServer.Mounts {
			logger.Printf("file %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
		}
		logger.Printf("listening on %s", *addr)
		errc <- srv.ListenAndServe()
	}()

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)
	// Shutdown gracefully with a 30s timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	logger.Println("exited")
}

// ErrorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func ErrorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Printf("[%s] ERROR: %s", id, err.Error())
	}
}
