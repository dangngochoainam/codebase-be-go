package dihelper

import (
	"sync"

	"github.com/sarulabs/di"
)

type DIBuilder func() []di.Def

var (
	buildOnce           sync.Once
	builder             *di.Builder
	container           di.Container
	ConfigsBuilder      DIBuilder
	HelpersBuilder      DIBuilder
	UseCasesBuilder     DIBuilder
	RepositoriesBuilder DIBuilder
	ControllersBuilder  DIBuilder
)

func BuildLibDIContainer() {
	buildOnce.Do(func() {
		builder, _ = di.NewBuilder()
		doBuild()
		container = builder.Build()
	})
}

func GetLibDependency(dependencyName string) interface{} {
	return container.Get(dependencyName)
}

func CleanDependency() error {
	return container.Clean()
}

func doBuild() {
	err := buildConfigs()
	if err != nil {
		panic(err)
	}

	err = buildHelpers()
	if err != nil {
		panic(err)
	}

	err = buildRepository()
	if err != nil {
		panic(err)
	}

	err = buildUseCases()
	if err != nil {
		panic(err)
	}

	err = buildControllers()
	if err != nil {
		panic(err)
	}
}

func buildConfigs() error {
	defs := []di.Def{}
	if ConfigsBuilder == nil {
		ConfigsBuilder = defaultBuilder
	}
	defs = ConfigsBuilder()
	err := builder.Add(defs...)
	if err != nil {
		return err
	}
	return nil
}

func buildHelpers() error {
	defs := []di.Def{}
	if HelpersBuilder == nil {
		HelpersBuilder = defaultBuilder
	}
	defs = HelpersBuilder()
	err := builder.Add(defs...)
	if err != nil {
		return err
	}
	return nil
}

func buildControllers() error {
	defs := []di.Def{}
	if ControllersBuilder == nil {
		ControllersBuilder = defaultBuilder
	}
	defs = ControllersBuilder()
	err := builder.Add(defs...)
	if err != nil {
		return err
	}
	return nil
}

func buildUseCases() error {
	defs := []di.Def{}
	if UseCasesBuilder == nil {
		UseCasesBuilder = defaultBuilder
	}
	defs = UseCasesBuilder()
	err := builder.Add(defs...)
	if err != nil {
		return err
	}
	return nil
}

func buildRepository() error {
	defs := []di.Def{}
	if RepositoriesBuilder == nil {
		RepositoriesBuilder = defaultBuilder
	}
	defs = RepositoriesBuilder()
	err := builder.Add(defs...)
	if err != nil {
		return err
	}
	return nil
}

func defaultBuilder() []di.Def {
	return []di.Def{}
}
