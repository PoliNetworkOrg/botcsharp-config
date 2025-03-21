package main

import (
	"fmt"

	"github.com/PoliNetworkOrg/botcsharp-config/pkg/env"
	"github.com/PoliNetworkOrg/botcsharp-config/pkg/writer"
)

type BotConfig struct {
	BotTypeApi             int64   `json:"botTypeApi"`
	AcceptedMessages       bool    `json:"acceptedMessages"`
	Token                  string  `json:"token"`
	OnMessages             string  `json:"onMessages"`
	Website                *string `json:"website"`
	ContactString          *string `json:"contactString"`
	SessionUserId          *string `json:"SessionUserId"`
	ApiHash                *string `json:"apiHash"`
	NumberCountry          *string `json:"NumberCountry"`
	NumberNumber           *string `json:"NumberNumber"`
	PasswordToAuthenticate *string `json:"passwordToAuthenticate"`
	Method                 *string `json:"method"`
	UserId                 *int64  `json:"userId"`
	ApiId                  *int64  `json:"apiId"`
}

type BotConfigJson struct {
	Bots []BotConfig `json:"bots"`
}

type DbConfig struct {
	User     string `json:"User"`
	Password string `json:"Password"`
	Database string `json:"Database"`
	Host     string `json:"Host"`
	Port     int64  `json:"Port"`
}

type MatConfig struct {
	Password string `json:"Password"`
	RootDir  string `json:"RootDir"`
}

var ENV = env.Env{}

const (
	BotConfigFilename = "bots_info.json"
	DbConfigFilename  = "dbconfig.json"
	MatConfigFilename = "materialbotconfig.json"
)

func createBotConfig(outDir string) {
	writer, err := writer.NewWriter[BotConfigJson](outDir)
	if err != nil {
		panic(fmt.Sprintf("OUT_DIR %s does not exist", outDir))
	}

	config := BotConfig{
		Token:                  ENV.GetStringPanic("BOT_TOKEN"),
		OnMessages:             ENV.GetStringPanic("BOT_ON_MESSAGES"),
		BotTypeApi:             *ENV.GetInt("BOT_TYPE_API", i64Ptr(1)),
		AcceptedMessages:       *ENV.GetBool("BOT_ACCEPTED_MESSAGES", boolPtr(true)),
		Website:                ENV.GetString("BOT_WEBSITE", nil),
		ContactString:          ENV.GetString("BOT_CONTACT_STRING", nil),
		SessionUserId:          ENV.GetString("BOT_SESSION_USER_ID", nil),
		ApiHash:                ENV.GetString("BOT_API_HASH", nil),
		NumberCountry:          ENV.GetString("BOT_NUMBER_COUNTRY", nil),
		NumberNumber:           ENV.GetString("BOT_NUMBER_NUMBER", nil),
		PasswordToAuthenticate: ENV.GetString("BOT_PASSWORD_TO_AUTHENTICATE", nil),
		Method:                 ENV.GetString("BOT_METHOD", nil),
		UserId:                 ENV.GetInt("BOT_USER_ID", nil),
		ApiId:                  ENV.GetInt("BOT_API_ID", nil),
	}

	jsonConfig := BotConfigJson {
		Bots: []BotConfig{config},
	}

	writer.JsonWrite(BotConfigFilename, jsonConfig, true)
}

func createDbConfig(outDir string) {
	writer, err := writer.NewWriter[DbConfig](outDir)
	if err != nil {
		panic(fmt.Sprintf("OUT_DIR %s does not exist", outDir))
	}

	config := DbConfig{
		User:     ENV.GetStringPanic("DB_USER"),
		Password: ENV.GetStringPanic("DB_PASSWORD"),
		Database: ENV.GetStringPanic("DB_DATABASE"),
		Host:     ENV.GetStringPanic("DB_HOST"),
		Port:     *ENV.GetInt("DB_PORT", i64Ptr(3306)),
	}

	writer.JsonWrite(DbConfigFilename, config, true)
}

func createMatConfig(outDir string) {
	writer, err := writer.NewWriter[MatConfig](outDir)
	if err != nil {
		panic(fmt.Sprintf("OUT_DIR %s does not exist", outDir))
	}

	config := MatConfig{
		RootDir:  ENV.GetStringPanic("MAT_ROOT_DIR"),
		Password: ENV.GetStringPanic("MAT_PASSWORD"),
	}

	writer.JsonWrite(MatConfigFilename, config, true)
}

func main() {
	outDir := ENV.GetStringPanic("OUT_DIR")

	if *ENV.GetBool("CREATE_BOT_CONFIG", boolPtr(true)) {
		createBotConfig(outDir)
	}

	if *ENV.GetBool("CREATE_DB_CONFIG", boolPtr(false)) {
		createDbConfig(outDir)
	}

	if *ENV.GetBool("CREATE_MAT_CONFIG", boolPtr(false)) {
		createMatConfig(outDir)
	}

	fmt.Println("Finished.")
}

func boolPtr(b bool) *bool {
	return &b
}

func i64Ptr(i int64) *int64 {
	return &i
}
