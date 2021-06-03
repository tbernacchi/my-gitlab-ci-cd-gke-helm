FROM python:3.7-slim

RUN mkdir -p /app && apt-get update -y && apt-get install curl vim -y
WORKDIR /app/
COPY . /app/

RUN pip3 install -r /app/src/frontend/requirements.txt && chmod a+rwx /app/src/frontend/frontend.py 

EXPOSE 8000

CMD [ "python", "/app/src/frontend/frontend.py" ] 
