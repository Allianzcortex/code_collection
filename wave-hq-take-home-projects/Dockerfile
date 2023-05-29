FROM python:3.6

EXPOSE 5000

WORKDIR /app

MAINTAINER Allianzcortex

COPY init.sql /docker-entrypoint-initdb.d/
COPY . ./

# build dependency 
RUN pip install -U pipenv \
    && pipenv install --deploy --ignore-pipfile \
    && export FLASK_APP=app \
    && mkdir -p static \

CMD ["echo", "Payroll System RUN SUCCESFULLY ! ! ! !  ! ! !"]
CMD ["pipenv", "run", "python", "app.py"]
