package concluir_modulo

import (
	"context"
	pb "ensina-renda/adapter/grpc/pb"
	"ensina-renda/controller/modulo/converter"
	"ensina-renda/domain/model"
	"errors"
	"strconv"

	"github.com/google/uuid"
)

type UsuarioModuloConverter struct {
	base any
}

func NewUsuarioModuloConverter(base any) converter.UsuarioModuloConverterInterface {
	return &UsuarioModuloConverter{
		base: base,
	}
}

func (uc *UsuarioModuloConverter) ToDomain(ctx context.Context) (*model.UsuarioModulo, error) {
	concluirModuloRequest, ok := uc.base.(*pb.ConcluirModuloRequest)
	if !ok {
		return nil, errors.New("nao foi possivel converter base para concluir_aula_request")
	}

	idUsuario, err := uuid.Parse(concluirModuloRequest.IdUsuario)
	if err != nil {
		return nil, errors.New("nao foi possivel converter o ID do usuario para UUID v4")
	}

	idModulo, err := strconv.Atoi(concluirModuloRequest.IdModulo)
	if err != nil {
		return nil, errors.New("nao foi possivel converter o ID do modulo para numero inteiro")
	}

	modelUsuarioModulo := &model.UsuarioModulo{
		IDUsuario: idUsuario,
		IDModulo:  idModulo,
		Concluido: false,
	}

	if err = modelUsuarioModulo.IsValid(); err != nil {
		return nil, err
	}

	return modelUsuarioModulo, nil
}
