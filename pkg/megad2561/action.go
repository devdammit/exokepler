package megad2561

import "fmt"

type Action struct {
	value string
}

// Add базовая команда которая генерирует действие на смену состояния.
//
// Значения command:
//
// 0 - выключить
//
// 1 - включить
//
// 2 - изменить состояние на противоположное (переключить), т.е. если было включено выключить и наоборот.
//
// 3 - прямая синхронизация выхода со входом (кнопка нажата - лампа включена; кнопка отпущена - лампа выключена).
//
// 4 - обратная синхронизация выхода со входов (кнопка нажата - лампа выключена; кнопка отпущена - лампа включена)
//
// [0..255] - в случае с диммируемым портами, установить значение диммера (яркости света).
func (action *Action) Add(port int16, value int8) *Action {
	command := fmt.Sprintf("%v:%v", port, value)
	action.addCommand(command)

	return action
}

// AddPause В сценариях контроллер поддерживает работу с паузами.
//
// value 1 - 100 милисекунд
//
// Паузы в полном объеме и без ограничений работают только в сценариях по умолчанию (Action).
// Начиная с версии прошивки 4.16b8 паузы также поддерживаются и в командах, поступающих извне.
// Но в этом случае одновременно может выполняться только один сценарий, содержащий паузы.
//
// При выполнении сценария, содержащего паузу, работа контроллера не блокируется. Паузы выполняются в фоновом режиме.
func (action *Action) AddPause(value int16) *Action {
	command := fmt.Sprintf("p%v", value)

	action.addCommand(command)

	return action
}

// AddRepeat. Повтор сценария.
//
//
// Повтор записанного сценария несколько раз.
//
// Включить порт 7; пауза 0,5с; выключить порт 7; пауза 0,5с; повторить все это с самого начала еще 4 раза
// Таким образом порт включится и выключится 5 раз.
// Такую команду можно использовать для более компактной записи сложных сценариев.
func (action *Action) AddRepeat(value int8) *Action {
	command := fmt.Sprintf("r%v", value)

	action.addCommand(command)

	return action
}

// AddGlobal Управление всеми выходами.
//
// Например: value = 0 (выключить все выходы), value = 1 (включить все выходы).
func (action *Action) AddGlobal(value int8) *Action {
	command := fmt.Sprintf("a:%v", value)

	action.addCommand(command)

	return action
}

// GetValue получаем состояние собранного Action.
func (action *Action) GetValue() string {
	return action.value
}

/**
 * Private methods
**/

func (action *Action) addCommand(str string) {
	if len(action.value) > 0 {
		action.value += ";"
	}

	action.value += str
}
