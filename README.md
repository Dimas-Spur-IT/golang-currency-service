# golang-currency-service

Для более быстрой разработки использован фреймворк beego


1)поднять образ mysql под докером

docker run --name mysql -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=currency_service -p 3306:3306 -d mysql:latest

2)накатить миграции

bee migrate -conn="root:root@tcp(127.0.0.1:3306)/currency_service"

3) запустить скрипт парсинга курсов валют

go run scripts/update_rates.go

4) добавить крон команду для выполения апдейта рейтов каждый день

0 8 * * * cd /path_to_project && go run scripts/update_rates.go


5) запустить приложение 

bee run 