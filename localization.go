package main

import "github.com/gin-gonic/gin"

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
	"edBuild1":     {"edBuild1", false, 3, []BuildingInfoType{{"Корпус №1", "Опис №1"}, {"Educational building №1", "Description №1"}}},
	"edBuild2":     {"edBuild2", false, 3, []BuildingInfoType{{"Корпус №2", "Опис №2"}, {"Educational building №2", "Description №2"}}},
	"edBuild3":     {"edBuild3", false, 3, []BuildingInfoType{{"Корпус №3", "Опис №3"}, {"Educational building №3", "Description №3"}}},
	"edBuild4":     {"edBuild4", false, 4, []BuildingInfoType{{"Корпус №4", "Опис №4"}, {"Educational building №4", "Description №4"}}},
	"edBuild5":     {"edBuild5", false, 2, []BuildingInfoType{{"Корпус №5", "Опис №5"}, {"Educational building №5", "Description №5"}}},
	"edBuild7":     {"edBuild7", false, 2, []BuildingInfoType{{"Корпус №7", "Опис №7"}, {"Educational building №7", "Description №7"}}},
	"edBuild7a":    {"edBuild7a", false, 1, []BuildingInfoType{{"Корпус №7-а", "Опис №7-a"}, {"Educational building №7-а", "Description №7-a"}}},
	"edBuild8":     {"edBuild8", false, 2, []BuildingInfoType{{"Корпус №8", "Опис №8"}, {"Educational building №8", "Description №8"}}},
	"edBuild9":     {"edBuild9", false, 2, []BuildingInfoType{{"Корпус №9", "Опис №9"}, {"Educational building №9", "Description №9"}}},
	"edBuild10":    {"edBuild10", false, 3, []BuildingInfoType{{"Корпус №10", "Опис №10"}, {"Educational building №10", "Description №10"}}},
	"edBuild11":    {"edBuild11", false, 4, []BuildingInfoType{{"Корпус №11", "Опис №11"}, {"Educational building №11", "Description №11"}}},
	"edBuild15":    {"edBuild15", false, 2, []BuildingInfoType{{"Корпус №15", "Опис №15"}, {"Educational building №15", "Description №15"}}},
	"edBuild17":    {"edBuild17", false, 2, []BuildingInfoType{{"Корпус №17", "Опис №17"}, {"Educational building №17", "Description №17"}}},
	"stadium":      {"stadium", false, 4, []BuildingInfoType{{"Стадіон", "Опис Стадіон"}, {"Stadium", "Description Stadium"}}},
	"candle":       {"candle", false, 2, []BuildingInfoType{{"Свічка", "Опис Свічка"}, {"Candle", "Description Candle"}}},
	"edBuild1Terr": {"edBuild1Terr", false, 2, []BuildingInfoType{{"Площа перед корпусом №1", "Опис Площа перед корпусом №1"}, {"Area in front of the building №1", "Description Площа перед корпусом №1"}}},
	"miniPitch":    {"miniPitch", false, 3, []BuildingInfoType{{"Міні поле", "Опис Міні поле"}, {"Mini football pitch", "Description miniPitch"}}},
	"military":     {"military", false, 2, []BuildingInfoType{{"Військова кафедра", "Опис Військова кафедра"}, {"Military Department", "Description military"}}},
	"loveBench":    {"loveBench", false, 1, []BuildingInfoType{{"Лава закоханих", "Опис Лава закоханих"}, {"Love-bench", "Description loveBench"}}},
	"studSquare":   {"studSquare", false, 3, []BuildingInfoType{{"Студентський сквер", "Опис Студентський сквер"}, {"Students' square", "Description studSquare"}}},
	"rectors":      {"rectors", false, 2, []BuildingInfoType{{"Алея ректорів", "Опис Алея ректорів"}, {"Alley of rectors", "Description rectors"}}},
	"areaOfFame":   {"areaOfFame", false, 3, []BuildingInfoType{{"Площа слави", "Опис Площа слави"}, {"Area of fame", "Description areaOfFame"}}},
	"botanicalGarden": {"botanicalGarden", false, 5, []BuildingInfoType{{"Ботанічний сад",
		"Ботанічний сад Національного університету біоресурсів та природокористування України — частина Голосіївського лісу, що є навчальним підрозділом Національного університету біоресурсів і природокористування України.",
	}, {"Botanical garden", "Description botanicalGarden"}}},
	"parkArea": {"parkArea", false, 3, []BuildingInfoType{{"Паркова зона", "Опис Паркова зона"}, {"Park area", "Description parkArea"}}},
	"kulin":    {"kulin", false, 1, []BuildingInfoType{{"Кулиничі біля корпусу №3", "Опис Кулиничі біля корпусу №3"}, {"Kulinichi near the building №3", "Description kulin2"}}},
	// other buildings
	"trc":     {"trc", true, 4, []BuildingInfoType{{"Торгово-розважальний центр", "Опис Торгово-розважальний центр"}, {"Shopping center", "Description Shopping center"}}},
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
	"firstQ3":           {"Бакалавр 3-4", "Bachelor 3-4"},
	"firstQ4":           {"Магістратура", "Magistracy"},
	"firstQ5":           {"Вже не навчаюсь", "I don't study anymore"},
	"firstQ6":           {"Я викладач", "I am a tutor"},
	"secondQ":           {"А як щодо університету?", "And what about university?"},
	"secondQ1":          {"Я абітурієнт", "I am an entrant"},
	"secondQ2":          {"Я викладач в НУБіП", "I am a tutor at NULES"},
	"secondQ3":          {"Я студент в НУБіП", "I am a student at NULES"},
	"secondQ4":          {"Раніше навчався в НУБІП", "Previously studied at NULES"},
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
	val := fromMap[key][getLocIndex()]
	return val
}
