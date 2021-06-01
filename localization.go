package main

import (
	"github.com/gin-gonic/gin"
)

// TerritoryCache - cache of all territories
var TerritoryCache = map[string]Territory{
	"edBuild1":        {"edBuild1", 3, false},
	"edBuild2":        {"edBuild2", 3, false},
	"edBuild3":        {"edBuild3", 3, false},
	"edBuild4":        {"edBuild4", 4, false},
	"edBuild5":        {"edBuild5", 2, false},
	"edBuild7":        {"edBuild7", 2, false},
	"edBuild7a":       {"edBuild7a", 1, false},
	"edBuild8":        {"edBuild8", 2, false},
	"edBuild9":        {"edBuild9", 2, false},
	"edBuild10":       {"edBuild10", 3, false},
	"edBuild11":       {"edBuild11", 4, false},
	"edBuild15":       {"edBuild15", 2, false},
	"edBuild17":       {"edBuild17", 2, false},
	"stadium":         {"stadium", 4, true},
	"candle":          {"candle", 2, false},
	"edBuild1Terr":    {"edBuild1Terr", 2, true},
	"miniPitch":       {"miniPitch", 3, true},
	"military":        {"military", 2, true},
	"loveBench":       {"loveBench", 1, true},
	"studSquare":      {"studSquare", 3, true},
	"rectors":         {"rectors", 2, true},
	"areaOfFame":      {"areaOfFame", 3, true},
	"botanicalGarden": {"botanicalGarden", 5, true},
	"parkArea":        {"parkArea", 3, true},
	"kulin":           {"kulin", 1, true},
}

// GetAllowedBuildings - gets buildings that can be placed on territory
func GetAllowedBuildings(tid string) []Building {
	answer := []Building{}
	ter, ok := TerritoryCache[tid]
	if !ok || !ter.Allow {
		return answer
	}

	for _, info := range BuildingCache {
		if info.BSize <= ter.TSize && (info.Added || info.BID == tid) {
			answer = append(answer, info)
		}
	}
	return answer
}

// BuildingCache - cache of all buildings
var BuildingCache = map[string]Building{
	"edBuild1":        {"edBuild1", false, 3, []BuildingInfoType{{"Корпус №1", "Тут знаходиться опис першого корпусу!"}, {"Educational building №1", "Here is a description of the first building!"}}},
	"edBuild2":        {"edBuild2", false, 3, []BuildingInfoType{{"Корпус №2", "Опис №2"}, {"Educational building №2", "Description №2"}}},
	"edBuild3":        {"edBuild3", false, 3, []BuildingInfoType{{"Корпус №3", "Тут знаходиться опис третього корпусу!"}, {"Educational building №3", "Here is a description of the third building!"}}},
	"edBuild4":        {"edBuild4", false, 4, []BuildingInfoType{{"Корпус №4", "Опис №4"}, {"Educational building №4", "Description №4"}}},
	"edBuild5":        {"edBuild5", false, 2, []BuildingInfoType{{"Корпус №5", "Опис №5"}, {"Educational building №5", "Description №5"}}},
	"edBuild7":        {"edBuild7", false, 2, []BuildingInfoType{{"Корпус №7", "Опис №7"}, {"Educational building №7", "Description №7"}}},
	"edBuild7a":       {"edBuild7a", false, 1, []BuildingInfoType{{"Корпус №7-а", "Опис №7-a"}, {"Educational building №7-а", "Description №7-a"}}},
	"edBuild8":        {"edBuild8", false, 2, []BuildingInfoType{{"Корпус №8", "Опис №8"}, {"Educational building №8", "Description №8"}}},
	"edBuild9":        {"edBuild9", false, 2, []BuildingInfoType{{"Корпус №9", "Опис №9"}, {"Educational building №9", "Description №9"}}},
	"edBuild10":       {"edBuild10", false, 3, []BuildingInfoType{{"Корпус №10", "Опис №10"}, {"Educational building №10", "Description №10"}}},
	"edBuild11":       {"edBuild11", false, 4, []BuildingInfoType{{"Корпус №11", "Опис №11"}, {"Educational building №11", "Description №11"}}},
	"edBuild15":       {"edBuild15", false, 2, []BuildingInfoType{{"Корпус №15", "Опис №15"}, {"Educational building №15", "Description №15"}}},
	"edBuild17":       {"edBuild17", false, 2, []BuildingInfoType{{"Корпус №17", "Опис №17"}, {"Educational building №17", "Description №17"}}},
	"stadium":         {"stadium", false, 4, []BuildingInfoType{{"Стадіон", "Фішка стадіону НУБіП - в екологічному розташуванні в одному з найзеленіших місць Києва. Тиша, чисте повітря, дерева навколо, старе покриття в досить непоганому стані і м'яка трава на футбольному полі, на якій так приємно повалятися після тренування."}, {"Stadium", "The feature of the NUBiP stadium is its environmentally friendly location in one of the greenest places in Kiev. Silence, clean air, trees around, the old coating is in good enough condition and soft grass on the football field, which is so nice to lie on after training."}}},
	"candle":          {"candle", false, 2, []BuildingInfoType{{"Свічка", "Опис Свічка"}, {"Candle", "Description Candle"}}},
	"edBuild1Terr":    {"edBuild1Terr", false, 2, []BuildingInfoType{{"Площа перед корпусом №1", "Опис Площа перед корпусом №1"}, {"Area in front of the building №1", "Description Площа перед корпусом №1"}}},
	"miniPitch":       {"miniPitch", false, 3, []BuildingInfoType{{"Міні поле", "Опис Міні поле"}, {"Mini football pitch", "Description miniPitch"}}},
	"military":        {"military", false, 2, []BuildingInfoType{{"Військова кафедра", "Опис Військова кафедра"}, {"Military Department", "Description military"}}},
	"loveBench":       {"loveBench", false, 1, []BuildingInfoType{{"Лава закоханих", "Опис Лава закоханих"}, {"Love-bench", "Description loveBench"}}},
	"studSquare":      {"studSquare", false, 3, []BuildingInfoType{{"Студентський сквер", "Опис Студентський сквер"}, {"Students' square", "Description studSquare"}}},
	"rectors":         {"rectors", false, 2, []BuildingInfoType{{"Алея ректорів", "Опис Алея ректорів"}, {"Alley of rectors", "Description rectors"}}},
	"areaOfFame":      {"areaOfFame", false, 3, []BuildingInfoType{{"Площа слави", "Опис Площа слави"}, {"Area of fame", "Description areaOfFame"}}},
	"botanicalGarden": {"botanicalGarden", false, 5, []BuildingInfoType{{"Ботанічний сад", "Ботанічний сад Національного університету біоресурсів та природокористування України — частина Голосіївського лісу, що є навчальним підрозділом Національного університету біоресурсів і природокористування України."}, {"Botanical garden", "Description botanicalGarden"}}},
	"parkArea":        {"parkArea", false, 3, []BuildingInfoType{{"Паркова зона", "Опис Паркова зона"}, {"Park area", "Description parkArea"}}},
	"kulin":           {"kulin", false, 1, []BuildingInfoType{{"Кулиничі біля корпусу №3", "Опис Кулиничі біля корпусу №3"}, {"Kulinichi near the building №3", "Description kulin2"}}},
	// other buildings
	"trc":     {"trc", true, 4, []BuildingInfoType{{"Торгово-розважальний центр", "Торговий центр (ТЦ) інколи торгово-розважальний центр (ТРЦ), галерея, пасаж, торговельний комплекс, торговельно-розважальний комплекс - універсальна крамниця чи комплекс крамниць, що зазвичай включає підприємства побутового обслуговування, громадського харчування та розважальні заклади."}, {"Shopping center", "A shopping center (mall) is sometimes a shopping and entertainment center (mall), gallery, passage, shopping mall, shopping and entertainment complex - a department store or complex of stores, which usually includes consumer services, catering and entertainment establishments."}}},
	"auchan":  {"auchan", true, 3, []BuildingInfoType{{"Ашан", "Опис Ашан"}, {"Auchan", "Description Auchan"}}},
	"atb":     {"atb", true, 2, []BuildingInfoType{{"АТБ", "Опис АТБ"}, {"АТБ", "Description ATB"}}},
	"fora":    {"fora", true, 2, []BuildingInfoType{{"Фора", "Опис Фора"}, {"Фора", "Description Fora"}}},
	"silpo":   {"silpo", true, 2, []BuildingInfoType{{"Сільпо", "Опис Сільпо"}, {"Сільпо", "Description Silpo"}}},
	"theatre": {"theatre", true, 3, []BuildingInfoType{{"Театр", "Опис Театр"}, {"Theatre", "Description Theatre"}}},
	"cinema":  {"cinema", true, 3, []BuildingInfoType{{"Кінотеатр", "Опис Кінотеатр"}, {"Cinema", "Description Cinema"}}},
	"sakura":  {"sakura", true, 2, []BuildingInfoType{{"Алея сакури", "Опис Алея сакури"}, {"Alley of sakura", "Description Alley of sakura"}}},
	"empty":   {"empty", true, 0, []BuildingInfoType{{"Пустка", "Пуста територія: лиш трава, каміння, поодинокі польові квіти."}, {"Barren", "Empty territory: only grass, stones, single wildflowers."}}},
}

// // PlayerBuildings - cache for standart objects. Key got from URL
// var PlayerBuildings = map[string][]BuildingInfoType{}

// // FillPlayerBuildings - fills map of building info
// func FillPlayerBuildings() {
// 	for i, v := range TerritoryCache {
// 		PlayerBuildings[v.TID] = BuildingCache[i].Info
// 	}
// 	if LoggedIn {
// 		// TODO

// 	}
// }

// func getBuildingInfo() {
// 	if !LoggedIn {

// 		return
// 	}
// }

// GinLabels - cache for labels on page
var GinLabels = map[string][]string{
	"back":         {"Назад", "Back"},
	"close":        {"Закрити", "Close"},
	"main":         {"Повернутись на головну", "Go to Main page"},
	"showAreas":    {"Показати всі об'єкти", "Show all objects"},
	"reg":          {"Реєстрація", "Registration"},
	"signIn":       {"Вхід", "Sign in"},
	"logout":       {"Вихід", "Logout"},
	"email":        {"Пошта", "Email"},
	"login":        {"Логін", "Login"},
	"password":     {"Пароль", "Password"},
	"remember":     {"Запам'ятати мене", "Remember me"},
	"changeTerr":   {"Змінити територію", "Change territory"},
	"send":         {"Відправити", "Send"},
	"choose":       {"Обрати", "Choose"},
	"other":        {"Інше", "Other"},
	"infographics": {"Інфографіка", "Infographics"},
	"updateCharts": {"Оновити діаграми", "Update charts"},
	"filterCharts": {"Фільтри", "Filters"},
	"bar":          {"Гістограма", "Bar"},
	"pie":          {"Кругова", "Pie"},
	"radar":        {"Радіолокаційна", "Radar"},
	"word":         {"Хмара слів", "Word Cloud"},
	"tech":         {"Технічні роботи", "Technical works"},
	"clear":        {"Очистити", "Clear"},
	"details":      {"Деталі", "Details"},
	"history":      {"Історія", "History"},
	"photo":        {"Фото", "Photo"},
}

// Errors - cache for error sentences
var Errors = map[string][]string{
	"errorTitle":             {"Помилка", "Error"},
	"territoryNotFound":      {"Обрана територія - не знайдена. Перевірте правильність посилання або спробуйте пізніше.", "Selected territory - not found. Please check that the link is correct or try again later."},
	"invalidLogin":           {"Помилка входу: Неправильний логін або пароль", "Login error: Invalid login or password"},
	"successfulLogin":        {"Успішний вхід", "Successful login"},
	"successfulRegistration": {"Успішна реєстрація", "Successful registration"},
	"emptyLoginField":        {"Помилка входу: Треба заповнити всі поля", "Login error: You need to fill all the fields"},
	"emptyRegField":          {"Помилка реєстрації: Треба заповнити всі поля", "Register error: You need to fill all the fields"},
	"notUniqueLogin":         {"Помилка реєстрації: Такий логін вже зайнятий", "Register error: This username is already used"},
	"notUniqueEmail":         {"Помилка реєстрації: Така пошта вже використовується", "Register error: This email is already used"},
	"invalidEmail":           {"Помилка реєстрації: Невірна адреса електронної пошти", "Register error: Invalid email address"},
	"unauth":                 {"Помилка: Дія доступна лише авторизованим користувачам", "Error: Action available only for authorized users"},
	"successfulTerrChanging": {"Успішна зміна території", "Successful territory changing"},
	"invalidTerrChanging":    {"Помилка зміни території: Будь-ласка спробуйте пізніше", "Territory changing error: Please try again later"},
	"internalError":          {"Помилка: На даний момент сервіс недоступний, будь-ласка спробуйте пізніше", "Error: The service isn`t available now, please try again later"},
	"chartClear":             {"Очищено", "Cleared"},
	"chartOk":                {"Вибірка здійснена", "Sampling performed"},
}

// Phrases - cache for phrases in survey and other sentences
var Phrases = map[string][]string{
	"surveyTitle":       {"Опитування", "Survey"},
	"infographicsTitle": {"Інфографіка", "Infographics"},
	"surThanks":         {"Дякую за те, що пройшли опитування, тепер Ви маєте доступ до інфографіки.", "Thank you for the passing the survey, you have access to the infographics now."},
	"surPhrase":         {"Ви нам дуже допоможете, якщо пройдете невеличке опитування. Це займе не більше хвилини Вашого часу!", "You will help us a lot, if you pass a small survey. It will take no more than a minute of your time!"},
	"surConf":           {"Інформація - повністю конфіденційна.", "The information is completely confidential."},
	"firstQ":            {"На якому курсі Ви навчаєтесь?", "Which year are you studying?"},
	"firstQ1":           {"Я абітурієнт", "I am an entrant"},
	"firstQ2":           {"Бакалавр 1-2", "Bachelor 1-2"},
	"firstQ3":           {"Бакалавр 3-4 / Скорочений термін навчання", "Bachelor 3-4 / Reduced training period"},
	"firstQ4":           {"Магістратура", "Magistracy"},
	"firstQ5":           {"Вже не навчаюсь", "I don't study anymore"},
	"firstQ6":           {"Я викладач", "I am a tutor"},
	"secondQ":           {"А як щодо університету?", "And what about university?"},
	"secondQ1":          {"Я абітурієнт", "I am an entrant"},
	"secondQ2":          {"Я викладач в НУБіП", "I am a tutor at NULES"},
	"secondQ3":          {"Я студент в НУБіП", "I am a student at NULES"},
	"secondQ4":          {"Раніше навчався в НУБІП", "Previously studied at NULES"},
	"secondQ5":          {"Інше", "Other"},
	"thirdQ":            {"Як би Ви оцінили територію кампусу (0 - 10)?", "How would you rate campus (0 - 10)?"},
	"fourthQ":           {"Як би Ви оцінили корпуси, їх зовнішній вигляд та внутрішнє обладнання (0 - 10)?", "How would you rate the buildings, their appearance and interior equipment (0 - 10)?"},
	"fifthQ":            {"Яку б частину Ви хотіли б змінити? (можна залишити пустим)", "Which part would you like to change? (can be left blank)"},
	"fifthQ1":           {"Лісову та ботанічний сад, вони потребують змін", "Forest and botanical garden, they need to be changed"},
	"fifthQ2":           {"Корпуси, їх потрібно оновити", "Buildings need to be updated"},
	"fifthQ3":           {"Територію між корпусами: алеї, сквери, дороги та інше", "The area between the buildings: alleys, squares, roads and other"},
	"sixthQ":            {"Чого не вистачає території біля кампусу? (можна залишити пустим)", "What is missing near the campus? (can be left blank)"},
	"sixthQ1":           {"Потрібно більше будівель розважального характеру (кінотеатри, музеї)", "Need more entertainment buildings (cinemas, museums)"},
	"sixthQ2":           {"Потрібно більше продовольчих магазинів (АТБ, Фора, Сільпо)", "Need more grocery stores (АТБ, Фора, Сільпо)"},
	"sixthQ3":           {"Потрібно більше магазинів з одягом та взуттям (LC Waikiki, New Yorker, O'Stin)", "Need more clothing and footwear stores (LC Waikiki, New Yorker, O'Stin)"},
	"sixthQ4":           {"Потрібно більше аптек/магазинів з косметикою (АНЦ, Eva, Watsons)", "Need more pharmacies/cosmetics stores (АНЦ, Eva, Watsons)"},
	"sixthQ5":           {"Потрібен гіпермаркет (АШАН, Fozzy, METRO)", "Need a hypermarket (Auchan, Fozzy, METRO)"},
	"seventhQ":          {"Чи потребують змін гуртожитки НУБіП України?", "Do NULES dormitories need changes?"},
	"seventhQ1":         {"Не знаю / Не живу в гуртожику", "I don't know / I don't live in dormitory"},
	"seventhQ2":         {"Ні: Все добре", "Not, everything is OK"},
	"seventhQ3":         {"Скоріше ні, але існують деякі неприємні моменти", "Probably not, but there are some unpleasant moments"},
	"seventhQ4":         {"Так, гуртожитки потребують змін", "Yes, dormitories need changes"},
	"seventhQ5":         {"Обов'язково. Потрібні термінові радикальні зміни", "Necessarily. Urgent radical changes are needed"},
	"pieBasic":          {" Основна кругова діаграма", " Basic pie chart"},
	"barBasic":          {" Основна гістограма", " Basic bar chart"},
	"pieRose":           {" Трояндова кругова діаграма", " Rose pie chart"},
	//"radar":             {" Радіолокаційна діаграма", " Radar chart"},
	//"wordCloud":         {" Діаграма хмари слів", " Word cloud chart"},
	"infoFirstQ":     {"Яка стара територія найчастіше змінювана?", "Which old territory is the most often changed?"},
	"infoSecondQ":    {"Яка нова територія найчастіше встановлена?", "Which new territory is the most often established?"},
	"resultFirstQ":   {"Результати на 1 питання", "Results for the 1 question"},
	"resultSecondQ":  {"Результати на 2 питання", "Results for the 2 question"},
	"resultThirdQ":   {"Результати на 3 питання", "Results for the 3 question"},
	"resultFourthQ":  {"Результати на 4 питання", "Results for the 4 question"},
	"resultFifthQ":   {"Результати на 5 питання", "Results for the 5 question"},
	"resultSixthQ":   {"Результати на 6 питання", "Results for the 6 question"},
	"resultSeventhQ": {"Результати на 7 питання", "Results for the 7 question"},
}

func getLocalization(c *gin.Context) string {
	if cookieVal, err := c.Cookie("loc"); err == nil {
		return cookieVal
	}
	return "UA"
}

func getLocIndex() int {
	if Localization == "EN" {
		return 1
	}
	return 0
}

func getLoc(key string, fromMap map[string][]string) string {
	valWOIndex, ok := fromMap[key]
	if !ok {
		Error.Println("localization.go -> getLoc: !ok with key =", key)
		return ""
	}
	val := valWOIndex[getLocIndex()]

	return val
}

// History - cache of history
var History = map[string][]string{
	"edBuild1":        {"У 1926-1931 рр. зодчий Д. Дяченко проектує комплекс споруд Київського сільськогосподарського інституту (нині — Національний університет біоресурсів і природокористування України) в Голосіївському районі столиці. Дослідники відзначають високу художню якість споруди Лісотехнічного факультету (1926-1927).", "In 1926-1931, architect D. Dyachenko designed a complex of buildings of the Kyiv Agricultural Institute (now the National University of Life and Environmental Sciences of Ukraine) in the Holosiivskyi district of the capital. Researchers note the high artistic quality of the building of the Faculty of Forestry (1926-1927)."},
	"edBuild2":        {"", ""},
	"edBuild3":        {"Історія велика та цікава! Він навіть згорав до тла!", "History is big and interesting! It was even burned out!"},
	"edBuild4":        {"", ""},
	"edBuild5":        {"", ""},
	"edBuild7":        {"", ""},
	"edBuild7a":       {"", ""},
	"edBuild8":        {"", ""},
	"edBuild9":        {"", ""},
	"edBuild10":       {"", ""},
	"edBuild11":       {"", ""},
	"edBuild15":       {"", ""},
	"edBuild17":       {"", ""},
	"stadium":         {"14 травня 2019 року на стадіоні НУБіП України розпочались змагання з легкої атлетики 62-ї спартакіади студентів. Сонячна і тепла погода супроводжувала найкращих спортсменів факультетів і ННІ в перший змагальний день на дистанціях 100м, 400м, 1500м і стрибках в довжину з розбігу.", "On May 14, 2019, the athletics competitions of the 62nd Student Games began at the NULES Stadium in Ukraine. Sunny and warm weather accompanied the best athletes of faculties and research institutes on the first day of competition in the distances of 100m, 400m, 1500m and long jumps from the run."},
	"candle":          {"", ""},
	"edBuild1Terr":    {"", ""},
	"miniPitch":       {"", ""},
	"military":        {"", ""},
	"loveBench":       {"", ""},
	"studSquare":      {"", ""},
	"rectors":         {"", ""},
	"areaOfFame":      {"", ""},
	"botanicalGarden": {"29 грудня 1988 року ректор Української сільськогосподарської академії, академік Д. О. Мельничук підписав наказ № 410 «Про створення ботанічного саду Української сільськогосподарської академії», згідно з яким до нової структури увійшли землі Боярської лісової дослідної станції і сільгоспакадемії, зайняті плодовим садом, ділянки пасіки, навчально-дослідної лабораторії конярства. Статус Державного ботанічного сад отримав згідно з Постановою Ради Міністрів УРСР від 13.02.1989 р. № 53, а у 1992 році набув статусу Ботанічного саду загальнодержавного значення (Постанова Верховної Ради України від 16.06.1992 р. № 2457-ХІІ).", "On December 29, 1988, the rector of the Ukrainian Agricultural Academy, Academician DO Melnychuk signed Order № 410 «On the establishment of the Botanical Garden of the Ukrainian Agricultural Academy», according to which the new structure included the lands of the Boyarka Forest Research Station and Agricultural Academy occupied by orchards. , training and research laboratory of horse breeding. It received the status of the State Botanical Garden in accordance with the Resolution of the Council of Ministers of the USSR of 13.02.1989 № 53, and in 1992 acquired the status of the Botanical Garden of national importance (Resolution of the Verkhovna Rada of Ukraine of 16.06.1992 № 2457-XII)."},
	"parkArea":        {"", ""},
	"kulin":           {"", ""},
	"trc":             {"Перші торгові центри такого типу стали з'являтися в XIX столітті. Найбільшим у світі є «Мол ов Арабія»", "The first shopping centers of this type began to appear in the XIX century. The largest in the world is «Mol ov Arabia»"},
	"auchan":          {"", ""},
	"atb":             {"", ""},
	"fora":            {"", ""},
	"silpo":           {"", ""},
	"theatre":         {"", ""},
	"cinema":          {"", ""},
	"sakura":          {"", ""},
	"empty":           {"Пусті території завжди були, є та будуть.", "Empty territories have always been, are and will be."},
}

// Details - cache of Details
var Details = map[string][]string{
	"edBuild1":        {"В першому корпусі знаходиться приймальна комісія, через яку проходять всі абітурієнти, а надалі - студенти НУБіП України.", "In the first building there is an admission commission through which all entrants, and further - students of NULES of Ukraine pass."},
	"edBuild2":        {"", ""},
	"edBuild3":        {"Увага! В навчальному корпусі №3 на виніс працює їдальня. З 26 травня 2020 у навчальному корпусі №3 у режимі на виніс працює їдальня. Страви відпускаються в боксах. У меню салати, другі страви, гарніри та напої.", "Warning! There is a canteen in the take-away building №3. From May 26, 2020, a canteen will be open in the take-away building №3. Meals are served in boxes. The menu includes salads, main courses, side dishes and drinks."},
	"edBuild4":        {"", ""},
	"edBuild5":        {"", ""},
	"edBuild7":        {"", ""},
	"edBuild7a":       {"", ""},
	"edBuild8":        {"", ""},
	"edBuild9":        {"", ""},
	"edBuild10":       {"", ""},
	"edBuild11":       {"", ""},
	"edBuild15":       {"", ""},
	"edBuild17":       {"", ""},
	"stadium":         {"Поруч із спортивним корпусом знаходиться стадіон з футбольним полем (розмір поля 99 м на 70 м) легкоатлетичними біговими доріжками і трибунами на 1500 місць.", "Next to the sports building is a stadium with a football field (field size 99 m by 70 m) with athletics treadmills and grandstands for 1500 seats."},
	"candle":          {"", ""},
	"edBuild1Terr":    {"", ""},
	"miniPitch":       {"", ""},
	"military":        {"", ""},
	"loveBench":       {"", ""},
	"studSquare":      {"", ""},
	"rectors":         {"", ""},
	"areaOfFame":      {"", ""},
	"botanicalGarden": {"Ботанічний сад складається з 6 наукових лабораторій: дендрології та лісової селекції, плодово-овочевих культур, квітникарства, екології рослин, зоології, зеленого будівництва. Загальна кількість видів, різновидів, гібридів, культиварів та форм які зростають на території Ботанічного саду становить 1499 таксономічних одиниці. З них колекція деревних рослин незахищеного ґрунту становить 604 колекційних одиниці (389 види, 4 різновиди, 33 гібриди, 2 форми та 176 культиварів), колекція лікарських рослин — 164 одиниці (163 види, 1 гібрид), колекція квітково-декоративних рослин відкритого ґрунту — 154 одиниці (95 види, 21 гібрид та 38 культиварів), колекція плодово-ягідних рослин — 155 одиниць (27 видів, 2 гібриди та 126 культиварів), колекція рослин захищеного ґрунту — 184 одиниці (162 види, 4 різновиди, 5 гібридів та 13 культиварів), колекція дикорослих та адвентивних трав'янистих рослин — 307 одиниць (306 видів, 1 гібрид), колекція мохоподібних рослин представлена 50 видами.", "The Botanical Garden consists of 6 scientific laboratories: dendrology and forest breeding, fruit and vegetable crops, floriculture, plant ecology, zoology, green building. The total number of species, varieties, hybrids, cultivars and forms that grow in the Botanical Garden is 1499 taxonomic units. Of these, the collection of woody plants of unprotected soil is 604 collection units (389 species, 4 varieties, 33 hybrids, 2 forms and 176 cultivars), the collection of medicinal plants - 164 units (163 species, 1 hybrid), the collection of flowering and ornamental plants of open ground - 154 units (95 species, 21 hybrids and 38 cultivars), collection of fruit and berry plants - 155 units (27 species, 2 hybrids and 126 cultivars), collection of protected soil plants - 184 units (162 species, 4 varieties, 5 hybrids and 13 cultivars), the collection of wild and adventitious herbaceous plants - 307 units (306 species, 1 hybrid), the collection of moss-like plants is represented by 50 species."},
	"parkArea":        {"", ""},
	"kulin":           {"", ""},
	"trc":             {"Сучасний торговий центр може бути великим торговельно-розважальним комплексом — багатоповерхова будівля в стилі гай-тек, в якому окрім крамниць можуть знаходитися також кав'ярні, бари, казино, кінотеатр. Як правило, комплекс обладнаний ескалаторами, забезпечений безкоштовним парковкуванням для особистого транспорту покупців, і розташований біля станцій метро і зупинок громадського транспорту.", "A modern shopping center can be a large shopping and entertainment complex - a high-rise building in the style of high-tech, which in addition to shops can also be coffee shops, bars, casinos, cinemas. As a rule, the complex is equipped with escalators, provided with free parking for personal transport of buyers, and is located near metro stations and public transport stops."},
	"auchan":          {"", ""},
	"atb":             {"", ""},
	"fora":            {"", ""},
	"silpo":           {"", ""},
	"theatre":         {"", ""},
	"cinema":          {"", ""},
	"sakura":          {"", ""},
	"empty":           {"Це птах? Це літак? Це Супермен? Ні! Це просто камінь на траві, бо тут нічого нема.", "Is this a bird? Is this a plane? Is this Superman? No! It's just a stone on the grass, because there's nothing here."},
}
