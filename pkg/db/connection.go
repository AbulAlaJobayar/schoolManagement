// package db

// import (
// 	"log"
// 	"os"
// 	"os/user"
// 	"schoolmanagement/internal/order"
// 	"schoolmanagement/internal/product"
// 	"schoolmanagement/pkg/config"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// type DbInstance struct {
// 	Db *gorm.DB
// }

// var Database DbInstance

// func ConnectDb() {
// 	db, err := gorm.Open(postgres.Open(config.AppConfig.Database), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal("failed to connect to the Database!\n", err.Error())
// 		os.Exit(2)
// 	}
// 	log.Println("Connect to The Database successfully")
// 	db.Logger = logger.Default.LogMode(logger.Info)
// 	log.Println("Running migration")
// 	// TODO: add migration

// 	db.AutoMigrate(&user.User{}, &product.Product{}, &order.Order{})
// 	Database = DbInstance{Db: db}
// }

// package db

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"log/slog"
// 	"os"
// 	"time"
// 	"schoolmanagement/pkg/config"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// // Database represents the database instance with additional context
// type Database struct {
// 	DB        *gorm.DB
// 	ctx       context.Context
// 	cancel    context.CancelFunc
// 	isHealthy bool
// }

// // Instance holds the global database instance
// var Instance *Database

// // Connect initializes the database connection with proper configuration
// func Connect() error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
// 	defer cancel()

// 	instance := &Database{
// 		ctx:    ctx,
// 		cancel: cancel,
// 	}

// 	if err := instance.initialize(); err != nil {
// 		return fmt.Errorf("failed to initialize database: %w", err)
// 	}

// 	Instance = instance
// 	return nil
// }

// // initialize sets up the database connection with proper configuration
// func (db *Database) initialize() error {
// 	// Validate database configuration
// 	if config.AppConfig.Database == "" {
// 		return fmt.Errorf("database connection string is empty")
// 	}

// 	// Create custom logger for GORM
// 	gormLogger := logger.New(
// 		slog.NewLogLogger(slog.Default().Handler(), slog.LevelInfo),
// 		logger.Config{
// 			SlowThreshold:             time.Second,
// 			LogLevel:                  getGormLogLevel(),
// 			IgnoreRecordNotFoundError: true,
// 			Colorful:                  false,
// 		},
// 	)

// 	// Configure GORM
// 	gormConfig := &gorm.Config{
// 		Logger:                                   gormLogger,
// 		PrepareStmt:                              true,
// 		DisableForeignKeyConstraintWhenMigrating: false,
// 		SkipDefaultTransaction:                   true,
// 		NowFunc: func() time.Time {
// 			return time.Now().UTC()
// 		},
// 	}

// 	// Establish database connection
// 	slog.Info("Connecting to database", "timeout", "30s")

// 	dbConn, err := gorm.Open(postgres.Open(config.AppConfig.Database), gormConfig)
// 	if err != nil {
// 		return fmt.Errorf("failed to open database connection: %w", err)
// 	}

// 	// Configure connection pool
// 	sqlDB, err := dbConn.DB()
// 	if err != nil {
// 		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
// 	}

// 	configureConnectionPool(sqlDB)
// 	db.DB = dbConn

// 	// Verify connection
// 	if err := db.verifyConnection(); err != nil {
// 		return fmt.Errorf("database connection verification failed: %w", err)
// 	}

// 	slog.Info("Database connection established successfully")
// 	return nil
// }

// // configureConnectionPool sets optimal connection pool settings
// func configureConnectionPool(sqlDB *sql.DB) {
// 	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
// 	sqlDB.SetMaxIdleConns(10)

// 	// SetMaxOpenConns sets the maximum number of open connections to the database.
// 	sqlDB.SetMaxOpenConns(100)

// 	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
// 	sqlDB.SetConnMaxLifetime(time.Hour)

// 	// SetConnMaxIdleTime sets the maximum amount of time a connection may be idle.
// 	sqlDB.SetConnMaxIdleTime(30 * time.Minute)
// }

// // verifyConnection pings the database to ensure connectivity
// func (db *Database) verifyConnection() error {
// 	sqlDB, err := db.DB.DB()
// 	if err != nil {
// 		return fmt.Errorf("failed to get sql.DB for ping: %w", err)
// 	}

// 	ctx, cancel := context.WithTimeout(db.ctx, 5*time.Second)
// 	defer cancel()

// 	if err := sqlDB.PingContext(ctx); err != nil {
// 		return fmt.Errorf("database ping failed: %w", err)
// 	}

// 	db.isHealthy = true
// 	return nil
// }

// // HealthCheck returns the health status of the database connection
// func (db *Database) HealthCheck() bool {
// 	if db == nil || db.DB == nil || !db.isHealthy {
// 		return false
// 	}

// 	sqlDB, err := db.DB.DB()
// 	if err != nil {
// 		db.isHealthy = false
// 		return false
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	if err := sqlDB.PingContext(ctx); err != nil {
// 		db.isHealthy = false
// 		return false
// 	}

// 	return true
// }

// // Close gracefully closes the database connection
// func (db *Database) Close() error {
// 	if db == nil || db.DB == nil {
// 		return nil
// 	}

// 	if db.cancel != nil {
// 		db.cancel()
// 	}

// 	sqlDB, err := db.DB.DB()
// 	if err != nil {
// 		return fmt.Errorf("failed to get sql.DB for closing: %w", err)
// 	}

// 	db.isHealthy = false
// 	if err := sqlDB.Close(); err != nil {
// 		return fmt.Errorf("failed to close database connection: %w", err)
// 	}

// 	slog.Info("Database connection closed successfully")
// 	return nil
// }

// // WithContext returns a database instance with context
// func (db *Database) WithContext(ctx context.Context) *gorm.DB {
// 	if db == nil || db.DB == nil {
// 		return nil
// 	}
// 	return db.DB.WithContext(ctx)
// }

// // Transaction executes a function within a database transaction
// func (db *Database) Transaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
// 	if db == nil || db.DB == nil {
// 		return fmt.Errorf("database not initialized")
// 	}

// 	return db.DB.WithContext(ctx).Transaction(fn)
// }

// // getGormLogLevel returns the appropriate GORM log level based on environment
// func getGormLogLevel() logger.LogLevel {
// 	env := os.Getenv("APP_ENV")
// 	switch env {
// 	case "production":
// 		return logger.Warn
// 	case "test":
// 		return logger.Silent
// 	default:
// 		return logger.Info
// 	}
// }

// // Migrate runs database migrations
// func (db *Database) Migrate(models ...interface{}) error {
// 	if db == nil || db.DB == nil {
// 		return fmt.Errorf("database not initialized")
// 	}

// 	slog.Info("Running database migrations")

// 	if err := db.DB.AutoMigrate(models...); err != nil {
// 		return fmt.Errorf("failed to run migrations: %w", err)
// 	}

// 	slog.Info("Database migrations completed successfully")
// 	return nil
// }

// // GetStats returns database connection statistics
// func (db *Database) GetStats() (*sql.DBStats, error) {
// 	if db == nil || db.DB == nil {
// 		return nil, fmt.Errorf("database not initialized")
// 	}

// 	sqlDB, err := db.DB.DB()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get sql.DB for stats: %w", err)
// 	}

// 	stats := sqlDB.Stats()
// 	return &stats, nil
// }
// Proper error handling with wrapped errors and specific error types

// Context support for timeouts and cancellation

// Connection pooling with optimal settings

// Health checks for monitoring

// Structured logging using slog

// Graceful shutdown support

// Transaction management helper

// Environment-aware logging levels

// Configuration validation

// Comprehensive documentation

// Connection statistics

// Thread-safe design

// Proper resource cleanup

// Type-safe database operations
package db

import (
	"log"
	"os"
	"schoolmanagement/pkg/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Database *gorm.DB

func ConnectDb() {
	db, err := gorm.Open(postgres.Open(config.AppConfig.Database), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
		os.Exit(2)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	Database = db
	log.Println(" Database connected successfully")
}