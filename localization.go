package main

import "github.com/gin-gonic/gin"

// TerritoryCache - cache of all territories
var TerritoryCache = []Territory{
	{"edBuild1", 3, false},
	{"edBuild2", 3, false},
	{"edBuild3", 3, false},
	{"edBuild4", 4, false},
	{"edBuild5", 2, false},
	{"edBuild7", 2, false},
	{"edBuild7a", 1, false},
	{"edBuild8", 2, false},
	{"edBuild9", 2, false},
	{"edBuild10", 3, false},
	{"edBuild11", 4, false},
	{"edBuild15", 2, false},
	{"edBuild17", 2, false},
	{"stadium", 4, true},
	{"candle", 2, false},
	{"edBuild1Terr", 2, true},
	{"miniPitch", 3, true},
	{"military", 2, false},
	{"loveBench", 1, true},
	{"studSquare", 3, true},
	{"ecoMarket", 1, false},
	{"rectors", 2, true},
	{"areaOfFame", 3, true},
	{"botanicalGarden", 5, true},
	{"parkArea", 3, true},
	{"kulin1", 1, true},
	{"kulin2", 1, true},
}

// BuildingInfoCache - cache of info of all buildings
var BuildingInfoCache = [][]BuildingInfoType{
	{{"Корпус №1", "Опис №1"}, {"Educational building №1", "Description №1"}},
	{{"Корпус №2", "Опис №2"}, {"Educational building №2", "Description №2"}},
	{{"Корпус №3", "Опис №3"}, {"Educational building №3", "Description №3"}},
	{{"Корпус №4", "Опис №4"}, {"Educational building №4", "Description №4"}},
	{{"Корпус №5", "Опис №5"}, {"Educational building №5", "Description №5"}},
	{{"Корпус №7", "Опис №7"}, {"Educational building №7", "Description №7"}},
	{{"Корпус №7-а", "Опис №7-a"}, {"Educational building №7-а", "Description №7-a"}},
	{{"Корпус №8", "Опис №8"}, {"Educational building №8", "Description №8"}},
	{{"Корпус №9", "Опис №9"}, {"Educational building №9", "Description №9"}},
	{{"Корпус №10", "Опис №10"}, {"Educational building №10", "Description №10"}},
	{{"Корпус №11", "Опис №11"}, {"Educational building №11", "Description №11"}},
	{{"Корпус №15", "Опис №15"}, {"Educational building №15", "Description №15"}},
	{{"Корпус №17", "Опис №17"}, {"Educational building №17", "Description №17"}},
	{{"Стадіон", "Опис Стадіон"}, {"Stadium", "Description Stadium"}},
	{{"Свічка", "Опис Свічка"}, {"Candle", "Description Candle"}},
	{{"Площа перед корпусом №1", "Опис Площа перед корпусом №1"}, {"Area in front of the building №1", "Description Площа перед корпусом №1"}},
	{{"Міні поле", "Опис Міні поле"}, {"Mini football pitch", "Description miniPitch"}},
	{{"Військова кафедра", "Опис Військова кафедра"}, {"Military Department", "Description military"}},
	{{"Лава закоханих", "Опис Лава закоханих"}, {"Love-bench", "Description loveBench"}},
	{{"Студентський сквер", "Опис Студентський сквер"}, {"Students' square", "Description studSquare"}},
	{{"ЕкоМаркет", "Опис ЕкоМаркет"}, {"EcoMarket", "Description ecoMarket"}},
	{{"Алея ректорів", "Опис Алея ректорів"}, {"Alley of rectors", "Description rectors"}},
	{{"Площа слави", "Опис Площа слави"}, {"Area of fame", "Description areaOfFame"}},
	{{"Ботанічний сад", "Опис Ботанічний сад"}, {"Botanical garden", "Description botanicalGarden"}},
	{{"Паркова зона", "Опис Паркова зона"}, {"Park area", "Description parkArea"}},
	{{"Кулиничі біля корпусу №5", "Опис Кулиничі біля корпусу №5"}, {"Kulinichi near the building №5", "Description kulin1"}},
	{{"Кулиничі біля корпусу №3", "Опис Кулиничі біля корпусу №3"}, {"Kulinichi near the building №3", "Description kulin2"}},
	// next will be

}

// BuildingInfo - cache for standart objects. Key got from URL
var BuildingInfo = map[string][]BuildingInfoType{}

// FillBuildingInfoMap - fills map of building info
func FillBuildingInfoMap() {
	// if !auth {
	for i, v := range TerritoryCache {
		BuildingInfo[v.TID] = BuildingInfoCache[i]
	}
	// } else {
	//
	// }
}

// GinLabels - cache for labels on page
var GinLabels = map[string][]string{
	"back":      {"Назад", "Back"},
	"close":     {"Закрити", "Close"},
	"main":      {"Повернутись на головну", "Go to Main page"},
	"showAreas": {"Показати всі об'єкти", "Show all objects"},
	"reg":       {"Реєстрація", "Registration"},
	"signIn":    {"Вхід", "Sign in"},
	"logout":    {"Вихід", "Logout"},
	"email":     {"Пошта", "Email"},
	"login":     {"Логін", "Login"},
	"password":  {"Пароль", "Password"},
	"remember":  {"Запам'ятати мене", "Remember me"},
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
	"internalError":          {"Помилка: На даний момент сервіс недоступний, будь-ласка спробуйте пізніше", "Error: The service isn`t available now, please try again later"},
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
