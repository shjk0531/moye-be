package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shjk0531/moye-be/internal/global/config"
	"github.com/shjk0531/moye-be/internal/global/routes"
)

func main() {
	// 환경 설정 초기화 (configs/config.yaml 사용)
	config.InitConfig()

	// Gin 엔진 생성
	router := gin.Default()

	// 전역 라우터 등록 (각 도메인 컨트롤러 연결)
	routes.RegisterRoutes(router)

	// 서버 실행 (환경설정에서 포트 사용 가능)
	router.Run(":8080")
}
