Вот **чёткий, простой и практичный план**, как **подружить Go с .ipynb в VS Code на Windows**, **без Docker**, **на основе двух статей**:

* [Статья с ByteGoblin](https://bytegoblin.io/blog/setting-up-jupyter-notebook-kernels-for-golang-and-node-js)
* [Статья с Хабра](https://habr.com/ru/companies/skillfactory/articles/556938/)

---

## 🎯 Что ты получишь

✅ Возможность запускать **Go-код в Jupyter .ipynb прямо в VS Code**
✅ Без Docker
✅ Без лишних зависимостей
✅ Используешь **VS Code как полноценную среду анализа данных на Go** с интерактивными ячейками.

---

## 🛠️ Шаги

### 1️⃣ Установи Go

Если уже установлен, пропусти.

* Скачай с [https://go.dev/dl/](https://go.dev/dl/)
* Установи, проверь:

  ```powershell
  go version
  ```

---

### 2️⃣ Установи Jupyter

Так как ты используешь VS Code:

* Установи Python (если нет) или Miniconda.
* Установи Jupyter:

  ```powershell
  pip install notebook jupyter
  ```
* Установи расширение **Jupyter** в VS Code.

---

### 3️⃣ Установи gophernotes

Это ядро, позволяющее запускать Go в Jupyter.

#### 3.1 Установка

В PowerShell:

```powershell
go install github.com/gopherdata/gophernotes@latest
```

После установки `gophernotes.exe` окажется в:

```
C:\Users\<ТВОЙ_ПОЛЬЗОВАТЕЛЬ>\go\bin\gophernotes.exe
```

---

### 4️⃣ Создай папку ядра Go для Jupyter

Мы используем **методику с Хабра, так как она стабильнее**:

#### 4.1 Создай папку ядра

Создай:

```
C:\Users\<ТВОЙ_ПОЛЬЗОВАТЕЛЬ>\gophernotes_kernelspec
```

#### 4.2 Создай `kernel.json`

Создай в этой папке файл `kernel.json` со следующим содержимым:

```json
{
 "argv": [
  "C:/Users/<ТВОЙ_ПОЛЬЗОВАТЕЛЬ>/go/bin/gophernotes.exe",
  "{connection_file}"
 ],
 "display_name": "Go",
 "language": "go"
}
```

**Замени `<ТВОЙ_ПОЛЬЗОВАТЕЛЬ>` на твой реальный логин.**

---

### 4.3 Установи ядро

В PowerShell:

```powershell
jupyter-kernelspec install C:\Users\<ТВОЙ_ПОЛЬЗОВАТЕЛЬ>\gophernotes_kernelspec --user
```

Ты увидишь:

```
Installed kernelspec Go in C:\Users\<ТВОЙ_ПОЛЬЗОВАТЕЛЬ>\AppData\Roaming\jupyter\kernels\go
```

---

### 5️⃣ Проверка установки ядра

Проверь список:

```powershell
jupyter kernelspec list
```

Должно быть:

```
Available kernels:
  python3    ...
  go         C:\Users\<ТВОЙ_ПОЛЬЗОВАТЕЛЬ>\AppData\Roaming\jupyter\kernels\go
```

---

### 6️⃣ Использование в VS Code

✅ Убедись, что установлены:

* **Jupyter extension**
* **Jupyter Keymap** (по желанию)
* **Go extension** (для редактирования .go файлов)

Теперь:

1. Открой VS Code.
2. Создай новый файл:

   ```
   test_go.ipynb
   ```
3. В правом верхнем углу выбери **Go kernel** (он будет отображаться как `Go`).
4. В ячейке можно писать:

   ```go
   import "fmt"

   fmt.Println("Hello from Go in VS Code Jupyter!")
   ```

   И запускать `Shift + Enter`.

---

## ❌ Частые ошибки и как их избежать

✅ **Ошибка `flag provided but not defined: -kernel`:**
Ты используешь неправильный `kernel.json`.
Используй:

```json
"argv": [
  "C:/Users/<ТВОЙ_ПОЛЬЗОВАТЕЛЬ>/go/bin/gophernotes.exe",
  "{connection_file}"
]
```

**Без `--kernel`.**

✅ **VS Code не видит ядро Go:**

* Перезапусти VS Code.
* Убедись, что установлен `jupyter` в активном Python окружении.
* Проверь `jupyter kernelspec list` на наличие `go`.

✅ **Не нужно писать `func main()`.**
В `gophernotes` ячейки **исполняются как Go-скрипты, поэтому `main` не нужен**.
Используй:

```go
import "fmt"

fmt.Println("It works!")
```

---

## 📌 Итог:

✅ Установил `gophernotes`
✅ Создал `kernel.json` с корректным вызовом
✅ Установил ядро через `jupyter-kernelspec`
✅ Используешь `.ipynb` в **VS Code с Go kernel без Docker**
✅ Работаешь с Go как с интерактивной средой анализа данных и прототипирования.

---

Если хочешь, могу:
✅ Подготовить **готовый ZIP с папкой ядра** для вставки и установки командой.
✅ Прислать **рабочий `.ipynb` пример для Go** для теста в VS Code.
✅ Дать **краткий чеклист и настройку VS Code для Go+Jupyter для твоего Obsidian**.

Сообщи, если нужно, чтобы ускорить твою настройку без лишних движений.
