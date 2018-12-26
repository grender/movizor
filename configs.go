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
