package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var quotes = []string{
	"You called in a fake propane emergency? That's a $50 fine after I report it.",
	"Dangit, Bobby!",
	"No got dang way, Bobby!",
	"An F in English? Bobby, you speak English!",
	"A poodle? Why don't you just get me a cat and a sex change operation?",
	"I sell propane and propane accessories.",
	"There better be a naked cheerleader under your bed!",
	"What the hell kind of country is this where I can only hate a man if he’s white?",
	"Why would anyone smoke weed when they could just mow a lawn?",
	"Bobby, if you weren’t my son, I’d hug you.",
	"He’s a football player, and football players know how to treat a woman right.",
	"I’m gonna kick your ass!",
	"If your sorority has to sell jam to buy beer, you’re drinking too much.",
	"BWAAAAHH!",
}

func Handler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(int64(time.Now().Nanosecond()))
	fmt.Fprintf(w, "%s", quotes[rand.Intn(len(quotes))])
}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}
