FROM python:3.11-slim
ENV PYTHONUNBUFFERED=1
WORKDIR /bot
RUN apt-get update && apt-get install -y --no-install-recommends gcc libffi-dev \
  && rm -rf /var/lib/apt/lists/*
COPY requirements.txt ./
RUN pip install --no-cache-dir --upgrade pip && pip install --no-cache-dir -r requirements.txt
COPY . .
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8080", "--access-log", "--log-level", "info"]