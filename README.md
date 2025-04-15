# Waku Project

A minimal starter template for building web applications with [Waku](https://waku.gg/), [React 19](https://react.dev/), [Bun](https://bun.sh/), [TypeScript](https://www.typescriptlang.org/), and [Tailwind CSS](https://tailwindcss.com/).

## Features

- Minimal setup with Waku (React framework)
- TypeScript support
- Bun as the runtime and package manager
- Tailwind CSS for styling
- Example Counter component
- Basic routing (Home and About pages)
- Static rendering

## Getting Started

### Prerequisites

- [Bun](https://bun.sh/) installed (v1.0 or later recommended)

### Install dependencies

```sh
cd apps/waku-project
bun install
```

### Development

Start the development server:

```sh
bun run dev
```

The app will be available at [http://localhost:3000](http://localhost:3000) by default.

### Build

To build the app for production:

```sh
bun run build
```

### Start (Production)

To start the production server:

```sh
bun run start
```

## Project Structure

```
.
├── apps/
│   └── waku-project/
│       ├── public/             # Static assets (e.g., images, favicon)
│       ├── src/
│       │   ├── components/     # React components (Counter, Header, Footer)
│       │   ├── pages/          # App pages (index.tsx, about.tsx, _layout.tsx)
│       │   └── styles.css      # Tailwind CSS styles
│       ├── package.json        # Project metadata and scripts
│       ├── postcss.config.js   # PostCSS configuration
│       ├── tsconfig.json       # TypeScript configuration
│       ├── bun.lock            # Bun lockfile
│       ├── moon.yml            # Moonrepo project configuration
│       └── README.md           # App-specific documentation (if any)
├── packages/                   # (Optional) Shared libraries or code
└── README.md                   # Project documentation (this file)
```

## About

This project demonstrates a minimal setup for building modern web applications using Waku, React, Bun, and TypeScript. It includes basic routing, a counter example, and uses Tailwind CSS for styling.

Feel free to use this as a starting point for your own projects!