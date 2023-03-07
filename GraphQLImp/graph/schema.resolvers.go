package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.25

import (
	model1 "GraphQL/graph/model"
	"GraphQLImp/graph/model"
	"context"
	"fmt"
)

// CreatePaciente is the resolver for the createPaciente field.
func (r *mutationResolver) CreatePaciente(ctx context.Context, input model1.NewPaciente) (*model1.Paciente, error) {
	panic(fmt.Errorf("not implemented: CreatePaciente - createPaciente"))
}

// UpdatePaciente is the resolver for the updatePaciente field.
func (r *mutationResolver) UpdatePaciente(ctx context.Context, id string, nombre string, apellido string, paga int, vsemana int) (*model1.Paciente, error) {
	panic(fmt.Errorf("not implemented: UpdatePaciente - updatePaciente"))
}

// DeletePaciente is the resolver for the deletePaciente field.
func (r *mutationResolver) DeletePaciente(ctx context.Context, id string) (*model1.Paciente, error) {
	panic(fmt.Errorf("not implemented: DeletePaciente - deletePaciente"))
}

// Pacientes is the resolver for the pacientes field.
func (r *queryResolver) Pacientes(ctx context.Context) ([]*model1.Paciente, error) {
	panic(fmt.Errorf("not implemented: Pacientes - pacientes"))
}

// Paciente is the resolver for the paciente field.
func (r *queryResolver) Paciente(ctx context.Context, id string) (*model1.Paciente, error) {
	panic(fmt.Errorf("not implemented: Paciente - paciente"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }


func (r *Resolver) Pacientes() []*model.Paciente {
	return r.pacientes
}
func (r *Resolver) CreatePaciente(input model.NewPaciente) *model.Paciente {
	paciente := &model.Paciente{
		ID:       input.ID,
		Nombre:   input.Nombre,
		Apellido: input.Apellido,
		Paga:     input.Paga,
		Vsemana:  input.Vsemana,
	}
	r.pacientes = append(r.pacientes, paciente)
	return paciente
}



type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.