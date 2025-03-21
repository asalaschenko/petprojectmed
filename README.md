📈 Web-приложение на основе fiber: запись пациентов к докторам в клинике.
=

📝 ***Мой домашний проект, находится в процессе разработки.***

Технологии:

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens)

Endpoints:
-
localhost:3000/

    GET     /doctor/list=all
    GET     /doctor/list=filter&specializations={:values}&datesOfBirth={:values}
    GET     /doctor/{:id}
    PUT     /doctor/:id
    DELETE  /doctor/:id
    POST    /doctor

    GET     /patient/list=all
    GET     /patient/list=filter&phoneNumbers={:values}&datesOfBirth={:values}
    GET     /patient/{:id}
    PUT     /patient/:id
    DELETE  /patient/:id
    POST    /patient

    GET     /schedule/list=all
    GET     /schedule/list=filter&doctorID={:values}&patientID={:values}&dateAppointment={:values}
    DELETE  /schedule/:id
    POST    /schedule


Примеры запросов POST.
-
BODY RAW:
-
    /doctor
    {
        "name": "Дмитрий",
        "family": "Петров",
        "dateOfBirth": "1995-01-12",
        "specialization": "хирург",
        "cabinet": 50
    }
    
    /patient
    {
        "name": "Дмитрий",
        "family": "Петров",
        "dateOfBirth": "1995-01-12",
        "gender": "мужской",
        "phoneNumber": "79039558644"
    }

    /schedule
    {
    "doctorID" : 16,
    "patientID": 10,
    "date": "2025-03-06",
    "time": "13:00"
    }

_Примечание:_
* _для запуска необходима переменная окружения PORT_GOLANG со значением номера порта, на котором будет запускаться приложение_
