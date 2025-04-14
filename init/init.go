package init

import (
	"adnpa/id-generater/internal/global"
	"adnpa/id-generater/internal/utils"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func init() {
	InitSnowFlake()
}

func InitSnowFlake() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	mid, err := strconv.ParseInt(os.Getenv("MACHINE_ID"), 10, 64)
	if err != nil {
		panic(err)
	}
	epoch, err := strconv.ParseInt(os.Getenv("EPOCH"), 10, 64)
	if err != nil {
		panic(err)
	}
	nsf, err := utils.NewSnowflake(mid, epoch)
	if err != nil {
		panic(err)
	}
	global.Snowflake = nsf
	log.Println("init snowflake succ")
}
