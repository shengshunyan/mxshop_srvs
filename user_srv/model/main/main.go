package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"io"
	"strings"
)

func genMd5(code string) string {
	Md5 := md5.New()
	_, err := io.WriteString(Md5, code)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(Md5.Sum(nil))
}

func main() {
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	//	logger.Config{
	//		SlowThreshold: time.Second, // Slow SQL threshold
	//		LogLevel:      logger.Info, // Log level
	//		//IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
	//		//ParameterizedQueries:      true,          // Don't include params in the SQL log
	//		Colorful: true, // Disable color
	//	},
	//)
	//// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	//dsn := "shane:4163054Shun.@tcp(43.139.50.150:3306)/mxshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	//if err != nil {
	//	panic("failed to connect database")
	//}
	//fmt.Println("connect database success")
	//
	//err = db.AutoMigrate(&model.User{})
	//if err != nil {
	//	panic("failed to auto migrate database")
	//}

	//value := genMd5("sheng627")
	//fmt.Println(value)

	// 加密
	options := &password.Options{16, 100, 32, sha256.New}
	salt, encodedPwd := password.Encode("generic password", options)
	newPassword := fmt.Sprintf("$pbkdf2-sha256$%s$%s", salt, encodedPwd)
	fmt.Println(newPassword)

	// 验证
	passwordInfo := strings.Split(newPassword, "$")
	check := password.Verify("generic password", passwordInfo[2], passwordInfo[3], options)
	fmt.Println(check) // true
}
