package movizor

import "encoding/json"

const (
	// DefaultAPIMovizorEndpoint is default API Movizor endpoint
	DefaultAPIMovizorEndpoint = "https://movizor.ru/api"
	// APIMovizorEndpoint is API Movizor endpoint suffix to pass method and parameters
	// /%s/%s = /project/action
	APIMovizorEndpointSuffix = "/%s/%s"
)

type Operator string

const (
	OperatorMTS     Operator = "mts"
	OperatorMegafon Operator = "megafon"
	OperatorBeeline Operator = "beeline"
	OperatorTele2   Operator = "tele2"
)

type TariffType json.Number

const (
	TariffManual   TariffType = "0"   // Вручную
	TariffOnline   TariffType = "1"   // Онлайн
	TariffOneMonth TariffType = "3"   // Пакет на 1 месяц
	TariffEvery15  TariffType = "15"  // Каждые 15 мин
	TariffEvery30  TariffType = "30"  // Каждые 30 мин
	TariffEvery60  TariffType = "60"  // Каждые 60 мин
	TariffEvery180 TariffType = "180" // Каждые 3 часа
)

type Status string

const (
	StatusNew                Status = "new"          // Новый
	StatusWaitOk             Status = "wait"         // Ожидание подтверждения
	StatusOk                 Status = "ok"           // Подтвержден и доступен для мониторинга
	StatusWaitOff            Status = "off_new"      // В процессе отключения от мониторинга
	StatusOff                Status = "off"          // Отключен от мониторинга
	StatusRejected           Status = "rejected"     // Абонент отказался от мониторинга
	StatusNotConfirmed       Status = "no_confirmed" // Абонент не подтвердил подключение
	StatusError              Status = "error"        // Оператор не поддерживается
	StatusOperatorNotAllowed Status = "wrong"        // Оператор не поддерживается
	StatusTrackCollision     Status = "error_exists" // Номер отслеживается другой компанией у оператора
	StatusTrackingIsOff      Status = "limited"      // У абонента ограничение на подключение к услуге
)

type ETAStatus string

const (
	NoETAStatus ETAStatus = "" // статус не указан
	// NullETAStatus     ETAStatus = "null"     // статус не указан - значение null не парсится, но и ошибок нет.
	OkETAStatus       ETAStatus = "ok"       // успевает
	LateETAStatus     ETAStatus = "late"     // опаздывает
	FinishedETAStatus ETAStatus = "finished" // прибыл
)

type EventType string

const (
	AddEvent                    EventType = "add"             // - добавлен объект
	AutoOffEvent                EventType = "auto_off"        // - автоматическое отключение
	OffEvent                    EventType = "off"             // - отключение
	ConfirmEvent                EventType = "confirm"         // - объект подтвердил подключение
	RejectEvent                 EventType = "reject"          // - объект отказался от подключения
	RequestOkEvent              EventType = "request"         // - запрос: успешно
	RequestErrorEvent           EventType = "request_error"   // - запрос: ошибка
	RequestObjectOfflineEvent   EventType = "request_offline" // - запрос: телефон недоступен
	RequestObjectInRoamingEvent EventType = "request_roaming" // - запрос: телефон в роуминге
	ReactivateEvent             EventType = "reactivate"      // - повторное подключение
	ChangeTariffEvent           EventType = "tariff_auto"     // - смена тарифного плана
	InTimeEvent                 EventType = "pos_ok"          // - объект начал успевать
	LateEvent                   EventType = "pos_late"        // - объект начал опаздывать
	FinishedEvent               EventType = "pos_finished"    // - объект прибывает
	CallToDriverEvent           EventType = "autoinform"      // - автоинформатор
	NoConfirmationEvent         EventType = "no_confirm"      // - объект не подтвердил подключение
	ObjectLimitedEvent          EventType = "limit"           // - у объекта стоит ограничение
	OnRouteEvent                EventType = "onroute"         // - встал на маршрут
	ReturnRouteEvent            EventType = "returnroute"     // - вернулся на маршрут
	LeftRouteEvent              EventType = "leftroute"       // - отклонился от маршрута
	NotRouteEvent               EventType = "notroute"        // - не на маршруте
	OnParkingEvent              EventType = "onparking"       // - встал на парковку
	OffParkingEvent             EventType = "offparking"      // - начал движение
	MStopEvent                  EventType = "mstop"           // - приложение остановлено
	MStartEvent                 EventType = "mstart"          // - приложение запущено
)

type NotificationType string

const (
	SMSNotification      NotificationType = "sms"
	EMailNotification    NotificationType = "email"
	TelegramNotification NotificationType = "telegram"
)

//type Weekday uint8
//
//const (
//	Monday Weekday = 1
//	Tuesday
//	Wednesday
//	Thursday
//	Friday
//	Saturday
//	Sunday
//)
