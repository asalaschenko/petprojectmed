package routes

import (
	"petprojectmed/common"
	doctorservices "petprojectmed/doctorServices"
	"petprojectmed/doctors"
	patientservices "petprojectmed/patientServices"
	"petprojectmed/patients"
	"petprojectmed/schedule"
	scheduleservices "petprojectmed/scheduleServices"
	"petprojectmed/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func RegisterRoutesDoctors(app *fiber.App, port string, conn *pgx.Conn) {
	var ListServiceDoctor = doctorservices.NewListService(storage.NewIDofDoctors(conn), storage.NewDoctorsByID(conn))
	var FilterServiceDoctor = doctorservices.NewFilterService(storage.NewAllDoctors(conn), storage.NewDoctorsIDandSpecializations(conn), storage.NewDoctorsIDandDateOfBirth(conn), storage.NewDoctorsByID(conn))
	var CreateServiceDoctor = doctorservices.NewCreateService(storage.NewDoctorCreate(conn))
	var UpdateServiceDoctor = doctorservices.NewUpdateService(storage.NewDoctorFamilyByID(conn), storage.NewDoctorNameByID(conn), storage.NewDoctorDateOfBirthByID(conn), storage.NewDoctorSpecializationByID(conn), storage.NewDoctorCabinetByID(conn))
	var DeleteServiceDoctor = doctorservices.NewDeleteService(storage.NewDoctorsByID(conn), storage.NewDoctorByID(conn))
	var ControllerGetDoctors = doctors.NewControllerGetDoctors(FilterServiceDoctor, ListServiceDoctor)
	var ControllerCreateDoctors = doctors.NewControllerCreateDoctor(CreateServiceDoctor)
	var ControllerUpdataDeleteDoctors = doctors.NewControllerUpdateDeleteDoctor(UpdateServiceDoctor, DeleteServiceDoctor)

	if port == "3000" {
		doctor := app.Group("/doctor")
		doctor.Get("/", func(c *fiber.Ctx) error {
			return ControllerGetDoctors.GetDoctors(c, common.FILTER)
		})
		doctor.Get("/:id", func(c *fiber.Ctx) error {
			return ControllerGetDoctors.GetDoctors(c, common.LIST_ID)
		})
		doctor.Post("/", ControllerCreateDoctors.CreateDoctor)
		doctor.Delete("/:id", func(c *fiber.Ctx) error {
			return ControllerUpdataDeleteDoctors.UpdateDeleteDoctor(c, common.DELETE)
		})
		doctor.Put("/:id", func(c *fiber.Ctx) error {
			return ControllerUpdataDeleteDoctors.UpdateDeleteDoctor(c, common.UPDATE)
		})
	}
}

func RegisterRoutesPatients(app *fiber.App, port string, conn *pgx.Conn) {
	var ListServicePatient = patientservices.NewListService(storage.NewIDofPatients(conn), storage.NewPatientsByID(conn))
	var FilterServicePatient = patientservices.NewFilterService(storage.NewAllPatients(conn), storage.NewPatientIDandPhoneNumbers(conn), storage.NewPatientIDandDateOfBirth(conn), storage.NewPatientsByID(conn))
	var CreateServicePatient = patientservices.NewCreateService(storage.NewPatientCreate(conn))
	var UpdateServicePatient = patientservices.NewUpdateService(storage.NewPhoneNumberOfPatients(conn), storage.NewPatientFamilyByID(conn), storage.NewDoctorNameByID(conn), storage.NewDoctorDateOfBirthByID(conn), storage.NewPatientPhoneNumberByID(conn), storage.NewPatientGenderByID(conn))
	var DeleteServicePatient = patientservices.NewDeleteService(storage.NewPatientsByID(conn), storage.NewDoctorByID(conn))
	var ControllerGetPatients = patients.NewControllerGetPatients(FilterServicePatient, ListServicePatient)
	var ControllerCreatePatients = patients.NewControllerCreatePatients(CreateServicePatient)
	var ControllerUpdateDeletePatients = patients.NewControllerUpdateDeletePatient(UpdateServicePatient, DeleteServicePatient)

	if port == "3000" {
		doctor := app.Group("/doctor")
		doctor.Get("/", func(c *fiber.Ctx) error {
			return ControllerGetPatients.GetPatients(c, common.FILTER)
		})
		doctor.Get("/:id", func(c *fiber.Ctx) error {
			return ControllerGetPatients.GetPatients(c, common.LIST_ID)
		})
		doctor.Post("/", ControllerCreatePatients.CreatePatient)
		doctor.Delete("/:id", func(c *fiber.Ctx) error {
			return ControllerUpdateDeletePatients.UpdateDeletePatient(c, common.DELETE)
		})
		doctor.Put("/:id", func(c *fiber.Ctx) error {
			return ControllerUpdateDeletePatients.UpdateDeletePatient(c, common.UPDATE)
		})
	}
}

func RegisterRoutesSchedule(app *fiber.App, port string, conn *pgx.Conn) {
	var FilterServiceSchedule = scheduleservices.NewFilterService(storage.NewAllAppointment(conn), storage.NewScheduleIDandDoctorID(conn), storage.NewScheduleIDandPatientID(conn), storage.NewScheduleIDandDateAppointment(conn), storage.NewAppointmentsByID(conn))
	var CreateServiceSchedule = scheduleservices.NewCreateService(storage.NewScheduleDateAppointmentsByDoctorID(conn), storage.NewAppointmentCreate(conn))
	var DeleteServiceSchedule = scheduleservices.NewDeleteService(storage.NewAppointmentsByID(conn), storage.NewAppointmentByID(conn))
	var ControllerGetAppointments = schedule.NewControllerGetAppointment(FilterServiceSchedule)
	var ControllerCreateAppointment = schedule.NewControllerCreateAppointment(CreateServiceSchedule)
	var ControllerDeleteAppointment = schedule.NewControllerDeleteAppointment(DeleteServiceSchedule)

	if port == "3000" {
		Schedule := app.Group("/schedule")
		Schedule.Get("/", ControllerGetAppointments.GetAppointment)
		Schedule.Post("/", ControllerCreateAppointment.CreateAppointment)
		Schedule.Delete("/:id", ControllerDeleteAppointment.DeleteAppointment)
	}
}
