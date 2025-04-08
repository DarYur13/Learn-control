package service_provider

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/DarYur13/learn-control/internal/config"
	"github.com/DarYur13/learn-control/internal/logger"
	pb "github.com/DarYur13/learn-control/pkg/learn_control"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// App is the application struct
type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
	httpServer      *http.Server
}

// NewApp creates a new App
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}

	return a, nil
}

// Run starts the app (both gRPC and HTTP)
func (a *App) Run(ctx context.Context) error {
	group, _ := errgroup.WithContext(ctx)

	group.Go(func() error {
		list, err := net.Listen("tcp", fmt.Sprintf(":%s", config.ApiGrpcPort()))
		if err != nil {
			return fmt.Errorf("failed to listen: %w", err)
		}

		logger.Info("gRPC сервер запущен на порту", config.ApiGrpcPort())

		return a.grpcServer.Serve(list)
	})

	group.Go(func() error {
		logger.Info("HTTP сервер запущен на порту", config.ApiHttpPort())

		return a.httpServer.ListenAndServe()
	})

	// Ждём завершения или ошибки
	return group.Wait()
}

// Shutdown gracefully stops servers
func (a *App) Shutdown(ctx context.Context) {
	if a.grpcServer != nil {
		logger.Info("Остановка gRPC сервера...")
		a.grpcServer.GracefulStop()
	}

	if a.httpServer != nil {
		logger.Info("Остановка HTTP сервера...")
		if err := a.httpServer.Shutdown(ctx); err != nil {
			logger.Error("ошибка при остановке HTTP сервера: ", err)
		}
	}
}

func (a *App) initDeps(ctx context.Context) error {
	initFns := []func(context.Context) error{
		a.initConfig,
		a.initLogger,
		a.initServiceProvider,
		a.initGrpcServer,
		a.initHTTPServer,
	}

	for _, fn := range initFns {
		if err := fn(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	config.LoadAll()

	return nil
}

func (a *App) initLogger(_ context.Context) error {
	lg, err := logger.New(config.LogFilePath())
	if err != nil {
		return fmt.Errorf("logger setup error: %w", err)
	}

	logger.SetLogger(lg)
	if err := logger.SetLogLevel(config.LogLevel()); err != nil {
		return fmt.Errorf("log level error: %w", err)
	}

	logger.Sync()

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()

	return nil
}

func (a *App) initGrpcServer(ctx context.Context) error {
	s := grpc.NewServer()

	pb.RegisterLearnControlServer(s, a.serviceProvider.getLearnControl(ctx))
	reflection.Register(s)

	a.grpcServer = s

	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()} // ⚠️ Не для продакшна

	if err := pb.RegisterLearnControlHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%s", config.ApiGrpcPort()), opts); err != nil {
		return err
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:3000", "https://6dnqnvhj-5173.uks1.devtunnels.ms"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
		Debug:            true,
	})

	a.httpServer = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.ApiHost(), config.ApiHttpPort()),
		Handler: c.Handler(mux),
	}

	return nil
}
