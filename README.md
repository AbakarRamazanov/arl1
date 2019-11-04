# Hi
Здесь реализация наивного классификатора Байеса

## How prepare to run
Структуризируем csv файл:
```cd sms_spam_collection
go run .\structuring_data.go <csvfilename>```
csvfilename - имя файла содержащего классы и текст для дальнейшей обработки в формате "class,text".
Будет создано несколько файлов.

## Run
Запустим на классификацию:
```cd sms_spam_collection
go run .\discrimination_by_bayes.go <filename>```
filename - имя файла с текстом. Результатом работы программы явится сообщение какой это класс.

## Example
```cd sms_spam_collection
go run .\structuring_data.go .\spam.csv
go run .\discrimination_by_bayes.go .\textHam.txt
go run .\discrimination_by_bayes.go .\textSpam.txt```
