package handler

import (
	"deepsearch/database"
	"deepsearch/pkg"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"

)

func Search(c *fiber.Ctx) error {
	encodedToken := c.Params("token")
	if encodedToken == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Token is required",
		})
	}

	token, err := url.QueryUnescape(encodedToken)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid token parameter",
		})
	}

	result, err := database.RunSearch(token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	if result == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "no_content",
			"message": "No results found",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   result,
	})
}
type WebSocketMessage struct {
	Status   string      `json:"status"`   
	Message  string      `json:"message"`  
	Progress int         `json:"progress"` 
	Data     interface{} `json:"data"`    
}

func SearchWebSocket(c *websocket.Conn) {

	log.Printf("Yeni WebSocket istemcisi bağlandı: %s", c.RemoteAddr())

	defer func() {
		log.Printf("WebSocket istemci bağlantısı kapandı: %s", c.RemoteAddr())
		c.Close()
	}()

	
	for {
		
		_, msg, err := c.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Beklenmedik WebSocket hatası: %v", err)
			}
			break
		}

		searchTerm := string(msg)
		log.Printf("'%s' için arama isteği alındı from %s", searchTerm, c.RemoteAddr())

		go runSearchAndReportProgress(c, searchTerm)
	}
}

func runSearchAndReportProgress(c *websocket.Conn, term string) {
	sendMessage := func(status, message string, progress int, data interface{}) {
		msg := WebSocketMessage{Status: status, Message: message, Progress: progress, Data: data}
		if err := c.WriteJSON(msg); err != nil {
			log.Printf("WebSocket'e yazma hatası: %v", err)
		}
	}


	sendMessage("progress", fmt.Sprintf("Veritabanında '%s' aranıyor...", term), 10, nil)
	results, err := database.PerformSearch(term)
	if err != nil {
		errMsg := fmt.Sprintf("Veritabanı arama hatası: %v", err)
		log.Println(errMsg)
		sendMessage("error", errMsg, 100, nil)
		return 
	}

	
	if len(results) > 0 {
		log.Printf("'%s' için sonuç veritabanında bulundu.", term)
		sendMessage("complete", "Sonuç veritabanında bulundu.", 100, results[0].Content)
		return
	}

	
	sendMessage("progress", "Veritabanında bulunamadı. Web'de aranıyor...", 30, nil)
	time.Sleep(500 * time.Millisecond) 

	geminiSummary, geminiErr := pkg.Gemini(term)
	if geminiErr != nil {
		errMsg := fmt.Sprintf("Gemini servisinden özet alınamadı: %v", geminiErr)
		log.Println(errMsg)
		sendMessage("error", errMsg, 100, nil)
		return
	}
	if geminiSummary == "" {
		errMsg := "Gemini servisinden boş yanıt döndü. İşlem durduruldu."
		log.Println(errMsg)
		sendMessage("error", errMsg, 100, nil)
		return
	}


	sendMessage("progress", "Web'den gelen sonuç veritabanına kaydediliyor...", 75, nil)
	if err := database.AddSearchEntry(term, geminiSummary); err != nil {
		errMsg := fmt.Sprintf("Yeni kayıt veritabanına eklenemedi: %v", err)
		log.Println(errMsg)
		sendMessage("error", errMsg, 100, nil)
		return
	}

	log.Printf("'%s' için işlem başarıyla tamamlandı.", term)
	sendMessage("complete", "Web araması sonucu başarıyla alındı ve kaydedildi.", 100, geminiSummary)
}