package movizor

import (
	"encoding/json"
	"net/url"
	"regexp"
	"time"
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

//type Tariff struct {
//	AbonentPayment json.Number `json:"abon"`    // Абоненская плата
//	RequestCost    json.Number `json:"request"` // Стоимость запроса
//	TariffTittle   string      `json:"title"`   // Название тарифа
//}

type Balance struct {
	Balance      json.Number                `json:"balance"` // Текущий остаток средств на балансе
	Credit       json.Number                `json:"credit"`  // Сумма кредитных средств на балансе
	ContractType string                     `json:"type"`    // Тип договора
	TariffPlans  map[string]json.RawMessage `json:"tariff"`  // Список операторов с их тарифами и доп тарифы
}

// Номер подключаемого абонента в формате MSISDN (например, 79210010203).
// Возможно так же передавать номер при добавлении в систему в следующих форматах:( +7 (921) 001-02-03; 8-921-001-02-03)
type Object string

// Stringer returns clean format of cell number.
// Casting string(Object) gives Original value.
// fmt.Println(v), fmt.Printf("%s",Object), fmt.Printf("%v",Object) return clean format.
func (o Object) String() string {
	return regexp.MustCompile("[^0-9]").ReplaceAllString(string(o), "")
}

func (o Object) values() url.Values {
	return url.Values{"phone": {o.String()}}
}

type ObjectInfo struct {
	Phone         Object            `json:"phone"`
	Status        Status            `json:"status"`
	Confirmed     bool              `json:"confirmed"`
	Title         string            `json:"title"`
	Tariff        string            `json:"tariff"`
	TariffNew     string            `json:"tariff_new"`
	LastTimestamp json.Number       `json:"last_timestamp"`
	AtRequest     bool              `json:"at_request"`
	CurrentLon    json.Number       `json:"current_lon"`
	CurrentLat    json.Number       `json:"current_lat"`
	Place         string            `json:"place"`
	Distance      json.Number       `json:"distance"`
	ETA           json.Number       `json:"distance_forecast_time"`
	ETAStatus     string            `json:"distance_forecast_status"`
	OnParking     bool              `json:"on_parking"`
	Destination   []Destination     `json:"destination"`
	OfflineTime   json.Number       `json:"offline_time"`
	PosError      bool              `json:"pos_error"`
	TimestampOff  json.Number       `json:"timestamp_off"`
	TimestampAdd  json.Number       `json:"timestamp_add"`
	Metadata      []json.RawMessage `json:"metadata"`
}

func (oi *ObjectInfo) TimeOff() time.Time {
	v, _ := oi.TimestampOff.Int64()
	return time.Unix(v, 0)
}

func (oi *ObjectInfo) TimeAdded() time.Time {
	v, _ := oi.TimestampAdd.Int64()
	return time.Unix(v, 0)
}

type Destination struct {
	Text   string      `json:"text"`
	Lat    json.Number `json:"lat"`
	Lon    json.Number `json:"lon"`
	Time   string      `json:"time"`
	Status string      `json:"status"`
}

type ObjectsWithStatus []ObjectStatus

func (os ObjectsWithStatus) Len() int           { return len(os) }
func (os ObjectsWithStatus) Swap(i, j int)      { os[i], os[j] = os[j], os[i] }
func (os ObjectsWithStatus) Less(i, j int) bool { return os[i].Phone < os[j].Phone }

type ObjectStatus struct {
	Phone  Object `json:"phone"`  // Номер телефона абонента
	Status Status `json:"status"` // Статус добавления для отслеживания
}

//type ObjectsWithStatusChannel <-chan ObjectStatus
//
//// Clear discards all unprocessed incoming elements of ObjectStatus.
//func (ch ObjectsWithStatusChannel) Clear() {
//	for len(ch) != 0 {
//		<-ch
//	}
//}

type ObjectPositions []ObjectPosition

type ObjectPosition struct {
	Phone     Object      `json:"phone"`                  // Номер телефона абонента
	Lon       json.Number `json:"lon"`                    // Широта
	Lat       json.Number `json:"lat"`                    // Долгота
	Timestamp json.Number `json:"timestamp"`              // Время
	Deviation json.Number `json:"radius"`                 // Радиус погрешности (м)
	Place     string      `json:"place"`                  // Населенный пункт местоположения
	Distance  json.Number `json:"distance"`               // Остаток в км до конечной точки
	ETA       json.Number `json:"distance_forecast_time"` // Прогноз оставшегося времени до конечной точки
	// Прогноз строится в зависимости от наличия информации
	// о конечном пункте назначения и времени прибытия.
	// Если этой информации нет, значения элементов будут пустыми.
	ETAStatus string `json:"distance_forecast_status"` // Прогноз успеваемости до конечной точки.
	// Прогноз строится в зависимости от наличия информации
	// о конечном пункте назначения и времени прибытия.
	// Если этой информации нет, значения элементов будут пустыми.
}

func (op *ObjectPosition) LastTimeUpdated() time.Time {
	v, _ := op.Timestamp.Int64()
	return time.Unix(v, 0)
}

//type ObjectPositionsChannel <-chan ObjectPosition
//
//// Clear discards all unprocessed incoming elements of ObjectPosition.
//func (ch ObjectPositionsChannel) Clear() {
//	for len(ch) != 0 {
//		<-ch
//	}
//}

type OperatorInfo struct {
	Operator Operator `json:"operator"`
	Title    string   `json:"title"`
	Region   string   `json:"region"`
}
