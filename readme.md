# 🐹 Go Problem Solving

Repositorio de ejercicios prácticos para aprender y dominar Go desde cero.

## 🎯 Objetivo

Este repositorio documenta mi proceso de aprendizaje de Go con el objetivo de ingresar a una big tech como Mercado Libre. El enfoque es aprender resolviendo problemas reales, progresando desde la sintaxis básica hasta arquitecturas limpias, concurrencia y sistemas distribuidos.

**Stack objetivo:** Go · Clean Architecture · Docker · Kubernetes

---

## 📁 Estructura del proyecto

```
go-problem-solving/
├── VariablesProblems/
│   ├── problem1/    # Calculadora de área
│   ├── problem2/    # Conversor de temperatura
│   ├── problem3/    # Swap sin variable temporal
│   ├── problem4/    # Segundos a horas y minutos
│   ├── problem5/    # Dígitos de un número
│   └── problem6/    # Calculadora de propina
├── IfProblems/      # 🔒 Próximamente
├── FunctionProblems/ # 🔒 Próximamente
└── README.md
```

Cada problema vive en su propia carpeta con su propio `main.go` y se ejecuta de forma independiente:

```bash
go run VariablesProblems/problem1/main.go
```

---

## ✅ Progreso

### 🟢 Variables — Completados (6/6)

| # | Problema | Estado |
|---|---|---|
| 1 | Calculadora de área de un círculo | ✅ |
| 2 | Conversor de temperatura (Celsius → Fahrenheit y Kelvin) | ✅ |
| 3 | Swap sin variable temporal | ✅ |
| 4 | Segundos a horas, minutos y segundos | ✅ |
| 5 | Extraer dígitos de un número de 3 cifras | ✅ |
| 6 | Calculadora de propina por comensal | ✅ |

### 🟡 Condicionales `if` — Pendientes (0/10)

| # | Problema | Estado |
|---|---|---|
| 1 | Par o impar | ⏳ |
| 2 | Mayor de tres números | ⏳ |
| 3 | Calculadora de IMC con categoría | ⏳ |
| 4 | Año bisiesto | ⏳ |
| 5 | Semáforo (switch de colores) | ⏳ |
| 6 | Clasificación de triángulo ⭐ HackerRank | ⏳ |
| 7 | Calculadora básica con operadores ⭐ LeetCode | ⏳ |
| 8 | Piedra, papel o tijera ⭐ HackerRank | ⏳ |
| 9 | Conversor de notas numéricas a letras ⭐ UVA | ⏳ |
| 10 | Validador de contraseña ⭐ LeetCode | ⏳ |

### 🔵 Funciones — Pendientes (0/10)

| # | Problema | Estado |
|---|---|---|
| 1 | Calculadora de área refactorizada con funciones | ⏳ |
| 2 | Conversor de unidades con funciones separadas | ⏳ |
| 3 | Función con múltiples retornos (min y max) | ⏳ |
| 4 | Función que calcula potencias sin `math.Pow` | ⏳ |
| 5 | Swap con retornos nombrados | ⏳ |
| 6 | FizzBuzz con función ⭐ HackerRank | ⏳ |
| 7 | Palíndromo ⭐ LeetCode | ⏳ |
| 8 | Suma de dígitos ⭐ HackerRank | ⏳ |
| 9 | Número de Fibonacci ⭐ LeetCode | ⏳ |
| 10 | Máximo Común Divisor y MCM ⭐ UVA | ⏳ |

---

## 🗺️ Roadmap de aprendizaje

- [x] Variables y tipos de datos
- [x] `fmt.Scan`, `fmt.Print`, `fmt.Printf`, `fmt.Sprintf`
- [x] Operadores aritméticos y módulo
- [ ] Condicionales `if / else if / else`
- [ ] `switch`
- [ ] Loops `for`
- [ ] Funciones y múltiples retornos
- [ ] Structs e interfaces
- [ ] Manejo de errores
- [ ] Concurrencia (goroutines y channels)
- [ ] Clean Architecture en Go
- [ ] APIs REST con Go
- [ ] Docker
- [ ] Kubernetes

---

> ⭐ Los problemas marcados con esta estrella siguen el estilo de plataformas competitivas como HackerRank, LeetCode o UVA, con casos de prueba definidos.