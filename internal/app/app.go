package service_provider

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	filesDownload "github.com/DarYur13/learn-control/internal/adapter/controller/files_download"
	"github.com/DarYur13/learn-control/internal/config"
	"github.com/DarYur13/learn-control/internal/logger"
	worker "github.com/DarYur13/learn-control/internal/worker/notification"
	pb "github.com/DarYur13/learn-control/pkg/learn_control"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	serviceProvider    *serviceProvider
	grpcServer         *grpc.Server
	httpServer         *http.Server
	notificationWorker worker.NotificationWorker
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run(ctx context.Context) error {
	group, _ := errgroup.WithContext(ctx)

	group.Go(func() error {
		list, err := net.Listen("tcp", fmt.Sprintf(":%s", config.ApiGrpcPort()))
		if err != nil {
			return fmt.Errorf("failed to listen: %w", err)
		}

		logger.Infof("gRPC server started at port: %s", config.ApiGrpcPort())

		return a.grpcServer.Serve(list)
	})

	group.Go(func() error {
		logger.Infof("HTTP server started at port: %s", config.ApiHttpPort())

		return a.httpServer.ListenAndServe()
	})

	group.Go(func() error {
		logger.Info("Notification worker started")
		a.notificationWorker.StartNotify(ctx)
		return nil
	})

	return group.Wait()
}

func (a *App) Shutdown(ctx context.Context) {
	if a.grpcServer != nil {
		logger.Info("Stopping gRPC server...")
		a.grpcServer.GracefulStop()
	}

	if a.httpServer != nil {
		logger.Info("Stopping HTTP server...")
		if err := a.httpServer.Shutdown(ctx); err != nil {
			logger.Error("failed to stop HTTP server: %s", err)
		}
	}
}

func (a *App) initDeps(ctx context.Context) error {
	initFns := []func(context.Context) error{
		a.initConfig,
		a.initLogger,
		a.initServiceProvider,
		a.initNotificationWorker,
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
	grpcMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()} // не для продакшна

	if err := pb.RegisterLearnControlHandlerFromEndpoint(ctx, grpcMux, fmt.Sprintf(":%s", config.ApiGrpcPort()), opts); err != nil {
		return err
	}

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
		Debug:            true,
	})

	filesMux := http.NewServeMux()
	filesMux.Handle("/files/download", filesDownload.New(a.serviceProvider.getService(ctx)))

	mainMux := http.NewServeMux()
	mainMux.Handle("/", corsHandler.Handler(grpcMux))
	mainMux.Handle("/files/download", filesMux)

	a.httpServer = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.ApiHost(), config.ApiHttpPort()),
		Handler: mainMux,
	}

	return nil
}

func (a *App) initNotificationWorker(ctx context.Context) error {
	a.notificationWorker = worker.New(
		a.serviceProvider.getEmplRepo(ctx),
		a.serviceProvider.getTrainingsRepo(ctx),
		a.serviceProvider.getNotificationsRepo(ctx),
		a.serviceProvider.getDocsGenerator(ctx),
		a.serviceProvider.getNotifier(ctx),
		time.Duration(config.NotificationWorkerQueueCheckPeriod())*time.Minute,
	)

	logger.Infof("Notification worker initialized. Interval: %v minutes", a.notificationWorker.Interval())

	return nil
}
