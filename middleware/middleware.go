package middleware
import (
    "github.com/gofiber/fiber"
    "github.com/gofiber/basicauth"
    "github.com/tdwiid/logrocket/config"
)

// AuthReq middleware
func AuthReq() func(*fiber.Ctx) {
    cfg := basicauth.Config{
        Users: map[string]string{
          config.Config("USERNAME"): config.Config("PASSWORD"),
        },
      }
    err := basicauth.New(cfg);
    return err
}
