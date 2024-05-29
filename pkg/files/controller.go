package files

import (
	"fmt"
	"io"
	golog "log"
	"net/http"
	"os"
	// "path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"file-system/pkg/helpers"
	"file-system/pkg/response"
)

type FileController struct {
	service FileService
}

func NewFileController(service FileService) FileController {
	return FileController{
		service: service,
	}
}

func (controller *FileController) Create(ctx *gin.Context) {
	log.Info().Msg("create file")
	// createFileRequest := CreateFileRequest{}
	// err := ctx.ShouldBindJSON(&createFileRequest)
	// helpers.ErrorHelper(err)

	ctx.Request.ParseMultipartForm(10 << 20)

	file, handler, err := ctx.Request.FormFile("file")

	if err != nil {
		helpers.ErrorHelper(err)
		return
	}

	golog.Println("upload")

	defer file.Close()

	fmt.Println("File Name: ", handler.Filename)
	fmt.Println("File Size: ", handler.Size)
	fmt.Println("MIME Header: ", handler.Header)

	tempFile, err := os.CreateTemp("./files", handler.Filename)

	if err != nil {
		helpers.ErrorHelper(err)
		return
	}

	defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)

	if err != nil {
		helpers.ErrorHelper(err)

	}

	tempFile.Write(fileBytes)

	fmt.Fprint(tempFile, handler.Filename, " was uploaded")

	// controller.service.Create()

	Response := response.Response{
		Code:   http.StatusNoContent,
		Status: "success",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusNoContent, Response)
}

func (controller *FileController) Delete(ctx *gin.Context) {}

func (controller *FileController) List(ctx *gin.Context) {}

func (controller *FileController) Retrieve(ctx *gin.Context) {}

func (controller *FileController) Update(ctx *gin.Context) {}
