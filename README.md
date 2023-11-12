![Darius Logo](https://github.com/atridadl/Darius-Go/assets/88056492/56fd051f-075b-430a-b29b-5ba0a638946a)

# ✨ Darius ✨

🚀 A Web Application Template Powered by HTMX + Go Fiber + Tailwind 🚀

## Directory Stucture

- public/ - Static assets
- pages/ - Page route handlers
- lib/ - Libraries for use in route handlers
- templates/ - HTML Templates for use in route handlers
- templates/partials - HTML partials for use in templates
- templates/layouts - HTML layouts for use in templates
- api/ - API route handlers

## CSS Changes

Any CSS changes will require running the following command to generate the new CSS file:

```bash
go generate
```

## Development

To start the development server run:

```bash
go run main.go
```

Open http://localhost:3000/ with your browser to view the development server.
