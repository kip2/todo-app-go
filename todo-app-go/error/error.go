package errorpkg

import "log"

/*
エラーハンドリング用
*/
func CheckError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
