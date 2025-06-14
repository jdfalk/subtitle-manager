// file: webui/src/components/DatabaseSettings.jsx
import { Box, Button, MenuItem, TextField, Typography } from "@mui/material";
import { useEffect, useState } from "react";

/**
 * DatabaseSettings allows selection of the storage backend and connection details.
 *
 * @param {Object} props - Component properties
 * @param {Object} props.config - Current configuration values
 * @param {Function} props.onSave - Callback invoked with updated values
 */
export default function DatabaseSettings({ config, onSave }) {
  const [backend, setBackend] = useState("pebble");
  const [dbPath, setDbPath] = useState("");
  const [sqliteFile, setSqliteFile] = useState("subtitle-manager.db");
  const [pgHost, setPgHost] = useState("");
  const [pgPort, setPgPort] = useState("");
  const [pgDB, setPgDB] = useState("");
  const [pgUser, setPgUser] = useState("");
  const [pgPass, setPgPass] = useState("");

  useEffect(() => {
    if (config) {
      setBackend(config.db_backend || "pebble");
      setDbPath(config.db_path || "");
      setSqliteFile(config.sqlite3_filename || "subtitle-manager.db");
      const pg = (config.database && config.database.postgres) || {};
      setPgHost(pg.host || "");
      setPgPort(pg.port || "");
      setPgDB(pg.database || "");
      setPgUser(pg.username || "");
      setPgPass(pg.password || "");
    }
  }, [config]);

  const handleSave = () => {
    const values = {
      db_backend: backend,
      db_path: dbPath,
      sqlite3_filename: sqliteFile,
      "database.postgres.host": pgHost,
      "database.postgres.port": pgPort,
      "database.postgres.database": pgDB,
      "database.postgres.username": pgUser,
      "database.postgres.password": pgPass,
    };
    onSave(values);
  };

  return (
    <Box sx={{ maxWidth: 500 }}>
      <Typography variant="h6" gutterBottom>
        Database Settings
      </Typography>
      <TextField
        select
        label="Backend"
        fullWidth
        sx={{ mb: 2 }}
        value={backend}
        onChange={(e) => setBackend(e.target.value)}
      >
        {['sqlite', 'pebble', 'postgres'].map((opt) => (
          <MenuItem key={opt} value={opt}>
            {opt}
          </MenuItem>
        ))}
      </TextField>
      <TextField
        label="Database Path"
        fullWidth
        sx={{ mb: 2 }}
        value={dbPath}
        onChange={(e) => setDbPath(e.target.value)}
      />
      {backend === 'sqlite' && (
        <TextField
          label="SQLite Filename"
          fullWidth
          sx={{ mb: 2 }}
          value={sqliteFile}
          onChange={(e) => setSqliteFile(e.target.value)}
        />
      )}
      {backend === 'postgres' && (
        <>
          <TextField label="Host" fullWidth sx={{ mb: 2 }} value={pgHost} onChange={(e) => setPgHost(e.target.value)} />
          <TextField label="Port" fullWidth sx={{ mb: 2 }} value={pgPort} onChange={(e) => setPgPort(e.target.value)} />
          <TextField label="Database" fullWidth sx={{ mb: 2 }} value={pgDB} onChange={(e) => setPgDB(e.target.value)} />
          <TextField label="Username" fullWidth sx={{ mb: 2 }} value={pgUser} onChange={(e) => setPgUser(e.target.value)} />
          <TextField label="Password" type="password" fullWidth sx={{ mb: 2 }} value={pgPass} onChange={(e) => setPgPass(e.target.value)} />
        </>
      )}
      <Button variant="contained" onClick={handleSave}>
        Save
      </Button>
    </Box>
  );
}
