package movizor

import (
	"encoding/json"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

// TODO: Добавить описание всех типов
// Ответ от сервиса с описанием типа сообщения и сегмента с данными.
type APIResponse struct {
	Result      string          `json:"result"`                  // "success" or "error" expected
	ResultCode  string          `json:"code"`                    // "OK" expected
	MessageType string          `json:"message"`                 // "Balance info", ...
	Data        json.RawMessage `json:"data,omitempty"`          // optional Payload of response
	ErrorCode   string          `json:"error_code,omitempty"`    // Код ошибка
	ErrorText   string          `json:"error_text,omitempty"`    // Текст ошибки
	ErrorTextRU string          `json:"error_text_ru,omitempty"` // optional Текст ошибки на русском
}

//type Tariff struct {
//	AbonentPayment json.Number `json:"abon"`    // Абоненская плата
//	RequestCost    json.Number `json:"request"` // Стоимость запроса
//	TariffTittle   string      `json:"title"`   // Название тарифа
//}

// Текущий баланс по договору и список подключенных тарифов по мобильным операторам.
type Balance struct {
	Balance      json.Number                `json:"balance"` // Текущий остаток средств на балансе
	Credit       json.Number                `json:"credit"`  // Сумма кредитных средств на балансе
	ContractType string                     `json:"type"`    // Тип договора
	TariffPlans  map[string]json.RawMessage `json:"tariff"`  // Список операторов с их тарифами и доп тарифы
}

// Номер подключаемого абонента в формате MSISDN (например, 79210010203).
// Возможно так же передавать номер при добавлении в систему в следующих форматах:( +7 (921) 001-02-03; 8-921-001-02-03).
type Object string

// Stringer returns clean format of cell number.
// Casting string(Object) gives Original value.
// fmt.Println(v), fmt.Printf("%s",Object), fmt.Printf("%v",Object) return clean format.
func (o Object) String() string {
	// ToDo: Переписать на что-то более надежное
	return regexp.MustCompile("[^0-9]").ReplaceAllString(string(o), "")
}

func (o Object) values() url.Values {
	return url.Values{"phone": {o.String()}}
}

// Почти полная информация по объекту.
// ToDo: Сделать обертки для ETA и ETAStatus - там могут быть null
type ObjectInfo struct {
	Phone         Object        `json:"phone"`                              // Номер абонента
	Status        Status        `json:"status"`                             // status type
	Confirmed     bool          `json:"confirmed"`                          // Получено подтверждение от абонента
	Title         string        `json:"title"`                              // Имя абонента (название объекта)
	Tariff        string        `json:"tariff"`                             // Текущий тарифный план
	TariffNew     string        `json:"tariff_new,omitempty"`               // Новый тарифный план со следующего дня
	LastTimestamp json.Number   `json:"last_timestamp"`                     // Время последнего запроса на определение местоположения
	AtRequest     bool          `json:"at_request,omitempty"`               // Производится определение местоположения в данный момент
	CurrentLon    json.Number   `json:"current_lon"`                        // Широта последнего местоположения
	CurrentLat    json.Number   `json:"current_lat"`                        // Долгота последнего местоположения
	Place         string        `json:"place,omitempty"`                    // Населенный пункт последнего местоположения
	Distance      json.Number   `json:"distance,omitempty"`                 // Остаток в км до конечной точки
	ETA           json.Number   `json:"distance_forecast_time,omitempty"`   // Прогноз оставшегося времени до конечной точки
	ETAStatus     string        `json:"distance_forecast_status,omitempty"` // Прогноз успеваемости до конечной точки
	OnParking     bool          `json:"on_parking,omitempty"`
	Destination   []Destination `json:"destination,omitempty"`
	OfflineTime   json.Number   `json:"offline_time,omitempty"` // Время последнего известного местоположения
	PosError      bool          `json:"pos_error,omitempty"`    // Последнее местоположение не удалось определить
	TimestampOff  json.Number   `json:"timestamp_off"`          // Время автоматического отключения от мониторинга
	TimestampAdd  json.Number   `json:"timestamp_add"`
	// ToDo: заменить на map[string]string и протестировать
	Metadata []json.RawMessage `json:"metadata,omitempty"` // Метаинформация объекта, массив
}

// TimeOff возвращает время отколючения объекта от мониторинга в типе time.
func (oi *ObjectInfo) TimeOff() time.Time {
	v, _ := oi.TimestampOff.Int64()
	return time.Unix(v, 0)
}

// TimeAdded возвращает время подключения объекта к мониторингу в типе time.
func (oi *ObjectInfo) TimeAdded() time.Time {
	v, _ := oi.TimestampAdd.Int64()
	return time.Unix(v, 0)
}

// Список точек назначения, которые должен посетить Водитель.
// ToDo: Протестировать работу
type Destination struct {
	Text   string      `json:"text"`
	Lat    json.Number `json:"lat"`
	Lon    json.Number `json:"lon"`
	Time   string      `json:"time"`
	Status ETAStatus   `json:"status"`
}

// Список объектов с их статусами.
type ObjectsWithStatus []ObjectStatus

func (os ObjectsWithStatus) Len() int           { return len(os) }
func (os ObjectsWithStatus) Swap(i, j int)      { os[i], os[j] = os[j], os[i] }
func (os ObjectsWithStatus) Less(i, j int) bool { return os[i].Phone < os[j].Phone }

// Текущий статус объекта трекинга.
type ObjectStatus struct {
	Phone  Object `json:"phone"`  // Номер телефона абонента
	Status Status `json:"status"` // Статус добавления для отслеживания
}

// Список местоположений
type Positions []Position

// Информация о последнем зафиксированном в системе местоположении
type Position struct {
	Lon              json.Number `json:"lon"`                              // Долгота
	Lat              json.Number `json:"lat"`                              // Широта
	Timestamp        json.Number `json:"timestamp"`                        // Время получения координат для этой точки
	TimestampRequest json.Number `json:"timestamp_request,omitempty"`      // Время создания запроса на получение координат
	Deviation        json.Number `json:"radius"`                           // Радиус погрешности (м)
	Distance         json.Number `json:"distance"`                         // Остаток в км до конечной точки
	ETA              json.Number `json:"distance_forecast_time,omitempty"` // Прогноз оставшегося времени до конечной точки
	// Прогноз строится в зависимости от наличия информации о конечном пункте назначения и времени прибытия.
	// Если этой информации нет, значения элементов будут пустыми.
	ETAStatus ETAStatus `json:"distance_forecast_status,omitempty"` // Прогноз успеваемости до конечной точки.
	// Прогноз строится в зависимости от наличия информации о конечном пункте назначения и времени прибытия.
	// Если этой информации нет, значения элементов будут пустыми.
	Place string `json:"place"` // Населенный пункт местоположения.
}

// LastTimeUpdated возвращает время последнего обновления координат в time.
func (p *Position) LastTimeUpdated() time.Time {
	v, _ := p.Timestamp.Int64()
	return time.Unix(v, 0)
}

// Список объектов с координатами, последним временем обновления координат, текущим местонахождением и ETA.
type ObjectPositions []ObjectPosition

// Координаты, последнее временя обновления координат, текущее местонахождение и ETA объекта.
type ObjectPosition struct {
	Phone     Object      `json:"phone"`                            // Номер телефона абонента
	Lon       json.Number `json:"lon"`                              // Широта
	Lat       json.Number `json:"lat"`                              // Долгота
	Timestamp json.Number `json:"timestamp"`                        // Время
	Deviation json.Number `json:"radius"`                           // Радиус погрешности (м)
	Place     string      `json:"place"`                            // Населенный пункт местоположения
	Distance  json.Number `json:"distance,omitempty"`               // Остаток в км до конечной точки
	ETA       json.Number `json:"distance_forecast_time,omitempty"` // Прогноз оставшегося времени до конечной точки
	// Прогноз строится в зависимости от наличия информации
	// о конечном пункте назначения и времени прибытия.
	// Если этой информации нет, значения элементов будут пустыми.
	ETAStatus ETAStatus `json:"distance_forecast_status,omitempty"` // Прогноз успеваемости до конечной точки.
	// Прогноз строится в зависимости от наличия информации
	// о конечном пункте назначения и времени прибытия.
	// Если этой информации нет, значения элементов будут пустыми.
}

// LastTimeUpdated возвращает время последнего обновления координат в time.
func (op *ObjectPosition) LastTimeUpdated() time.Time {
	v, _ := op.Timestamp.Int64()
	return time.Unix(v, 0)
}

type PositionRequest struct {
	RequestID uint64 `json:"request_id"`
}

func (pr PositionRequest) values() url.Values {
	return url.Values{"id": {strconv.FormatUint(pr.RequestID, 10)}}
}

// Информация по сотовому оператору
type OperatorInfo struct {
	Operator Operator `json:"operator"`         // Внутренний идентификатор оператора
	Title    string   `json:"title"`            // Название оператора
	Region   string   `json:"region,omitempty"` // Домашний регион абонента
}

// Список событий по объектам
type ObjectEvents []ObjectEvent

// ObjectEvent содержит информацию о событиях, которые происходили с объектом.
// Такие как: подтверждение трекинга, отклонение трекига, отклонения от маршрута следования и т.д.
type ObjectEvent struct {
	EventID   json.Number `json:"id"`        // Идентификатор события (возрастающий номер события)
	Timestamp json.Number `json:"timestamp"` // Время возникновения события
	Phone     Object      `json:"phone"`     // Номер телефона абонента, по которому произошло событие
	Event     EventType   `json:"type"`      // Тип события
}

// Список подписок на события
type SubscribedEvents []SubscribedEvent

// SubscribedEvent содержит информацию о подписке на одно из событий.
// ToDo: Обернуть IsAllPhoneSubscribed и IsTelegram в bool
type SubscribedEvent struct {
	SubscriptionID       json.Number `json:"id"`         // Идентификатор события (возрастающий номер события)
	IsAllPhoneSubscribed int         `json:"phones_all"` // Уведомление о событии для всех объектов (в том числе добавляемых в будущем)
	PhonesSubscribed     []Object    `json:"phones"`     // Список телефонов (объектов)
	Timestamp            json.Number `json:"timestamp"`  // Время возникновения события
	Event                EventType   `json:"type"`       // Тип события, на которые зарегистрирована подписка
	PhoneSubscribed      Object      `json:"phone"`      // Номер телефона абонента, по которому отправляются уведомления
	EMail                string      `json:"email"`      // Email, по которому отправляются уведомления
	IsTelegram           int         `json:"telegram"`   // Уведомления отправляются на аккаунт telegram указанный в настройках аккаунта
}
