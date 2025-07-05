import os
import asyncio
import logging
from http import HTTPStatus

import httpx
import uvicorn
from fastapi import FastAPI
from starlette.responses import JSONResponse
from dotenv import load_dotenv
from pydantic import BaseModel

load_dotenv()

BOT_TOKEN = os.getenv("BOT_TOKEN")
CHAT_ID = os.getenv("CHAT_ID")

chat_endpoint = f'https://api.telegram.org/bot{BOT_TOKEN}/sendMessage'


class UpdationBody(BaseModel):
    updation_type: str
    updation_name: str


def get_app() -> FastAPI:
    app = FastAPI()
    return app

app = get_app()


@app.post("/new_item/")
async def receive_new_item(updation: UpdationBody):
    await send_message(updation)
    return JSONResponse(status_code=HTTPStatus.CREATED, content={'result': 'ok'})


async def send_message(updation: UpdationBody):
    bot_message = f"📥 Received new item for collection {updation.updation_type}: {updation.updation_name}"
    async with httpx.AsyncClient() as client:
        try:
            await client.post(chat_endpoint, data={"chat_id": CHAT_ID, "text": bot_message})
        except Exception as ex:
            logging.error(ex)