package entry

/*
-Событие происходит в конкретный момент времени
-Событие имеет одно или более решений и требует выбор из доступных вариантов
-Событие должно быть завершено чтобы попасть в историю
-Событие всегда записывается в конец истории
-Событие может иметь причину

-реконструкция истории событий, должно приводить к одному и тому же результату


-решение может иметь описание
-решение может требовать проверку успеха
-решение может иметь описание успеха
-решение может иметь описание провала
-решение должно иметь описание успеха или описание провала
-решение должно приниматься пользователем, случайно по весу или по большему весу

*/
type EventDescr struct { //Событие происходит в конкретный момент времени
	Key         string `json:"Event"`
	Description string `json:"Description,omitempty"`
	Outcome     string `json:"Outcome"`
	Consequance string `json:"Consequance,omitempty"`
}
