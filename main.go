package main

import "fmt"

type Observer interface {
	update(string)
	getID() string
}

type Subject interface {
	subscribe(observer Observer)
	unsubscribe(observer Observer)
}

type Follower struct {
	id string
}

func (f *Follower) update(itemName string) {
	fmt.Printf("Sending email to customer %s for magazine %s\n", f.id, itemName)
}

func (f *Follower) getID() string {
	return f.id
}

type Magazine struct {
	observerList []Observer
	name         string
	published    bool
}

func newMagazine(name string) *Magazine {
	return &Magazine{
		name: name,
	}
}

func (m *Magazine) notifyAll() {
	for _, observer := range m.observerList {
		observer.update(m.name)
	}
}

func (m *Magazine) updateAvailability() {
	fmt.Printf("Item %s is now published\n", m.name)
	m.published = true
	m.notifyAll()
}

func (m *Magazine) subscribe(o Observer) {
	m.observerList = append(m.observerList, o)
}

func (m *Magazine) unsubscribe(o Observer) {
	m.observerList = removeFromList(m.observerList, o)
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

func main() {
	newMag := newMagazine("Forbes")

	observerFirst := &Follower{id: "murmillo228@gmail.com"}
	observerSecond := &Follower{id: "sagittarius2282286@gmail.com"}

	newMag.subscribe(observerFirst)
	newMag.subscribe(observerSecond)

	newMag.updateAvailability()
}
