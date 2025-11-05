package repository

import (
	"context"
	"ensina-renda/config/database"
	"ensina-renda/domain/model"

	"github.com/google/uuid"
)

type AulaRepository struct {
}

func NewAulaRepository() *AulaRepository {
	return &AulaRepository{}
}

func (r *AulaRepository) CompletarAula(ctx context.Context, id_aula int, id_usuario uuid.UUID) error {
	err := database.GetDB(ctx).Table("usuario_aula").
		Where("id_aula = ? ", id_aula).
		Where("id_usuario = ?", id_usuario).
		Updates(map[string]interface{}{"concluido": true}).
		Error
	if err != nil {
		return err
	}

	return nil
}

func (r *AulaRepository) GetUsuarioAula(ctx context.Context, id_aula int, id_usuario uuid.UUID) (*model.UsuarioAula, error) {
	var usuarioAula *model.UsuarioAula

	err := database.GetDB(ctx).
		Table("usuario_aula").
		Where("id_aula  = ?", id_aula).
		Where("id_usuario = ?", id_usuario).
		Find(&usuarioAula).Error

	if err != nil {
		return nil, err
	}

	return usuarioAula, nil
}

func (r *AulaRepository) ListarUsuarioModuloAulas(ctx context.Context, id_usuario uuid.UUID) ([]*model.UsuarioModuloAula, error) {
	var moduloAulas []*model.UsuarioModuloAula

	err := database.GetDB(ctx).
		Raw(`
				SELECT
				ma.id_modulo,
				ua.id_usuario,
				ua.id_aula,
				ua.concluido AS aula_concluida,
				um.concluido AS modulo_concluido
					FROM
						usuario_aula AS ua
					JOIN
						modulo_aula AS ma ON ua.id_aula = ma.id_aula
					JOIN
						usuario_modulo AS um ON ua.id_usuario = um.id_usuario AND ma.id_modulo = um.id_modulo
					WHERE
						ua.id_usuario = ?
					ORDER BY
						ma.id_modulo, ua.id_aula;
`, id_usuario).
		Find(&moduloAulas).Error
	if err != nil {
		return nil, err
	}

	return moduloAulas, nil
}

func (r *AulaRepository) ListarAulas(ctx context.Context) ([]*model.Aula, error) {
	var aulas []*model.Aula

	err := database.GetDB(ctx).Table("aula").Find(&aulas).Error
	if err != nil {
		return nil, err
	}

	return aulas, nil
}
