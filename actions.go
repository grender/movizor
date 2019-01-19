package movizor

import (
	"fmt"
	"net/url"
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

//type SchedulingOptions struct {
//	// sw1 string Включить расписание на понедельник
//	// sw2 string Включить расписание на вторник
//	// sw3 string Включить расписание на среду
//	// sw4 string Включить расписание на четверг
//	// sw5 string Включить расписание на пятницу
//	// sw6 string Включить расписание на субботу
//	// sw7 string Включить расписание на воскресенье
//	FireAt []time.Time // st Массив времени в расписании. Передается в многомерном массиве, каждый вложенный элемент является временем для срабатывания расписания в формате hh:mm
//}

type ObjectOptions struct {
	Title   string     //title - Название объекта
	Tags    []string   //tags - Список меток через запятую
	DateOff time.Time  //dateoff - Дата и время автоматического отключения абонента (dd.mm.yyyy hh:mm:ss)
	Tariff  TariffType //tariff - Id-тарифного плана
	//PackageProlong string     //package_prolong - Автоматически продлевать пакет (при использовании пакетного тарифа)
	//Destinations   []DestinationOptions
	//Schedules      SchedulingOptions
	Metadata map[string]string // metadata Массив с дополнительной информацией по объекту для отображения в событиях и отчетах.
	// Каждый элемент обозначает одну запись метаинформации для объекта. Имена элементов и значения произвольные:
	// metadata[Исполнитель] = Петров;
	// metadata[Склад] = Восточный;
	// autoinform integer Включить услугу автоинформатора.
}

func (o *ObjectOptions) addValuesTo(v *url.Values) {
	if o.Title != "" {
		v.Add("title", o.Title)
	}
	if !o.DateOff.IsZero() {
		v.Add("dateoff", o.DateOff.Format("02.01.2006 15:04:05"))
	}
	if o.Tariff != "" {
		v.Add("tariff", string(o.Tariff))
	}
	if len(o.Tags) > 0 {
		v.Add("tags", strings.Join(o.Tags, ","))
	}
	for key, val := range o.Metadata {
		v.Add(fmt.Sprintf("metadata[%s]", key), val)
	}

	return
}

//type ObjectAddOptions struct {
//	// account integer Идентификатор аккаунта подчинённого кабинета в который добавляется объект.
//}
//
//type ObjectEditOptions struct {
//	// activate string Сразу активировать новый тариф
//}
