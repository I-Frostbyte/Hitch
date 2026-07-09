package users

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/I-Frostbyte/users/users"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	// "github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	logger := zerolog.New(os.Stderr).With().Timestamp().Caller().Logger()

	err := run(ctx, logger)
	if err != nil {
		logger.Err(err).Msg("failed to run grpc service")
		os.Exit(1)
	}
}

func run(ctx context.Context, logger zerolog.Logger) error {
	logger.Info().Msg("Starting grpc service...")

	config := model.Config{}

	err := config.LoadConfig()
	if err != nil {
		logger.Err(err).Msg("failed to load config")
		return err
	}

	logger.Info().Msgf("Successfully loaded config...: %+v", config)

	logLevel, err := zerolog.ParseLevel(config.LogLevel)
	if err != nil {
		return fmt.Errorf("failed to parse log level: %w", err)
	}
	logger = logger.Level(logLevel)

	/*
	dbConnectionURL := getPostgresConnectionURL(config.DB)
	db, err := pgxpool.New(ctx, dbConnectionURL)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	defer db.Close()

	err = repo.Migrate(dbConnectionURL, config.MigrationPath, logger)
	if err != nil {
		logger.Err(err).Msg("Migration not successful...")
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	usersRepo := repo.NewUsersRepo(db)
	
	*/
	// First start the gRPC server in a separate goroutine
	svrOpts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			grpc_recovery.UnaryServerInterceptor(),
		),
	}
	

	grpcServer := grpc.NewServer(svrOpts...)
	reflection.Register(grpcServer)

	usersgrpc.RegisterUsersServiceServer(grpcServer, users.NewUsersService(logger))
	logger.Info().Msg("Successfully registered UsersServiceServer...")

}
