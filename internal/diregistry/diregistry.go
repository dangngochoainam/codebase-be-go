package diregistry

import (
	"example/config"
	"example/internal/common/helper/copyhepler"
	"example/internal/common/helper/cronschedulerhelper"
	"example/internal/common/helper/dihelper"
	"example/internal/common/helper/logwriterhelper"
	"example/internal/common/helper/sqlormhelper"
	"example/internal/common/helper/validatehelper"
	"example/internal/controller"
	"example/internal/repository"
	"example/internal/usecase"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/sarulabs/di"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

const (
	ConfigDIName                string = "Config"
	ValidateDIName              string = "Validate"
	ModelConverterDIName        string = "ModelConverter"
	CronSchedulerDIName         string = "CronScheduler"
	SqlGormLogHelperDIName      string = "SqlGormLogHelper"
	SqlGormPostgresHelperDIName string = "SqlGormPostgresHelper"
	LogsFileWriterHelperDIName  string = "LogsFileWriterHelper"
	UserRepositoryDIName        string = "UserRepository"
	ProductRepositoryDIName     string = "ProductRepository"
	UserUseCaseDIName           string = "UserUseCase"
	ExampleUseCaseDIName        string = "ExampleUseCase"
	UserControllerDIName        string = "UserController"
	ExampleControllerDIName     string = "ExampleController"
)

func BuildDIContainer() {
	initBuilder()
	dihelper.BuildLibDIContainer()
}

func GetDependency(name string) interface{} {
	return dihelper.GetLibDependency(name)
}

func initBuilder() {
	dihelper.ConfigsBuilder = func() []di.Def {
		arr := []di.Def{}
		arr = append(arr, di.Def{
			Name:  ConfigDIName,
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				cfg, err := config.LoadEnvironment()
				if err != nil {
					return nil, err
				}
				return cfg, nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		})
		return arr
	}

	dihelper.HelpersBuilder = func() []di.Def {
		arr := []di.Def{}
		arr = append(arr, di.Def{
			Name:  ValidateDIName,
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return validatehelper.NewValidate(), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		}, di.Def{
			Name:  ModelConverterDIName,
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				modelConverter := copyhepler.NewModelConverter()
				return modelConverter, nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		}, di.Def{
			Name:  SqlGormLogHelperDIName,
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				cfg := ctn.Get(ConfigDIName).(*config.Config)
				gormConfig := &gorm.Config{}
				if cfg.DatabaseLog.LoggingEnabled {
					var logWriter io.Writer = os.Stdout
					if cfg.DatabaseLog.UseLoggingDb {
						sqlGormLogDb := ctn.Get(SqlGormLogHelperDIName).(sqlormhelper.SqlGormDatabase)
						logWriter = logwriterhelper.NewSqlPostgresAuditLogWriter(sqlGormLogDb)
					} else if cfg.DatabaseLog.UseLoggingFile {
						logWriter = ctn.Get(LogsFileWriterHelperDIName).(io.Writer)
					}
					gormLogger := gormLogger.New(
						log.New(logWriter, "\r\n", log.LstdFlags), // io writer
						gormLogger.Config{
							SlowThreshold:             time.Millisecond * 200, // Slow SQL threshold
							LogLevel:                  gormLogger.Warn,        // Log level
							IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
							ParameterizedQueries:      true,                   // Don't include params in the SQL log
							Colorful:                  false,                  // Disable color
						},
					)
					gormConfig.Logger = gormLogger
				}
				// password := confighelper.DescrytEnv(cfg.EncryptionKey, cfg.DatabaseLog.Password)
				// if password == "" {
				// 	loghelper.Logger.Panic("password cannot empty")
				// }
				return sqlormhelper.NewGormPostgresqlDB(
					&sqlormhelper.GormConnectionOptions{
						Host:               cfg.DatabaseLog.Host,
						Port:               int(cfg.DatabaseLog.Port),
						Username:           cfg.DatabaseLog.Username,
						Password:           cfg.DatabaseLog.Password,
						Database:           cfg.DatabaseLog.Database,
						Schema:             cfg.DatabaseLog.Schema,
						GormConfig:         gormConfig,
						UseTls:             cfg.DatabaseLog.UseTls,
						TlsMode:            cfg.DatabaseLog.TlsMode,
						TlsRootCACertFile:  cfg.DatabaseLog.TlsRootCACertFile,
						TlsKeyFile:         cfg.DatabaseLog.TlsKeyFile,
						TlsCertFile:        cfg.DatabaseLog.TlsCertFile,
						InsecureSkipVerify: cfg.DatabaseLog.InsecureSkipVerify,
						MaxOpenConns:       cfg.DatabaseLog.MaxOpenConns,
						// Timezone:           string(timehelper.Timezone_UTC),
					},
				), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		}, di.Def{
			Name:  SqlGormPostgresHelperDIName,
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				cfg := ctn.Get(ConfigDIName).(*config.Config)

				gormConfig := &gorm.Config{}
				if cfg.DatabasePostgres.LoggingEnabled {
					var logWriter io.Writer = os.Stdout
					if cfg.DatabasePostgres.UseLoggingDb {
						sqlGormLogDb := ctn.Get(SqlGormLogHelperDIName).(sqlormhelper.SqlGormDatabase)
						logWriter = logwriterhelper.NewSqlPostgresAuditLogWriter(sqlGormLogDb)
					} else if cfg.DatabasePostgres.UseLoggingFile {
						logWriter = ctn.Get(LogsFileWriterHelperDIName).(io.Writer)
					}
					gormLogger := gormLogger.New(
						log.New(logWriter, "\r\n", log.LstdFlags), // io writer
						gormLogger.Config{
							SlowThreshold:             time.Millisecond * 200, // Slow SQL threshold
							LogLevel:                  gormLogger.Warn,        // Log level
							IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
							ParameterizedQueries:      true,                   // Don't include params in the SQL log
							Colorful:                  false,                  // Disable color
						},
					)
					gormConfig.Logger = gormLogger
				}
				return sqlormhelper.NewGormPostgresqlDB(&sqlormhelper.GormConnectionOptions{
					Host:               cfg.DatabasePostgres.Host,
					Port:               int(cfg.DatabasePostgres.Port),
					Username:           cfg.DatabasePostgres.Username,
					Password:           cfg.DatabasePostgres.Password,
					Database:           cfg.DatabasePostgres.Database,
					Schema:             cfg.DatabasePostgres.Schema,
					GormConfig:         gormConfig,
					UseTls:             cfg.DatabasePostgres.UseTls,
					TlsMode:            cfg.DatabasePostgres.TlsMode,
					TlsRootCACertFile:  cfg.DatabasePostgres.TlsRootCACertFile,
					TlsKeyFile:         cfg.DatabasePostgres.TlsKeyFile,
					TlsCertFile:        cfg.DatabasePostgres.TlsCertFile,
					InsecureSkipVerify: cfg.DatabasePostgres.InsecureSkipVerify,
					MaxOpenConns:       cfg.DatabasePostgres.MaxOpenConns,
				}), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		}, di.Def{
			Name:  LogsFileWriterHelperDIName,
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return logwriterhelper.NewRotatingFileWriter(), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		})
		return arr
	}

	dihelper.RepositoriesBuilder = func() []di.Def {
		arr := []di.Def{}
		arr = append(arr, di.Def{
			Name:  UserRepositoryDIName,
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				cfg := ctn.Get(ConfigDIName).(*config.Config)
				postgresOrmDb := ctn.Get(SqlGormPostgresHelperDIName).(sqlormhelper.SqlGormDatabase)
				return repository.NewUserRepository(cfg, postgresOrmDb), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		}, di.Def{
			Name:  ProductRepositoryDIName,
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				cfg := ctn.Get(ConfigDIName).(*config.Config)
				return repository.NewProductRepository(cfg), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		})
		return arr
	}

	dihelper.UseCasesBuilder = func() []di.Def {
		arr := []di.Def{}
		arr = append(arr, di.Def{
			Name:  UserUseCaseDIName,
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				userRepository := ctn.Get(UserRepositoryDIName).(repository.UserRepository)
				modelConverter := ctn.Get(ModelConverterDIName).(copyhepler.ModelConverter)
				return usecase.NewUserUseCase(userRepository, modelConverter), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		}, di.Def{
			Name:  ExampleUseCaseDIName,
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				mutex := &sync.Mutex{}
				return usecase.NewExampleUseCase(mutex), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		})
		return arr
	}

	dihelper.ControllersBuilder = func() []di.Def {
		arr := []di.Def{}
		arr = append(arr, di.Def{
			Name:  UserControllerDIName,
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				userUseCase := ctn.Get(UserUseCaseDIName).(usecase.UserUseCase)
				modelConverter := ctn.Get(ModelConverterDIName).(copyhepler.ModelConverter)
				return controller.NewUserController(userUseCase, modelConverter), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		}, di.Def{
			Name:  ExampleControllerDIName,
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				exampleUseCase := ctn.Get(ExampleUseCaseDIName).(usecase.ExampleUseCase)
				return controller.NewExampleController(exampleUseCase), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		})
		return arr
	}

	dihelper.CronSchedulerBuilder = func() []di.Def {
		arr := []di.Def{}
		arr = append(arr, di.Def{
			Name:  CronSchedulerDIName,
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				exampleUseCase := ctn.Get(ExampleUseCaseDIName).(usecase.ExampleUseCase)
				jobs := []*cronschedulerhelper.Job{
					{
						Spec: "@every 20h00m10s",
						Cmd:  exampleUseCase.CronScheduler,
					},
				}
				return cronschedulerhelper.NewCronSchedulerHelper(jobs), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		})
		return arr
	}
}
