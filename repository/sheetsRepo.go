package repository

import (
	"context"
	"fmt"
	"io"
	"log"
	"myHttpServer/models"
	"myHttpServer/utils"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var SheetsRepo = NewSheetsRepo()

type sheetsRepo struct {
	sheetsService  *sheets.Service
	masterSheet    *string
	userSheetTitle *string
	taskSheetTitle *string
}

func NewSheetsRepo() *sheetsRepo {
	ctx := context.Background()
	//Get Password from dotenv
	godotenv.Load()
	serviceAccountCredPath := os.Getenv(utils.SHEETS_SERVICE_ACCOUNT_KEY_PATH)
	scope := "https://www.googleapis.com/auth/spreadsheets"
	credJson := readJSON(serviceAccountCredPath)
	credentials, _ := google.CredentialsFromJSON(ctx, credJson, scope)

	sheetsService, err := sheets.NewService(ctx, option.WithCredentials(credentials))
	if err != nil {
		log.Println(err)
	}

	masterSheet := os.Getenv(utils.MASTER_SHEET)
	userSheet, _ := strconv.Atoi(os.Getenv(utils.USER_SHEET))
	taskSheet, _ := strconv.Atoi(os.Getenv(utils.TASK_SHEET))

	var userSheetTitle string
	var taskSheetTitle string

	sheet, _ := sheetsService.Spreadsheets.Get(masterSheet).Fields("sheets(properties(sheetId,title))").Do()
	log.Println(sheet)
	for _, value := range sheet.Sheets {
		prop := value.Properties
		if prop.SheetId == int64(userSheet) {
			userSheetTitle = prop.Title
			continue
		}
		if prop.SheetId == int64(taskSheet) {
			taskSheetTitle = prop.Title
			continue
		}
	}

	return &sheetsRepo{
		sheetsService:  sheetsService,
		masterSheet:    &masterSheet,
		userSheetTitle: &userSheetTitle,
		taskSheetTitle: &taskSheetTitle,
	}
}

// Save User Row
func (r *sheetsRepo) SaveUser(*models.User) {

}

// Save Task Row
func (s *sheetsRepo) SaveTask(task *models.Task) {

	//Append value to the sheet.
	row := &sheets.ValueRange{
		Values: [][]interface{}{{"1", "ABC", "abc1@gmail.com", task.Username, task.Type}},
	}

	response2, err := s.sheetsService.Spreadsheets.Values.Append(*s.masterSheet, *s.taskSheetTitle, row).ValueInputOption("USER_ENTERED").Context(context.Background()).Do()
	if err != nil || response2.HTTPStatusCode != 200 {
		log.Println(err)
		return
	}
}

func readJSON(filePath string) []byte {
	jsonFile, err := os.Open(filePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened " + filePath)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := io.ReadAll(jsonFile)
	log.Println("Read" + string(byteValue))
	return byteValue
}
