package backend

import (
	"context"
	"embed"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"reflect"

	"github.com/99designs/keyring"
	"github.com/cjhuaxin/ElasticDesktopManager/backend/connection"
	"github.com/cjhuaxin/ElasticDesktopManager/backend/models"
	"github.com/cjhuaxin/ElasticDesktopManager/backend/resource"
	"github.com/thanhpk/randstr"
	"github.com/wailsapp/wails/v2/pkg/options"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Bind interface {
	Init(ctx *models.EdmContext) error
}

// App struct
type App struct {
	ctx *models.EdmContext
}

func WailsInit(assets embed.FS) *options.App {
	// Create an instance of the app structure
	app := &App{}
	bindList := make([]interface{}, 0)
	extraBindList := extraBinds()
	bindList = allBinds(app, extraBindList)

	return &options.App{
		Title:            resource.AppTitle,
		Width:            1024,
		Height:           768,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			err := app.onStart(ctx, extraBindList...)
			if err != nil {
				os.Exit(-1)
			}
		},
		Bind: bindList,
	}
}

func (a *App) onStart(ctx context.Context, binds ...Bind) error {
	a.ctx = models.NewContext(ctx)
	//init directory for app
	err := a.initDirectoryStructure()
	if err != nil {
		return err
	}
	//init log
	a.initLog()
	for _, bind := range binds {
		err = bind.Init(a.ctx)
		if err != nil {
			return err
		}
	}
	//init keyring
	err = a.initKeyring()
	if err != nil {
		return err
	}

	return nil
}

func allBinds(app *App, extraBinds []Bind) []interface{} {
	all := make([]interface{}, 0, len(extraBinds)+1)
	all = append(all, app)

	for _, bind := range extraBinds {
		all = append(all, bind)
	}

	return all
}

func extraBinds() []Bind {
	return []Bind{
		connection.NewConnection(),
	}
}

func (a *App) initDirectoryStructure() error {
	u, err := user.Current()
	if err != nil {
		return err
	}
	homeDir := filepath.Join(u.HomeDir, ".edm")
	fmt.Printf("user home is: %s\n", homeDir)
	a.ctx.SetPath(&models.Paths{
		HomeDir: homeDir, // Home directory of the user
		ConfDir: filepath.Join(homeDir, "conf"),
		DbDir:   filepath.Join(homeDir, "db"),
		LogDir:  filepath.Join(homeDir, "log"),
		TmpDir:  filepath.Join(homeDir, "tmp"),
	})

	//create folder
	err = a.createFolderIfNotExists()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) createFolderIfNotExists() error {
	pathValue := reflect.ValueOf(*a.ctx.GetPath())
	for i := 0; i < pathValue.NumField(); i++ {
		path := pathValue.Field(i).String()
		if _, err := os.Stat(path); os.IsNotExist(err) {
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

func (a *App) initLog() {
	infoLogger := &lumberjack.Logger{
		Filename: filepath.Join(a.ctx.GetPath().LogDir, "edm.log"),
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	infoCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(infoLogger),
		zapcore.InfoLevel,
	)
	logger := zap.New(infoCore, zap.AddStacktrace(zapcore.FatalLevel))
	a.ctx.SetLog(logger.Sugar())
	defer a.ctx.Log().Sync()
}

func (a *App) initKeyring() error {
	ring, err := keyring.Open(keyring.Config{
		ServiceName: resource.AppTitle,
	})
	if err != nil {
		return err
	}
	_, err = ring.Get(resource.EncryptAESKey)
	if err != nil {
		if err == keyring.ErrKeyNotFound {
			// init aes key if not exists
			token := randstr.String(32)
			ring.Set(keyring.Item{
				Key:  resource.EncryptAESKey,
				Data: []byte(token),
			})
			return nil
		}
		return err
	}
	a.lo

	return nil
}
