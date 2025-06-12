import { useEffect, useState } from "react";
import "./App.css";
import Dashboard from "./Dashboard.jsx";
import Extract from "./Extract.jsx";
import History from "./History.jsx";
import Settings from "./Settings.jsx";
import Setup from "./Setup.jsx";
import System from "./System.jsx";
import Wanted from "./Wanted.jsx";

function App() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [status, setStatus] = useState("");
  const [authed, setAuthed] = useState(false);
  const [setupNeeded, setSetupNeeded] = useState(false);
  const [page, setPage] = useState("dashboard");

  useEffect(() => {
    fetch("/api/config").then((res) => {
      if (res.ok) setAuthed(true);
    });
    fetch("/api/setup/status")
      .then((r) => r.json())
      .then((d) => setSetupNeeded(d.needed));
  }, []);

  const login = async () => {
    const form = new URLSearchParams({ username, password });
    const res = await fetch("/api/login", { method: "POST", body: form });
    if (res.ok) {
      setStatus("logged in");
      setAuthed(true);
    } else {
      setStatus("login failed");
    }
  };

  if (!authed) {
    if (setupNeeded) {
      return <Setup />;
    }
    return (
      <div className="login">
        <h1>Subtitle Manager</h1>
        <input
          placeholder="Username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <button onClick={login}>Login</button>
        <p>{status}</p>
      </div>
    );
  }

  return (
    <div>
      <nav>
        <button onClick={() => setPage("dashboard")}>Dashboard</button>
        <button onClick={() => setPage("settings")}>Settings</button>
        <button onClick={() => setPage("extract")}>Extract</button>
        <button onClick={() => setPage("history")}>History</button>
        <button onClick={() => setPage("system")}>System</button>
        <button onClick={() => setPage("wanted")}>Wanted</button>
      </nav>
      {page === "settings" ? (
        <Settings />
      ) : page === "extract" ? (
        <Extract />
      ) : page === "history" ? (
        <History />
      ) : page === "system" ? (
        <System />
      ) : page === "wanted" ? (
        <Wanted />
      ) : (
        <Dashboard />
      )}
    </div>
  );
}

export default App;
