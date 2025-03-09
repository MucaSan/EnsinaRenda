import os
import httpx
from fastapi import FastAPI, HTTPException
from dotenv import load_dotenv

# Carregar variáveis de ambiente do .env
load_dotenv()

app = FastAPI()

# Configurações da API OpenAI
OPENAI_API_KEY = os.getenv("OPENAI_API_KEY")
OPENAI_API_URL = "https://api.openai.com/v1/chat/completions"

if not OPENAI_API_KEY:
    raise ValueError("A variável OPENAI_API_KEY não está definida no ambiente!")

# Definir o prompt
PROMPT = """
O que é renda fixa e como ela funciona?
Quais são os principais tipos de investimentos em renda fixa?
Quais os principais fatores que influenciam a rentabilidade da renda fixa?
"""

async def fetch_openai_response(prompt: str):
    """Função para obter resposta da API OpenAI"""
    headers = {
        "Authorization": f"Bearer {OPENAI_API_KEY}",
        "Content-Type": "application/json"
    }
    data = {
        "model": "gpt-3.5-turbo",
        "messages": [
            {"role": "system", "content": "Você é um assistente financeiro."},
            {"role": "user", "content": prompt}
        ],
        "temperature": 0.7
    }

    async with httpx.AsyncClient(timeout=30.0) as client:
        try:
            response = await client.post(OPENAI_API_URL, headers=headers, json=data)
            response.raise_for_status()  # Levanta exceções para erros HTTP
            return response.json()
        except httpx.HTTPStatusError as exc:
            raise HTTPException(status_code=exc.response.status_code, detail=exc.response.text)
        except httpx.RequestError as exc:
            raise HTTPException(status_code=500, detail=f"Erro de conexão: {str(exc)}")

@app.get("/chatgpt-response")
async def get_chatgpt_response():
    """Endpoint para obter resposta do ChatGPT"""
    return await fetch_openai_response(PROMPT)
