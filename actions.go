package movizor

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Destinations array [string] массив конечных точек маршрута.
// Передается в многомерном массиве, каждый вложенный массив обозначает одну конечную (или промежуточную) точку:
// destination[0] = array(text = Казань, // название пункта coord = 55.760419,49.190294, // координаты (lat,lon) time = 10.05.2016 18:00, // дата и время прибытия (dd.mm.yyyy hh:mm);
// destination[1] = array(text = Москва, coord = 55.7098009,37.0536908, time = 12.05.2016 18:00,);
//type DestinationOptions struct {
//	Text         string
//	Lon          float32
//	Lat          float32
//	ExpectedTime time.Time
//}

// Расписание запросов на определение координат объекта
type SchedulingOptions struct {
	weekdays [7]bool
	FireAt   []time.Time // st Массив времени в расписании. Передается в многомерном массиве, каждый вложенный элемент является временем для срабатывания расписания в формате hh:mm
}

func (s SchedulingOptions) MondayOn() {
	s.weekdays[0] = true
}

func (s SchedulingOptions) TuesdayOn() {
	s.weekdays[1] = true
}

func (s SchedulingOptions) WednesdayOn() {
	s.weekdays[2] = true
}

func (s SchedulingOptions) ThursdayOn() {
	s.weekdays[3] = true
}

func (s SchedulingOptions) FridayOn() {
	s.weekdays[4] = true
}

func (s SchedulingOptions) SaturdayOn() {
	s.weekdays[5] = true
}

func (s SchedulingOptions) SundayOn() {
	s.weekdays[6] = true
}

func (s SchedulingOptions) MondayOff() {
	s.weekdays[0] = false
}

func (s SchedulingOptions) TuesdayOff() {
	s.weekdays[1] = false
}

func (s SchedulingOptions) WednesdayOff() {
	s.weekdays[2] = false
}

func (s SchedulingOptions) ThursdayOff() {
	s.weekdays[3] = false
}

func (s SchedulingOptions) FridayOff() {
	s.weekdays[4] = false
}

func (s SchedulingOptions) SaturdayOff() {
	s.weekdays[5] = false
}

func (s SchedulingOptions) SundayOff() {
	s.weekdays[6] = false
}

func (s SchedulingOptions) IsMondayOn() bool {
	return s.weekdays[0]
}

func (s SchedulingOptions) IsTuesdayOn() bool {
	return s.weekdays[1]
}

func (s SchedulingOptions) IsWednesdayOn() bool {
	return s.weekdays[2]
}

func (s SchedulingOptions) IsThursdayOn() bool {
	return s.weekdays[3]
}

func (s SchedulingOptions) IsFridayOn() bool {
	return s.weekdays[4]
}

func (s SchedulingOptions) IsSaturdayOn() bool {
	return s.weekdays[5]
}

func (s SchedulingOptions) IsSundayOn() bool {
	return s.weekdays[6]
}

func (s *SchedulingOptions) addValuesTo(v *url.Values) error {
	if len(s.FireAt) <= 0 {
		return errors.New("time to fire scheduling is not set (set FireAt property)")
	}

	if v == nil {
		v = new(url.Values)
	}

	// sw1 string Включить расписание на понедельник
	// sw2 string Включить расписание на вторник
	// sw3 string Включить расписание на среду
	// sw4 string Включить расписание на четверг
	// sw5 string Включить расписание на пятницу
	// sw6 string Включить расписание на субботу
	// sw7 string Включить расписание на воскресенье

	chk := false
	for idx, val := range s.weekdays {
		if val {
			chk = true
			v.Add(fmt.Sprintf("sw%d", idx+1), "1")
		}
	}
	if !chk {
		return errors.New("no single weekday to fire scheduling is set")
	}

	for _, val := range s.FireAt {
		v.Add("st[]", val.Format("15:04"))
	}
	return nil
}

// ObjectOptions предоставляет опции для AddObject (add_object) и EditObject (edit_object)
type ObjectOptions struct {
	Title          string     //title - Название объекта
	Tags           []string   //tags - Список меток через запятую
	DateOff        time.Time  //dateoff - Дата и время автоматического отключения абонента (dd.mm.yyyy hh:mm:ss)
	Tariff         TariffType //tariff - Id-тарифного плана
	PackageProlong bool       //package_prolong - Автоматически продлевать пакет (при использовании пакетного тарифа)
	//Destinations   []DestinationOptions // destination[] - массив конечных точек маршрута.
	Schedules    SchedulingOptions
	Metadata     map[string]string // metadata Массив с дополнительной информацией по объекту для отображения в событиях и отчетах.
	CallToDriver bool              // autoinform integer Включить услугу автоинформатора.
}

func (o *ObjectOptions) addValuesTo(v *url.Values) {
	if v == nil {
		v = new(url.Values)
	}
	if o.Title != "" {
		v.Add("title", o.Title)
	}
	if !o.DateOff.IsZero() {
		v.Add("dateoff", o.DateOff.Format("02.01.2006 15:04:05"))
	}
	if o.Tariff != "" {
		v.Add("tariff", string(o.Tariff))
	}
	if o.PackageProlong {
		v.Add("package_prolong", "1")
	}
	if len(o.Tags) > 0 {
		v.Add("tags", strings.Join(o.Tags, ","))
	}
	// Каждый элемент обозначает одну запись метаинформации для объекта. Имена элементов и значения произвольные:
	// metadata[Исполнитель] = Петров;
	// metadata[Склад] = Восточный;
	for key, val := range o.Metadata {
		v.Add(fmt.Sprintf("metadata[%s]", key), val)
	}
	if o.CallToDriver {
		v.Add("autoinform", "1")
	}
	// ToDo:добавить обработку ошибки
	_ = o.Schedules.addValuesTo(v)

	return
}

//type ObjectAddOptions struct {
//	// account integer Идентификатор аккаунта подчинённого кабинета в который добавляется объект.
//}
//
//type ObjectEditOptions struct {
//	// activate string Сразу активировать новый тариф
//}

// ObjectEventsOptions предоставляет опции для GetEvents.
type ObjectEventsOptions struct {
	RequestLimit uint64
	AfterEventID uint64
}

func (eo ObjectEventsOptions) values() url.Values {
	v := url.Values{}
	if eo.RequestLimit != 0 {
		v.Add("req_limit", strconv.FormatUint(eo.RequestLimit, 10))
	}
	if eo.AfterEventID != 0 {
		v.Add("afterid", strconv.FormatUint(eo.AfterEventID, 10))
	}
	return v
}

// SubscribeEventOptions предоставляет опции подписки на ноификацию по событиям. Если установлен признак AllPhones,
// то список Phones игнорируется.
type SubscribeEventOptions struct {
	AllPhones bool
	Phones    []Object
	Event     EventType
	notifyTo  NotificationType
	smsPhone  Object
	email     string
}

// SetSMSNotification устанавливает нотификацию на указанный телефон по СМС. Работает только та нотификация,
// которая была установлена последней в данной подписке. Это особенности API Movizor.
func (se SubscribeEventOptions) SetSMSNotification(phone Object) error {
	se.notifyTo = SMSNotification
	// ToDo: Переписать на что-то более надежное
	if phone.String() == "" {
		return fmt.Errorf("invalid phone number %s", string(phone))
	}

	se.smsPhone = phone
	return nil
}

// SetEMailNotification устанавливает нотификацию на указанный почтовый адрес. Работает только та нотификация,
// которая была установлена последней в данной подписке. Это особенности API Movizor.
func (se SubscribeEventOptions) SetEMailNotification(mail string) error {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !re.MatchString(mail) {
		return fmt.Errorf("%s is not valid email address", mail)
	}
	se.notifyTo = EMailNotification
	se.email = mail
	return nil
}

// SetTelegramNotification устанавливает нотификацию на Телеграм указанные в профиле держателя аккаута.
// Работает только та нотификация, которая была установлена последней в данной подписке. Это особенности API Movizor.
func (se SubscribeEventOptions) SetTelegramNotification() {
	se.notifyTo = TelegramNotification
}

func (se SubscribeEventOptions) values() (url.Values, error) {
	if !se.AllPhones && len(se.Phones) == 0 {
		return url.Values{}, errors.New("no single phone is set to subscribe for event")
	}
	if string(se.Event) == "" {
		return url.Values{}, errors.New("event to subscribe is not set")
	}
	if string(se.notifyTo) == "" {
		return url.Values{}, errors.New("notification type is not set")
	}

	v := url.Values{}

	if se.AllPhones {
		v.Add("phones_all", "1")
	} else {
		for _, val := range se.Phones {
			v.Add("phones[]", val.String())

		}
	}

	v.Add("events", string(se.Event))
	v.Add("notify_type", string(se.notifyTo))
	switch se.notifyTo {
	case SMSNotification:
		v.Add("notify_value", se.smsPhone.String())
	case EMailNotification:
		v.Add("notify_value", se.email)
	}

	return v, nil
}
