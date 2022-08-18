package services

import (
	"context"

	"github.com/Permify/permify/internal/commands"
	"github.com/Permify/permify/internal/entities"
	"github.com/Permify/permify/pkg/dsl/schema"
	"github.com/Permify/permify/pkg/tuple"
)

// IPermissionService -
type IPermissionService interface {
	Check(ctx context.Context, subject tuple.Subject, action string, entity tuple.Entity, version string, d int32) (response commands.CheckResponse, err error)
	Expand(ctx context.Context, entity tuple.Entity, action string, version string) (response commands.ExpandResponse, err error)
}

// ISchemaService -
type ISchemaService interface {
	All(ctx context.Context, version string) (sch schema.Schema, err error)
	Read(ctx context.Context, name string, version string) (sch schema.Schema, err error)
	Write(ctx context.Context, configs entities.EntityConfigs) (version string, err error)
}