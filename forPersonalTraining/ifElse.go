package forPersonalTraining

import "fmt"

/*Реализуйте функцию DomainForLocale(domain, locale string) string,
которая добавляет язык locale как поддомен к домену domain.
Язык может прийти пустым, тогда нужно добавить поддомен en..
*/

func DomainForLocale(domain, locale string) string {
	if locale == "" {
		locale = "en"
	}
	return fmt.Sprintf("%s.%s", locale, domain)
}
