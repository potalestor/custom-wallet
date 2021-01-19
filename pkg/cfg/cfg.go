package cfg

import (
	"fmt"
	"os"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

const (
	formatConnectionDB = `host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`
	formatConnection   = `host=%s port=%d user=%s password=%s sslmode=disable`

	formatURI = `postgres://%s:%s@%s:%d/%s?sslmode=disable`
)

type Config struct {
	Database  Database
	Migration Migration
	Logger    Logger
	Web       Web
}

func (c Config) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Database),
		validation.Field(&c.Migration),
	)
}

type Web struct {
	Adddress string
}

type Logger struct {
	File bool
}

type Database struct {
	User     string
	Password string
	DB       string
	Host     string
	Port     int
}

func (d Database) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.User, validation.Required),
		validation.Field(&d.Password, validation.Required),
		validation.Field(&d.Host, validation.Required, is.Host),
		validation.Field(&d.Port, validation.Required, validation.NotNil),
	)
}

func (d Database) URI() string {
	return fmt.Sprintf(
		formatURI,
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.DB,
	)
}

func (d Database) String() string {
	if d.DB == "" {
		return fmt.Sprintf(formatConnection,
			d.Host,
			d.Port,
			d.User,
			d.Password,
		)
	}
	return fmt.Sprintf(formatConnectionDB,
		d.Host,
		d.Port,
		d.User,
		d.Password,
		d.DB,
	)
}

type Migration struct {
	Enabled bool
	Path    string
}

func (m Migration) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Path, validation.When(m.Enabled,
			validation.Required, validation.By(
				func(value interface{}) error {
					path, _ := value.(string)

					if _, err := os.Stat(path); os.IsNotExist(err) {
						return err
					}

					return nil
				}))),
	)
}
