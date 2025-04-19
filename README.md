# Monorepo Training Project

This repository is a monorepo setup for experimenting with various web and data technologies. It utilizes [Moon](https://moonrepo.dev/) for managing the workspace, [Bun](https://bun.sh/) for JavaScript/TypeScript projects, and [UV](https://github.com/astral-sh/uv) for Python projects.

## Workspace Structure

```
.
├── .moon/              # Moon configuration
├── apps/
│   ├── fastify-react/  # Fastify backend with React frontend (Vite)
│   ├── next-fifteen/   # Next.js 15 application
│   ├── py-project/     # Python project using Polars
│   └── waku-project/   # Waku (React Server Components) application
├── packages/           # Shared packages (if any)
├── package.json        # Root Node.js package configuration (workspaces)
├── pyproject.toml      # Root Python project configuration (UV workspace)
├── tsconfig.json       # Root TypeScript configuration
├── bun.lock            # Bun lockfile
├── uv.lock             # UV lockfile
└── README.md           # This file
```

## Technologies Used

- **Monorepo Management:** [Moon](https://moonrepo.dev/)
- **JavaScript Runtime/Toolkit:** [Bun](https://bun.sh/)
- **Python Package Management:** [UV](https://github.com/astral-sh/uv)
- **Frontend Frameworks:**
    - [React](https://react.dev/)
    - [Next.js](https://nextjs.org/)
    - [Waku](https://waku.gg/)
- **Backend Frameworks:**
    - [Fastify](https://fastify.dev/)
- **Styling:** [Tailwind CSS](https://tailwindcss.com/)
- **Data Processing (Python):** [Polars](https://pola.rs/)
- **Language:** [TypeScript](https://www.typescriptlang.org/), [Python](https://www.python.org/)

## Getting Started

### Prerequisites

- [Bun](https://bun.sh/) installed
- [Python](https://www.python.org/) >= 3.13 installed
- [UV](https://github.com/astral-sh/uv) installed (`pip install uv` or `brew install uv`)

### Installation

Install dependencies for all projects:

```sh
# Install Node.js dependencies
bun install

# Install Python dependencies
uv sync
```

### Running Projects

You can run individual projects using Moon commands from the root directory.

**Example: Run the Fastify/React app development server:**

```sh
moon run fastify-react:dev
```

**Example: Run the Next.js app development server:**

```sh
moon run next-fifteen:dev
```

**Example: Run the Waku app development server:**

```sh
moon run waku-project:dev
```

**Example: Run the Python script:**

```sh
moon run py-project:dev
```

Refer to the `moon.yml` file within each `apps/*` directory for available tasks.

## Individual App READMEs

For more specific details about each application, please refer to their respective README files:

- [`apps/fastify-react/README.md`](apps/fastify-react/README.md) (Note: Needs creation)
- [`apps/next-fifteen/README.md`](apps/next-fifteen/README.md)
- [`apps/py-project/README.md`](apps/py-project/README.md)
- [`apps/waku-project/README.md`](apps/waku-project/README.md)