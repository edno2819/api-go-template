package setting

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type App struct {
	JwtSecret string
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string
	QrCodeSavePath string
	FontSavePath   string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Database struct {
	Url      string
	Database string
}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var AppSetting = &App{}
var ServerSetting = &Server{}
var RedisSetting = &Redis{}
var DatabaseSetting = &Database{}

func Setup() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DatabaseSetting.Url = os.Getenv("MONGO_URL")
	DatabaseSetting.Database = os.Getenv("MONGO_DB")
	AppSetting.JwtSecret = os.Getenv("JWT_KEY")

}
