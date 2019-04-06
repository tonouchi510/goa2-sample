package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	goa2sample "github.com/tonouchi510/goa2-sample/controller"
	admin "github.com/tonouchi510/goa2-sample/gen/admin"
	users "github.com/tonouchi510/goa2-sample/gen/users"
	viron "github.com/tonouchi510/goa2-sample/gen/viron"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "localhost", "Server host (valid values: localhost)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
		db *sql.DB
		err error
	)
	{
		logger = log.New(os.Stderr, "[goa2sample] ", log.Ltime)
		db, err = sql.Open("mysql", "test:test@/sampledb")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer db.Close()
	}

	// Initialize the services.
	var (
		usersSvc users.Service
		vironSvc viron.Service
		adminSvc admin.Service
	)
	{
		usersSvc = goa2sample.NewUsers(logger, db)
		vironSvc = goa2sample.NewViron(logger)
		adminSvc = goa2sample.NewAdmin(logger, db)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		usersEndpoints *users.Endpoints
		vironEndpoints *viron.Endpoints
		adminEndpoints *admin.Endpoints
	)
	{
		usersEndpoints = users.NewEndpoints(usersSvc)
		vironEndpoints = viron.NewEndpoints(vironSvc)
		adminEndpoints = admin.NewEndpoints(adminSvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "localhost":
		{
			addr := "http://localhost:8080"
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *httpPortF
			} else if u.Port() == "" {
				u.Host += ":80"
			}
			handleHTTPServer(ctx, u, usersEndpoints, vironEndpoints, adminEndpoints, &wg, errc, logger, *dbgF)
		}

	default:
		fmt.Fprintf(os.Stderr, "invalid host argument: %q (valid hosts: localhost)", *hostF)
	}

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}
