package env

import (
	"os"
)

var (
	AllowOrigin   string
	AdminPassword string
)

func init() {
	// .env 파일을 읽으려고 시도하던 부분을 완전히 제거했습니다.
	// 이제 App Engine의 app.yaml 파일에 적어둔 환경 변수를 직접 읽어옵니다.
	AllowOrigin = os.Getenv("ALLOW_ORIGIN")
	AdminPassword = os.Getenv("ADMIN_PASSWORD")
}