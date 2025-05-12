from fastapi import APIRouter, Depends
from sqlalchemy.orm import Session
from app.controllers.usuario_controller import UsuarioController
from app.schemas.usuario_schema import UsuarioCreate, UsuarioResponse
from app.database import get_db

router = APIRouter(tags=["Usuários"])

@router.post("/usuarios/", response_model=UsuarioResponse)
def criar_usuario(
    usuario: UsuarioCreate,
    db: Session = Depends(get_db)
):
    controller = UsuarioController(db)
    return controller.criar_usuario(usuario)

@router.get("/usuarios/{usuario_id}", response_model=UsuarioResponse)
def obter_usuario(usuario_id: int, db: Session = Depends(get_db)):
    controller = UsuarioController(db)
    return controller.obter_usuario(usuario_id)

# Outras rotas...