/* Реализовать паттерн проектирования «Адаптер» на любом примере.
Описание: паттерн Adapter позволяет сконвертировать интерфейс
одного класса в интерфейс другого, который ожидает клиент.
Продемонстрируйте на простом примере в Go: у вас есть существующий
интерфейс (или структура) и другой, несовместимый по интерфейсу
потребитель — напишите адаптер, который реализует нужный интерфейс
и делегирует вызовы к встроенному объекту.
Поясните применимость паттерна, его плюсы и минусы, а также
приведите реальные примеры использования. */

package main

import (
	"fmt"
)

type Generator struct{}

func (gen *Generator) GeneratePower() {
	fmt.Println("Generator generates some power.")
}

type Transformer struct{}

func (transformer *Transformer) ConvertCurrent() {
	fmt.Println("Transformer converts voltage level.")
	
}

type EnergyAdapter interface {
	MakeNoise()
}

type GenAdapter struct {
	*Generator
}

func (adapter *GenAdapter) MakeNoise() {
	adapter.GeneratePower()
}

func NewGenAdapter(gen *Generator) EnergyAdapter {
	return &GenAdapter{gen}
}

type TransformerAdapter struct {
	*Transformer
}

func (adapter *TransformerAdapter) MakeNoise() {
	adapter.ConvertCurrent()
}

func NewTransformerAdapter(transformer *Transformer) EnergyAdapter {
	return &TransformerAdapter{transformer}
}

 func main() {
	station := [2]EnergyAdapter{NewGenAdapter(&Generator{}), NewTransformerAdapter(&Transformer{})}

    for _, object := range station {
        object.MakeNoise()
    }
}


/* Паттерн Адаптер используется для интеграции несовместимых интерфейсов.

Плюсы

- Гибкость: интегрирует legacy-код и сторонние API без переписывания.

- Изоляция: вся конвертация в одном месте, клиент видит чистый интерфейс.

- Расширяемость: новые адаптеры не ломают существующий код.

Минусы

- Сложность: лишний уровень абстракции усложняет отладку.

- Накладные расходы: делегация + конвертация данных.

- Риск перепроектирования: куча мелких адаптеров трудно поддерживать. */
