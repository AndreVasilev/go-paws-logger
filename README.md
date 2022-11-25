#
Специализированный логгер на основе (zap)[https://github.com/uber-go/zap]
```
pawsLogger "github.com/altatec-sources/go-paws-logger"

logger, err := pawsLogger.CreateLogger("logs", "app", cfg.Log.Level)

if err != nil {
	panic(err)
}
defer logger.Sync()
//palin
logger.Info("failed to fetch URL",
  zap.String("url", "http://example.com"),
  zap.Int("attempt", 3),
  zap.Duration("backoff", time.Second),
)
//sugar
sugar := logger.Sugar()
defer sugar.Sync()
sugar.Infof("Test message %d",1)
```
Функции передаются путь к папке хранения логоd - path , имя логгера - fileName, и желаемый уровень логирования.
- debug
- info
- error
- warn

В логгере применена суточная ротация.
Выходной файл: %path%/2022-11-25/%fileName%.json, где 2022-11-25 заменяется на текущее дату.
Логи автомитически удаляются с диска через 60 дней.

Функция возвращает *zap.Logger который можно испльзовать в plain режиме со строгим указанием типов логируемых параметров.
или в sugar режиме позволяющем испольщовать  функции форматирования логов.
plain режим обеспечивает максимальную скорость логирования, sugar - удобные методы логирования.

Sync() принудительно синхронизирует данные в логгере на диск. 
Если риск потери логов данных при останоке приложения не критичен Sync() можно не использовать.


