package app

// import (
// 	"database/sql"
// 	"log"
// 	"os/exec"
// 	"strings"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"

// 	"github.com/stretchr/testify/require"

// 	"github.com/go-park-mail-ru/2022_2_VDonate/internal/config"

// 	"github.com/labstack/echo/v4"
// )

// func TestApp_init(t *testing.T) {
// 	e := echo.New()
// 	cfg := config.New()
// 	err := cfg.Open("../../cmd/api/configs/config_local.yaml")
// 	require.NoError(t, err)

// 	log.Println("Enabling containers with all environment...")

// 	cmd := exec.Command("docker-compose", "-f", "../../test/test.docker-compose.yaml", "up", "-d")

// 	out, err := cmd.CombinedOutput()

// 	log.Println("\n" + string(out))

// 	if err != nil {
// 		exitErr, ok := err.(*exec.ExitError)
// 		if ok {
// 			assert.NoError(t, err, string(exitErr.Stderr))
// 		} else {
// 			log.Fatalf("cmd.Wait: %v", err)
// 		}
// 	}

// 	log.Println("Successfully enabled")

// 	log.Println("Initializing server...")

// 	postgres, err := sql.Open(cfg.DB.Driver, cfg.DB.URL)
// 	assert.NoError(t, err, "Postgres not UP")

// 	log.Println("waiting for postgres to start")
// 	for postgres.Ping() != nil {
// 		log.Println("...")
// 		time.Sleep(time.Second * 2)
// 	}

// 	// Change local config for all services to start
// 	server := New(e, cfg)
// 	assert.NoError(t, server.init())
// 	log.Println("Successfully initialized")

// 	log.Println("Removing all containers...")
// 	cmd = exec.Command("docker", "ps", "-aq")

// 	out, err = cmd.CombinedOutput()

// 	if err != nil {
// 		exitErr, ok := err.(*exec.ExitError)
// 		if ok {
// 			require.NoError(t, err, string(exitErr.Stderr))
// 		} else {
// 			log.Fatalf("cmd.Wait: %v", err)
// 		}
// 	}

// 	containersToRemove := strings.Split(string(out), "\n")
// 	containersToRemove = containersToRemove[:len(containersToRemove)-1]
// 	cmd = exec.Command("docker", append(append([]string{"rm"}, "-f"), containersToRemove...)...)

// 	out, err = cmd.CombinedOutput()
// 	log.Println("\n" + string(out))

// 	if err != nil {
// 		exitErr, ok := err.(*exec.ExitError)
// 		if ok {
// 			require.NoError(t, err, string(exitErr.Stderr))
// 		} else {
// 			log.Fatalf("cmd.Wait: %v", err)
// 		}
// 	}

// 	log.Println("Successfully removed")
// }
