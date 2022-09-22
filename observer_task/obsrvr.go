// i was inspired by the alerts that were played on every phone recently
package main

import (
	"fmt"
)

type Subject interface {
	follow(observer Observer)
	unfollow(observer Observer)
	notifyAll()
}

type Item struct {
	observerList []Observer
	alert_name   string
	inStock      bool
}

func newAlert(alert_name string) *Item {
	return &Item{
		alert_name: alert_name,
	}
}

func (i *Item) updateAlert() {
	fmt.Printf("There will be strong %s in Astana city\n\n", i.alert_name)
	i.inStock = true
	i.notifyAll()
}

func (i *Item) follow(observer Observer) {
	i.observerList = append(i.observerList, observer)
}

func (i *Item) unfollow(o Observer) {
	i.observerList = removeFromList(i.observerList, o)
}

func (i *Item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.alert_name)
	}
}

func removeFromList(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

type Observer interface {
	update(string)
	getID() string
}

type Citizen struct {
	id string
}

func (c *Citizen) update(itemalert_name string) {
	fmt.Printf("Sending alert to Citizen %s for dangerous weather %s\n", c.id, itemalert_name)
}

func (c *Citizen) getID() string {
	return c.id
}

func main() {
	alertNotify := newAlert("rain")

	observerFirst := &Citizen{"Adilkhan"}
	observerSecond := &Citizen{"Martin"}

	alertNotify.follow(observerFirst)
	alertNotify.follow(observerSecond)

	alertNotify.updateAlert()
}
