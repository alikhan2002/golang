package main

import (
	"assignment2.alikhan.net/internal/data"
	"assignment2.alikhan.net/internal/jsonlog"
	"assignment2.alikhan.net/internal/mailer"
	"context"
	"database/sql"
	"flag"
	_ "github.com/lib/pq"
	"os"
	"strings"
	"sync"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
	limiter struct {
		enabled bool
		rps     float64
		burst   int
	}
	smtp struct {
		host     string
		port     int
		username string
		password string
		sender   string
	}
	cors struct {
		trustedOrigins []string
	}
}

// Update the application struct to hold a new Mailer instance.
type application struct {
	config config
	logger *jsonlog.Logger
	models data.Models
	mailer mailer.Mailer
	wg     sync.WaitGroup
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	// flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("GREENLIGHT_DB_DSN"), "PostgreSQL DSN")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://alikhan:pa55word@postgres:5432/alikhan?sslmode=disable", "PostgreSQL DSN")

	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")
	flag.BoolVar(&cfg.limiter.enabled, "limiter-enabled", true, "Enable rate limiter")
	flag.Float64Var(&cfg.limiter.rps, "limiter-rps", 2, "Rate limiter maximum requests per second")
	flag.IntVar(&cfg.limiter.burst, "limiter-burst", 4, "Rate limiter maximum burst")

	flag.StringVar(&cfg.smtp.host, "smtp-host", "sandbox.smtp.mailtrap.io", "SMTP host")
	flag.IntVar(&cfg.smtp.port, "smtp-port", 25, "SMTP port")
	flag.StringVar(&cfg.smtp.username, "smtp-username", "46c668cde02d82", "SMTP username")
	flag.StringVar(&cfg.smtp.password, "smtp-password", "e00c4ce70e03c9", "SMTP password")
	flag.StringVar(&cfg.smtp.sender, "smtp-sender", "Greenlight <no-reply@greenlight.alexedwards.net>", "SMTP sender")

	flag.Func("cors-trusted-origins", "Trusted CORS origins (space separated)", func(val string) error {
		cfg.cors.trustedOrigins = strings.Fields(val)
		return nil
	})
	flag.Parse()
	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)
	db, err := openDB(cfg)
	if err != nil {
		logger.PrintFatal(err, nil)
	}
	defer db.Close()
	logger.PrintInfo("database connection pool established", nil)
	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
		mailer: mailer.New(cfg.smtp.host, cfg.smtp.port, cfg.smtp.username, cfg.smtp.password, cfg.smtp.sender),
	}
	err = app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}
func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	// db, err := sql.Open("postgres", "host=127.0.0.1 user=alikhan password=pa55word dbname=alikhan port=5432 sslmode=disable")

	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)
	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
