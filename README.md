# go-htmx-templ-demo

Just putting a quick page together just to see how tailwind, templ, htmx and go works as a stack.

To get this running locally you need the templ and standalone tailwindcss cmd line tools installed and then run:

```
tailwindcss -i ./assets/tailwind.css -o ./assets/dist/styles.css && templ generate && go run *.go
```