package entity

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var StorageRoot = "default"

const (
	data_countable = "countable"
	data_static    = "static"
)

/*

 */

type EntityFile struct {
	ID                  string             `json:"ID,omitempty"`
	Type                string             `json:"TYPE"`
	ClassifficationTree []string           `json:"Classification,omitempty"`
	Countable           map[string]float64 `json:"Countable"`
	Static              map[string]string  `json:"Static"`
	Events              []string           `json:"Events,omitempty"`
	/*
		*исчисляемое*
			сохраняет свойства/характеристики
			меняет колличество
			//деньги/статы/навыки
			map[string]int
		*статичное*
			сохраняет свойства/характеристики
			количество 1
		*история*
			фиксированная цепочка последовательных событий
			можно добавлять
			нельзя редактировать и удалять


	*/
}

type Entity interface {
	StaticData() map[string]string
	CountableData() map[string]float64
	//Type() string
	ID() string
	Classification() []string
	Events() []string
	//DestinationKey() string
	Validation() error
}

func New(entityType string, isTemplate bool, classes ...string) *EntityFile {
	e := EntityFile{}
	switch isTemplate {
	case true:
		e.ClassifficationTree = append(e.ClassifficationTree, "template")
	case false:
		id := time.Now().UnixNano()
		e.ID = fmt.Sprintf("%v", id)
	}
	if len(classes) > 0 {
		e.ClassifficationTree = append(e.ClassifficationTree, classes...)
	}
	e.Countable = make(map[string]float64)
	e.Static = make(map[string]string)
	e.Type = entityType
	return &e
}

func (e *EntityFile) WithData(data ...Data) error {
	for _, dt := range data {
		switch dt.dataType {
		case data_countable:
			if val, ok := e.Countable[dt.key]; ok {
				return fmt.Errorf("have %v value '%v' for key '%v'", dt.dataType, val, dt.key)
			}
			e.Countable[dt.key] = dt.flVal
		case data_static:
			if val, ok := e.Countable[dt.key]; ok {
				return fmt.Errorf("have %v value '%v' for key '%v'", dt.dataType, val, dt.key)
			}
			e.Static[dt.key] = dt.strVal
		default:
			return fmt.Errorf("bad argument: data{type: '%v'; key: '%v'; int: %v; str: '%v'}", dt.dataType, dt.key, dt.flVal, dt.strVal)
		}
	}
	return nil
}

type Data struct {
	key      string
	dataType string
	flVal    float64
	strVal   string
}

func DataCountable(key string, val float64) Data {
	return Data{
		key:      key,
		dataType: data_countable,
		flVal:    val,
		strVal:   "",
	}
}

func DataStatic(key string, val string) Data {
	return Data{
		key:      key,
		dataType: data_static,
		flVal:    0,
		strVal:   val,
	}
}

/*
путь расчитывается изходя из того именной это объект или шаблон
Шаблон - не имеет id (id=0) - помещается в дерево папок по классификации
и имеет имя файла равное статичному ключу Template
Именной - помещается в папку с главным типом и имеет имя файла id.json

/items/melee_weapons/swords/12344.json
/templates/melee_weapons/sword.json
*/
func (e *EntityFile) path() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic("can't get home: " + err.Error())
	}
	sep := string(filepath.Separator)
	root := home + sep + ".tabletoptools" + sep + "entity"
	tree := strings.Join(e.ClassifficationTree, fmt.Sprintf("s%v", sep))
	switch e.ID {
	case "":
		return root + sep + tree + ".json"
	default:
		return root + sep + e.Type + sep + fmt.Sprintf("%v.json", e.ID)
	}
}

func path(id string, classification []string) string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic("can't get home: " + err.Error())
	}
	sep := string(filepath.Separator)
	root := home + sep + ".tabletoptools" + sep + "entity"
	tree := strings.Join(classification, sep)
	switch id {
	case "":
		return root + sep + "templates" + sep + tree + ".json"
	default:
		return root + sep + tree + sep + fmt.Sprintf("%v.json", id)
	}
}

func Save(e Entity) error {
	path := path(e.ID(), e.Classification())
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("file creation failed: %v", err)
	}
	defer f.Close()
	bt, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling failed: %v", err)
	}
	_, err = f.Write(bt)
	if err != nil {
		return fmt.Errorf("writing failed: %v", err)
	}
	return nil
}

func Load(id string, classification []string) (*EntityFile, error) {
	path := path(id, classification)
	bt, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("file reading failed: %v", err)
	}
	e := &EntityFile{}
	if err := json.Unmarshal(bt, e); err != nil {
		return nil, fmt.Errorf("unmarshaling failed: %v", err)
	}
	return e, nil
}

func ToFile(e Entity) (*EntityFile, error) {
	if err := e.Validation(); err != nil {
		return nil, fmt.Errorf("validation failed: %v", err)
	}
	ef := EntityFile{}
	for _, event := range e.Events() {
		ef.Events = append(ef.Events, event)
	}
	ef.Countable = make(map[string]float64)
	for k, v := range e.CountableData() {
		ef.Countable[k] = v
	}
	ef.Static = make(map[string]string)
	for k, v := range e.StaticData() {
		ef.Static[k] = v
	}
	return &ef, nil
}
