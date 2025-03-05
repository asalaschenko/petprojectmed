### Web-приложение на основе fiber: запись пацентов к докторам в клинике.

***Мой домашний проект, находится в процессе разработки.***

endpoints:
localhost:3000/

    GET     /doctor/list=all
    GET     /doctor/list=filter&specializations={:values}&datesOfBirth={:values}
    GET     /doctor/:id
    PUT     /doctor/:id
    DELETE  /doctor/:id
    POST    /doctor

    GET     /patient/list=all
    GET     /patient/list=filter&phoneNumbers={:values}&datesOfBirth={:values}
    GET     /patient/:id
    PUT     /patient/:id
    DELETE  /patient/:id
    POST    /patient

    GET     /schedule/list=all
    GET     /schedule/list=filter&doctorID={:values}&patientID={:values}&dateAppointment={:values}
    DELETE  /schedule/:id
    POST    /schedule


_Примечание:_
* _для запуска необходима переменная окружения PORT_GOLANG со значением номера порта, на котором будет запускаться приложение_