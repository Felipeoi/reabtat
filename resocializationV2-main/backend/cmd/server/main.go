package main

import (
	"context"
	authHttp "github.com/jeje-gab/resocializationV2/backend/internal/collections/auth/delivery/http"
	authRepository "github.com/jeje-gab/resocializationV2/backend/internal/collections/auth/repository"
	authUseCase "github.com/jeje-gab/resocializationV2/backend/internal/collections/auth/usecase"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	userHttp "github.com/jeje-gab/resocializationV2/backend/internal/collections/user/delivery/http"
	userRepository "github.com/jeje-gab/resocializationV2/backend/internal/collections/user/repository"
	userUseCase "github.com/jeje-gab/resocializationV2/backend/internal/collections/user/usecase"

	inmatesHttp "github.com/jeje-gab/resocializationV2/backend/internal/collections/inmates/delivery/http"
	inmatesRepository "github.com/jeje-gab/resocializationV2/backend/internal/collections/inmates/repository"
	inmatesUseCase "github.com/jeje-gab/resocializationV2/backend/internal/collections/inmates/usecase"

	ufsHttp "github.com/jeje-gab/resocializationV2/backend/internal/collections/ufs/delivery/http"
	ufsRepository "github.com/jeje-gab/resocializationV2/backend/internal/collections/ufs/repository"
	ufsUseCase "github.com/jeje-gab/resocializationV2/backend/internal/collections/ufs/usecase"

	citiesHttp "github.com/jeje-gab/resocializationV2/backend/internal/collections/cities/delivery/http"
	citiesRepository "github.com/jeje-gab/resocializationV2/backend/internal/collections/cities/repository"
	citiesUseCase "github.com/jeje-gab/resocializationV2/backend/internal/collections/cities/usecase"

	matchHttp "github.com/jeje-gab/resocializationV2/backend/internal/collections/match/delivery/http"
	matchRepository "github.com/jeje-gab/resocializationV2/backend/internal/collections/match/repository"
	matchUseCase "github.com/jeje-gab/resocializationV2/backend/internal/collections/match/usecase"

	"github.com/jeje-gab/resocializationV2/backend/pkg/config"
	"github.com/jeje-gab/resocializationV2/backend/pkg/db"
	"github.com/jeje-gab/resocializationV2/backend/pkg/jwt"
)

func main() {
	loadEnv()

	cfg := config.Load()

	// ctx raiz que cancela com SIGINT/SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// DB
	pool, err := db.NewPGXPool(ctx, cfg.DB.DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	log.Printf("[BOOT] usando DB_DSN: %s", maskDSN(cfg.DB.DSN))

	// JWT
	jwtSvc := jwt.NewJWT(cfg.JWT.Secret, 12*time.Hour)

	// Repos / Usecases
	authRepo := authRepository.NewRepository(pool)
	userRepo := userRepository.NewRepository(pool)
	inmatesRepo := inmatesRepository.NewRepository(pool)
	ufsRepo := ufsRepository.NewRepository(pool)
	citiesRepo := citiesRepository.NewRepository(pool)
	matchRepo := matchRepository.NewRepository(pool)

	authUC := authUseCase.NewUseCase(authRepo, jwtSvc)
	userUC := userUseCase.NewUseCase(userRepo)
	inmatesUC := inmatesUseCase.NewUseCase(inmatesRepo)
	ufsUC := ufsUseCase.NewUseCase(ufsRepo)
	citiesUC := citiesUseCase.NewUseCase(citiesRepo)
	matchUC := matchUseCase.NewUseCase(matchRepo)

	// HTTP
	e := echo.New()
	e.HideBanner = true
	// Echo com NewHTTPError(code, "texto") serializa o corpo como JSON string; o front espera objeto { "error": "..." }.
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}
		he, ok := err.(*echo.HTTPError)
		if !ok {
			c.Logger().Error(err)
			_ = c.JSON(http.StatusInternalServerError, map[string]string{"error": "erro interno"})
			return
		}
		code := he.Code
		if code == 0 {
			code = http.StatusInternalServerError
		}
		switch m := he.Message.(type) {
		case string:
			_ = c.JSON(code, map[string]string{"error": m})
		case echo.Map:
			_ = c.JSON(code, map[string]interface{}(m))
		case map[string]interface{}:
			_ = c.JSON(code, m)
		case map[string]string:
			_ = c.JSON(code, m)
		default:
			if m == nil {
				_ = c.JSON(code, map[string]string{"error": http.StatusText(code)})
			} else {
				_ = c.JSON(code, map[string]string{"error": "requisição inválida"})
			}
		}
	}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS: localhost sempre; produção via FRONTEND_URL (ex.: https://reabtat-production.up.railway.app)
	allowedOrigins := corsAllowedOrigins(cfg.FrontendURL)
	if cfg.Env == "production" {
		log.Printf("[CORS] origens permitidas: %v", allowedOrigins)
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions, http.MethodPatch},
		AllowHeaders:     []string{"Authorization", "Content-Type", "Accept", "Origin", "X-Requested-With"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		MaxAge:           3600,
	}))

	// Health check endpoints
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status":  "ok",
			"service": "ressocializacao-advocacia-api",
			"version": "1.1.0",
		})
	})
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "healthy"})
	})

	// API Routes
	api := e.Group("/api")
	authHttp.RegisterAuthRoutes(api, authUC, jwtSvc.Middleware())
	userHttp.RegisterUserRoutes(api, userUC, jwtSvc.RequireRole("admin"))                  // apenas admin
	inmatesHttp.RegisterInmatesRoutes(api, inmatesUC, jwtSvc.RequireRole("user", "admin")) // user e admin
	ufsHttp.RegisterUFRoutes(api, ufsUC)
	citiesHttp.RegisterCityRoutes(api, citiesUC)
	matchHttp.RegisterMatchRoutes(api, matchUC, jwtSvc.RequireRole("user", "admin")) // user e admin

	// Start
	go func() {
		if err := e.Start(":" + cfg.HTTP.Port); err != nil && err != http.ErrServerClosed {
			log.Printf("server error: %v", err)
		}
	}()
	log.Printf("[BOOT] server started on port %s", cfg.HTTP.Port)

	// espera sinal
	<-ctx.Done()

	// shutdown gracioso
	shCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := e.Shutdown(shCtx); err != nil {
		log.Printf("shutdown error: %v", err)
	}
}

func corsAllowedOrigins(frontendURL string) []string {
	allowed := []string{
		"http://localhost:5173",
		"http://127.0.0.1:5173",
		"http://localhost:3000",
		"http://127.0.0.1:3000",
		"http://localhost:5174",
		"http://127.0.0.1:5174",
	}
	u := strings.TrimRight(strings.TrimSpace(frontendURL), "/")
	if u != "" && u != "http://localhost:5173" {
		allowed = append(allowed, u)
	}
	return allowed
}

func maskDSN(dsn string) string {
	i := strings.Index(dsn, "://")
	if i == -1 {
		return dsn
	}
	at := strings.Index(dsn[i+3:], "@")
	col := strings.Index(dsn[i+3:], ":")
	if at == -1 || col == -1 || col > at {
		return dsn
	}
	prefix := dsn[:i+3]
	rest := dsn[i+3:]
	return prefix + rest[:col] + ":****" + rest[at:]
}

func loadEnv() {
	// caminhos possíveis, ajuste conforme sua estrutura
	paths := []string{
		".env",                      // se você rodar no mesmo diretório do .env
		"dist/.env",                 // se o .env estiver em ./dist/.env
		filepath.Join("..", ".env"), // se você estiver dentro de cmd/server e o .env estiver em ../.env
		filepath.Join("..", "dist", ".env"),
		filepath.Join("..", "..", ".env"),
	}

	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			// arquivo existe, tenta carregar
			if err := godotenv.Load(p); err != nil {
				log.Printf("⚠️ Falha ao carregar .env em %s: %v", p, err)
				continue
			}
			log.Printf("✅ .env carregado de: %s", p)
			return
		}
	}

	log.Println("⚠️ Nenhum .env encontrado, usando apenas variáveis de ambiente do sistema")
}
