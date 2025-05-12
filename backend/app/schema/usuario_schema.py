from pydantic import BaseModel, EmailStr, Field
from datetime import date, datetime
from typing import Optional

class UsuarioBase(BaseModel):
    nome: str = Field(..., max_length=255)
    email: EmailStr = Field(..., max_length=64)
    data_nascimento: date

class UsuarioCreate(UsuarioBase):
    senha_hash: str = Field(..., 
                          min_length=64, 
                          max_length=64,
                          regex="^[a-f0-9]{64}$",  # Valida formato SHA-256
                          description="Hash SHA-256 da senha (64 caracteres hexadecimais)")

class UsuarioResponse(UsuarioBase):
    id: int
    mensagem: str

    class Config:
        orm_mode = True