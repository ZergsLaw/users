// +build integration

package repo_test

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"testing"
	"time"

	zergrepo "github.com/ZergsLaw/zerg-repo"
	"github.com/ZergsLaw/zerg-repo/zergrepo/fs"
	"github.com/jmoiron/sqlx"
	"github.com/zergslaw/boilerplate/internal/app"
	"github.com/zergslaw/boilerplate/internal/repo"
	"go.uber.org/zap"
)

var (
	Repo     *repo.Repo
	truncate func() error

	timeoutConnect = time.Second * 1000
)

func TestMain(m *testing.M) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutConnect)
	defer cancel()

	dbConn, err := zergrepo.Connect(ctx, "postgres")
	if err != nil {
		log.Fatal(fmt.Errorf("connect UserRepo: %w", err))
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(fmt.Errorf("connect zap: %w", err))
	}

	zp := repo.Connect(dbConn, logger.Named("test").Sugar(), "test")

	const dir = "../../migrate/"
	f := fs.New()
	migrates, err := f.Walk(dir)
	if err != nil {
		log.Fatal(fmt.Errorf("walk migrate: %w", err))
	}

	regMigrates := make([]zergrepo.Migrate, len(migrates))
	for i := range migrates {
		regMigrates[i] = zergrepo.Migrate{
			Version: migrates[i].Version,
			Up:      zergrepo.Query(migrates[i].Query.Up),
			Down:    zergrepo.Query(migrates[i].Query.Down),
		}
	}

	err = zergrepo.RegisterMetric(regMigrates...)
	if err != nil {
		log.Fatal(fmt.Errorf("register migrate: %w", err))
	}

	resetDB := func() {
		err := zp.Reset(ctx)
		if err != nil {
			log.Fatal(fmt.Errorf("migration reset: %w", err))
		}
	}
	// For convenient cleaning DB.
	resetDB()

	err = zp.Up(ctx)
	if err != nil {
		log.Fatal(fmt.Errorf("up migration: %w", err))
	}
	defer resetDB()

	Repo = repo.New(zp)
	truncate = func() error {
		return zp.Do(func(db *sqlx.DB) error {
			_, err := db.Exec("TRUNCATE users, sessions, notifications, recovery_code RESTART IDENTITY CASCADE")
			return err
		})
	}

	os.Exit(m.Run())
}

var (
	userGenerator = generatorUser()
	ctx           = context.Background()
	ip            = "192.100.10.4"
	origin        = app.Origin{
		IP:        net.ParseIP(ip),
		UserAgent: "UserAgent",
	}
)

func generatorUser() func() app.User {
	x := 0

	return func() app.User {
		x++
		return app.User{
			ID:        app.UserID(x),
			Email:     fmt.Sprintf("email%d@gmail.com", x),
			Name:      fmt.Sprintf("username%d", x),
			PassHash:  []byte("pass"),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	}
}
