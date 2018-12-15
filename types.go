package movizor

import (
	"encoding/json"
)

// TODO: Добавить описание всех типов
// Ответ от сервиса с описанием типа сообщения и сегмента с данными
type APIResponse struct {
	Result      string          `json:"result"`        // "success" or "error" expected
	ResultCode  string          `json:"code"`          // "OK" expected
	MessageType string          `json:"message"`       // "Balance info", ...
	Data        json.RawMessage `json:"data"`          // optional Payload of response
	ErrorCode   string          `json:"error_code"`    // Код ошибка
	ErrorText   string          `json:"error_text"`    // Текст ошибки
	ErrorTextRU string          `json:"error_text_ru"` // optional Текст ошибки на русском
}

type Tariff struct {
	AbonentPayment json.Number `json:"abon"`    // Абоненская плата
	RequestCost    json.Number `json:"request"` // Стоимость запроса
	TariffTittle   string      `json:"title"`   // Название тарифа
}

// Информация о балансе договора и тарифном плане
// TODO: Сделать enum
//	TariffManual *Tariff 	`json:"0"`			// Вручную
//	TariffOnline *Tariff 	`json:"1"`			// Онлайн
//	TariffOneMonth *Tariff	`json:"3"`			// Пакет на 1 месяц
//	TariffEvery15 *Tariff	`json:"15"`			// Каждые 15 мин
//	TariffEvery30 *Tariff	`json:"15"`			// Каждые 30 мин
//	TariffEvery60 *Tariff	`json:"15"`			// Каждые 60 мин
//	TariffEvery180 *Tariff	`json:"15"`			// Каждые 3 часа
type Balance struct {
	Balance      json.Number                `json:"balance"` // Текущий остаток средств на балансе
	Credit       json.Number                `json:"credit"`  // Сумма кредитных средств на балансе
	ContractType string                     `json:"type"`    // Тип договора
	TariffPlans  map[string]json.RawMessage `json:"tariff"`  // Список операторов с их тарифами и доп тарифы
}

type ObjectPositions []ObjectPosition

type ObjectPosition struct {
	Phone     json.Number `json:"phone"`                  // Номер телефона абонента
	Lon       json.Number `json:"lon"`                    // Широта
	Lat       json.Number `json:"lat"`                    // Долгота
	Timestamp int64       `json:"timestamp"`              // Время
	Deviation int64       `json:"radius"`                 // Радиус погрешности (м)
	Place     string      `json:"place"`                  // Населенный пункт местоположения
	Distance  string      `json:"distance"`               // Остаток в км до конечной точки
	ETA       string      `json:"distance_forecast_time"` // Прогноз оставшегося времени до конечной точки
	// Прогноз строится в зависимости от наличия информации
	// о конечном пункте назначения и времени прибытия.
	// Если этой информации нет, значения элементов будут пустыми.
	ETAStatus string `json:"distance_forecast_status"` // Прогноз успеваемости до конечной точки.
	// Прогноз строится в зависимости от наличия информации
	// о конечном пункте назначения и времени прибытия.
	// Если этой информации нет, значения элементов будут пустыми.
}
