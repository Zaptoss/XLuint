Бібліотека XLuint реалізує структуру великого числа та функції для роботи з цим типом.\
Структури: XLuint\
Методи структури: SetHex(hexString string) \[Для ініціалізації змінної], GetHex() -> string \[Для отримання значення зміної]\
Побітові операції: INV(number XLuint) -> XLuint \[інверсія], XOR(number1, number2 XLuint) -> XLuint, OR(number1, number2 XLuint) -> XLuint, AND(number1, number2 XLuint) -> XLuint, ShiftR(number XLuint, bits uint) -> XLuint, ShiftL(number XLuint, bits uint) -> XLuint\
Арифметичні операції: ADD(number1, number2 XLuint) -> XLuint, SUB(number1, number2 XLuint) -> XLuint, MUL(number1, number2 XLuint) -> XLuint, DIV(number1, number2 XLuint) -> XLuint, MOD(number, modulus XLuint) -> XLuint, POW(number, modulus XLuint, power uint) -> XLuint \[Піднесення в степінь за модулем]\

Зроблено в якості домашнього завдання для курсу Cryptography for Developers від Distributed Lab.
