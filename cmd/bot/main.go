package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Endropr/ai-programming-mentor/internal/domain"
	"github.com/Endropr/ai-programming-mentor/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

func main() {
	_ = godotenv.Load()

	// Подключение к БД
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("Ошибка БД:", err)
	}
	defer conn.Close(context.Background())
	repo := repository.NewPostgresRepo(conn)

	// Подключение к ТГ
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Fatal("Ошибка ТГ:", err)
	}
	updates := bot.GetUpdatesChan(tgbotapi.NewUpdate(0))

	// 3. Опен аи клиент
	aiClient := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	// Хранилище языков в серве
	userLanguages := make(map[int64]string)

	fmt.Println(" \u001b[37;1m--- Бот запущен и готов к общению! ---\033[0m ")

	for update := range updates {
		// Кнопки
		if update.CallbackQuery != nil {
			userID := update.CallbackQuery.From.ID
			callbackData := update.CallbackQuery.Data
			var responseText string
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "Выбран язык: "+userLanguages[userID])
			bot.Request(callback)

			// Выбор юзера
			switch callbackData {
			case "lang_go":
				userLanguages[userID] = "Go"
				responseText = "Go — это компилируемый язык программирования с открытым исходным кодом, разработанный в Google для создания высокоэффективных и масштабируемых сервисов. Он сочетает в себе производительность C++ с простотой Python, что делает его идеальным выбором для работы с облачными технологиями и микросервисами.\n\n" +
					"Давай приступим! С чего хочешь начать изучение языка?"
			case "lang_python":
				userLanguages[userID] = "Python"
				responseText = "Python — это высокоуровневый язык программирования с акцентом на читаемость кода и продуктивность разработчика. Благодаря понятному синтаксису, напоминающему английский язык, он считается идеальным выбором для новичков и золотым стандартом в сфере Data Science и AI.\n\n" +
					"Давай приступим! С чего хочешь начать изучение языка?"
			case "lang_js":
				userLanguages[userID] = "JS"
				responseText = "JavaScript (JS) — это мультипарадигменный язык программирования, ставший стандартом де-факто для создания интерактивных веб-страниц. Первоначально созданный для браузеров, сегодня он позволяет разрабатывать серверные приложения (Node.js), мобильные и десктопные программы.\n\n" +
					"Давай приступим! С чего хочешь начать изучение языка?"
			case "lang_c++":
				userLanguages[userID] = "C++"
				responseText = "C++ — это высокопроизводительный компилируемый язык программирования, который предоставляет разработчику полный контроль над аппаратными ресурсами и памятью компьютера. Он является прямым наследником языка C, дополняя его объектно-ориентированными возможностями и мощными абстракциями.\n\n" +
					"Давай приступим! С чего хочешь начать изучение языка?"
			case "lang_php":
				userLanguages[userID] = "PHP"
				responseText = "PHP — это серверный скриптовый язык программирования, специально разработанный для веб-разработки. Он является фундаментом для более чем 75% всех сайтов в интернете, включая такие гиганты, как Wikipedia и системы управления контентом вроде WordPress.\n\n" +
					"Давай приступим! С чего хочешь начать изучение языка?"
			case "lang_html":
				userLanguages[userID] = "HTML/CSS"
				responseText = "HTML и CSS — это неразлучный дуэт технологий, на которых держится весь визуальный интернет. HTML отвечает за структуру и скелет страницы (заголовки, списки, кнопки), а CSS — за её внешний вид, стиль и адаптивность под разные экраны (цвета, шрифты, сетки и анимации).\n\n" +
					"Давай приступим! С чего хочешь начать изучение языка?"
			}

			bot.Request(tgbotapi.NewCallback(update.CallbackQuery.ID, ""))
			bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, responseText))
			continue
		}

		if update.Message == nil {
			continue
		}

		// Обработка старта и кнопки
		if update.Message.IsCommand() && update.Message.Command() == "start" {
			msgText := "<b>🚀 Добро пожаловать в IT Mentor Bot!</b>\n" +
				"Я - бот на основе искусственного интеллекта, который помогает в изучении языков программирования.\n" +
				"Готов помочь разобраться в сложном синтаксисе, сделаю для тебя код-ревью и составлю план обучения под твои цели.\n\n" +
				"<b>❓Чем я могу быть полезен?</b>\n" +
				"— Объясню любую тему «на пальцах».\n" +
				"— Проверю твой код и подскажу, как его улучшить.\n" +
				"— Подберу практические задачи твоего уровня.\n\n" +
				"<b>Выбери направление ниже, и начнём кодить!</b> 👇"

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)

			// ВОТ ЭТА СТРОКА ВКЛЮЧАЕТ ЖИРНЫЙ ШРИФТ И ТЕГИ
			msg.ParseMode = "HTML"

			keyboard := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Golang", "lang_go"),
					tgbotapi.NewInlineKeyboardButtonData("Python", "lang_python"),
					tgbotapi.NewInlineKeyboardButtonData("JavaScript", "lang_js"),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("C++", "lang_cpp"),
					tgbotapi.NewInlineKeyboardButtonData("PHP", "lang_php"),
					tgbotapi.NewInlineKeyboardButtonData("HTML/CSS", "lang_html"),
				),
			)
			msg.ReplyMarkup = keyboard
			bot.Send(msg)
			continue
		}

		// Обработка дефолт сообщений

		// Определяем текущий язык
		currentLang := userLanguages[update.Message.From.ID]
		if currentLang == "" {
			currentLang = "Не выбран"
		}

		// Сейвим вопрос юзера в бд (с учетом языка)
		userMsg := domain.Message{
			UserID:           update.Message.From.ID,
			Role:             "user",
			Content:          update.Message.Text,
			SelectedLanguage: currentLang,
		}
		repo.SaveMessage(context.Background(), userMsg)

		// Получаем ответ от ИИ
		aiReply := getAIResponse(aiClient, update.Message.Text)

		// Сейв ответ ИИ в бд
		botMsg := domain.Message{
			UserID:           update.Message.From.ID,
			Role:             "assistant",
			Content:          aiReply,
			SelectedLanguage: currentLang,
		}
		repo.SaveMessage(context.Background(), botMsg)

		// Отправляем ответ в телегу
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, aiReply)
		bot.Send(msg)
	}
}

func getAIResponse(client *openai.Client, userText string) string {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "Ты — ментор по кодингу. Отвечай на простом и понятном юзеру языке (без заумных слов), по делу и с юмором.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: userText,
				},
			},
		},
	)

	if err != nil {
		log.Printf("Ошибка OpenAI: %v", err)
		return "Сорян, бот щас тебе не ответит, потому что жадный разраб зажал 5 баксов и не купил ВДС для бота с подпиской OpenAI"
	}

	return resp.Choices[0].Message.Content
}
