package movizor

import (
	"encoding/json"
	"fmt"
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
	Balance      float64                    `json:"balance"` // Текущий остаток средств на балансе
	Credit       float64                    `json:"credit"`  // Сумма кредитных средств на балансе
	ContractType string                     `json:"type"`    // Тип договора
	TariffPlans  map[string]json.RawMessage `json:"tariff"`  // Список операторов с их тарифами и доп тарифы
}

func (b *Balance) UnmarshalJSON(data []byte) (err error) {
	type Alias Balance
	aux := &struct {
		Balance json.Number `json:"balance"`
		Credit  json.Number `json:"credit"`
		*Alias
	}{
		Alias: (*Alias)(b),
	}
	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if b.Balance, err = aux.Balance.Float64(); err != nil {
		return err
	}

	if b.Credit, err = aux.Credit.Float64(); err != nil {
		return err
	}

	return nil
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

type Coordinate float32

func (c Coordinate) Float32() float32 {
	return float32(c)
}

func (c Coordinate) String() string {
	return fmt.Sprintf("%.8f", c.Float32())
}

func (c *Coordinate) UnmarshalJSON(data []byte) (err error) {
	var num json.Number
	err = json.Unmarshal(data, &num)
	if err != nil {
		return
	}

	var val float32
	val, err = jsonNumberToFloat32(num)
	if err != nil {
		return
	}

	*c = Coordinate(val)
	return nil
}

type Time time.Time

func (t Time) Time() time.Time {
	return time.Time(t)
}

func (t *Time) UnmarshalJSON(data []byte) error {
	var num json.Number
	err := json.Unmarshal(data, &num)
	if err != nil {
		return err
	}

	val, err := num.Int64()
	if err != nil {
		return err
	}

	*t = Time(time.Unix(val, 0))
	return nil
}

// Почти полная информация по объекту.
type ObjectInfo struct {
	Phone         Object            `json:"phone"`                              // Номер абонента
	Status        Status            `json:"status"`                             // status type
	Confirmed     bool              `json:"confirmed"`                          // Получено подтверждение от абонента
	Title         string            `json:"title"`                              // Имя абонента (название объекта)
	Tariff        TariffType        `json:"tariff"`                             // Текущий тарифный план
	TariffNew     *TariffType       `json:"tariff_new,omitempty"`               // Новый тарифный план со следующего дня
	LastTimestamp Time              `json:"last_timestamp"`                     // Время последнего запроса на определение местоположения
	AtRequest     bool              `json:"at_request,omitempty"`               // Производится определение местоположения в данный момент
	CurrentLon    *Coordinate       `json:"current_lon"`                        // Широта последнего местоположения
	CurrentLat    *Coordinate       `json:"current_lat"`                        // Долгота последнего местоположения
	Place         string            `json:"place,omitempty"`                    // Населенный пункт последнего местоположения
	Distance      int64             `json:"distance,omitempty"`                 // Остаток в км до конечной точки
	ETA           *Time             `json:"distance_forecast_time,omitempty"`   // Прогноз оставшегося времени до конечной точки
	ETAStatus     *string           `json:"distance_forecast_status,omitempty"` // Прогноз успеваемости до конечной точки
	OnParking     *bool             `json:"on_parking,omitempty"`               // Находится ли объект на парковке
	Destination   []Destination     `json:"destination,omitempty"`              // Список точек назначения, которые должен посетить Водитель.
	OfflineTime   Time              `json:"offline_time,omitempty"`             // Время последнего известного местоположения
	PosError      bool              `json:"pos_error,omitempty"`                // Последнее местоположение не удалось определить
	TimestampOff  Time              `json:"timestamp_off"`                      // Время автоматического отключения от мониторинга
	TimestampAdd  Time              `json:"timestamp_add"`                      // Время добавления объекта в Мовизор
	Metadata      map[string]string `json:"metadata,omitempty"`                 // Метаинформация объекта, массив
}

func (oi *ObjectInfo) UnmarshalJSON(data []byte) (err error) {
	type Alias ObjectInfo
	aux := &struct {
		Distance json.Number     `json:"distance,omitempty"`
		Metadata json.RawMessage `json:"metadata,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(oi),
	}

	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if oi.Distance, err = aux.Distance.Int64(); err != nil {
		return err
	}

	var probe []interface{}
	if err = json.Unmarshal(aux.Metadata, &probe); err == nil {
		return nil
	}

	if err = json.Unmarshal(aux.Metadata, &oi.Metadata); err != nil {
		return err
	}

	return nil
}

// Список точек назначения, которые должен посетить Водитель.
type Destination struct {
	Text   string     `json:"text"`
	Lat    Coordinate `json:"lat"`
	Lon    Coordinate `json:"lon"`
	Time   string     `json:"time"`
	Status ETAStatus  `json:"status"`
}

// Текущий статус объекта трекинга.
type ObjectStatus struct {
	Phone  Object `json:"phone"`  // Номер телефона абонента
	Status Status `json:"status"` // Статус добавления для отслеживания
}

// Список объектов с их статусами.
type ObjectsWithStatus []ObjectStatus

func (os ObjectsWithStatus) Len() int           { return len(os) }
func (os ObjectsWithStatus) Swap(i, j int)      { os[i], os[j] = os[j], os[i] }
func (os ObjectsWithStatus) Less(i, j int) bool { return os[i].Phone < os[j].Phone }

func (os ObjectsWithStatus) IsObjectIn(o Object) bool {
	for _, os := range os {
		if os.Phone == o {
			return true
		}
	}
	return false
}

// Список местоположений
type Positions []Position

// Информация о последнем зафиксированном в системе местоположении
type Position struct {
	Lon              Coordinate `json:"lon"`                              // Долгота
	Lat              Coordinate `json:"lat"`                              // Широта
	Timestamp        Time       `json:"timestamp"`                        // Время получения координат для этой точки
	TimestampRequest Time       `json:"timestamp_request,omitempty"`      // Время создания запроса на получение координат
	Deviation        int64      `json:"radius,omitempty"`                 // Радиус погрешности (м)
	Distance         int64      `json:"distance"`                         // Остаток в км до конечной точки
	ETA              Time       `json:"distance_forecast_time,omitempty"` // Прогноз оставшегося времени до конечной точки
	// Прогноз строится в зависимости от наличия информации о конечном пункте назначения и времени прибытия.
	// Если этой информации нет, значения элементов будут пустыми.
	ETAStatus ETAStatus `json:"distance_forecast_status,omitempty"` // Прогноз успеваемости до конечной точки.
	// Прогноз строится в зависимости от наличия информации о конечном пункте назначения и времени прибытия.
	// Если этой информации нет, значения элементов будут пустыми.
	Place string `json:"place"` // Населенный пункт местоположения.
}

func (p *Position) UnmarshalJSON(data []byte) (err error) {
	type Alias Position
	aux := &struct {
		Deviation json.Number `json:"radius"`
		Distance  json.Number `json:"distance"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if p.Deviation, err = aux.Deviation.Int64(); err != nil {
		return err
	}

	if p.Distance, err = aux.Distance.Int64(); err != nil {
		return err
	}

	return nil
}

// Список объектов с координатами, последним временем обновления координат, текущим местонахождением и ETA.
type ObjectPositions []ObjectPosition

// Координаты, последнее временя обновления координат, текущее местонахождение и ETA объекта.
type ObjectPosition struct {
	Phone     Object      `json:"phone"`                            // Номер телефона абонента
	Lon       json.Number `json:"lon"`                              // Широта
	Lat       json.Number `json:"lat"`                              // Долгота
	Timestamp Time        `json:"timestamp"`                        // Время
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

type PositionRequest struct {
	RequestID int64 `json:"request_id"`
}

func (pr PositionRequest) values() url.Values {
	return url.Values{"id": {strconv.FormatInt(pr.RequestID, 10)}}
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
	EventID   int64     `json:"id"`        // Идентификатор события (возрастающий номер события)
	Timestamp Time      `json:"timestamp"` // Время возникновения события
	Phone     Object    `json:"phone"`     // Номер телефона абонента, по которому произошло событие
	Event     EventType `json:"type"`      // Тип события
}

func (oe *ObjectEvent) UnmarshalJSON(data []byte) (err error) {
	type Alias ObjectEvent
	aux := &struct {
		EventID json.Number `json:"id"`
		*Alias
	}{
		Alias: (*Alias)(oe),
	}
	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if oe.EventID, err = aux.EventID.Int64(); err != nil {
		return err
	}

	return nil
}

// Список подписок на события
type SubscribedEvents []SubscribedEvent

// SubscribedEvent содержит информацию о подписке на одно из событий.
type SubscribedEvent struct {
	SubscriptionID         int64     `json:"id"`         // Идентификатор события (возрастающий номер события)
	IsAllObjectsSubscribed bool      `json:"phones_all"` // Уведомление о событии для всех объектов (в том числе добавляемых в будущем)
	ObjectsSubscribed      []Object  `json:"phones"`     // Список телефонов (объектов)
	Timestamp              Time      `json:"timestamp"`  // Время возникновения события
	Event                  EventType `json:"type"`       // Тип события, на которые зарегистрирована подписка
	Phone                  Object    `json:"phone"`      // Номер телефона абонента, по которому отправляются уведомления
	EMail                  string    `json:"email"`      // Email, по которому отправляются уведомления
	IsTelegram             bool      `json:"telegram"`   // Уведомления отправляются на аккаунт telegram указанный в настройках аккаунта
}

func (se *SubscribedEvent) UnmarshalJSON(data []byte) (err error) {
	type Alias SubscribedEvent
	aux := &struct {
		SubscriptionID         json.Number `json:"id"`
		IsAllObjectsSubscribed int         `json:"phones_all"`
		IsTelegram             int         `json:"telegram"`
		*Alias
	}{
		Alias: (*Alias)(se),
	}
	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if se.SubscriptionID, err = aux.SubscriptionID.Int64(); err != nil {
		return err
	}

	if aux.IsAllObjectsSubscribed == 1 {
		se.IsAllObjectsSubscribed = true
	}

	if aux.IsTelegram == 1 {
		se.IsTelegram = true
	}

	return nil
}

func (se SubscribedEvent) MakeOptions() (seo SubscribeEventOptions, err error) {
	seo = SubscribeEventOptions{}
	seo.Event = se.Event
	if se.IsAllObjectsSubscribed {
		seo.AllObjects = true
	} else {
		seo.Objects = se.ObjectsSubscribed
	}
	switch {
	case se.Phone.String() != "":
		err = seo.SetSMSNotification(se.Phone)
	case se.EMail != "":
		err = seo.SetEMailNotification(se.EMail)
	default:
		seo.SetTelegramNotification()
	}
	return
}
