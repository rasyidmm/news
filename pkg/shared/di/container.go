package di

import "github.com/sarulabs/di"

// MetaData :
type MetaData struct {
	ClientName string
	ClientIP   string
	UUID       string
	ActivityID string
}

// Container :
type Container struct {
	ctn di.Container
}

func NewContainer() *Container {
	builder, _ := di.NewBuilder()
	_ = builder.Add([]di.Def{}...)
	return &Container{
		ctn: builder.Build(),
	}

}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}
