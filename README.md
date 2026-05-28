**Game Concept**

The catch mechanic introduces randomness tied to a Pokemon's BaseExperience, making stronger Pokemon harder to catch. This is a simple but effective game design pattern - using existing data to drive probabilistic outcomes without needing a separate difficulty system.

The key skills here are struct design, map lookups, API consumption, caching, and building extensible CLI tools - all transferable to real backend work.

**Architecture**

The Pokedex is structured as a REPL (Read-Eval-Print Loop) - a pattern common in CLIs, database shells (like psql), and language interpreters. The core loop reads user input, parses it, dispatches to a command handler, and prints output. Shared state (the API client, cache, and caught Pokemon map) lives in a config struct that gets passed to every command, acting as a lightweight dependency injection pattern.

**CLI Command Interactivity**

Each command is a function with a consistent signature registered in a map. This is the "command pattern" - new commands can be added without touching the core loop. Input is cleaned and split into a command name and arguments, making it easy to support commands like catch pidgey or inspect charizard with a uniform dispatch mechanism.

**API and Caching**

The app fetches data from a real REST API (PokeAPI) and caches responses to avoid redundant network calls. Once a Pokemon is caught, its full data is stored locally in memory - so inspect never needs to hit the network again. This mirrors how real applications separate data fetching from data display.
