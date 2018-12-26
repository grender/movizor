package movizor

import (
	"net/url"
)

type ObjectAddOptions struct {
	Title string //title - Название объекта
	//Tags []string //tags - Список меток через запятую
	DateOff string     //dateoff - Дата и время автоматического отключения абонента (dd.mm.yyyy hh:mm:ss)
	Tariff  TariffType //tariff - Id-тарифного плана
	// PackageProlong string //package_prolong - Автоматически продлевать пакет (при использовании пакетного тарифа)
	// Destinations array [string] массив конечных точек маршрута.
	// Передается в многомерном массиве, каждый вложенный массив обозначает одну конечную (или промежуточную) точку:
	// destination[0] = array(text = Казань, // название пункта coord = 55.760419,49.190294, // координаты (lat,lon) time = 10.05.2016 18:00, // дата и время прибытия (dd.mm.yyyy hh:mm);
	// destination[1] = array(text = Москва, coord = 55.7098009,37.0536908, time = 12.05.2016 18:00,);
	// sw1 string Включить расписание на понедельник
	// sw2 string Включить расписание на вторник
	// sw3 string Включить расписание на среду
	// sw4 string Включить расписание на четверг
	// sw5 string Включить расписание на пятницу
	// sw6 string Включить расписание на субботу
	// sw7 string Включить расписание на воскресенье
	// st[] array [string] Массив времени в расписании. Передается в многомерном массиве, каждый вложенный элемент является временем для срабатывания расписания в формате hh:mm:
	// metadata[] array [string] массив с дополнительной информацией по объекту для отображения в событиях и отчетах. Каждый элемент обозначает одну запись метаинформации для объекта. Имена элементов и значения произвольные:
	// metadata[Исполнитель] = Петров;
	// metadata[Склад] = Восточный;
	// autoinform integer Включить услугу автоинформатора.
	// account integer Идентификатор аккаунта подчинённого кабинета в который добавляется объект.
}

func (oa *ObjectAddOptions) addValuesTo(v *url.Values) {
	if oa.Title != "" {
		v.Add("title", oa.Title)
	}
	if oa.DateOff != "" {
		v.Add("dateoff", oa.DateOff)
	}
	if oa.Tariff != "" {
		v.Add("tariff", string(oa.Tariff))
	}

	return
}

type ObjectEditOptions struct {
	Title string //title - Название объекта
	//Tags []string //tags - Список меток через запятую
	DateOff string     //dateoff - Дата и время автоматического отключения абонента (dd.mm.yyyy hh:mm:ss)
	Tariff  TariffType //tariff - Id-тарифного плана
	// PackageProlong string //package_prolong - Автоматически продлевать пакет (при использовании пакетного тарифа)
	// Destinations array [string] массив конечных точек маршрута.
	// Передается в многомерном массиве, каждый вложенный массив обозначает одну конечную (или промежуточную) точку:
	// destination[0] = array(text = Казань, // название пункта coord = 55.760419,49.190294, // координаты (lat,lon) time = 10.05.2016 18:00, // дата и время прибытия (dd.mm.yyyy hh:mm);
	// destination[1] = array(text = Москва, coord = 55.7098009,37.0536908, time = 12.05.2016 18:00,);
	// activate string Сразу активировать новый тариф
	// sw1 string Включить расписание на понедельник
	// sw2 string Включить расписание на вторник
	// sw3 string Включить расписание на среду
	// sw4 string Включить расписание на четверг
	// sw5 string Включить расписание на пятницу
	// sw6 string Включить расписание на субботу
	// sw7 string Включить расписание на воскресенье
	// st[] array [string] Массив времени в расписании. Передается в многомерном массиве, каждый вложенный элемент является временем для срабатывания расписания в формате hh:mm:
	// metadata[] array [string] массив с дополнительной информацией по объекту для отображения в событиях и отчетах. Каждый элемент обозначает одну запись метаинформации для объекта. Имена элементов и значения произвольные:
	// metadata[Исполнитель] = Петров;
	// metadata[Склад] = Восточный;
	// autoinform integer Включить услугу автоинформатора.
}

func (oa *ObjectEditOptions) addValuesTo(v *url.Values) {
	if oa.Title != "" {
		v.Add("title", oa.Title)
	}
	if oa.DateOff != "" {
		v.Add("dateoff", oa.DateOff)
	}
	if oa.Tariff != "" {
		v.Add("tariff", string(oa.Tariff))
	}

	return
}
