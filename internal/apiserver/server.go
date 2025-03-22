package apiserver

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/tsukiyoz/knowlith/internal/apiserver/handler"
	mw "github.com/tsukiyoz/knowlith/internal/pkg/middlware"
	genericoptions "github.com/tsukiyoz/knowlith/pkg/options"
)

type Config struct {
	HttpOptions *genericoptions.HttpOptions `json:"http" mapstructure:"http"`
}

type Server struct {
	cfg *Config
	srv *http.Server
}

func (cfg *Config) NewServer() (*Server, error) {
	app := fiber.New(fiber.Config{
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		DisableStartupMessage: true,
	})

	app.Use(recover.New(), mw.NoCache, mw.Cors, mw.RequestID)

	cfg.InstallRESTAPI(app)
	srv := &http.Server{
		Addr:    cfg.HttpOptions.Addr,
		Handler: adaptor.FiberApp(app),
	}

	return &Server{cfg: cfg, srv: srv}, nil
}

func (cfg *Config) InstallRESTAPI(srv *fiber.App) {
	handler := handler.NewHandler()
	srv.All("/", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusNotFound, "Not Found")
	})

	v1 := srv.Group("/v1")
	{
		v1.Post("/prompt", handler.Prompt)
	}
}

func (s *Server) Run() error {
	slog.Info("APIServer is running ...", slog.String("Addr", s.cfg.HttpOptions.Addr))
	go func() {
		if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error(err.Error())
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("APIServer is shutting down ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		slog.Error("Insecure Server forced to shutdown", "err", err)
		return err
	}

	slog.Info("APIServer exit.")

	return nil
}
