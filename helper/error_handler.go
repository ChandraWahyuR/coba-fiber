package helper

import (
	"log"
	"net/http"
	"presensi/constant"

	"github.com/gofiber/fiber/v2"
)

func ConverResponse(err error) int {
	log.Printf("Received error: %v", err)
	switch err {
	// General errors
	case constant.ErrBadRequest:
		return http.StatusBadRequest
	case constant.ErrUnauthorized:
		return http.StatusUnauthorized
	case constant.ErrEmptyOtp:
		return http.StatusBadRequest
	case constant.ErrDataNotfound, constant.ErrKursusNotfound, constant.ErrInstrukturNotFound, constant.ErrKategoriNotFound, constant.ErrUserNotFound:
		return http.StatusNotFound
	case constant.ErrGetData, constant.ErrGetInstruktur, constant.ErrGetID:
		return http.StatusNotFound
	case constant.ErrEmptyId:
		return http.StatusBadRequest

	// JWT errors
	case constant.ErrGenerateJWT, constant.ErrValidateJWT:
		return http.StatusUnauthorized

	// Update and validator errors
	case constant.ErrUpdate, constant.ErrHashPassword:
		return http.StatusInternalServerError

	// Register errors
	case constant.ErrEmptyEmailRegister, constant.ErrEmptyNameRegister, constant.ErrEmptyPasswordRegister, constant.ErrPasswordNotMatch, constant.ErrInvalidEmail, constant.ErrInvalidUsername, constant.ErrInvalidPhone:
		return http.StatusBadRequest

	// Login errors
	case constant.ErrEmptyLogin, constant.ErrEmptyPasswordLogin, constant.ErrInvalidPassword, constant.ErrLenPassword:
		return http.StatusBadRequest

	// Admin errors
	case constant.ErrAdminNotFound, constant.ErrAdminUserNameEmpty, constant.ErrAdminPasswordEmpty, constant.ErrEmptyGender, constant.ErrGenderChoice:
		return http.StatusBadRequest

	// Instruktur errors
	case constant.ErrInstrukturNotFound, constant.ErrInstrukturID, constant.ErrGenderInstruktorRmpty, constant.ErrEmptyNameInstuktor, constant.ErrEmptyEmailInstuktor, constant.ErrEmptyAlamatInstuktor, constant.ErrEmptyNumbertelponInstuktor, constant.ErrEmptyDescriptionInstuktor:
		return http.StatusBadRequest

	// Kategori errors
	case constant.ErrKategoriNotFound, constant.ErrEmptyNamaKategori, constant.ErrEmptyImageUrlKategori, constant.ErrEmptyDeskripsiKategori:
		return http.StatusBadRequest

	// Kursus errors
	case constant.ErrKursusNotFound, constant.ErrJadwal, constant.ErrJadwalFormat, constant.ErrGambarKursus, constant.ErrKategoriKursus, constant.ErrMateriPembelajaran, constant.ErrDekripsiKursus, constant.ErrHargaKursus:
		return http.StatusBadRequest

	// GCS errors
	case constant.ErrOpeningFile, constant.ErrUploadGCS:
		return http.StatusInternalServerError

	// Voucher errors
	case constant.ErrVoucherNotFound, constant.ErrVoucherFailedCreate, constant.ErrVoucherUsed, constant.ErrNameVoucher, constant.ErrDekripsiVoucher, constant.ErrDiscountVoucher, constant.ErrExpriedAtVoucher:
		return http.StatusBadRequest
	case constant.ErrVoucherIDNotFound:
		return http.StatusNotFound

	// Transaksi errors
	case constant.ErrTransaksiNotFound, constant.ErrValidateDokumenUser, constant.ErrSameKursusValid:
		return http.StatusBadRequest
	// Default case for internal server errors
	default:
		return http.StatusInternalServerError
	}
}
func HandleFiberError(c *fiber.Ctx, err error) (int, string) {
	if err != nil {
		// Fiber has built-in error handling
		return fiber.StatusBadRequest, constant.BadInput
	}
	return fiber.StatusBadRequest, constant.BadInput
}

func UnauthorizedError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(FormatResponse(false, constant.Unauthorized, nil))
}

func InternalServerError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(FormatResponse(false, constant.InternalServerError, nil))
}

func JWTErrorHandler(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(FormatResponse(false, constant.Unauthorized, nil))
}
