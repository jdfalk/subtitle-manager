# Importing Settings from Bazarr

Existing Bazarr users can migrate their configuration to Subtitle Manager to
avoid re-entering preferences. Bazarr exposes a REST API implemented with Flask.
The following endpoints are useful:

- `/api/system/settings` – returns the complete settings object (general
  options, proxy, authentication, notifier URLs, etc.).
- `/api/system/languages` – lists all languages with `enabled` flags.
- `/api/providers` – returns provider status and names.
- `/api/webhooks/radarr` and `/api/webhooks/sonarr` – trigger subtitle searches
  from Sonarr/Radarr events.

### Proposed Synchronization

1. **Fetch settings** from `/api/system/settings` using the user-provided Bazarr
   API key.
2. **Map fields** to the Viper configuration keys used by Subtitle Manager:
   - Network options (bind address, port, URL base) → `web.*` configuration.
   - Authentication and API key → `auth.*` fields.
   - Enabled providers and languages → `providers.*` and `languages.*` sections.
3. **Save the mapped configuration** via the `/api/config` endpoint or by
   writing the YAML file directly.
4. **Optional CLI command** `import-bazarr` to perform the import and start a
   migration wizard.

### Benefits

- Quick onboarding for users already running Bazarr.
- Ensures subtitle provider credentials and language profiles remain consistent.
- Reduces configuration errors by reusing verified settings.

Further implementation details should consider error handling for unreachable
Bazarr instances and partial imports when settings do not have direct
equivalents.

### Import Command Usage

Run the CLI tool to automatically migrate your configuration:

\```bash subtitle-manager import-bazarr http://localhost:6767 MY_API_KEY \```

The command fetches `/api/system/settings`, maps the values to Subtitle
Manager's configuration keys and writes them to your current config file.
