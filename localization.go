package main

// ObjectInfo - cache for objects. Key got from URL
var ObjectInfo = map[string]ObjectInfoType{
	"edBuild1":        {"Корпус №1", "Description №1"},
	"edBuild2":        {"Корпус №2", "Description №2"},
	"edBuild3":        {"Корпус №3", "Description №3"},
	"edBuild4":        {"Корпус №4", "Description №4"},
	"edBuild5":        {"Корпус №5", "Description №5"},
	"edBuild7":        {"Корпус №7", "Description №7"},
	"edBuild7a":       {"Корпус №7-а", "Description №7-a"},
	"edBuild8":        {"Корпус №8", "Description №8"},
	"edBuild9":        {"Корпус №9", "Description №9"},
	"edBuild10":       {"Корпус №10", "Description №10"},
	"edBuild11":       {"Корпус №11", "Description №11"},
	"edBuild15":       {"Корпус №15", "Description №15"},
	"edBuild17":       {"Корпус №17", "Description №17"},
	"stadium":         {"Стадіон", "Description Stadium"},
	"candle":          {"Свічка", "Description Candle"},
	"edBuild1Terr":    {"Площа перед корпусом №1", "Description Площа перед корпусом №1"},
	"miniPitch":       {"Міні поле", "Description miniPitch"},
	"military":        {"Військова кафедра", "Description military"},
	"loveBench":       {"Лава закоханих", "Description loveBench"},
	"studSquare":      {"Студентський сквер", "Description studSquare"},
	"ecoMarket":       {"ЕкоМаркет", "Description ecoMarket"},
	"rectors":         {"Алея ректорів", "Description rectors"},
	"areaOfFame":      {"Площа слави", "Description areaOfFame"},
	"botanicalGarden": {"Ботанічний сад", "Description ecoMarket"},
	"parkZone":        {"Паркова зона", "Description parkZone"},
	"kulin1":          {"Кулиничі біля корпусу №5", "Description kulin1"},
	"kulin2":          {"Кулиничі біля корпусу №3", "Description kulin2"},
}

// ObjectInfoEN - cache for objects in English. Key got from URL
var ObjectInfoEN = map[string]ObjectInfoType{
	"edBuild1":        {"Educational building №1", "Description №1"},
	"edBuild2":        {"Educational building №2", "Description №2"},
	"edBuild3":        {"Educational building №3", "Description №3"},
	"edBuild4":        {"Educational building №4", "Description №4"},
	"edBuild5":        {"Educational building №5", "Description №5"},
	"edBuild7":        {"Educational building №7", "Description №7"},
	"edBuild7a":       {"Educational building №7-а", "Description №7-a"},
	"edBuild8":        {"Educational building №8", "Description №8"},
	"edBuild9":        {"Educational building №9", "Description №9"},
	"edBuild10":       {"Educational building №10", "Description №10"},
	"edBuild11":       {"Educational building №11", "Description №11"},
	"edBuild15":       {"Educational building №15", "Description №15"},
	"edBuild17":       {"Educational building №17", "Description №17"},
	"stadium":         {"Stadium", "Description Stadium"},
	"candle":          {"\"Candle\"", "Description Candle"},
	"edBuild1Terr":    {"Area in front of the building №1", "Description Площа перед корпусом №1"},
	"miniPitch":       {"Mini football pitch", "Description miniPitch"},
	"military":        {"Military Department", "Description military"},
	"loveBench":       {"Love-bench", "Description loveBench"},
	"studSquare":      {"Students' square", "Description studSquare"},
	"ecoMarket":       {"EcoMarket", "Description ecoMarket"},
	"rectors":         {"Alley of rectors", "Description rectors"},
	"areaOfFame":      {"Area of fame", "Description areaOfFame"},
	"botanicalGarden": {"Botanical garden", "Description botanicalGarden"},
	"parkArea":        {"Park area", "Description parkArea"},
	"kulin1":          {"Kulinichi near the building №5", "Description kulin1"},
	"kulin2":          {"Kulinichi near the building №3", "Description kulin2"},
}

// Buttons - cache for button words
var Buttons = map[string]string{
	"back": "Назад",
	"main": "Повернутись на головну",
}

// ButtonsEN - cache for button words in English
var ButtonsEN = map[string]string{
	"back": "Back",
	"main": "Go to Main page",
}

// Errors - cache for error sentences
var Errors = map[string]string{
	"errorTitle":     "Помилка",
	"objectNotFound": "Обраний об'єкт - не знайдено. Перевірте правильність посилання або спробуйте пізніше.",
}

// ErrorsEN - cache for error sentences in English
var ErrorsEN = map[string]string{
	"errorTitle":     "Error",
	"objectNotFound": "Selected object - not found. Please check that the link is correct or try again later.",
}
