from sqlalchemy import Column, Integer, String, Date
from sqlalchemy.sql import func
from backend.app.config.database.banco import Base

class Usuario(Base):
    __tablename__ = "usuario"
    
    id = Column(Integer, primary_key=True, index=True)
    nome = Column(String(255), nullable=False)
    email = Column(String(64), nullable=False, unique=True)
    senha = Column(String(64), nullable=False)  
    data_nascimento = Column(Date, nullable=False)