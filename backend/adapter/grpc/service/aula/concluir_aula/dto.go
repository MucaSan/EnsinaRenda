package concluir_aula

import (
	"context"
	pb "ensina-renda/adapter/grpc/pb"
	"ensina-renda/domain/model"
	"errors"

	"github.com/google/uuid"
)

type UsuarioAulaConverter struct {
	base any
}

func NewUsuarioAulaConverter(base any) *UsuarioAulaConverter {
	return &UsuarioAulaConverter{
		base: base,
	}
}

func (uc *UsuarioAulaConverter) ToDomain(ctx context.Context) (*model.UsuarioAula, error) {
	concluirAulaRequest, ok := uc.base.(*pb.ConcluirAulaRequest)
	if !ok {
		return nil, errors.New("nao foi possivel converter base para cadastrar_aluno_request")
	}

	idUsuario, err := uuid.Parse(concluirAulaRequest.IdUsuario)
	if err != nil {
		return nil, errors.New("nao foi possivel converter o ID do usuario para UUID v4")
	}

	modelUsuarioAula := &model.UsuarioAula{
		IDUsuario: idUsuario,
		IDAula:    int(concluirAulaRequest.IdAula),
		Concluido: false,
	}

	if err = modelUsuarioAula.IsValid(); err != nil {
		return nil, err
	}

	return modelUsuarioAula, nil

}
