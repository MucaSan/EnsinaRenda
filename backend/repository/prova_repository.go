package repository

import (
	"context"
	"ensina-renda/config/auth"
	"ensina-renda/config/database"
	"ensina-renda/domain/model"
)

type ProvaRepository struct {
}

func NewProvaRepository() *ProvaRepository {
	return &ProvaRepository{}
}

func (r *ProvaRepository) GetProvaBase(ctx context.Context, idModulo string) (*model.ProvaBase, error) {
	var provaBase *model.ProvaBase

	err := database.GetDB(ctx).
		Table("provas_base").
		Where("id_modulo  = ?", idModulo).
		Find(&provaBase).Error

	if err != nil {
		return nil, err
	}

	return provaBase, nil
}

func (r *ProvaRepository) SalvarProva(ctx context.Context, provaBase *model.ProvaUsuario) error {
	err := database.GetDB(ctx).
		Table("provas_usuario").Where("").Save(provaBase).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *ProvaRepository) GetProvaUsuario(ctx context.Context, idModulo string) (*model.ProvaUsuario, error) {
	var provaUsuario *model.ProvaUsuario

	err := database.GetDB(ctx).
		Table("provas_usuario").
		Where("id_usuario = ?", auth.GetUserUuidPeloContext(ctx)).
		Where("id_modulo = ?", idModulo).
		First(&provaUsuario).Error
	if err != nil {
		return nil, err
	}

	return provaUsuario, nil
}

func (r *ProvaRepository) SalvarCorrecaoProva(ctx context.Context, correcaoProva *model.CorrecaoProva) error {
	err := database.GetDB(ctx).
		Table("correcao_prova").Save(correcaoProva).Error
	if err != nil {
		return err
	}

	return nil
}
