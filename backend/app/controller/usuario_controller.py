from typing import Optional
from sqlalchemy.orm import Session
from app.repositories.usuario_repository import UsuarioRepository
from app.schemas.usuario_schema import UsuarioCreate, UsuarioUpdate, UsuarioResponse
from fastapi import HTTPException

class UsuarioController:
    def __init__(self, db: Session):
        self.repository = UsuarioRepository(db)

    def criar_usuario(self, usuario_data: UsuarioCreate) -> UsuarioResponse:
        """Controller para criação de usuário"""
        try:
            # Validações de negócio (exemplo)
            if not self._validar_idade_minima(usuario_data.data_nascimento):
                raise HTTPException(
                    status_code=400,
                    detail="Usuário deve ter pelo menos 18 anos"
                )

            # Verifica se email já existe
            if self.repository.obter_usuario_por_email(usuario_data.email):
                raise HTTPException(
                    status_code=400,
                    detail="Email já cadastrado"
                )

            # Chama o repositório
            db_usuario = self.repository.criar_usuario(usuario_data)
            return UsuarioResponse.from_orm(db_usuario)

        except Exception as e:
            # Log de erro pode ser adicionado aqui
            raise HTTPException(
                status_code=500,
                detail=f"Erro ao criar usuário: {str(e)}"
            )

    def _validar_idade_minima(self, data_nascimento) -> bool:
        """Validação de negócio específica"""
        from datetime import date
        idade = (date.today() - data_nascimento).days // 365
        return idade >= 18

    # Outros métodos do controller...
    def obter_usuario(self, usuario_id: int) -> Optional[UsuarioResponse]:
        db_usuario = self.repository.obter_usuario_por_id(usuario_id)
        if not db_usuario:
            raise HTTPException(status_code=404, detail="Usuário não encontrado")
        return UsuarioResponse.from_orm(db_usuario)

    def atualizar_usuario(self, usuario_id: int, usuario_data: UsuarioUpdate) -> UsuarioResponse:
        # Lógica similar com validações
        pass