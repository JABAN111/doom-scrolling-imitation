FROM python:3.13-slim

# Устанавливаем Java
RUN apt-get update && apt-get install -y openjdk-17-jdk

WORKDIR /spark-app

COPY requirements.txt .

RUN pip install -r requirements.txt

COPY . .

ENV JAVA_HOME=/usr/lib/jvm/java-17-openjdk-amd64
ENV PATH="$JAVA_HOME/bin:$PATH"

CMD ["fastapi","dev", "api.py"]
